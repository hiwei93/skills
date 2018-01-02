# How to Write Go Code

## 介绍

该文档展示了开发一个GO报并介绍了GO工具，这是获取、构建和安装GO包和命令的标准方法。

Go工具需要一特定的方式组织代码。请仔细阅读本文，该文档解释了使用Go安装进行启动和运行的最简单方法。

可以通过[视频](https://www.youtube.com/watch?v=XCsL89YtqCs)了解类似说明。

## 组织代码

### 概述

- GO程序猿通常将所有GO代码保存在一个workspace中；
- 一个workspace会有多个版本管理库；
- 每一个版本管理库包含一个或多个包；
- 每个包有由一个或多个Go源文件组成；
- 包的目录路径决定了它的导入路径。

注意：与其他的编程环境不同，其他的编程环境中，每个项目都有独立的workspace并且workspace与版本管理库紧密相关。

### 工作区（Workspaces）

工作区是一个目录层次结构，其根目录下有三个目录：

- `src`由Go源文件组成；
- `pkg`由包对象组成；
- `bin`由可执行命令组成。

Go工具构建源代码包，并将生成的二进制文件安装到pkg和bin目录中。

`src`子目录同通常包含多个版本管理库，用于跟踪一个或多个源代码包的开发。

以下展示了工作区在实践中的结构：

``` bash
bin/
    hello                          # 可执行命令
    outyet                         # 可执行命令
pkg/
    linux_amd64/
        github.com/golang/example/
            stringutil.a           # 包对象
src/
    github.com/golang/example/
        .git/                      # git库元数据
    hello/
        hello.go               # 命令源码
    outyet/
        main.go                # 命令源码
        main_test.go           # 测试源码
    stringutil/
        reverse.go             # 包源码
        reverse_test.go        # 测试源码
    golang.org/x/image/
        .git/                      # git库元数据
    bmp/
        reader.go              # 包源码
        writer.go              # 包源码
    ... (省略其他的远程库和包) ...
```

上面的树状结构包含两个远程库的工作区（example and image）。example远程库包含两个命令（hello and outyet）和一个库（stringutil）。image远程库包含bmp包和[其他组成部分](https://godoc.org/golang.org/x/image)。

典型的工作区包由多个包含多个包和命令的源代码库组成。大多数Go程序员将所有Go源码和依赖保存到一个工作区中。

命令和库是有不同类型的源码包构建而成的，稍后将会讨论两者的区别。

### GOPATH环境变量

GOPATH环境变量指定工作间的位置。改环境变量默认为主目录中名为go的目录，在Unix上为`$HOME/go`，在Plan 9上为`$home/go`，在Windows上为`％USERPROFILE％\go`通常为`C:\Users\YourName\go`）。

如果希望在其他的位置工作，需要将GOPATH设置为目标目录的路径。（另一个命令设置是设置`GOPATH=$HOME`。）注意：GOPATH必须于Go的安装路径（即GOROUT的环境变量不同）不同。

`go env GOPATH`命令会打印出当前有效的GOPATH；如果未设置相应的环境变量，则会显示默认的路径。（`go env`命令会输出所有Go相关的环境参数。）

方便起见，将工作区的bin子目录添加到PATH中：

``` bash
export PATH=$PATH:$(go env GOPATH)/bin
```

简洁起见，余下文档的脚本中，将会使用`$GOPATH`代替`$(go env GOPATH)`。如果还没有设置GOPATH，为了使脚本以书面形式运行，你可以使用这些命令替换`$HOME/go`或者运行：

``` bash
export GOPATH=$(go env GOPATH)
```

获取更多有关GOPATH环境变量的信息，参看[go help gopath](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable)

使用自定义的工作间路径，[请设置GOPATH环境变量](https://golang.org/wiki/SettingGOPATH)。

### 导入路径

导入路径是唯一表示包的字符串。程序包的导入路径对于其在工作区内或远程库中的位置（解释如下）。

标准库中的软件包拥有短的倒入路径，如 "fmt" 和 "net/http"。对于用户自己的软件包，必须选择一个基本路径，这个路径不会在将来添加到标准库或其他外部库时发生冲突。

如果将代码保存到某个源代码库中，则应该使用该源代码库的根节点作为基本路径。例如，如果你在github.com/user有一个GitHub账号，github.com/user应该就是你的基本路径。

注意：在构建代码前，不需要将代码发不到远程库。培养组织代码的好习惯，假设你某天会发布它。实践中，只要路径名称是标准库和更大GO生态系统中所独有的，你可以选择任意路径名称。

我们会使用 github.com/user 作为基本路径。在你的工作区内创建一个目录来保存源码：

``` bash
mkdir -p $GOPATH/src/github.com/user
```

### 第一个Go程序

要编译和运行一个简单的程序，首先选择一个包路径（这里使用github.com/user/hello），在工作区内创建一个对应的包目录：

``` bash
mkdir $GOPATH/src/github.com/user/hello
```

接下来，在hello目录中创建一个名为hello.go的目录，该文件中包含以下内容：

``` go
package main

import "fmt"

func main() {
    fmt.Printf("Hello, world.\n")
}
```

现在可以，可以使用go工具来构建（build）和安装（install）该程序。

``` bash
go install github.com/user/hello
```

注意：可以从系统的任何位置运行该命令。Go工具通过在GOPATH指定指定的工作区内查找github.com/user/hello包来查找源代码。

如果从包目录运行go install，则可以省略软件包路径：

``` bash
cd $GOPATH/src/github.com/user/hello
go install
```

go install命令会构建hello命令，产生一个可执行的二进制文件。然后，会将产生的二进制文件安装到工作区的bin目录（在windows下生成hello.exe）。在我们的例子中，产生的文件为$GOPATH/bin/hello，也就是$HOME/go/bin/hello。

只有有错误发生时，Go 工具才会打印输出，所以如果运行go install命令没有产生输出，说明命令已经成功执行了。

现在，可以在命令行输入二进制文件的完整路径来执行该程序了：

``` bash
$GOPATH/bin/hello
Hello, world.
```

如果在PATH环境变量添加了$GOPATH/bin，可以直接输入二进制文件名称：

``` bash
hello
Hello, world.
```

如果正在使用源码管理系统，现在是初始化库的最好时机了，添加文件并提交第一次更改。再次强调，这一步是可选的：弄不需要使用远吗控制系统编写Go代码。

``` bash
$ cd $GOPATH/src/github.com/user/hello
$ git init
Initialized empty Git repository in /home/user/work/src/github.com/user/hello/.git/
$ git add hello.go
$ git commit -m "initial commit"
[master (root-commit) 0b4507d] initial commit
 1 file changed, 1 insertion(+)
  create mode 100644 hello.go
```

将代码推送到远程库作为读者的联系。

### 第一个Go库

接下来写一个库，并在hello程序中使用。

再次强调，第一步是选择一个包路径（这里使用github.com/user/stringutil）并创建包目录：

``` bash
mkdir $GOPATH/src/github.com/user/stringutil
```

接下来，在刚穿件的目录中创建一个名为reverse.go的文件，该文件中包含以下内容：

``` go
// 包stringutil包含用于处理字符串的使用函数。
package stringutil

// 将传入的字符串参数翻转
func Reverse(s string) string {
    r := []rune(s)
    for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
```

现在，测试一下编译包的内容：

``` bash
go build github.com/user/stringutil
```

或者，在存放包源代码的目录中，只需执行：

``` bash
go build
```

执行该命令不会生成任何文件。因此，必须使用go install，这个命令将包对象放到工作区的pkg目录中。

确认stringutil包构建之后，修改hello.go文件（在$GOPATH/src/github.com/user/hello目录中）使用新生成的包：

``` go
package main

import (
    "fmt"

    "github.com/user/stringutil"
)

func main() {
    fmt.Printf(stringutil.Reverse("!oG ,olleH"))
}
```

每当Go工具安装一个包或二进制文件，它也会安装它所有的依赖关系。因此在按轧辊hello程序时：

``` bash
go install github.com/user/hello
```

stringutil包也会自动安装。

运行新的程序，会看到新的、反向的信息：

``` bash
$ hello
Hello, Go!
```

经过上面的步骤，工作去应该看起来下面这样：

``` bash
bin/
    hello                 # 可执行命令
pkg/
    linux_amd64/          # 这个目录取决操作系统及其使用架构
        github.com/user/
            stringutil.a  # 包对象
src/
    github.com/user/
        hello/
            hello.go      # 命令源码
        stringutil/
            reverse.go    # 包源码
```

注意：go install将stringutil.a对象放到了pkg/linux_amd64目录下的目录中，该目录映射包源码所在目录。这是为了使将来的Go工具的调用可以找到包对象，并避免不必要的重复编译包。`linux_amd64`部分是为了帮助交叉编译，并体现操作系统的类型和体系结构。

Go可执行命令是静态的链接；在运行Go程序时，包对象无需存在。

### 包名

Go源码文件第一行语句必须是：

``` go
package name
```

其中name是导入包的默认名称。（包中的文件必须使用相同的name）

Go的惯例是，包名是倒入路径的最后一个元素：导入为 "crypto/rot13" 的包应该命名为rot13。

可执命令必须使用package main。

链接到单个的二进制文件的所有包，没有要求包名必须是唯一的，只有倒入路径（包源文件的完整路径）要求是唯一的。

参阅[Effective Go](https://golang.org/doc/effective_go.html#names)获取更多的命名约定。

### 测试

Go拥有轻量级的测试框架，框架由go test命令和testing包组成。

可以通过创建名称以 _test.go 结尾的文件来编写测试，该文件包含名为TestXXX方法，该方法带有 `func (t *testing.T)` 签名（signature）。测试框架会运行每个方法；如果方法调用失败函数，比如`t.Error`或`t.Fail`，则认为测试失败。

通过创建包含以下代码的文件：$GOPATH/src/github.com/user/stringutil/reverse_test.go，添加一个测试到stringutil包中：

``` go
package stringutil

import "testing"

func TestReverse(t *testing.T) {
    cases := []struct {
        in, want string
    }{
        {"Hello, world", "dlrow ,olleH"},
        {"Hello, 世界", "界世 ,olleH"},
        {"", ""},
    }
    for _, c := range cases {
        got := Reverse(c.in)
        if got != c.want {
            t.Errorf("Reverse(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}
```

接着使用go test命令运行测试：

``` bash
$ go test github.com/user/stringutil
ok  github.com/user/stringutil 0.165s
```

一如往常，如果从包目录下运行Go工具，就可以忽略包路径：

``` bash
$ go test
ok  github.com/user/stringutil 0.165s
```

运行[go help test](https://golang.org/cmd/go/#hdr-Test_packages)并查看[testing package documentation](https://golang.org/pkg/testing/)获取更多信息。

### 远程包

导入路径可以描述如何使用像Git或Mercurial这样的版本管理工具获取包源码。Go 工具使用这个特性从远程库自动获取包。例如，这个文档所描述的示例也保存在GitHub托管的Git库中`github.com/golang/example`。如果在包的导入路径中包含存储哭的URL，go get命令会自动获取，构建和安装它。

``` bash
$ go get github.com/golang/example/hello
$ $GOPATH/bin/hello
Hello, Go examples!
```

如果指定的软件包不再工作区中，go get会把这个软件包放到GOPATH指定的第一个工作区内。（如果包已经存在，go get会跳过远程获取，直接执行，相当于go install）

发出go get命令后，工作区目录层次结构应该如下所示：

``` bash
bin/
    hello                           # 可执行命令
pkg/
    linux_amd64/
        github.com/golang/example/
            stringutil.a            # 包对象
        github.com/user/
            stringutil.a            # 包对象
src/
    github.com/golang/example/
    .git/                       # Git仓库元数据
        hello/
            hello.go                # 命令源码
        stringutil/
            reverse.go              # 包源码
            reverse_test.go         # 测试源码
    github.com/user/
        hello/
            hello.go                # 命令源码
        stringutil/
            reverse.go              # 包源码
            reverse_test.go         # 测试源码

```

托管在GitHub伤的hello命令依赖于同一仓库中的stringutil包。hello.go文件中的导入使用相同的导入路径约定，所以go get命令也能够定位和安装依赖包。

``` bash
import "github.com/golang/example/stringutil"
```

这个约定是使你的Go包可供其他人使用的最简单的方式。`Go Wiki`和`godoc.org`提供外部Go项目列表。

查看[go help importpath](https://golang.org/cmd/go/#hdr-Remote_import_paths)获取更多有关通过GO工具使用远程库的信息。

## 下一步

订阅[golang-announce](https://groups.google.com/group/golang-announce)邮件列表，在Go发布新的稳定版本时收到通知

查看[Effective Go](https://golang.org/doc/effective_go.html)获取有关编写清晰，本土化的GO代码的提示。

参看[A Tour of Go]学习Go 语言本身。

访问[documentation page](https://golang.org/doc/#articles)，获取关于Go 语言及其库和工具的深入文章。

## 获取帮助

要获取及时帮助，请询问 [Freenode](http://freenode.net/) IRC 上 #go-nuts 中的 Gopher 们。

Go 语言的官方邮件列表[Go Nuts](https://groups.google.com/group/golang-nuts)。

使用[Go issue tracker](https://golang.org/issue)报告错误。