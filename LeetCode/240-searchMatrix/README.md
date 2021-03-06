## 搜索二维矩阵 II
示例:

现有矩阵 matrix 如下：
```
[
  [1,   4,  7, 11, 15],
  [2,   5,  8, 12, 19],
  [3,   6,  9, 16, 22],
  [10, 13, 14, 17, 24],
  [18, 21, 23, 26, 30]
]
```
给定 target = 5，返回 true。

给定 target = 20，返回 false。
方法一：暴力法
对于每一行我们可以像搜索未排序的一维数组——通过检查每个元素来判断是否有目标值。

算法：
这个算法并没有做到聪明的事情。我们循环数组，依次检查每个元素。如果，我们找到了，我们返回 true。否则，对于搜索到末尾都没有返回的循环，我们返回 false。此算法在所有情况下都是正确的答案，因为我们耗尽了整个搜索空间。
复杂度分析

时间复杂度：O(mn)O(mn)。因为我们在 n\times mn×m 矩阵上操作，总的时间复杂度为矩阵的大小
空间复杂度：O(1)O(1)，暴力法分配的额外空间不超过少量指针，因此内存占用是恒定的。

方法二：二分法搜索
矩阵已经排过序，就需要使用二分法搜索以加快我们的算法。

算法：
首先，我们确保矩阵不为空。那么，如果我们迭代矩阵对角线，从当前元素对列和行搜索，我们可以保持从当前 (row,col) 对开始的行和列为已排序。 因此，我们总是可以二分搜索这些行和列切片。我们以如下逻辑的方式进行 : 在对角线上迭代，二分搜索行和列，直到对角线的迭代元素用完为止（意味着我们可以返回 false ）或者找到目标（意味着我们可以返回 true ）。binary search 函数的工作原理和普通的二分搜索一样,但需要同时搜索二维数组的行和列。
复杂度分析

时间复杂度：O(lg(n!))
空间复杂度：O(1)