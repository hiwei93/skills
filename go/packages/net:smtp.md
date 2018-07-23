# net/smtp

## func SendMail

SendMail 链接addr所在的服务器，如果可能的话将会转换成TLS，如果可能先泽可选机制a惊醒身份验证，从地址from发送一封邮件到地址to，发送信息为msg。addr必须包含端口，如：“mail.example.com:smtp”。

参数中的邮件地址都是SMTP RCPT地址。

msg参数必须符合RFC 822邮件形式：头部信息，空白行，后面是消息体。msg的行应该是CRLF终止的（即换行符为`\r\n`）.msg头部信息通常包含："From", "To", "Subject", and "Cc"等字段。发送 Bcc 消息是通过在to参数中包含电子邮件地址但该邮件地址不存在于msg头部信息来完成的。
