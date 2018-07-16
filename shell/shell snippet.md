# shell snippet

## 文件中搜索指定字符串

``` bash
grep target_string example.txt
```

### 1. 在目录中所有文件中查找

``` bash
grep -r update /etc/example
```

### 2. 统计匹配字符串所在行数

``` bash
grep -c target_string example.txt
```

### 3. 将找到的行写入到其他文件中

``` bash
grep target_string example.txt > other_file.txt
```

- `other_file.txt`文件可以不存在。
- 涉及到[Shell 输入/输出重定向](http://www.runoob.com/linux/linux-shell-io-redirections.html)。

## 计算文件的Byte数、字数、或是列数

### 计算文件行数

``` bash
wc -l filename
```

## 将文件分割成若干份

### 按行分割

``` bash
split -行数 需要分割的文件名 输出文件的前缀名
```