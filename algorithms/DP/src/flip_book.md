# 题目：翻书问题或者跳台阶问题。

## 问题分析

```
Fibonacci(100) = Fibonacci(99) + Fibonacci(98) = Fibonacci(98) + Fibonacci(97) + Fibonacci(98)
```
## 解法
参考 flip_book.go


## 性能分析
递归实现的复杂度：O(2^n)
循环实现的复杂度：O(n)

**原因：**



## 单元测试结果
**1、递归实现**
```
e:\GoProject\src\it_skill_system\数据结构与算法\算法\DP\src>go test -v
=== RUN   TestFlipBook
1836311903
--- PASS: TestFlipBook (29.86s)
PASS
ok      it_skill_system/数据结构与算法/算法/DP/src      30.004s

```
**2.循环实现**
```
e:\GoProject\src\it_skill_system\数据结构与算法\算法\DP\src>go test -v
=== RUN   TestFlipBook
1836311903
--- PASS: TestFlipBook (0.00s)
PASS
ok      it_skill_system/数据结构与算法/算法/DP/src      0.224s
```
**总结**
循环性能更优。推荐使用循环
