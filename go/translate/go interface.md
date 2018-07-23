# Go中如何使用interface

`http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go`

在我开始使用Go语言前，我大部分的工作都是使用Python。作为一个Python开发者，我发现学习使用Go中的interface相当的困难。尽管interface的基础内容简单，而且我也知道如何在标准库中使用interface，但是在我了解如何设计自己的interface前，还是花费了很多时间、做了很多的练习。在这片文章中，我会讨论Go的类型系统，从而解释如何有效的使用interface。

## interface介绍

所以，什么是interface呢？interface是如下两个事物：interface是一系列方法的集合，它也是一个类型。我们先来从“方法集合”方面了解interface。

通常，我们会使用一些例子来介绍interface。我们先从定义了Animal数据类型的应用开始，因为这是我们经常会遇到的实际问题。Animal类型将会是一个interface，我们会定义一个Animal为任何能够说话的东西。这是Go类型系统的核心概念：我们不是根据类型可以接受什么样的数据来设计抽象类型，而是根据类型可以执行的行为来设计抽象数据。

我们先来定义Animal interface：

``` go
type Animal interface {
    Speak() string
}
```

很简单：我们定一个一个Animal，可以成为任何带有Speak方法的类型。Speak方法没有参数且返回一个字符串。任何实现了Speak方法的类型都可以认为是满足Animal interface的。GO语言中是没有implements关键字的；一个类型是否满足一个interface是自动确定的。我们创建几个满足Animal interface的类型：

``` go
type Dog struct {
}

func (d Dog) Speak() string {
    return "Woof!"
}

type Cat struct {
}

func (c Cat) Speak() string {
    return "Meow!"
}

type Llama struct {
}

func (l Llama) Speak() string {
    return "?????"
}

type JavaProgrammer struct {
}

func (j JavaProgrammer) Speak() string {
    return "Design patterns!"
}
```

现在，我们有了四个不同类型的Aniaml：Dog，Cat，Llama 和 JavaProgrammer。在main()函数中，我可以创建一个Animal的Slice，将上述的每个类型各放一个到Slice中，看看每个Animal会所些什么。我们来实现：

``` go
func main() {
    animals := []Animal{Dog{}, Cat{}, Llama{}, JavaProgrammer{}}
    for _, animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```

好了，现在你知道如何使用Interface了，我不需要在讨论这个问题了，对不对？其实，我们还需要更深入的去了解。让我们来看看对新gopher并不是很明显的事情。

## `interface{}`类型

interface{}类型是一个空的interface，是一切的混乱之源。Interface{}类型是没有函数的interface。因为没有implements关键字，所有的类型都至少实现了0个方法，而且是否实现interface是自动判断的，所以所有的类型都实现了空的interface。这就是说，如果你写了一个接受interfac{}作为参数值的函数，那么这个函数就可以接受任何类型作为参数了。所以，这个函数：

``` go
func DoSomething(v interface{}) {
   // ...
}
```

可以接受任何类型作为参数。

这就是导致困惑的地方了：在DoSomething方法体中，v的类型是什么呢？刚入门的gopher可能会在上面内容的引导下认为“v是任何类型”，但这是错误的。v不是任何类型；v就是interface{}类型的。等会儿，啥？当传值到DoSomething方法中，Go运行时将会执行一个类型转换（如果需要的话），将传入值转换成一个interface{}类型的值。在运行时，所有值只有一种类型，v的静态类型就是interface{}。

这应该会让你好奇：如果正在进行转换，那么传递给接受interface{}作为参数值的函数的是什么呢？（或者说，[]Animal slice中实际上存储的是什么呢？）一个interface值由两个数据字（two words of data）构成；一个数据字用于指向一个值的基础类型的方法表，另一个数据字用于指向该值所持有的实际数据。我不想过分的纠结于此。如果你明白一个interface的值是两个words的宽度，并且它包含了一个指向基础类型的指针，那通常足以避免常见的陷阱。如果你想了解更多的接口实现的信息，我觉得[Russ Cox’s description of interfaces](https://research.swtch.com/interfaces)非常有用。

在我们之前的例子中，当我们创建了一个Animal的slice的时候，我们不用做多余的事情（比如`Animal(Dog{})`）来将一个类型为Dog的值放到Animal slice中，因为转换已经自动的为我们处理了。在Animal slice中，每个元素都是Animal类型的，但是不同的值有着不同的基础类型。

那么，为什么这很重要？了解interface如何在内存中表示就使得一些看似困惑的事情变得清晰明白。比如，一旦了解了interface在内存中如何表示，问题“[can I convert a []T to an []interface{}](https://golang.org/doc/faq#convert_slice_of_interface)”就容易回答了。下面是一个对interface{}类型常见误解的示例：

``` go
package main

import (
    "fmt"
)

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    PrintAll(names)
}
```

运行这段代码，你会得到这样的错误信息： cannot use names (type []string) as type []interface {} in function argument.如果想要这样实现，我们就必须将[]string转换成[]interface{}：

``` go
package main

import (
    "fmt"
)

func PrintAll(vals []interface{}) {
    for _, val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"stanley", "david", "oscar"}
    vals := make([]interface{}, len(names))
    for i, v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}
```

这样线的代码非常难看，但是“c'est la vie”（这就是生活）。没有事情是完美的。（实际上，这种用法并不常见，应为[]interface{}并不像你认为的那样有用）

## 指针和interface

interface另外的一个微妙之处在于，interface定义没有规定实现者是否应该使用指针接收器或者值接收器来实现接口。当获得一个interface值时，你无法判断基础类型是不是指针。在之前的例子中，我们定义所有的方法使用值接收器，而且我们把相应的值放到Animal Slice中。下面我们做一些改变，将Cat的Speak()方法使用指针接收器：

``` go
func (c *Cat) Speak() string {
    return "Meow!"
}
```

你再次运行程序，将会得到以下错误：

prog.go:40: cannot use Cat literal (type Cat) as type Animal in array element:
    Cat does not implement Animal (Speak method requires pointer receiver)

老实说，这个错误信息一开始可能会令人困惑。他所描述的并不是指Animal接口需要你定义你的方法为指针接收器（*这里没翻译通？？*），而是指你已经尝试将Cat结构体转换为Animal接口值，但是只有`*Cat`满足这个接口。你可以通过传入一个`*Cat`指针（而不是Cat值）到Animal slice中来修复这个bug（你也可以传入`&Cat{}`，我只是喜欢`new(Cat)`这样的外观）：

``` go
animals := []Animal{Dog{}, new(Cat), Llama{}, JavaProgrammer{}}
```

现在，我们的例子又可以执行了。

我们现在从相反的方向思考：我们传入一个`*Dog`指针（而不是Dog值），但是这一次我们并没有改变Dog类型方法的定义：

``` go
animals := []Animal{new(Dog), new(Cat), Llama{}, JavaProgrammer{}}
```

这样依然执行通过，但是需要注意一个微妙的差异：我们不需要改变Speak方法的接收器。这样也能够运行是因为指针类型可以访问其关联指类型的方法集，但反过来就不行了。也就是说，`*Dog`值可以使用Dog上定义的Speak方法，但是正如我们之前所看到的，Cat值无法访问`*Cat`上定义的Speak方法。

这听起来很神秘，但是当你记住下面的内容时，上面的内容看起开就合理了：
Go中的所有内容都是按值传递的（everything in Go is passed by value）。每当你调用一个方法的时候，你参入方法的值都是复制过的。对于带有值接收器的方法，在调用方法的是时候复制值。当你明白以下下面的签名方法时，上述内容就更明显了：

``` go
func (t T)MyMethod(s string) {
    // ...
}
```

这个签名方法是类型为`func(T, string)`的函数；方法接收器像方法的其他参数一样通过值传入方法。

对于定义在一个值类型上的方法（比如：`func (d Dog) Speak() {...}`）内部进行的任何更改，对调用者来说都是不可见的，因为调用者在使用一个完全独立的Dog值。由于所有内容都是按值传递的，为什么一个`*Cat`方法不可以用Cat值调用；任何一个Cat值都有可能有任意数量的`*Cat`指针指向它。如果我们尝试使用一个Cat值调用一个`*Cat`方法，永远也不会有可以用来使用的`*Cat`指针。反之，如果我们有一个Dog类型的方法，还有一个`*Dog`指针，在调用这个方法的时候，我们确切的知道使用的是那个Dog值，因为`*Dog`指针指向具体的一个Dog值；Go的运行时会在必要的时候取消起关联Dog值的指针。也就是说，给定一个`*Dog`值d和一个Dog类型的方法Speak，我们可以使用`d.Speak()`；部门不必要使用在其他语言中类似`d->Speak()`这样的形式。

## 真实的案例：从Twitter API中获取适当的时间戳

Twitter API使用以下格式的字符串表示时间戳：

`"Thu May 31 00:00:01 +0000 2012"`

当然，由于时间戳不是JSON规范的一部分，在JSON文档中可以使用多种方式来表示时间戳。简洁起见，我不会把推文的整个JSON都放过来，但是让我们看看如何使用`encoding/json`处理`created_at`字段：

``` go
package main

import (
    "encoding/json"
    "fmt"
    "reflect"
)

// start with a string representation of our JSON data
var input = `
{
    "created_at": "Thu May 31 00:00:01 +0000 2012"
}
`

func main() {
    // our target will be of type map[string]interface{}, which is a
    // pretty generic type that will give us a hashtable whose keys
    // are strings, and whose values are of type interface{}
    var val map[string]interface{}

    if err := json.Unmarshal([]byte(input), &val); err != nil {
        panic(err)
    }

    fmt.Println(val)
    for k, v := range val {
        fmt.Println(k, reflect.TypeOf(v))
    }
}
```

运行程序，可以得到以下输出：

``` bash
map[created_at:Thu May 31 00:00:01 +0000 2012]
created_at string
```

可以看到，我们已经正确访问到key，但是字符串格式的时间戳并不是很有用。如果我们想比较时间戳开查看那个时间更早，或者看看给定时间后或丹铅时间后经过了多长时间，使用普通字符串是做不到的。

我们来天真地尝试一下将字符串解析称为time.Time值（time.Time是表示时间的标准库），然后来看看我们会得到怎样的错误。进行以下修改：

``` go
   var val map[string]time.Time

    if err := json.Unmarshal([]byte(input), &val); err != nil {
        panic(err)
    }
```

运行代码，会得到如下错误：

``` bash
parsing time ""Thu May 31 00:00:01 +0000 2012"" as ""2006-01-02T15:04:05Z07:00"":
    cannot parse "Thu May 31 00:00:01 +0000 2012"" as "2006"
```

这个有些令人困惑的错误消息来自于Go处理time.Time值与字符串之间的转换。简而言之，这意味着我们提供的字符串表示与标准时间格式不匹配（因为Twitter的API使用Ruby写的，Ruby默认的时间格式和Go默认的时间格式是不一样的）。我们需要定义自己的类型来正确解析这个字符串时间戳。encoding/json包会查看传递给json.Unmarshal的值是否满足json.Unmarshaler接口，json.Unmarshaler接口看起来像这样：

``` go
type Unmarshaler interface {
    UnmarshalJSON([]byte) error
}
```

此处引用了文档此处：`http://golang.org/pkg/encoding/json/#Unmarshaler`

所以，我们需要一个带有`UnmarshalJSON([]byte) error`方法的time.Time值：

``` go
type Timestamp time.Time

func (t *Timestamp) UnmarshalJSON(b []byte) error {
    // ...
}
```

通过实现`UnmarshalJSON([]byte) error`方法，我们满足了json.Unmarshaler接口，当遇到Timestamp值的时候，json.Unmarshal会调用我们自定义的解析代码。在本例中，我们使用了一个指针方法，因为我们希望调用者可以看到对接收器的修改。为了设置指针指向的值，我们使用`*`操作符手动取消引用指针。在UnmarshalJSON方法内部，t表示指向Timestamp值的指针。通过`*t`，我们取消了指针t的引用，使得我们可以访问到t指向的值。记住：Go中的所有内容都是按值传递的。这意味着在UnmarshalJSON方法内部，指针t是复制得到的与调用方法的上下文中的指针是不相同的。（在方法内部）如果你直接将t分配给另一个值，你只是重新分配了函数本地指针；调用者不会看到变化。但是，在调用方法内部指针指向的数据与其调用范围作用域内的指针指向的数据是一样的；通过取消指针引用，使得我们的更改对调用者可见。

我们可以使用time.Parse方法，该方法具有签名`func(layout, value string) (Time, error)`。也就是说，这个方法需要接收两个字符串：第一个字符串是一个“布局”字符串，用来描述我们如何格式化时间戳；第二个字符串是我们希望解析的值。这个方法返回一个time.Time值和一个错误值（以防应为某些原因在解析时间戳时失败）。你可以在[time包的文档](https://t.umblr.com/redirect?z=http%3A%2F%2Fgolang.org%2Fpkg%2Ftime%2F&t=MjBjODdjZjg0YzVhYTJjMTk1NjM0OTE0NGFhODY2Y2U1MmU3ZDIwOSxoZmZyTU04Sg%3D%3D&b=t%3AWvVgq8ST2uhqerqL8LopZw&p=http%3A%2F%2Fjordanorelli.com%2Fpost%2F32665860244%2Fhow-to-use-interfaces-in-go&m=1)中获取更多有关layout字符串的语义信息，但是在本例中，我们不需要手动计算layout字符串，应为这个layout字符串已经作为值time.RubyDate存在于标准库中。所以，实际上我们通过调用函数`time.Parse(time.RubyDate, "Thu May 31 00:00:01 +0000 2012")`将字符串“Thu May 31 00:00:01 +0000 2012”处理成 time.Time值。我们获得time.Time类型的值。在我们的例子中，我们关注Timestamp类型的值。我们可以使用Timestamp(v)将time.Time值转换成Timestamp值（v是time.Time值）。最后，我们的UnmarshalJSON函数看起来像这样：

``` go
func (t *Timestamp) UnmarshalJSON(b []byte) error {
    v, err := time.Parse(time.RubyDate, string(b[1:len(b)-1]))
    if err != nil {
        return err
    }
    *t = Timestamp(v)
    return nil
}
```

因为传入的byte slice是JSON元素的原始数据（`\"Thu May 31 00:00:01 +0000 2012\"`），且包含字符串两侧的引号，所以我们获取传入的byte slice的子切片；我们希望在传给time.Parse前将多余的引号剔除。