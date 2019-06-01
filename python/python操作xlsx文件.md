# python操作xlsx文件

[Working with Excel Files in Python](http://www.python-excel.org/)

在网上浏览了一下，比较了一下，python使用`openpyxl`库处理xlsx文件是比较好的选择。

[openpyxl document](http://openpyxl.readthedocs.io/en/latest/index.html)

目录：

<!-- TOC -->

- [python操作xlsx文件心得](#python操作xlsx文件心得)
    - [安装](#安装)
    - [基本使用](#基本使用)
        - [读取文件](#读取文件)
            - [导入`load_workbook`方法](#导入load_workbook方法)
            - [导入文件获取workbook](#导入文件获取workbook)
            - [获取worksheet](#获取worksheet)
        - [操作数据](#操作数据)
            - [批量获取](#批量获取)
            - [操作cell](#操作cell)
        - [添加数据](#添加数据)
            - [添加行](#添加行)
            - [添加列](#添加列)
        - [删除数据](#删除数据)
            - [删除指定行](#删除指定行)
            - [删除指定列](#删除指定列)
        - [写入到文件](#写入到文件)
        - [为单元格添加样式](#为单元格添加样式)
            - [导入PatternFill](#导入patternfill)
            - [创建填充实例](#创建填充实例)
            - [填充单元格](#填充单元格)
            - [保存更改](#保存更改)
            - [例子](#例子)
    - [常用方法](#常用方法)
        - [workbook module常用属性](#workbook-module常用属性)
        - [worksheet module 常用属性](#worksheet-module-常用属性)
        - [cell model 常用属性](#cell-model-常用属性)

<!-- /TOC -->

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

### 读取文件

#### 导入`load_workbook`方法

``` python
from openpyxl import load_workbook
```

#### 导入文件获取workbook

``` python
wb = load_workbook('file_name')
```

#### 获取worksheet

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

### 操作数据

获取worksheet后，获取其中的数据：

#### 批量获取

1. 获取指定行

   ```  python
   sheet[1]
   ```

   输入所需行号：1, 2 ,3, ...

2. 获取指定列

   ```  python
   sheet['A']
   ```

   输入所需列标：'A', 'B', 'C', ...

3. 获取指定单元格

    ``` python
    sheet['A1']
    ```

4. 获取指定范围的单元格

    ``` python
    sheet['A2:D2']  # 获取行
    sheet['A2:A4']  # 获取列
    sheet['A1:D4']  # 获取4行4列
    ```

    - 注意：通过这样的方式得到的元素是个二元组

#### 操作cell

1. 获取值: `cell.value`

   > 注意：空单元格值为`None`

2. 获取cell所在行: `cell.row`

3. 获取cell所在列：`cell.column`

### 添加数据

#### 添加行

末尾增加一行

```python
ws.append(['cell1', 'cell 2'])
```

- [worksheet.append](https://openpyxl.readthedocs.io/en/latest/api/openpyxl.worksheet.worksheet.html#openpyxl.worksheet.worksheet.Worksheet.append)

指定位置插入一行

```python
ws.insert_rows(7)
for cell in ws[7]:
    cell.value = 'cell value'
```

#### 添加列

指定位置插入列

```python
ws.insert_cols(2)
for i in ws.max_row:
    cell = ws.cell(i, 2)
    cell.value = 'cell value'
```

[Insert column using openpyxl](https://stackoverflow.com/questions/15826305/insert-column-using-openpyxl)

### 删除数据

#### 删除指定行

```python
ws.delete_rows(2)
```

#### 删除指定列

```python
ws.delete_cols(2)
```

### 写入到文件

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

### 为单元格添加样式

样式相关的模块都在[openpyxl.styles package](http://openpyxl.readthedocs.io/en/stable/api/openpyxl.styles.html)

该例只添加单元格的底色：

#### 导入PatternFill

``` python
from openpyxl.styles import PatternFill
```

#### 创建填充实例

``` python
style = PatternFill("solid", fgColor="E2EFDA")
```

#### 填充单元格

``` python
cell.fill = style
```

#### 保存更改

``` python
wb.save(path)
```

#### 例子

有时候期望填充一行而不是单个单元格，则可以进行如下操作：

``` python
from openpyxl import Workbook, load_workbook
from openpyxl.styles import PatternFill

path = './test.xlsx'
wb = load_workbook(filename=path)
ws = wb['Sheet1']
style = PatternFill("solid", fgColor="E2EFDA")
for cell in ws[1]:
    cell.fill = style
wb.save(path)
wb.close() // maybe Unnecessary
```

- [styles.fills module](https://openpyxl.readthedocs.io/en/latest/api/openpyxl.styles.fills.html#module-openpyxl.styles.fills)

## 常用方法

### workbook module常用属性

1. `sheetnames`：获取工作表名称列表

### worksheet module 常用属性

1. `dimensions`：返回包含数据的所有单元格的最小边界范围

   - TODO: 不清楚会不会自动忽略空值行

2. `max_column`：包含数据的最大列数

3. `max_row`：包含数据的最大行数

4. `values`：按行获取所有单元格的值，返回一个生成器

### cell model 常用属性

1. `row`：单元格的行号（从1开始）

2. `column`：单元格的列索引字母（从A开始）

3. `col_idx`：单元格的列号（从1开始）

4. `column_letter`：单元格的字母索引（A, B, ...）

5. `coordinate`: 单元格的坐标（比如：A1）

6. `parent`：单元格所在的worksheet

7. `data_type`：单元格值的类型

   ``` python
   TYPE_STRING = 's'
   TYPE_FORMULA = 'f'
   TYPE_NUMERIC = 'n'
   TYPE_BOOL = 'b'
   TYPE_NULL = 'n'
   TYPE_INLINE = 'inlineStr'
   TYPE_ERROR = 'e'
   TYPE_FORMULA_CACHE_STRING = 'str'
   ```