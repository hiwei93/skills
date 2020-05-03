# python 异步IO

本文为个人学习python asyncio模块的内容，可以看为python异步编程的入门。本文将介绍使用asyncio模块需要了解的*3个主要的awaitable对象*和*3个运行协程的机制*，并介绍个人在实际开发过程中使用到的两个异步开发的第三方包。

## Awaitables

awaitable对象是可以用于await表达式的对象。有三个主要的awaitable对象：Coroutines，Tasks和Futures。

### 1. Coroutines

协程的概念在异步编程中很重要，常常作为异步运行的顶层入口。有两类协程：

1. 协程函数：使用`async def`声明的函数
2. 协程对象：协程函数返回的对象

### 2. Tasks

Task用于协程的同时调度。使用`asyncio.create_task()`类似的函数将协程包裹到Task中时，协程会立即自动调度运行。

创建Task：

``` python
import asyncio

# for all Python version
task = asyncio.ensure_future(coroutine_func())

# for Python 3.7+
task = asyncio.create_task(coroutine_func())

await task
```

### 3. Futures

Future表示异步操作的最终结果，属于底层的awaitable对象。

一般不需要在应用层创建Future对象。不过asyncio有些API会接受或者返回Future对象。

[loop.run_in_executor()](https://docs.python.org/3.7/library/asyncio-eventloop.html#asyncio.loop.run_in_executor)函数就是一个返回Future对象的函数。

## 运行协程

直接调用协程函数会返回一个协程对象，不会直接运行该协程。想要运行协程需要将其通过各种机制添加到事件循环中，事件循环会对其进行调度与运行。

asyncio提供了3个运行协程的主要机制：

### 对于顶层的协程

协程的调度与运行是基于事件循环（even loop）的，因此对于顶层的协程，需要为其找一个事件循环进行运行。

python 3.7之前，运行顶层的协程如下：

``` python
if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(top_level_coroutine_func())
    loop.close()
```

python 3.7之后，对事件循环的获取和关闭等上下文操作进行了封装，协程的操作更加简单：

``` python
if __name__ == '__main__':
    asyncio.run(top_level_coroutine_func())
```

### 2. 使用await表达式

await表达式会暂停当前协程运行，等待表达式中awaitable对象返回。

await表达式只能用在协程函数内，可以将多个协程串联起来。

### 3. 将协程作为Task运行

使用`asyncio.ensure_future()`方法（对于所有Python版本）或者`asyncio.create_task()`方法（对于Python 3.7+）将协程作为Task同时运行。

请看[Awaitables: Tasks](#tasks)

如果需要同时运行多个协程，可以使用`asyncio.gather()`函数，常用的模式如下：

``` python
import asyncio

async def count(num):
    print(f"this is number: {num}")
    await asyncio.sleep(1)  # do something should be awaited

async def main(loop):
    tasks = []
    for i in range(5):
        tasks.append(count(i))

    await asyncio.gather(*tasks)

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())
    loop.close()
```

`asyncio.gather()`函数会将传入的协程包装成Task进行调度。如果通过该函数传入的协程有返回值，`asyncio.gather()`函数在运行完后，返回结果包含所有传入协程的结果，结果的顺序与对应协程传入顺序一致。详细了解`asyncio.gather()`请看[Running Tasks Concurrently](https://docs.python.org/3.7/library/asyncio-task.html#id6)。

## 常用第三方协程库

### 异步网络-aiohttp

[aiohttp 文档](https://docs.aiohttp.org/en/stable/client_quickstart.html)

``` python
import aiohttp
import asyncio

async def fetch(session, url):
    async with session.get(url) as response:
        return await response.text()

async def main():
    async with aiohttp.ClientSession() as session:
        html = await fetch(session, 'http://python.org')
        print(html)

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(main())
    loop.close()
```

### 异步访问MongoDB-Motor

[Motor 文档](https://motor.readthedocs.io/en/stable/tutorial-asyncio.html)

``` python
import asyncio
import motor.motor_asyncio

client = motor.motor_asyncio.AsyncIOMotorClient()
db = client['test_database']
collection = db['test_collection']

async def do_insert():
    document = {'key': 'value'}
    result = await db.test_collection.insert_one(document)
    print('result %s' % repr(result.inserted_id))

if __name__ == '__main__':
    loop = asyncio.get_event_loop()
    loop.run_until_complete(do_insert())
    loop.close()

```

### 其他

- 异步文件操作：[aiofiles](https://github.com/Tinche/aiofiles)

## 注意事项

1. 异步的执行是在线程上进行的，asyncio模块提供的大部分API不是线程安全的，如果需要跨线程执行可以参看[Scheduling From Other Threads](https://docs.python.org/3.7/library/asyncio-task.html#scheduling-from-other-threads)
2. 异步编程适用于IO密集型的程序，使用异步编程运行CPU密集型的程序可能会导致效率不佳，可以使用python提供的Executor根据不同的程序特点选择使用线程或者进程运行程序，详细可以参看[Running Blocking Code](https://docs.python.org/3.7/library/asyncio-dev.html#running-blocking-code)，[Executing code in thread or process pools](https://docs.python.org/3.7/library/asyncio-eventloop.html#executing-code-in-thread-or-process-pools)

## 结论

asyncio模块提供了基于事件驱动的异步编程模式和API，随着python版本的迭代，asyncio API也更加容易理解和使用。合理使用异步模式可以高效利用单线程的计算能力，大幅提高I/O密集型程序的执行速度。但是不合理的使用，比如不合适的await和计算密集型的程序都可能会影响异步编程的表现。

会用和合理使用之间还有很多的路要走，还需要深入了解事件驱动的机制，了解coroutine的调度及运行的方式，明白Task和Future出现的根本原因等等。

## 参考

- [python docs: Coroutines and Tasks](https://docs.python.org/3.7/library/asyncio-task.html)
- [python docs: Developing with asyncio](https://docs.python.org/3.7/library/asyncio-dev.html)
- [python docs: Event Loop](https://docs.python.org/3.7/library/asyncio-eventloop.html#asyncio.loop.run_in_executor)
