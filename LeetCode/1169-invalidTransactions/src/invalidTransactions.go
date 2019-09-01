package src

import (
	"fmt"
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	if nil == transactions || len(transactions) < 1 {
		return nil
	}
	inval := []string{}
	stringMap := make(map[string]string)
	for i := 0; i < len(transactions); i++ {
		s := strings.Split(transactions[i], ",")
		sToi, _ := strconv.Atoi(s[2])
		if sToi > 1000 {
			inval = append(inval, transactions[i])
		}
		if ms, ok := stringMap[s[0]]; ok {
			mapEle := strings.Split(ms, ",")
			mapEle1, _ := strconv.Atoi(mapEle[1])
			mapEle2, _ := strconv.Atoi(mapEle[2])
			s2, _ := strconv.Atoi(s[1])
			if mapEle[3] != s[3] && abs(s2, mapEle1) < 60 {
				if sToi <= 1000 {
					inval = append(inval, transactions[i])
				}
				if mapEle2 <= 1000 {
					inval = append(inval, ms)
				}

			}
		}
		if sToi > 1000 {
			inval = append(inval, transactions[i])
		}
		stringMap[s[0]] = transactions[i]
	}
	return inval
}

func abs(a, b int) int {
	d := a - b
	if d > 0 {
		return d
	} else {
		return 0 - d
	}
}
func RemoveRepByMap(slc []string) []string {
	result := []string{}
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}
