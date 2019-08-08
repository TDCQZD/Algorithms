## 栈

栈（stack），是一种线性存储结构，它有以下几个特点：
1. 栈中数据是按照"后进先出（LIFO, Last In First Out）"方式进出栈的。
2. 向栈中添加/删除数据时，只能从栈顶进行操作。

栈通常包括的三种操作：push、peek、pop。
- push -- 向栈中添加元素。
- peek -- 返回栈顶元素。
- pop  -- 返回并删除栈顶元素的操作。

### 定义
栈是限定仅在表尾进行插入和删除的线性表。
- 允许插入和删除的一端称为栈顶(top),另一端称为栈底(bottom)。
- 不含任何数据元素的栈称为空栈。
- 栈又称为后进先出(Last In First Out)的线性表，简称LIFO结构
- 栈的插入操作，叫作进栈，也称压栈、入栈
- 栈的删除操作，叫作出栈，也称弹栈

### 栈的抽象数据类型
```
ADT 栈(stack)
Data
    同线性表。元素具有相同的类型，相邻元素具有前驱和后继关系。
Operation
    InitStack(*S):初始化操作，建立一个空栈S。
    DestroyStack(*S):若栈存在，则销毁它。
    ClearStack(*S):将栈清空。
    StackEmpty(*S):若栈为空，返回true，否则则返回false。
    GetTop(S,*e):若栈存在且非空，用e返回S的栈顶元素。
    Push(*S,e):若栈S存在，插入新元素e到栈S中并成为栈顶元素
    Pop(*S,*e):删除栈中栈顶元素，并用e返回其值
    StackLength(s):返回栈S的元素个数
endADT
```

### 示例
[ADT]()
