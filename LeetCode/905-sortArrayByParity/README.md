# 905. 按奇偶排序数组
给定一个非负整数数组 A，返回一个数组，在该数组中， A 的所有偶数元素之后跟着所有奇数元素。

你可以返回满足此条件的任何数组作为答案。

 

示例：
``` 
输入：[3,1,2,4]
输出：[2,4,3,1]
输出 [4,2,3,1]，[2,4,1,3] 和 [4,2,1,3] 也会被接受。
``` 

提示：

* 1 <= A.length <= 5000
* 0 <= A[i] <= 5000
## 方法 1：排序
### 思路
使用排序算法，按照模 2 的结果排序。
### 算法

### 复杂度分析

* 时间复杂度：O(NlogN)，其中 N 是 A 的长度。
* 空间复杂度：排序空间为 O(N)，取决于内置的 sort 函数实现。
## 方法 2：两边扫描
### 思路

第一遍输出偶数，第二遍输出奇数。
### 算法

### 复杂度分析

* 时间复杂度：O(N)O(N)，其中 NN 是 A 的长度。
* 空间复杂度：O(N)O(N)，存储结果的数组。
## 方法 3：原地算法
### 思路

如果希望原地排序，可以使用快排，一个经典的算法。
### 算法
维护两个指针 i 和 j，循环保证每刻小于 i 的变量都是偶数（也就是 A[k] % 2 == 0 当 k < i），所有大于 j 的都是奇数。

所以， 4 种情况针对 (A[i] % 2, A[j] % 2)：

- 如果是 (0, 1)，那么万事大吉 i++ 并且 j--。
- 如果是 (1, 0)，那么交换两个元素，然后继续。
- 如果是 (0, 0)，那么说明 i 位置是正确的，只能 i++。
- 如果是 (1, 1)，那么说明 j 位置是正确的，只能 j--。
通过这 4 种情况，循环不变量得以维护，并且 j-i 不断变小。最终就可以得到奇偶有序的数组。
### 复杂度分析
* 时间复杂度：O(N)，其中 N 是 A 的长度。循环的每一步都让 j-i 至少减少了一。（注意虽然快排的复杂度是 O(NlogN)，但是我们只需要一轮扫描就可以了）。
* 空间复杂度：O(1)，不需要额外空间。
