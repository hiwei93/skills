# python文件 & 文件夹处理

## 获取文件夹信息

使用[os.scandir](https://docs.python.org/3/library/os.html#os.scandir)获取文件夹信息：

``` python
with os.scandir(path) as it:
    for entry in it:
        if entry.is_file() and entry.name.endswith('.xlsx'):
```