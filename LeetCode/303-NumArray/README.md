## 区域和检索 - 数组不可变
给定一个整数数组  nums，求出数组从索引 i 到 j  (i ≤ j) 范围内元素的总和，包含 i,  j 两点。

示例：
```
给定 nums = [-2, 0, 3, -5, 2, -1]，求和函数为 sumRange()

sumRange(0, 2) -> 1
sumRange(2, 5) -> -1
sumRange(0, 5) -> -3
```
说明:

你可以假设数组不可变。
会多次调用 sumRange 方法。

## 方法三：缓存
### 复杂度分析

* 时间复杂度：每次查询的时间 O(1)，O(N) 预计算时间。由于累积和被缓存，每个sumrange查询都可以用 O(1) 时间计算。
* 空间复杂度：O(n).
