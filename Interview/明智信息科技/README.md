# 明智信息科技面试-区块链/GO
- 快速排序
- 字符串反转
- 反转字符串 II （输入: s = "abcdefg", k = 3 输出: "defgabc"）

## 快速排序
思想：选基准，分区，重复
```
func partition(arr []int, p, r int) int {
    x := arr[r]
    i := p - 1
    j := p
    for j >= P && j < r {
        if arr[j] <= x {
            i++
            arr[i],arr[j] = arr[j],arr[i]
        }
        j++
    }
    arr[i+1], arr[r] = arr[r], arr[i+1]
    return i+1
}

func quickSort(arr []int, p, r int){
    if p < r {
        q := partition(p,r)
        partition(arr,p,q-1)
        partition(arr,q+1,r)
    }

}
```
## 字符串反转
```
func reverString(str string) string {
    char := []runes[str]
    for i, j := 0, len(str) - 1;i < j; i++, j-- {
        char[i], char[j] = char[j], char[i]
    }

    return string(char)
}
```