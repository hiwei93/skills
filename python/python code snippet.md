# python code snippet

## 获取文件夹中所有文件、文件夹

``` python
import os
with os.scandir(path) as it:
    for entry in it:
        if entry.is_file() and entry.name.endswith('.xlsx'):
```

使用`os.scandir()`可以获取文件夹中的所有文件信息。

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

### 读取json文件

``` python
import json
with open(json_path, mode="r", encoding="utf-8") as f:
    content = json.load(f)
```

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
