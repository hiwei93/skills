# python code snippet

## 获取文件夹中所有文件、文件夹

``` python
with os.scandir(path) as it:
    for entry in it:
        if entry.is_file() and entry.name.endswith('.xlsx'):
```

使用`os.scandir()`可以获取文件夹中的所有文件信息。

## 创建文件夹

``` python
if not os.path.exists(path):
    os.mkdir(path)
```

## 读取json文件

``` python
with open(jsonPath, mode="r", encoding="utf-8") as f:
    content = json.load(f)
```

## 生成json文件

``` python
with open('./output/' + fileName, 'w+', encoding="utf-8") as f: 
    json.dump(content, f)
```