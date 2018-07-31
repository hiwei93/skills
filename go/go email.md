# 使用Go作为Email客户端

## 基本知识

1. 发送邮件:

    - 协议：[SMTP](https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol)
    - go标准包：[net/smtp](https://golang.google.cn/pkg/net/smtp/#example_SendMail)，可以查看[Wiki](https://github.com/golang/go/wiki/SendingMail)中使用smtp包的使用方法。

2. 接收邮件：

    - 协议：[IMAP](https://en.wikipedia.org/wiki/Internet_Message_Access_Protocol)
    - 协议：[MIME](https://en.wikipedia.org/wiki/MIME)
    - 第三方包[go-imap](https://github.com/emersion/go-imap)，建议查看[Wiki](https://github.com/emersion/go-imap/wiki)中的例子。
    - 可选包：go标准包[net/mail](https://golang.google.cn/pkg/net/mail/)，[net/textproto](https://golang.google.cn/pkg/net/textproto/)，[mime](https://golang.google.cn/pkg/mime/)，[mime/multipart](https://golang.google.cn/pkg/mime/multipart/)

## 发送邮件

可以查看[golang wiki: SendingMail](https://github.com/golang/go/wiki/SendingMail)查看`net/smtp`包使用方法，但是官方默认的是`STARTTLS`而不是`ssl/tls`，如果使用QQ邮箱的话，就无法发送成功，需要使用`ssl/tls`重新构建通信链接。

``` go
package main

import (
    "crypto/tls"
    "fmt"
    "log"
    "net"
    "net/smtp"
    "strings"
)

// Dial returns a new Client connected to an SMTP server at addr
func Dial(addr string) (*smtp.Client, error) {
    conn, err := tls.Dial("tcp", addr, nil)
    if err != nil {
        return nil, err
    }
    host, _, _ := net.SplitHostPort(addr)
    return smtp.NewClient(conn, host)
}

// SendMailUsingTLS uses TLS sending mail
func SendMailUsingTLS(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
    c, err := Dial(addr)
    if err != nil {
        return err
    }
    defer c.Close()

    if a != nil {
        if ok, _ := c.Extension("AUTH"); ok {
            if err = c.Auth(a); err != nil {
                return err
            }
        }
    }
    if err = c.Mail(from); err != nil {
        return err
    }
    for _, addr := range to {
        if err = c.Rcpt(addr); err != nil {
            return err
        }
    }
    w, err := c.Data()
    if err != nil {
        return err
    }
    _, err = w.Write(msg)
    if err != nil {
        return err
    }
    err = w.Close()
    if err != nil {
        return err
    }
    return c.Quit()
}

func send() {
    hostname := "example.server.qq.com" // 发送服务器
    port := 000 // 发送服务器端口号
    sender := "sender@example.com" // 发件人邮箱
    password := "123456" // 邮箱密码
    recipients := []string{"recipient@example.com"} // 收件人邮箱地址
    message := "From: sender@example.com\r\n" + // 注意message的结构，需要满足
        "To: " + strings.Join(recipients, ";") + "\r\n" +
        "Subject: test mail\r\n" +
        "Content-Type: text/html;chartset=UTF-8\r\n" +
        "\r\n" +
        "测试邮件\r\n"

    print(message)
    auth := smtp.PlainAuth(
        "",
        sender,
        password,
        hostname,
    )

    err = SendMailUsingTLS(
        fmt.Sprintf("%s:%d", hostname, port),
        auth,
        sender,
        recipients,
        []byte(message),
    )

    if err != nil {
        log.Fatal(err)
    }
}
```

- 注：发件信息的格式需要满足`RFC 822`的样式：头部信息，空白行，后面是消息体，行终止符为CRLF（即换行符为`\r\n`）。头部信息通常包含："From", "To", "Subject", and "Cc"等字段。发送 Bcc 消息是通过在to参数中包含电子邮件地址但该邮件地址不存在于msg头部信息来完成的。

## 接收邮件

接收邮件使用的是IMAP协议，目前Go还没有相关的官方包，这里使用第三方库`go-imap`，具体使用可以查看[Wiki](https://github.com/emersion/go-imap/wiki/Fetching-messages)中的例子。

该例中较全的展示了使用go-imap及其相关包获取文件内容或附件的方法，并且较为友好的为邮件解码。

<!--TODO Example and Delete mail-->