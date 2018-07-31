# golang code snippet

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
    print(encodeString)
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
    print(string(decodeBytes))
}
```

- 使用`simplifiedchinese.GB18030.NewDecoder().Reader(io.Reader)`可以对字节流解码。