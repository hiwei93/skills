# golang learning

- 官网文档：`https://golang.org/doc/`
- 中文网文档：`https://go-zh.org/doc/`

## 1. install

### 1. mac

1. 从官网下载安装包；
1. 点击运行即可，mac系统会安装到`/usr/local/go`文件夹下。
1. 安装过程会自动将`/usr/local/go/bin`添加到环境变量。

### 2. Ubuntu

1. 官网下载源码压缩包；
1. 使用以下命令解压到`/usr/local`目录下；

   ``` bash
   sudo tar -C /usr/local -xzf go1.9.2.linux-amd64.tar.gz
   ```

1. 添加go环境变量

   ``` bash
   export PATH=$PATH:/usr/local/go/bin
   ```

#### 常见问题

环境变量添加到~/.profile，重启terminal后无效。

解决办法：可以讲环境变量添加到~/.bashrc中。

## 2. uninstall

### 1. mac

1. 移除`/usr/local/go`文件夹；
1. 删除环境变量：移除文件`etc/paths.d/go`。

### 2. Ubuntu

1. 移除`/usr/local/go`文件夹；
1. 编辑文件`/etc/profile or $HOME/.profile`，移除环境变量。

## 3. 阅读[How to Write Go Code](https://golang.org/doc/code.html)

中文版：[如何使用Go编程](https://go-zh.org/doc/code.html)