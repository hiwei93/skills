# GOPATH environment variable

Go路径用于处理导入语句。Go路径由`go/build`时间并记录在文件中。

GOPATH环境变量列出了寻找Go代码的地方。在Unix上，GOPATH环境变量是以冒号`:`分隔的字符串。在Windows上，是以分毫`;`分隔的字符串。在Plan 9上，该值是一个列表。

如果还没有设置环境变量，GOPATH默认为用户主目录下名为go的子目录（Unix中为$HOME/go，windows中为%USERPROFILE%\go），另外就是包含Go分配的目录。运行"go env GOPATH"命令查看当前的GOPATH。

查看`https://golang.org/wiki/SettingGOPATH`自定义GOPATH。

GOPATH列表中的每个目录必须有一个规定的结构：

- src目录保存远吗。src下的路径决定了导入路径或可执行文件的名称。
- pkg目录保存已安装的包对象。与Go的树形结构一样，每个目标操作系统和体系结构都有自己的pkg（pkg/GOOS_GOARCH）子目录。
- 如果DIR在GOPATH列出的目录中，在DIR/src/foo/bar中包含源代码的包可以作为 "foo/bar"导入，并将其编译形式安装到"DIR/pkg/GOOS_GOARCH/foo/bar.a"。
- bin目录保存以编译的命令，每个命令都用其源码目录命名，但不是全部的路径，不包含最后的元素
- bin目录保存已编译的命令。每个命令的源目录都被命名，但只有最后一个元素，而不是整个路径。也就是说DIR/src/foo/quux中的源代码命令安装在DIR/bin/quux中，而不是DIR/bin/foo/quux中。foot/前缀被剥离以便可以将DIR/bin加到PATH环境变量以获取已安装的命令。如果已经设置GOBIN环境变量，命令会安装到它所指定的目录，而不是DIR/bin。GOBIN必须是绝对路径。

下面是一个目录布局的例子：

``` bash
GOPATH=/home/user/go

/home/user/go/
    src/
        foo/
            bar/               (bar包中的go代码)
                x.go
            quux/              (main包中的go代码)
                y.go
    bin/
        quux                   (已安装的命令)
    pkg/
        linux_amd64/
            foo/
                bar.a          (已安装的包对象)
```

Go会搜索GOPATH列出的每个目录以查找源代码，但是新包终会下载到列表中的第一个目录中。

查看`https://golang.org/doc/code.html`获取一个例子。

## 自定义GOPATH

注意：GOPATH不能包含Go的安装目录。

### Bash

编辑`~/.bash_profile`添加下面命令：

``` bash
export GOPATH=$HOME/work
```

保存更改。然后指定源`~/.bash_profile`:

``` bash
source ~/.bash_profile
```

注意：设置GOPIN环境变量在`go install`运行时生成二进制文件：

``` bash
export GOBIN=$HOME/work/bin
```

### Zsh

编辑`~/.zshrc`文件，添加以下命令：

``` bash
export GOPATH=$HOME/work
```

保存更改。然后指定源`~/.zshrc`:

``` bash
source ~/.zshrc
```

### 注意：

设置环境变量时使用绝对路径：

``` bash
/Users/**/work
```

不要使用参数：

``` bash
$HOME/work
```

不然只会匹配第一个路径，第二个路径将无效。