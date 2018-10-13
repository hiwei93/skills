# golang code snippet

<!-- TOC -->

- [golang code snippet](#golang-code-snippet)
    - [时间类型格式化](#时间类型格式化)
    - [文件读取](#文件读取)
        - [去除文件BOM头](#去除文件bom头)
    - [中文编码转换](#中文编码转换)
        - [中文字符编码](#中文字符编码)
        - [中文字符解码](#中文字符解码)
    - [异常处理](#异常处理)
        - [捕捉异常，打印异常栈](#捕捉异常打印异常栈)
    - [反射（TODO）](#反射todo)
    - [并发（TODO）](#并发todo)
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

> TODO: 说明为什么是 "2006-01-02"

## 文件读取

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

## 中文编码转换

需要导入官方包：`https://godoc.org/golang.org/x/text/encoding`

``` bash
go get -u golang.org/x/text/encoding
```

以下例子均以中文编码转换为例，其他语言的编码方法可以在[官方文档](https://godoc.org/golang.org/x/text/encoding#pkg-subdirectories)找到。

### 中文字符编码

``` go
package main

import (
    "encoding/base64"
    "golang.org/x/text/encoding/simplifiedchinese"
)

func encodeMessage(){
    str := `文字信息`
    encodeString, err := simplifiedchinese.GB18030.NewEncoder().String(str)
    if err != nil {
        panic(err)
    }
    encodeString = base64.StdEncoding.EncodeToString([]byte(encodeString))
    println(encodeString)
}
```

- 使用`simplifiedchinese.GB18030.NewEncoder().Writer(io.Writer)`可以对字节流编码。

### 中文字符解码

``` go
package main

import (
    "encoding/base64"
    "golang.org/x/text/encoding/simplifiedchinese"
)

func decodeMessage(){
    str := `zsTX1tDFz6I=`
    decodeBytes, err := base64.StdEncoding.DecodeString(str)
    if err != nil {
        panic(err)
    }
    decodeBytes, err = simplifiedchinese.GB18030.NewDecoder().Bytes(decodeBytes)
    if err != nil {
        log.Fatalln(err)
    }
    println(string(decodeBytes))
}
```

- 使用`simplifiedchinese.GB18030.NewDecoder().Reader(io.Reader)`可以对字节流解码。

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

## 并发（TODO）

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