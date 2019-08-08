# 二分查找-Binarysearch

## 算法思想
二分搜索是一种在有序数组中查找某一特定元素的搜索算法。搜索过程从数组的中间元素开始，如果中间元素正好是要查找的元素，则搜索过程结束；如果某一特定元素大于或者小于中间元素，则在数组大于或小于中间元素的那一半中查找，而且跟开始一样从中间元素开始比较。如果在某一步骤数组为空，则代表找不到。这种搜索算法每一次比较都使搜索范围缩小一半。

## 过程分析
1. 先找到这个 数组的中间的下标  `midIndex =  (leftIndex + rightIndex) / 2`
2.  开始比较  让 findVal  和  arr[midIndex ]  比较
3.  如果  findVal  >  arr[midIndex ]   , 则应当到 arr数组 [  midIndex + 1 : rightIndex  ]
4.  如果  findVal  <  arr[midIndex ]   , 则应当到 arr数组 [  leftIndex : midIndex -1   ]
5.  如果  findVal  ==  arr[midIndex ]   , 找到
6.  将 2, 3, 4, 5 的逻辑进行递归.
7.  什么情况下，表示找不到 当   leftIndex   > rightIndex  成立，则表示没有