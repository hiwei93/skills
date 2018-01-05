# A Tour of Go

<!-- TOC -->

- [A Tour of Go](#a-tour-of-go)
    - [包](#%E5%8C%85)
        - [导入包](#%E5%AF%BC%E5%85%A5%E5%8C%85)
        - [导出名](#%E5%AF%BC%E5%87%BA%E5%90%8D)
    - [函数](#%E5%87%BD%E6%95%B0)
    - [变量](#%E5%8F%98%E9%87%8F)
        - [初始化变量](#%E5%88%9D%E5%A7%8B%E5%8C%96%E5%8F%98%E9%87%8F)
        - [短声明变量](#%E7%9F%AD%E5%A3%B0%E6%98%8E%E5%8F%98%E9%87%8F)
        - [基本类型](#%E5%9F%BA%E6%9C%AC%E7%B1%BB%E5%9E%8B)
        - [零值](#%E9%9B%B6%E5%80%BC)
        - [类型转换](#%E7%B1%BB%E5%9E%8B%E8%BD%AC%E6%8D%A2)
        - [类型推导](#%E7%B1%BB%E5%9E%8B%E6%8E%A8%E5%AF%BC)
        - [常量](#%E5%B8%B8%E9%87%8F)
        - [数值常量](#%E6%95%B0%E5%80%BC%E5%B8%B8%E9%87%8F)
    - [流程控制](#%E6%B5%81%E7%A8%8B%E6%8E%A7%E5%88%B6)
        - [for循环](#for%E5%BE%AA%E7%8E%AF)
            - [以for代替while](#%E4%BB%A5for%E4%BB%A3%E6%9B%BFwhile)
            - [死循环](#%E6%AD%BB%E5%BE%AA%E7%8E%AF)
        - [if语句](#if%E8%AF%AD%E5%8F%A5)
            - [if的便捷语句](#if%E7%9A%84%E4%BE%BF%E6%8D%B7%E8%AF%AD%E5%8F%A5)
            - [if else语句](#if-else%E8%AF%AD%E5%8F%A5)
        - [switch语句](#switch%E8%AF%AD%E5%8F%A5)
            - [没有条件的 switch](#%E6%B2%A1%E6%9C%89%E6%9D%A1%E4%BB%B6%E7%9A%84-switch)
        - [defer](#defer)
            - [defer 栈](#defer-%E6%A0%88)
    - [复杂类型](#%E5%A4%8D%E6%9D%82%E7%B1%BB%E5%9E%8B)
        - [指针](#%E6%8C%87%E9%92%88)
        - [结构体](#%E7%BB%93%E6%9E%84%E4%BD%93)
            - [结构体指针](#%E7%BB%93%E6%9E%84%E4%BD%93%E6%8C%87%E9%92%88)
            - [结构体文法](#%E7%BB%93%E6%9E%84%E4%BD%93%E6%96%87%E6%B3%95)
        - [数组](#%E6%95%B0%E7%BB%84)
            - [slice](#slice)
            - [slice 的 slice](#slice-%E7%9A%84-slice)
            - [对 slice 切片](#%E5%AF%B9-slice-%E5%88%87%E7%89%87)
            - [构造 slice](#%E6%9E%84%E9%80%A0-slice)
            - [nil slice](#nil-slice)
            - [向 slice 添加元素](#%E5%90%91-slice-%E6%B7%BB%E5%8A%A0%E5%85%83%E7%B4%A0)
            - [range](#range)
        - [map](#map)
            - [map 的文法](#map-%E7%9A%84%E6%96%87%E6%B3%95)
            - [修改 map](#%E4%BF%AE%E6%94%B9-map)
        - [函数值](#%E5%87%BD%E6%95%B0%E5%80%BC)
            - [函数的闭包](#%E5%87%BD%E6%95%B0%E7%9A%84%E9%97%AD%E5%8C%85)
    - [方法](#%E6%96%B9%E6%B3%95)
        - [方法是函数](#%E6%96%B9%E6%B3%95%E6%98%AF%E5%87%BD%E6%95%B0)
        - [方法（续）](#%E6%96%B9%E6%B3%95%EF%BC%88%E7%BB%AD%EF%BC%89)
        - [receivers为指针](#receivers%E4%B8%BA%E6%8C%87%E9%92%88)
        - [指针和方法](#%E6%8C%87%E9%92%88%E5%92%8C%E6%96%B9%E6%B3%95)
        - [指针和方法，另一种情况](#%E6%8C%87%E9%92%88%E5%92%8C%E6%96%B9%E6%B3%95%EF%BC%8C%E5%8F%A6%E4%B8%80%E7%A7%8D%E6%83%85%E5%86%B5)
        - [选择值receiver还是指针receiver](#%E9%80%89%E6%8B%A9%E5%80%BCreceiver%E8%BF%98%E6%98%AF%E6%8C%87%E9%92%88receiver)
    - [接口](#%E6%8E%A5%E5%8F%A3)
        - [接口是隐式实现的](#%E6%8E%A5%E5%8F%A3%E6%98%AF%E9%9A%90%E5%BC%8F%E5%AE%9E%E7%8E%B0%E7%9A%84)
        - [接口值](#%E6%8E%A5%E5%8F%A3%E5%80%BC)
        - [带有nil底层值的接口值](#%E5%B8%A6%E6%9C%89nil%E5%BA%95%E5%B1%82%E5%80%BC%E7%9A%84%E6%8E%A5%E5%8F%A3%E5%80%BC)
        - [Nil接口值](#nil%E6%8E%A5%E5%8F%A3%E5%80%BC)
        - [空接口](#%E7%A9%BA%E6%8E%A5%E5%8F%A3)
        - [类型断言](#%E7%B1%BB%E5%9E%8B%E6%96%AD%E8%A8%80)
        - [类型switch](#%E7%B1%BB%E5%9E%8Bswitch)
    - [Stringers](#stringers)
    - [Errors](#errors)
    - [Readers](#readers)
    - [Web服务器](#web%E6%9C%8D%E5%8A%A1%E5%99%A8)
    - [图片](#%E5%9B%BE%E7%89%87)
    - [并发](#%E5%B9%B6%E5%8F%91)
        - [Goroutines](#goroutines)
        - [Channels](#channels)
            - [缓冲 channel](#%E7%BC%93%E5%86%B2-channel)
            - [range 和 close](#range-%E5%92%8C-close)
        - [select](#select)
            - [默认选择](#%E9%BB%98%E8%AE%A4%E9%80%89%E6%8B%A9)
    - [sync.Mutex](#syncmutex)

<!-- /TOC -->

## 包

- 每个 Go 程序都是由包组成的。
- 程序运行的入口是包 main 。
- 按照惯例，包名与导入路径的最后一个目录一致。例如，"math/rand" 包由 package rand 语句开始。

### 导入包

导入语句：

``` go
import "packageName"
```

打包导入：

``` go
import (
    "package1"
    "package2"
    ... ...
)
```

### 导出名

名字首字母大写的成员是被导出的。

在导入包之后，只能访问包所导出成员，任何未导出成员的是不能被包外的代码访问的。

## 函数

1. 函数可以没有参数或接受多个参数，参数类型在变量名之后：

   ``` go
   func add(x int, y int) int {
       return x + y
   }
   ```

1. 当两个或多个连续的函数参数是同一类型，则除了最后一个类型之外，其他都可以省略：

   ``` go
   func add(x, y int) int {
       return x + y
   }
   ```

1. 函数可以返回任意数量的返回值：

   ``` go
   func swap(x, y string) (string, string) {
        return y, x
   }
   ```

1. Go 的返回值可以被命名，并且就像在函数体开头声明的变量那样使用：

   ``` go
   func split(sum int) (x, y int) {
        x = sum * 4 / 9
        y = sum - x
        return
    }
   ```

   - 没有参数的 return 语句返回各个返回变量的当前值。这种用法被称作“裸”返回。
   - 直接返回语句仅应当用在上面那样的短函数中。在长的函数中它们会影响代码的可读性。

## 变量

var 语句定义了一个变量的列表，且类型在变量名后面：

   ``` go
   var c, python, java bool
   var i int
   ```
var 语句可以定义在包或函数级别。

### 初始化变量

可以一次为多个变量初始化：

``` go
var i, j int = 1, 2
```

如果初始化是使用表达式，则可以省略类型，变量会从初始值中获得类型：

``` go
var c, python, java = true, false, "no!"
```

### 短声明变量

**在函数中**， `:=`简洁赋值语句在明确类型的地方，可以用于替代 var 定义：

``` go
i, j := 1, 2
k := 3
```

> ⚠️注意：`:=`简洁赋值语句只能在函数中使用，函数外的每个语句都必须以关键字开始（ var 、 func 、等等）。

### 基本类型

``` go
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8 的别名

rune // int32 的别名
     // 代表一个Unicode码

float32 float64

complex64 complex128
```

- int，uint 和 uintptr 类型在32位的系统上一般是32位，而在64位系统上是64位。
- 当需要使用一个整数类型时，**应该首选 int**，仅当有特别的理由才使用定长整数类型或者无符号整数类型。

### 零值

变量在定义时没有明确的初始化时会赋值为 零值 ：

- 数值类型为 0 ，
- 布尔类型为 false ，
- 字符串为 "" （空字符串）

### 类型转换

表达式 T(v) 将值 v 转换为类型 T：

``` go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
```

或者，更加简单的形式：

``` go
i := 42
f := float64(i)
u := uint(f)
```

> ⚠️注意：Go 的在不同类型的变量间转换时需要**显式转换**。

### 类型推导

在定义一个变量却并不显式指定其类型时（使用 := 语法或者 var = 表达式语法）， 变量的类型由（等号）右侧的值推导得出。

### 常量

1. 使用 `const` 关键字修饰。
1. 常量可以是字符、字符串、布尔或数字类型的值。
1. 常量不能使用 `:=` 语法定义。

### 数值常量

数值常量是高精度的值:

``` go
const (
    Big   = 1 << 100
    Small = Big >> 99
)
```

## 流程控制

### for循环

Go 只有for 循环这一种循环结构:

``` go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

基本的 for 循环包含三个由分号分开的组成部分：

- 初始化语句：在第一次循环执行前被执行
- 循环条件表达式：每轮迭代开始前被求值
- 后置语句：每轮迭代后被执行

for 语句的三个组成部分 并不需要用括号括起来，但循环体必须用 `{ }` 括起来。

> 注意：初始化语句一般是一个短变量声明，这里声明的变量仅在整个 for 循环语句可见。

循环初始化语句和后置语句都是可选的：

``` go
sum := 1
for ; sum < 1000; {
    sum += sum
}
```

#### 以for代替while

while 在 Go 中用 for 实现：

``` go
sum := 1
for sum < 1000 {
    sum += sum
}
```

#### 死循环

省略了循环条件，循环就不会结束，简洁的形式实现了死循环：

``` go
for {
}
```

### if语句

Go 的 if 语句不要求用 `( )` 将条件括起来，但 `{ }` 是必须有的。

``` go
if x < 0 {
    return sqrt(-x) + "i"
}
```

#### if的便捷语句

if 语句可以在条件之前执行一个简单语句:

``` go
if v := math.Pow(x, n); v < lim {
    return v
}
```

由这个语句定义的变量的作用域仅在 if 范围之内。

#### if else语句

在 if 的便捷语句定义的变量同样可以在任何对应的 else 块中使用。

``` go
if v := math.Pow(x, n); v < lim {
    return v
} else {
    fmt.Printf("%g >= %g\n", v, lim)
}
// 这里开始就不能使用 v 了
```

### switch语句

``` go
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("OS X.")
case "linux":
    fmt.Println("Linux.")
default:
    // freebsd, openbsd,
    // plan9, windows...
```

#### 没有条件的 switch

没有条件的 switch 同 switch true 一样。

``` go
t := time.Now()
switch {
case t.Hour() < 12:
    fmt.Println("Good morning!")
case t.Hour() < 17:
    fmt.Println("Good afternoon.")
default:
    fmt.Println("Good evening.")
}
```

这样可以用更清晰的形式来编写长的 if-then-else 链。

### defer

defer 语句会延迟函数的执行直到上层函数返回。

延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用。

#### defer 栈

延迟的函数调用被压入一个栈中。当函数返回时， 会按照后进先出的顺序调用被延迟的函数调用。

``` go
for i := 0; i < 10; i++ {
    defer fmt.Println(i)
}
```

## 复杂类型

### 指针

指针保存了变量的内存地址。

1. 类型 `*T` 是指向类型 `T` 的值的指针。其零值是 `nil` 。

   ``` go
   var p *int
   ```

1. `&` 符号会生成一个指向其作用对象的指针。

   ``` go
   i := 42
   p = &i
   ```

1. `*` 符号表示指针指向的底层的值。

   ``` go
   fmt.Println(*p) // 通过指针 p 读取 i
   *p = 21         // 通过指针 p 设置 i
   ```

1. 这也就是通常所说的“间接引用”或“非直接引用”。

1. 与 C 不同，Go 没有指针运算。

### 结构体

一个结构体（ struct ）就是一个字段的集合。（而 type 的含义跟其字面意思相符。）

``` go
type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    v.X = 4
    fmt.Println(v.X)
}
```

实例化结构体后，可以使用`.`来访问结构体字段。

#### 结构体指针

结构体字段可以通过结构体指针来访问。

``` go
v := Vertex{1, 2}
p := &v
p.X = 1e9
```

通过指针间接的访问是透明的。

#### 结构体文法

结构体文法表示通过结构体字段的值作为列表来新分配一个结构体。

使用 `Name:` 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。

``` go
var (
    v1 = Vertex{1, 2}  // 类型为 Vertex
    v2 = Vertex{X: 1}  // Y:0 被省略
    v3 = Vertex{}      // X:0 和 Y:0
    p  = &Vertex{1, 2} // 类型为 *Vertex
)
```

### 数组

类型 `[n]T` 是一个有 n 个类型为 T 的值的数组。

``` go
var a [10]int
```

- 数组的长度是其类型的一部分，因此数组不能改变大小。
- 这看起来是一个制约，但是请不要担心，Go 提供了更加便利的方式来使用数组。

#### slice

一个 slice 会指向一个序列的值，并且包含了长度信息。

`[]T` 是一个元素类型为 T 的 slice。

`len(s)` 返回 slice s 的长度。

``` go
s := []int{2, 3, 5, 7, 11, 13}
len(s)
```

#### slice 的 slice

slice 可以包含任意的类型，当然也可以包含另一个 slice。

``` go
game := [][]string{
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
    []string{"_", "_", "_"},
}
game[2][0] = "X"
```

#### 对 slice 切片

slice 可以重新切片，创建一个新的 slice 值指向相同的数组。

``` go
s[lo:hi]
```

表示从 lo 到 hi-1 的 slice 元素。

#### 构造 slice

slice 由函数 `make` 创建。这会分配一个全是零值的数组并且返回一个 slice 指向这个数组：

``` go
a := make([]int, 5)  // len(a)=5
```

为了指定容量，可传递第三个参数到 make：

``` go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
```

- 第二个参数是初始化的slice元素个数
- 第三个参数是slice的最大容量。

`cap()`返回slice的最大容量。

#### nil slice

slice 的零值是 nil 。

一个 nil 的 slice 的长度和容量是 0。

#### 向 slice 添加元素

Go 提供了一个内建函数 `append` ：

``` go
func append(s []T, vs ...T) []T
```

- `append` 的第一个参数 `s` 是一个元素类型为 `T` 的 slice ，其余类型为 `T` 的值将会附加到该 slice 的末尾。
- `append` 的结果是一个包含原 slice 所有元素加上新添加的元素的 slice。
- 如果 s 的底层数组太小，而不能容纳所有值时，会分配一个更大的数组。 返回的 slice 会指向这个新分配的数组。

相关参考：

- [append文档](https://go-zh.org/pkg/builtin/#append)
- 博文：[Go 切片：用法和本质](https://blog.go-zh.org/go-slices-usage-and-internals)

#### range

for 循环的 range 格式可以对 slice 或者 map 进行迭代循环。

当使用 for 循环遍历一个 slice 时，每次迭代 range 将返回两个值。 第一个是当前下标（序号），第二个是该下标所对应元素的一个拷贝。

``` go
var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
    for i, v := range pow {
        fmt.Printf("2**%d = %d\n", i, v)
    }
}
```

- 可以通过赋值给 `_` 来忽略序号和值。
- 如果只需要索引值，去掉 `, value` 的部分即可。

### map

map 映射键到值。

map 在使用之前必须用 make 来创建；值为 nil 的 map 是空的，并且不能对其赋值。

``` go
type Vertex struct {
    Lat, Long float64
}

var m map[string]Vertex

func main() {
    m = make(map[string]Vertex)
    m["Bell Labs"] = Vertex{
        40.68433, -74.39967,
    }
    fmt.Println(m["Bell Labs"])
}
```

#### map 的文法

``` go
type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": Vertex{
        40.68433, -74.39967,
    },
    "Google": Vertex{
        37.42202, -122.08408,
    },
}
```

若顶级类型只是一个类型名，则可以在文法的元素中省略它。

``` go
type Vertex struct {
    Lat, Long float64
}

var m = map[string]Vertex{
    "Bell Labs": {40.68433, -74.39967},
    "Google":    {37.42202, -122.08408},
}
```

#### 修改 map

1. 在 map m 中插入或修改一个元素：

   ``` go
   m[key] = elem
   ```

1. 获得元素：

   ``` go
   elem = m[key]
   ```

1. 删除元素：

   ``` go
   delete(m, key)
   ```

1. 通过双赋值检测某个键存在：

   ``` go
   elem, ok = m[key]
   ```

   如果 key 在 m 中， ok 为 true。否则， ok 为 false，并且 elem 是 map 的元素类型的零值。

   同样的，当从 map 中读取某个不存在的键时，结果是 map 的元素类型的零值。

### 函数值

函数也是值，可以像其他值一样传递，比如，函数值可以作为函数的参数或者返回值。

``` go
hypot := func(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}

hypot(5, 12)
```

#### 函数的闭包

Go 函数可以是一个闭包。

- 闭包是一个函数值，它引用了函数体之外的变量。
- 这个函数可以对这个引用的变量进行访问和赋值；换句话说这个函数被“绑定”在这个变量上。

``` go
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}
```

## 方法

Go 没有类。但是可以在结构体类型上定义方法。

方法就是带有一个特殊receiver参数的函数。

receiver出现在func关键字和方法名称的其自己的参数列表中。

``` go
type Vertex struct {
    X, Y float64
}

func (v *Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

v := Vertex{3, 4}
v.Abs()
```

### 方法是函数

记住：方法就是一个带有receiver参数的函数。

### 方法（续）

可以在非结构类型上声明方法。

``` go
type MyFloat float64

func (f MyFloat) Abs() float64 {
    if f < 0 {
        return float64(-f)
    }
    return float64(f)
}

f := MyFloat(-math.Sqrt2)
f.Abs()
```

只可以使用与方法在同一个包中定义的类型的receiver来声明方法。

不能对来自其他包的类型或基础类型定义方法。

### receivers为指针

可以使用指针receivers来声明方法。

这意味着receivers类型对于某些类型 T 具有字面语法 `*T`；另外类型 T 本身不能成为指针，比如 `*int`

有两个原因需要使用指针receivers。

1. 首先避免在每个方法调用中拷贝值（如果值类型是大的结构体的话会更有效率）。
1. 其次，方法可以修改接收者指向的值。

因此方法需要改变类型的值时，或者类型的结构体比较大的话需要使用指针receivers。

### 指针和方法

带有指针参数的函数，调用时必须接受一个指针：

``` go
type Vertex struct {
    X, Y float64
}

func ScaleFunc(v *Vertex, f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

var v Vertex
ScaleFunc(v, 5)  // Compile error!
ScaleFunc(&v, 5) // OK
```

而，带有指针receivers的方法，调用时可以接受值或者指针：

``` go
type Vertex struct {
    X, Y float64
}

func (v *Vertex) Scale(f float64) {
    v.X = v.X * f
    v.Y = v.Y * f
}

var v Vertex
v.Scale(5)  // OK
p := &v
p.Scale(10) // OK
```

因为Scale方法有指针receivers，Go会将`v.Scale(5)`解释为`(&v).Scale(5)`。

### 指针和方法，另一种情况

带有值参数的函数，调用时只能接受具有具体类型的值：

``` go
type Vertex struct {
    X, Y float64
}

func AbsFunc(v Vertex) float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

var v Vertex
fmt.Println(AbsFunc(v))  // OK
fmt.Println(AbsFunc(&v)) // Compile error!
```

而，带有值receivers的方法，调用时可以接受值或者指针：

``` go
type Vertex struct {
    X, Y float64
}

func (v Vertex) Abs() float64 {
    return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

var v Vertex
fmt.Println(v.Abs()) // OK
p := &v
fmt.Println(p.Abs()) // OK
```

Go会将`p.Abs()`解释为`(*p).Abs()`。

### 选择值receiver还是指针receiver

有以下两个原因选择指针receiver：

1. 方法可以修改receiver指向的值；
1. 避免法制每个方法调用时拷贝值。在使用大的结构体时更有效率。

一般而言，所有给定类型的方法都应该具有有值receivers或指针receivers，而不是两者的混合。（后面将会看到原因。）

## 接口

接口类型由一组方法(methods)定义的集合。

接口类型的值可以存放实现这些方法的任何值。

``` go
type Abser interface {
    methodName() returnType
}
```

### 接口是隐式实现的

类型通过实现接口的方法来实现接口。没有显示声明的必要，也就没有"implements"关键字。

隐式接口将定义接口从其实现中分离出来，这样接下来在没有预先安排的情况下出现在任何包中。

``` go
type I interface {
    M()
}

type T struct {
    S string
}

// 该方法意味着类型 T 实现了接口 I
// 但并没有明确的说明是这样
func (t T) M() {
    fmt.Println(t.S)
}
```

### 接口值

接口值可以被看作是一个值和一个具体类型的元组：

``` go
(value, type)
```

接口值保存了一个特定的底层具体类型的值。

调用接口值上的方法，将会执行相同名称的基础类型上的方法。

### 带有nil底层值的接口值

如果具体值而不是接口本身为 nil，则方法将会被调用，改调用带有nil receiver。

在某些语言中，这样会除法一个空指针异常，但是在Go中，通常会编写一些可以优雅处理这类情况的方法，这样的方法就是nil receiver。

注意：存有具体nil值的接口值，其本身是非nil的。

### Nil接口值

nil接口值既不包含值，也不包含具体类型。

调用nil接口上的方法会导致运行时异常，因为接口元组中没有没有类型信息，所以没有具体的方法去调用。

### 空接口

没有方法的接口类型就是空接口：

``` go
interface{}
```

空接口可以持有任何类型的值。每种类型至少实现零个方法。

空接口用于处理未知类型的值。比如，`fmt.Print`可以接受任意数量的interface {}类型的参数。

### 类型断言

类型断言提供了对接口值的底层实现值的访问途径。

``` go
t := i.(T)
```

这个语句断言了接口值i存有具体类型T，并且将基础类型T的值赋给了变量t。

如果i没有存有T，该语句会触发panic。

为了测试接口值是否存有具体的类型，类型断言会返回两个值：基础类型的值和一个报告断言是否成功的布尔值。

``` go
t, ok := i.(T)
```

如果i存有T，则t就是基础类型的值、ok就是true。

如果没有存有T，则t就是基础类型T的零值，这样并不会发生panic。

注意：该语法与读取map语法的相似之处。

### 类型switch

类型switch是一个允许串联多个类型断言的结构。

类型switch与普通的switch类似，但是类型switch中的case是具体的类型（不是值），并且这些具体类型会与给定接口值保存的类型进行比较。

``` go
switch v := i.(type) {
case T:
    // here v has type T
case S:
    // here v has type S
default:
    // no match; here v has the same type as i
}
```

类型switch的声明语法与类型断言类似，但是类型断言中指定的类型T在类型switch中已经被关键字`type`代替。

类型switch的声明测试了接口值i是否存有类型T或者类型S。在类型T或者类型S的case中，变量v分别成为类型T或者类型S，并持有i持有的值。在default case（在没有匹配到的地方）中，变量v与i的接口类型和值相同。

## Stringers

一个最普遍存在的接口是定义在 `fmt` 包中的 `Stringers`。

``` go
type Stringer interface {
    String() string
}
```

Stringer是一个可以一字符串描述其自身的类型。`fmt` 包（和其他包）会查找这个借口来打印值。

``` go
type Person struct {
    Name string
    Age  int
}

func (p Person) String() string {
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}
```

## Errors

Go程序使用错误值来表示错误状态。

`error`类型是内建借口：

``` go
type error interface {
    Error() string
}
```

（与 `fmt.Stringer` 类似，`fmt` 包在打印值时也会试图匹配 `error`。）

通常方法会返回一个`error`值，调用它的代码应该通过判断`error`是否等于nil来处理`error`。

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

`error`为nil意味着成功；非nil`error`意味着失败。

## Readers

io包指定了`io.Reader`接口，该接口表示表示从数据流读取结束。

Go标准库中有这个接口的[多种实现](https://golang.org/search?q=Read#Global)，包括文件、网络连接、压缩、加密等等。

`io.Reader`接口有一个`Read`方法：

```go
func (T) Read(b []byte) (n int, err error)
```

`Read`方法用数据填充给定的byte slice，并返回填充的字节数和错误信息。在数据流结束时，会返回`io.EOF`错误。

## Web服务器

http包通过任何实现了 `http.Handler` 的值来响应 HTTP 请求：

``` go
package http

type Handler interface {
    ServeHTTP(w ResponseWriter, r *Request)
}
```

## 图片

image包定义了`Image`接口：

``` go
package image

type Image interface {
    ColorModel() color.Model
    Bounds() Rectangle
    At(x, y int) color.Color
}
```

注意：`Bounds`方法的返回值`Rectangle`其实是在image包中定义的`image.Rectangle`。

`color.Color` 和 `color.Model` 也是接口，但是通常因为直接使用预定义的实现 `image.RGBA` 和 `image.RGBAModel` 而被忽视了。这些接口和类型由`image/color` 包定义。

## 并发

### Goroutines

Goroutines是Go运行时环境管理的轻量级线程：

``` go
go f(x, y, z)
```

开启一个新的 goroutine 执行：

``` go
f(x, y, z)
```

f，x，y和z是在当前的goroutine中定义的，f的执行发生在新goroutine中。

Goroutines在相同的地址空间中运行，因此访问共享内存必须是同步的。[sync](https://go-zh.org/pkg/sync/)提供了这样的帮助，不过在 Go 中并不经常用到，因为有其他的办法。（在接下来的内容中会涉及到。）

### Channels

Channel是有，类型的管道，可以说使用Channel操作符`<-`发送或收取值。

``` go
ch <- v    // Send v to channel ch.
v := <-ch  // Receive from ch, and
           // assign value to v.
```

（箭头就是数据流动的方向。）

像map和slice，channel必须在使用前创建：

``` go
ch := make(chan int)
```

默认情况下，在另一端准备好以前，发送和接受都会阻塞。这样goroutines就可以在没有明确的锁或竞态变量的情况下进行同步。

#### 缓冲 channel

channel是可以带缓冲的。为 make 提供第二个参数作为缓冲长度来初始化一个缓冲 channel：

``` go
ch := make(chan int, 100)
```

向带缓冲的 channel 发送数据的时候，只有在缓冲区满的时候才会阻塞。 当缓冲区为空的时候接收操作会阻塞。

#### range 和 close

发送者可以关闭channel来表明没有值会被发送了。接受者可以通过接收表达式的第二个参数测试判断channel是否已经关闭：

``` go
v, ok := <-ch
```

如果已经没有值可接受并且channel是关闭的，ok将会被赋值位false。

循环`for i := range c`会不断从channel接受值，知道channel关闭。

注意：只有发送者应该关闭channel，而不是接受者。向已经关闭的channel发送数据会导致panic。

注意：channel与文件不同；通常无需关闭channel；只有在接收者必须被告知没有更多数据时才有必要关闭，比如中断一个`range`。

### select

select 语句使得一个goroutine在多个通讯操作上等待。

select会阻塞，直到条件分支的某个分支可以继续执行，接着就会执行那条分支。

#### 默认选择

当select中其他条件分支都没有准备好时，会执行default分支。

为了尝试非阻塞方式发送或接受，可以使用default分支：

``` go
select {
case i := <-c:
    // use i
default:
    // receiving from c would block
}
```

## sync.Mutex

我们已经看到channel 非常适合用来在各个 goroutine 间进行通信。

但是如果不需要通信呢？比如只是想确保在一个时刻，只有一个goroutine可以访问共享变量，以此来避免冲突。

这个概念交做互斥，通常使用mutex这个数据结构来实现。

GO的标准库提供了[sync.Mutex](https://golang.org/pkg/sync/#Mutex)和一下两个方法来实现互斥：

- `Lock`
- `Unlock`

可以通过在代码前调用`Lock`方法，在代码后调用`Unlock`方法来保证一段代码的互斥执行。参见Inc方法。

也可以使用defer来确保互斥锁一定会被解锁。参见Value方法。