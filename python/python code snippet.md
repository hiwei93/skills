# python code snippet

<!-- TOC -->

- [python code snippet](#python-code-snippet)
    - [获取文件夹中所有文件、文件夹](#%E8%8E%B7%E5%8F%96%E6%96%87%E4%BB%B6%E5%A4%B9%E4%B8%AD%E6%89%80%E6%9C%89%E6%96%87%E4%BB%B6%E3%80%81%E6%96%87%E4%BB%B6%E5%A4%B9)
    - [创建文件夹](#%E5%88%9B%E5%BB%BA%E6%96%87%E4%BB%B6%E5%A4%B9)
    - [读取文件](#%E8%AF%BB%E5%8F%96%E6%96%87%E4%BB%B6)
        - [按行读取文件](#%E6%8C%89%E8%A1%8C%E8%AF%BB%E5%8F%96%E6%96%87%E4%BB%B6)
        - [读取json文件](#%E8%AF%BB%E5%8F%96json%E6%96%87%E4%BB%B6)
        - [读取CSV文件](#%E8%AF%BB%E5%8F%96csv%E6%96%87%E4%BB%B6)
    - [写入文件](#%E5%86%99%E5%85%A5%E6%96%87%E4%BB%B6)
        - [生成json文件](#%E7%94%9F%E6%88%90json%E6%96%87%E4%BB%B6)
    - [正则表达式](#%E6%AD%A3%E5%88%99%E8%A1%A8%E8%BE%BE%E5%BC%8F)

<!-- /TOC -->

## 获取文件夹中所有文件、文件夹

``` python
import os
with os.scandir(path) as it:
    for entry in it:
        if entry.is_file() and entry.name.endswith('.xlsx'):
```

`os.scandir()`方法返回一个`DirEntry`对象的迭代器，通过`DirEntry`对象可以获取文件或文件夹的属性。

- 查看 [`os.scandir()`](https://docs.python.org/3.5/library/os.html?highlight=scandir#os.scandir)
- 查看 [`os.DirEntry`](https://docs.python.org/3.5/library/os.html?highlight=scandir#os.DirEntry)

## 创建文件夹

``` python
import os
if not os.path.exists(path):
    os.mkdir(path)
```

## 读取文件

使用`open`方法和`with`结构可以方便的读取文件，并释放资源：

``` python
with open(file_path, mode="r", encoding="utf-8") as f:
    content = f.read()
```

1. `open()`方法中`mode`参数为读取文件的方式，常用的有:

    - `r`：只读；
    - `r+`：既可读，也可写；
    - `rb`：以字节方式读取文件，此模式下不能设置编码格式，即不设置参数`encoding`。

2. `file`对象的`read()`方法，可以接受数字类型的参数值，如`r.read(1024)`：

   - 当`open()`的参数`mode=r`，则是读取的**字符**个数；
   - 当`open()`的参数`mode=rb`，则是读取的**字节**个数。

### 按行读取文件

有些情况下，文件内容是按行存放的，这些文件也可能比较大，这样的话按行读取就比较好处理了。

``` python
with open(file_path, mode="r", encoding="utf-8") as f:
   line = f.readline()
    while line:
        # do something ...
        line = f.readline()
```

``` python
with open(file_path, mode="r", encoding="utf-8") as f:
   lines = f.readlines()
    while lines:
        for line in lines:
            # do something ...
        lines = f.readlines()
```

上述情况，与下面的代码等价：

``` python
for line in open(file_path, mode="r", encoding="utf-8"):
    # do something ...
```

- `readline(size=-1)`和`readlines(hint=-1)`可以传入字节数来控制获取内容的大小。

### 读取json文件

``` python
import json
with open(json_path, mode="r", encoding="utf-8") as f:
    content = json.load(f)
```

### 读取CSV文件

[csv doc](https://docs.python.org/3.5/library/csv.html#id3)

1. 需要导入包：`csv`;
2. `open`函数需要设置参数`newline=''`;
3. `csv.reader(f, dialect='excel-tab')`设置方言（dialect）为excel-tab可以读取以tab键为分割的txt文件；
4. `csv.DictReader(f)`读取文件，可以根据header获取相应列的内容

## 写入文件

``` python
with open(file_path, mode="w", encoding="utf-8") as f:
    f.write(content)
```

1. `open()`方法中`mode`参数为读取文件的方式，常用的有:

    - `w`：只写，文件存在会被重写，文件不存在则创建；
    - `w+`：既可读，也可写；
    - `wb`：以字节方式写文件，此模式下不能设置编码格式，即不设置参数`encoding`；
    - `a`：追加模式，文件存在则追加文件内容，不存在则创建；
    - `ab`：字节方式的追加模式。

### 生成json文件

``` python
import json
with open('./output/' + file_name, 'w', encoding="utf-8") as f:
    json.dump(content, f)
```

- 设置`json.dump()`方法的参数`ensure_ascii=False`，就可以正常导出json中的中文了。
- 设置`json.dump()`方法的参数`separators=(',', ':')`，就可以将导出紧凑的json格式了。

## 正则表达式

1. 字符串替换匹配项：

    ``` python
    import re
    pattern = r''
    string = re.sub(pattern, "", string)
    ```

2. 使用正则表达式语法`(?P<name>...)`

    可以将匹配到的内容与`name`组成一个dict

    ``` python
    import re
    m = re.match(r"(?P<first_name>\w+) (?P<last_name>\w+)", "Malcolm Reynolds")
    m.group('first_name') # 'Malcolm'
    m.group('last_name') # 'Reynolds'
    m.groupdict() # {'first_name': 'Malcolm', 'last_name': 'Reynolds'}
    ```

3. 使用`start()`和`end()`截取字符串

    ``` python
    email = "tony@tiremove_thisger.net"
    m = re.search("remove_this", email)
    email[:m.start()] + email[m.end():] # 'tony@tiger.net'
    ```