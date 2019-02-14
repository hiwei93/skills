# golang code snippet

<!-- TOC -->

- [golang code snippet](#golang-code-snippet)
    - [时间类型格式化](#时间类型格式化)
    - [计算程序运行时间](#计算程序运行时间)
    - [IO操作](#io操作)
        - [获取指定格式的文件名](#获取指定格式的文件名)
        - [去除文件BOM头](#去除文件bom头)
        - [读取csv文件](#读取csv文件)
    - [字符编解码](#字符编解码)
        - [中文字符编解码](#中文字符编解码)
        - [base64方式编解码](#base64方式编解码)
    - [异常处理](#异常处理)
        - [捕捉异常，打印异常栈](#捕捉异常打印异常栈)
    - [反射（TODO）](#反射todo)
    - [并发](#并发)
        - [限制并发数量](#限制并发数量)
        - [等待所有线程运行完毕](#等待所有线程运行完毕)
        - [注意：循环变量快照问题（TODO）](#注意循环变量快照问题todo)
        - [注意：goroutine泄露（TODO）](#注意goroutine泄露todo)
    - [网络](#网络)
        - [根据url获取返回资源](#根据url获取返回资源)

<!-- /TOC -->

## 时间类型格式化

``` go
package main

import (
    "time"
)

func main() {
    now := time.Now()
    println(now.String())
    t := now.Format("2006-01-02") // Format("2006-01-02 15:04:05")
    println(t)
}
```

注意：golang使用特定时间`Mon Jan 2 15:04:05 MST 2006`作为格式化的时间
<!-- 使用这个时间的原因 https://stackoverflow.com/questions/20530327/origin-of-mon-jan-2-150405-mst-2006-in-golang-->

## 计算程序运行时间

``` go
package main

import (
    "fmt"
    "time"
)

func expensiveCall() {} // 目标程序

func main() {
    t0 := time.Now()
    expensiveCall()
    t1 := time.Now()
    fmt.Printf("The call took %v to run.\n", t1.Sub(t0))
}
```

参看：`https://golang.google.cn/pkg/time/#example_Duration`

## IO操作

### 获取指定格式的文件名

``` go
package main

import (
    "fmt"
    "path/filepath"
)

func matchPath() {
    fileNames, err := filepath.Glob("*.*")
    if err != nil {
        panic(err)
    }
    fmt.Printf("%q\n", fileNames)

    pattern := "*.log"
    for _, file := range fileNames {
        matched, err := filepath.Match(pattern, file)
        if err != nil {
            panic(err)
        }
        fmt.Printf("file %s match pattern %s: %t\n", file, pattern, matched)
    }
}
```

- `filepath.Glob()`：返回所有匹配模式匹配字符串*pattern*的文件或者*nil*（如果没有匹配的文件）；
- `filepath.Match()`：判断指定文件绝对路径是否与指定*pattern*匹配；
- 具体的*pattern*使用参看[官方文档](https://golang.google.cn/src/path/filepath/match.go?s=1226:1284#L34)，或者查看博客[Golang学习 - path/filepath 包](http://www.mamicode.com/info-detail-1546088.html)

### 去除文件BOM头

``` go
package main

import (
    "bytes"
    "io"
)
func removeBOM(r io.Reader) io.Reader {
    var buf bytes.Buffer
    if _, err := io.Copy(&buf, r); err != nil {
        return nil
    }
    noBOM := ignoreUTFBOM(buf.Bytes())
    return bytes.NewBuffer(noBOM)
}

func ignoreUTFBOM(data []byte) []byte {
    if data == nil {
        return nil
    }

    if len(data) >= 3 && data[0] == 0xef && data[1] == 0xbb && data[2] == 0xbf {
        return data[3:]
    }
    return data
}
```

更为简单的方式：

``` go
func trapBOM(fileBytes []byte) []byte {
    trimmedBytes := bytes.Trim(fileBytes, "\xef\xbb\xbf")
    return trimmedBytes
}
```

- 参看[Reading files with a BOM in Go](https://stackoverflow.com/questions/21371673/reading-files-with-a-bom-in-go)

以上两种方法的缺点相同，都需要将文件全部读到内存中，无法处理很大的文件。

### 读取csv文件

``` go
package main

import (
    "encoding/csv"
    "fmt"
    "log"
)

func readCsv(csvPath string) {
    file, err := os.OpenFile(csvPath, os.O_RDONLY, 0600)
    if err != nil {
        log.Panic(err)
    }
    csvReader := csv.NewReader(file)
    line := 0
    for {
        line++
        row, err := csvReader.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Panicf("execute file line [%d], get error %s", line, err)
        }
        // if csv file has header, then skip it
        if line == 1 {
            continue
        }
        log.Printf("%d row is %v", line, row)
    }
    err = file.Close()
    if err != nil {
        log.Panic(err)
    }
}
```

读取特殊分隔符的文件，比如使用`\t`作为分隔符，需要添加配置：

``` go
csvReader := csv.NewReader(file)
csvReader.Comma = '\t'
```

注意：

当`csv.Reader`的参数`TrimLeadingSpace`设为`true`时，连续的`\t`分隔符会被忽略，如下例：

``` go
in := "first_name\tlast_name\tusername\thappy\nRob\tPike\trob\ttrue\nKen\t\t\tfalse\nRobert\tGriesemer\tgri\ttrue"
r := csv.NewReader(strings.NewReader(in))
r.Comma = '\t'
r. TrimLeadingSpace = true

for {
    record, err := r.Read()
    if err == io.EOF {
        break
    }
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%#v\n", record)
}
```

会产生异常：`record on line 3: wrong number of fields`

但是`,;，`等分隔符不会产生这样的问题。

## 字符编解码

### 中文字符编解码

需要导入官方包：`https://godoc.org/golang.org/x/text/encoding`

``` bash
go get -u golang.org/x/text/encoding
```

以下例子均以中文编码（GB18030编码方式）转换为例，其他语言的编码方法可以在[官方文档](https://godoc.org/golang.org/x/text/encoding#pkg-subdirectories)找到。

``` go
package main

import "golang.org/x/text/encoding/simplifiedchinese"

func main(){
    str := `字符信息`
    // 对中文编码
    encodeBytes, err := simplifiedchinese.GB18030.NewEncoder().Bytes([]byte(str))
    if err != nil {
        panic(err)
    }
    // 解码为中文字符
    decodeBytes, err := simplifiedchinese.GB18030.NewDecoder().Bytes(encodeBytes)
    if err != nil {
        panic(err)
    }
    println(string(decodeBytes))
}
```

- 使用`simplifiedchinese.GB18030.NewEncoder().Writer(io.Writer)`可以对字节流编码；
- 使用`simplifiedchinese.GB18030.NewDecoder().Reader(io.Reader)`可以对字节流解码。

### base64方式编解码

``` go
package main

import "encoding/base64"

func main(){
    str := `字符信息`
    // base64 encode
    encodeString = base64.StdEncoding.EncodeToString([]byte(str))
    println(encodeString)
    // base64 decode
    decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
    if err != nil {
        panic(err)
    }
    println(string(decodeBytes))
}
```

## 异常处理

### 捕捉异常，打印异常栈

``` go
package main

import (
    "fmt"
    "log"
    "runtime/debug"
)

func task() {
    defer func() {
        if p := recover(); p != nil {
            stack := string(debug.Stack())
            err := fmt.Sprintf("panic: %v\n%s", p, stack)
            log.Print(err)
        }
    }()

    // do something
}
```

## 反射（TODO）

## 并发

1. 主函数返回时，所有的goroutine都会被直接打断，程序退出。

### 限制并发数量

``` go
package main

func main() {
    var timers = make(chan struct{}, 100)

    for {
        timers <- struct{}{}
        go func() {
            // do something
            <- timers
        }()
    }
}
```

注：异步方法中`timers`资源的释放应该放在`defer`中，否则可能会因为程序错误导致`timers`资源没有释放而造成的阻塞：

``` go
defer func() {
    <- timers
}()
```

### 等待所有线程运行完毕

``` go
package main

import "sync"

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 10; i ++ {
        wg.Add(1)
        // worker
        go func() {
            defer wg.Done()
            // do something
        }()
    }

    // closer
    go func() {
        wg.Wait()
    }()
}
```

### 注意：循环变量快照问题（TODO）

### 注意：goroutine泄露（TODO）

## 网络

### 根据url获取返回资源

``` go
package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
    "strings"
    "time"
)

func main() {
    var resp *http.Response
    var req *http.Request
    // TODO what is Transport ??
    tr := &http.Transport{
        MaxIdleConns:       10,
        IdleConnTimeout:    30 * time.Second,
        DisableCompression: true,
    }
    client := &http.Client{Transport: tr}

    url := "访问网址"
    reqBody := strings.NewReader("请求内容") // reqBody需要根据具体需要修改
    req, err := http.NewRequest("POST", url, reqBody) // POST 要大写
    // req, err := http.NewRequest("GET", "https://www.baidu.com", nil)
    if err != nil {
        return
    }

    // 添加头部信息
    req.Header.Add("Content-Type", "application/json")

    resp, err = client.Do(req)
    if err != nil {
        return
    }
    defer resp.Body.Close()
    respBody, err := ioutil.ReadAll(resp.Body)
    fmt.Printf("get data %s", respBody)
}
```

注：请求的方法都需要大写，如`GET`，`POST`等，不然会404错误。