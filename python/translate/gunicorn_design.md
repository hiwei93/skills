# Gunicorn 设计

原文：[Design](http://docs.gunicorn.org/en/stable/design.html)

本文简单的描述了Gunicorn的架构

## 服务器模块

Gunicorn基于 pre-fork worker （TODO: pre-fork worker 是什么？）模型。这意味着会有一个主进程管理一系列worker进程。主进程不会知道任何关于独立客户端（TODO：什么是individual client?）的情况。所有的请求和响应完全由worker进程处理。

### 主进程

主进程是一个简单的循环，这个循环监听多个进程的信号并且做出相应的响应。主进程通过监听信号（如：TTIN，TTOU和CHLD）来管理一系列的运行worker。

TTIN和TTOU来通知主进程增加或减少正在运行的worker数量。

CHLD表明一个子进程（TODO：与work process有什么区别么？）已经终止，这种情况下，主进程会自动重启失败的worker。

### 同步workers

最基本的、默认的worker类型是同步worker类，同步worker类一次处理一个请求。该模型最容易推断，因为任何错误最多只会影响一个请求。尽管正如在下面描述的那样，一次只处理一个请求需要一些有关应用程序如何编写的假设。

同步worker不支持持久连接，每个连接在响应发送后会关闭（即使程序中人工在请求头部添加了`Keep-Alive`或者`Connection: keep-alive`)

### 异步workers

异步worker基于Greenlets (via Eventlet and Gevent). Greenlets是一个适用于python的协作式多线程（TODO：什么是协作式多线程）。通常，一个应用需要能够无需更改就可以使用这些worker class（TODO：没懂）。

为了完全支持Greenlet的应用可能需要作息以恶调整。当使用Gevent和Psycopg时，应该保证安装并设置了psycogreen。

因为依赖于原始的未修补行为，这些应用程序可能完全不兼容。（一点都不懂）

### Tornado workers（龙卷风workers）

还有一种Tronado worker类，该worker可以使用Tornado框架来开发应用。

尽管Tornado worker可以用作WSGI应用的服务，但是并不建议这么做。

### 异步IO workers

异步IO Worker与Python 3兼容。

gthread worker 是一个thread worker。 该worker在main loop中接受一个连接，接受到的连接会作为连接工作添加到线程池中。
保持活动的连接会被放回到循环中等待一个事件。如果在连接超时前没有事件发生，连接将会关闭。

可以将你的应用使用aiohttp的Web.Application API，并且使用 aiohttp.worker.GunicornWebWorker worker。

## 选择合适的worker类型

默认的同步worker假设你的应用在CPU和网络带宽方面受资源限制。通常这意味着你的程序不需要做任何会花费不确定时间的事。比如，向互联网发送一个请求。有时，外部网络会因为客户端请求堆积在服务器，这会导致外部网络失败。因此，在这些场景下，任何以API方式接受请求访问的网络应用都会受益于同步worker。

正是基于资源限制的假设，在默认Gunicorn配置前需要一个缓存代理。如果将同步worker暴露到网络，通过创建负载，以“涓涓细流”的方式将数据传入到服务，从而将DOS攻击弱化到微不足道。如果对相关内容感到好奇，[Hey](https://github.com/rakyll/hey)是一种该类型的负载。

一些需要使用同步worker的场景：

- 会进行长时间阻塞调用的应用（比如请求外部Web服务）
- 接受直接来自于互联网的请求
- 流式请求和响应
- 长时间轮询（TODO：Long polling，不明白什么意思）
- Web socket
- Comet（TODO：什么意思？彗星？长尾？？）

## 设置多少workers合适？

不要将worker的数量设置为期望的客户端数量的比例。Gunicorn只需要4-12个工作进程，就可以每秒钟处理上百个或者上千个请求。

Gunicorn在处理请求时，依赖于操作系统来提供负载均衡。通常，推荐使用公式 `(2 * $num_cores) + 1` 计算worker的数量。虽然不太科学，但是公式基于这样的假设：对于给定计算核心数量，一个worker在读取或者写入socket的时候，另一个worker正在处理请求。

显然，对于特定的硬件和应用也会影响到最佳worker的数量。建议一开始使用上面的猜想，在应用在负载时，使用TTIN和TTOU信号进行调整。

请注意，过多的worker进程会影响系统资源，从而降低整个系统的吞吐量。

## 设置多少线程

从Gunicorn 19后，线程选项可以用于多线程处理请求。线程选项使用gthread worker。使用线程的一个好处是，请求在worker超时后请求还在处理过程中的时候后可以通知主进程该worker未冻结，不应该被杀死。根据同不同的系统，使用多线程、多worker进程，获取其他的组合方式，可能会产生最佳结果。比如，当使用线程的时候CPython的表现可能就不如Jython，因为他们对线程的实现是不同的。使用线程代替进程是减少Gunicorn内存使用的好方法，同时仍然允许应用使用重载信号进行升级因为应用的代码会在个worker之间共享，但是尽在worker进程中进行加载。（与使用预加载设置不同，预加载设置会在master进程中加载代码）

TODO：后面程序代码加载的部分没有看懂。
