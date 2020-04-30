# Python 分析器

python version 3.7

原文：[The Python Profilers](https://docs.python.org/3.7/library/profile.html#pstats.Stats.strip_dirs)

## 分析器介绍

cProfile和profile提供Python程序的确定性分析（deterministic profiling，TODO：什么叫确定性分析？）。*分析*是一列的统计，这一系列统计描述了程序的各个部分执行的频率和时间。这些统计信息通过pstats模块可以格式化为报告。

Python标准库提供了同一分析接口的两种不同的实现：

1. 建议大多数用户使用cProfile；该实现是具有合理开销的C扩展，适用于分析长时间运行的程序。基于由Brett Rosen和Ted Czotter贡献的lsprogf。
2. profile，一个纯Python模块（TODO：意思是纯用Python写的？），其接口被cProfile模。profile相较于cProfile增加了大量的分析程序的开销。如果你想尝试拓展分析器，使用这个模块会更容易一些。由Jim Roskind最初设计和编写。

> 注意：分析器模块是设计用来为给定程序提供执行分析的，而不是用来做基准测试的（如果需要进行基准测试，timeit可以提供合理的准确结果）。该情况在“对依赖C代码的Python代码进行基准测试”的情况下尤其突出：分析器引入Python代码的开销，但是对于C层次的方法没有开销，所以C代码看起来会比任何Python代码都快。

## 用户快速入门手册

本节为不想阅读手册的用户提供。本节提供了简短的概览，用户可以逐渐的将分析模块应用到已有的应用中。

对于只接受一个参数的方法，可以使用如下方法进行分析：

``` python
import cProfile
import re
cProfile.run('re.compile("foo|bar")')
```

（如果cProfile不能在你的系统运行，请使用profile）

### 分析结果说明

上面的代码将会运行`re.compile()`并且将分析结果打印出来，结果如下：

TODO:样例

第一行说明本次分析监控了197各方法调用。在这些方法调用中，192个是*原始调用（primitive）*，原始调用是指非递归的调用。接下来的一行：`Ordered by: standard name`，指明使用最右列的字符串文本对输入进行排序。列标题有：

- ncalls 调用的次数

- tottime 给定方法执行的时间（排除了子函数调用花费的时间）（TODO：没有翻译明白，到底包不包含子函数的调用）

- percall tottime 除以 ncalls 的商数（即平均每次调用的时间）

- cumtime 该方法及其子方法所花费的累计时间（即从调用到退出）。对于递归的方法，该计算也是准确的。

- percall cumtime 除以 原始调用的数量

- filename:lineno(function) 提供每个方法的相应数据。

如果第一列出现了两个数字（比如`3/1`），意味着该方法是递归的。第二个数字是原始调用的系数，第一个数字是总调用数量。注意，如果方法没有递归调用，前后两个数字是一样的，只有一个数字会打印出来。

### 分析结果输出到文件

可以在调用run()方法的时候指定一个filename，从而将分析结果保存到文件，而不是直接打印出来。

```python
import cProfile
import re
# restats 为文件名
cProfile.run('re.compile("foo|bar")', 'restats')
```

[pstats.Stats](https://docs.python.org/3.7/library/profile.html#pstats.Stats)类从文件中读取分析结果，并以多种方式格式化。

### 分析指定脚本

cPython文件也可以作为脚本进行调用去分析另外一个脚本。如下例：

```bash
python -m cProfile [-o output_file] [-s sort_order] (-m module | myscript.py)
```

- `-o`：将分析结果写到指定文件而不是标准输出中。
- `-s`：指定一种[sort_stats()](https://docs.python.org/3.7/library/profile.html#pstats.Stats.sort_stats)提供的排序值，用于对输出进行排序。只有当`-o`没有指定的时候才有效。
- `-m`：指定需要分析的模块而不是脚本文件。（`-m`选项是3.7版本新增的）

注：输出结果使用文本文件无法查看。

### 格式化分析结果文件

pstats模块的Stats类有多种方法可以用于处理并打印保存到文件中的分析结果：

``` python
import pstats
from pstats import SortKey

p = pstats.Stats('restats')  # 读取文件
p.strip_dirs().sort_stats(-1).print_stats()
```

- strip_dirs()方法移除所有模块名称中额外的路径信息。
- sort_stats()方法根据打印的标准格式module/line/name字符串对所有的条目进行排序。
- print_stats()方法打印所有的统计结果。

也可以尝试以下的排序方式：

``` python
p.sort_stats(SortKey.NAME)
p.print_stats()
```

第一个方法会使用方法名称对列表排序，第二个方法会打印统计的结果。下面的例子是有实验性质的调用：

``` python
p.sort_stats(SortKey.CUMULATIVE).print_stats(10)
```

该例子使用函数的累计运行时间对分析结果排序，并且只打印排在前10的行。如果你想知道什么算法最花时间，可以用上面的例子。

如果想查看哪些函数循环运行了多次，并且花了很多时间，可以运行：

``` python
p.sort_stats(SortKey.TIME).print_stats(10)
```

根据每个函数花费的时间排序，然后打印前10的函数。

可以尝试运行：

``` python
p.sort_stats(SortKey.FILENAME).print_stats('__init__')
```

上例会使用文件名进行排序，然后仅打印类的init方法的统计信息（因为这些方法拼写为__init__）。最后一个示例，可以尝试运行：

``` python
p.sort_stats(SortKey.TIME, SortKey.CUMULATIVE).print_stats(.5, 'init')
```

使用time作为主键、累计时间为次键进行排序，并输出一些统计信息。详细说，统计列表先缩减为原来的50%（由参数`.5`得到），然后只保留包含`init`的行，然后打印获得的子列表。

如果你想知道什么函数调用了上面的函数，可以这样运行：（p仍然按照上面的条件排序）

``` python
p.print_callers(.5, 'init')
```

然后可以得到每个列出函数的一系列调用者。

如果想使用更多的功能，就需要阅读手册了，或者猜一下下面函数的作用：

``` python
p.print_callees()  # 获取方法调用的子方法
p.add('restats')  # 将统计结果添加到分析文件中
```

作为脚本调用，pstats模型是一个读取和验证分析文件的统计浏览器。pstats由一个简单的命令行接口（使用cmd实现的）和交互式帮助。

## 什么是确定性分析（deterministic profiling）

*确定性分析*旨在反映这样一个事实：所有函数调用、函数返回和异常事件都被监控，并且对这些事件之间的间隔进行准确计时（在这些间隔中执行用户代码）。相反，*统计分析*（不是本模块完成的）随机抽样有效指令指针，并且推断时间都花费在什么地方。传统上，统计分析技术会有更少的开销（因为不需要检测代码，TODO 没懂），但是只会给出在什么地方花费了多长时间（TODO：那么确定性分析提供了而外的信息么？）

Python中，因为在程序执行过程中解释器一直处于活跃状态，确定性分析不需要instrumented code(TODO：什么是instrumented code？)的存在。Python为每个事件自动提供了一个钩子（可选的回调）。另外，Python的及时性语言特点增加了很多的开销，以至于确定性分析在典型应用上的开销显得微不足道。结论为，确定性分析并不会有很大开销，而且还提供了Python程序执行的大量的运行时统计。

- 调用次数统计（call count statistics）可以用于识别代码中的bug（意外计数，TODO：是指调用的次数过多，可以认为是死循环？），还可以用来识别内联扩展点（TODO：什么叫inline-expansion points？）（调用次数多）。

- 内部时间统计（internal time statistics）可以用来识别“热循环”（TODO：什么叫hot loops），需要仔细优化。

- 累计时间统计（cumulative time statistics）可以用于识别算法选择的高级别错误（TODO：如何识别？）。注意，在分析器中对累计时间的异常处理允许将递归算法的统计与迭代算法的统计直接比较。（如何比较？是指两种方法都实现，然后分别运行，获得分析结果再进行比较？？？）

## 局限

一个局限性与计时信息的准确性有关。对于准确性，确定性分析有一个根本的问题。最明显的限制就是底层的“clock”（通常）只以.001秒的速率计时。因此，没有任何其他的测量工具比地秤的clock更准。如果进行了足够的测量，“误差”将趋于平均。不幸的是，消除第一个错误会引起第二个错误。（TODO：完全不知其所云）


## The Stats Class

### print_stats(*restrictions)

Stats类的方法按照profile.run()定义中的描述打印报告。

（如果提供了）参数可以限制列表，输出重要的条目。初始状态，列表是全部的分析后的方法。每个限制要么是一个整型（用于限制行数），要么是一个介于0.0到1.0（包含1）之间的小数（用于选择一定比例的行），或者是一个字符串，该字符串将会解释为正则表达式（用于匹配标准名称进行输出）。如果提供了多个限制，则会依次实施。

### print_callers(*restrictions)

### print_callees(*restrictions)

## 重要的点

了解这些就能够执行分析了

1. 提供哪些分析方法，排序、过滤、调用关系。
2. 分析的统计参数有哪些？分别对应性能的什么地方？
3. 性能分析常关注哪些信息，如何获取这些信息？