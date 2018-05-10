# python操作xlsx文件心得

在网上浏览了一下，比较了一下，python使用`openpyxl`库处理xlsx文件是比较好的选择。

[openpyxl document](http://openpyxl.readthedocs.io/en/latest/index.html)

## 安装

``` bash
pip install openpyxl
```

## 基本使用

[Simple Usage](http://openpyxl.readthedocs.io/en/latest/usage.html)

基本的使用主要涉及以下三个模块：

![openpyxl主要模块](./images/main_entry.png)

- [workbook model](http://openpyxl.readthedocs.io/en/stable/api/openpyxl.workbook.workbook.html?highlight=workbook)

- [wooksheet model](http://openpyxl.readthedocs.io/en/stable/api/openpyxl.worksheet.worksheet.html?highlight=worksheet%20)

- [cell model](http://openpyxl.readthedocs.io/en/stable/api/openpyxl.cell.cell.html?highlight=cell)

### 1. 读取文件

#### 1. 导入`load_workbook`方法

``` python
from openpyxl import load_workbook
```

#### 1. 导入文件获取workbook

``` python
wb = load_workbook('file_name')
```

#### 1. 获取worksheet

1. 根据sheet名称获取

   ```  python
   ws = wb['sheet_name']
   ```

1. 遍历sheet获取

   ```  python
   for sheet in wb:
       ... ...
   ```

1. 操作完成后注意关闭workbook: `wb.close()`

### 2. 操作数据

获取worksheet后，获取其中的数据：

#### 1. 批量获取

1. 获取指定行

   ```  python
   sheet[1]
   ```

   输入所需行号：1, 2 ,3, ...

1. 获取指定列

   ```  python
   sheet['A']
   ```

   输入所需列标：'A', 'B', 'C', ...

1. 获取指定单元格

    ``` python
    sheet['A1']
    ```

1. 获取指定范围的单元格

    ``` python
    sheet['A2:D2']
    sheet['A2:A4']
    ```

#### 2. 操作cell

1. 获取值: `cell.value`

   > 注意：空表格 `cell.value != None`

1. 获取cell所在行: `cell.row`

1. 获取cell所在列：`cell.column`

### 3. 写入到文件

类似使用的例子：[Using filters and sorts](http://openpyxl.readthedocs.io/en/latest/filters.html?highlight=filter)

``` python
data = [
    ["Fruit", "Quantity"],
    ["Kiwi", 3],
    ["Grape", 15],
    ["Apple", 3]
]

wb = Workbook()
ws = wb.active
for r in date:
    ws.append(r)
wb.save('file_name')
```

- 写入文件的数据是**二维数组**；
- 使用`worksheet.append()`，填加一行内容，内容是一个数组；
- `save()`方法会以`'w'`的模式去写入文件，注意覆盖的问题。

### 4. 为单元格添加样式

样式相关的模块都在[openpyxl.styles package](http://openpyxl.readthedocs.io/en/stable/api/openpyxl.styles.html)

此处例子值添加单元格的底色：

#### 1. 导入PatternFill

``` python
from openpyxl.styles import PatternFill
```

#### 2. 创建填充实例

``` python
style = PatternFill("solid", fgColor="E2EFDA")
```

#### 3. 填充单元格

``` python
cell.fill = style
```

#### 4. 保存更改

``` python
wb.save(path)
```

#### 5. 例子

有时候期望填充一行而不是单个单元格，则可以进行如下操作：

``` python
from openpyxl import Workbook, load_workbook
from openpyxl.styles import PatternFill

path = './test.xlsx'
wb = load_workbook(filename=path)
ws = wb['Sheet1]
style = PatternFill("solid", fgColor="E2EFDA")
for cell in ws[1]
    cell.fill = style
wb.save(path)
wb.close() // maybe Unnecessary
```