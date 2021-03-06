# 922. 按奇偶排序数组 II
给定一个非负整数数组 A， A 中一半整数是奇数，一半整数是偶数。

对数组进行排序，以便当 A[i] 为奇数时，i 也是奇数；当 A[i] 为偶数时， i 也是偶数。

你可以返回任何满足上述条件的数组作为答案。

示例：
``` 
输入：[4,2,5,7]
输出：[4,5,2,7]
解释：[4,7,2,5]，[2,5,4,7]，[2,7,4,5] 也会被接受。
``` 

提示：

* 2 <= A.length <= 20000
* A.length % 2 == 0
* 0 <= A[i] <= 1000
## 方法一： 两次遍历
### 思路

遍历一遍数组把所有的偶数放进 ans[0]，ans[2]，ans[4]，依次类推。

再遍历一遍数组把所有的奇数依次放进 ans[1]，ans[3]，ans[5]，依次类推。
### 算法
### 复杂度分析
* 时间复杂度： O(N)，其中 N 是 A 的长度。

* 空间复杂度： O(N)。

## 方法二： 双指针
### 思路

我们可能会被面试官要求写出一种不需要开辟额外空间的解法。

在这个问题里面，一旦所有偶数都放在了正确的位置上，那么所有奇数也一定都在正确的位子上。所以只需要关注 A[0], A[2], A[4], ... 都正确就可以了。

将数组分成两个部分，分别是偶数部分 even = A[0], A[2], A[4], ... 和奇数部分 odd = A[1], A[3], A[5], ...。定义两个指针 i 和 j, 每次循环都需要保证偶数部分中下标 i 之前的位置全是偶数，奇数部分中下标 j 之前的位置全是奇数。

### 算法

让偶数部分下标 i 之前的所有数都是偶数。为了实现这个目标，把奇数部分作为暂存区，不断增加指向奇数部分的指针，直到找到一个偶数，然后交换指针 i，j 所指的数。


### 复杂度分析

* 时间复杂度： O(N)，其中 NN 是 A 的长度。

* 空间复杂度： O(1)。

