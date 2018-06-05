# Golang http

## 1. HTTP client implementations

### 1. Http(或 https)请求

``` go
resp, err := http.Get("http://example.com/")
...
resp, err := http.Post("http://example.com/upload", "image/jpeg", &buf)
...
resp, err := http.PostForm("http://example.com/form", url.Values{"key": {"Value"}, "id": {"123"}})
```

### 2. 当响应结束后client必须关闭response body

``` go
resp, err := http.Get("http://example.com/")
if err != nil {
   // handle error
}
defer resp.Body.Close()
body, err := ioutil.ReadAll(resp.Body)
// ...
```

### 3. 创建`Client`来控制HTTP头部信息，重定向策略和其他的设置

``` go
client := &http.Client{
    CheckRedirect: redirectPolicyFunc,
}

resp, err := client.Get("http://example.com")
// ...

req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
```

### 4. 创建`Transport`来控制代理，TLS配置，`keep-alives`，压缩和其他设置

``` go
tr := &http.Transport{
    MaxIdleConns:       10,
    IdleConnTimeout:    30 * time.Second,
    DisableCompression: true,
}
client := &http.Client{Transport: tr}
resp, err := client.Get("https://example.com")
```

### 5. 说明

1. `Client`和`Transport`是线程安全的，可以安全地由多个goroutines并发使用。
2. 为了高效`Client`和`Transport`只需创建一次，并且重复使用。