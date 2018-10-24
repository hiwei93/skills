# 翻译服务

当前常见翻译网站均限制单次搜索字符为5000个（包括标点符号，空格等），这些翻译网站也提供了自己的翻译服务（见[附录：主流翻译服务](#主流翻译服务)），大家可以按照自己的需要去选择。

本文主要介绍[百度通用翻译API](http://api.fanyi.baidu.com/api/trans/product/apidoc)的使用和Google翻译API非官方包[py-googletrans](https://github.com/ssut/py-googletrans)的使用。

说明：

- 该文档只针对词和独立句子的翻译，不涉及文章等复杂的翻译；
- 该文档中的代码均以python实现（版本：python 3.6）。

## 百度通用翻译API

1. 在百度[翻译开放平台](http://api.fanyi.baidu.com/api/trans/product/index)注册成为开发者；
2. 在管理控制台的[开发者信息](http://api.fanyi.baidu.com/api/trans/product/index)中找到*APP ID*和*密钥*；
3. 可以在文档中找到各种语言的demo，使用时需要注意一下，官方Python语言的demo使用的是Python2.7版本，可以按照下表进行包的转换到python3.x版本：

python 2.7 | python 3.*
----|----
httplib | http.client
md5 | hashlib.md5
urllib.quote | urllib.parse.quote

4. 更符合python3.x版本的Demo

``` python
import hashlib
import random
import urllib
import urllib.parse
import urllib.request

appid = '' #你的appid
secretKey = '' #你的密钥
url = 'http://api.fanyi.baidu.com/api/trans/vip/translate'

query = 'apple'
from_lang = 'en'
to_lang = 'zh'
salt = random.randint(32768, 65536)
sign = appid+query+str(salt)+secretKey
sign = hashlib.md5(sign.encode(encoding='UTF-8')).hexdigest()

values = {
    'q': query,
    'from': from_lang,
    'to': to_lang,
    'appid': appid,
    'salt': salt,
    'sign': sign
}

url_values = urllib.parse.urlencode(values)
full_url = url + '?' + url_values
with urllib.request.urlopen(full_url) as response:
    resp_data = response.read().decode("unicode-escape")
    print(resp_data)
```

注：最好将搜索结果保存到数据库，防止浪费搜索资源。

## Google翻译API（非官方包 py-googletrans）

通过翻译网页爬取翻译内容的难度与复杂度与日俱增，自己写一个爬虫的工作量也越来越大，幸而体有大神已经完成了这项艰巨的工作，[py-googletrans](https://github.com/ssut/py-googletrans)就是其中一个，虽然稳定性难以保证，但是可以满足基本的翻译需求，且速度也可以接受。

可以参看[API文档](https://py-googletrans.readthedocs.io/en/latest/)使用，简单的翻译使用：

``` python
from googletrans import Translator

translator = Translator(service_urls=['translate.google.cn'])
translate = translator.translate("apple", src='en', dest='zh-cn')
print(translate.text)
```

注意：

- 中国地区用户可指定`service_urls`为`translate.google.cn`（谷歌翻译的中国域名），不用翻墙就可以重用了。
- 语言选项需要在[googletrans.LANGUAGES](https://py-googletrans.readthedocs.io/en/latest/#googletrans-languages)列表中，比如简体中文为`zh-cn`（而不是`zh`），如不满足条件会报错`ValueError: invalid destination language`

### 使用时遇到的问题

使用时会出现错误：
`AttributeError: 'NoneType' object has no attribute 'group'`

这个问题已经有人提了`https://github.com/ssut/py-googletrans/issues/88`

stackoverflow上提供了临时解决方案，`https://stackoverflow.com/a/52487148/2231702`需要自己修改源码修复:

打开`py-googletrans`包中的`gtoken.py`文件，对应代码更新以下代码：

``` python
RE_TKK = re.compile(r'TKK=eval\(\'\(\(function\(\)\{(.+?)\}\)\(\)\)\'\);',
                        re.DOTALL)
    RE_RAWTKK = re.compile(r'TKK=\'([^\']*)\';',re.DOTALL)

    def __init__(self, tkk='0', session=None, host='translate.google.com'):
        self.session = session or requests.Session()
        self.tkk = tkk
        self.host = host if 'http' in host else 'https://' + host

    def _update(self):
        """update tkk
        """
        # we don't need to update the base TKK value when it is still valid
        now = math.floor(int(time.time() * 1000) / 3600000.0)
        if self.tkk and int(self.tkk.split('.')[0]) == now:
            return

        r = self.session.get(self.host)

        rawtkk = self.RE_RAWTKK.search(r.text)
        if rawtkk:
            self.tkk = rawtkk.group(1)
            return
```

或者替换成[py_translator](https://pypi.org/project/py-translator/)

## 附录

### 主流翻译服务

服务商 | 费用 | API文档
----|----|------
百度 | 通用翻译API，每月200万字符免费，超出后按照百万字符为单位收费，[详情](http://api.fanyi.baidu.com/api/trans/product/prodinfo#0) | [通用翻译API技术文档](http://api.fanyi.baidu.com/api/trans/product/apidoc)
有道 | 文本翻译，收费，新注册用户会给100元体验资金，[详情](http://ai.youdao.com/docs/doc-trans-price.s#p07) | [有道智云文本翻译API](http://ai.youdao.com/docs/doc-trans-api.s#p01)
Google | 翻译服务，收费，[详情](https://cloud.google.com/translate/pricing) | [API与参考](https://cloud.google.com/translate/docs/apis)
Microsoft | 文本翻译，每月200万字符免费，超出后按照百万字符为单位收费，[详情](https://azure.microsoft.com/zh-cn/pricing/details/cognitive-services/translator-text-api/) | [文本翻译API文档](https://docs.microsoft.com/zh-cn/azure/cognitive-services/translator/)

注：以上数据汇总日期为2018-10-21

来源于：[mr-wolverine](https://github.com/mr-wolverine/skills/blob/master/python/Translate%20Service.md)