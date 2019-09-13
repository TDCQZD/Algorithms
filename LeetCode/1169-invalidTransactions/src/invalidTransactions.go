package src

import (
	"strconv"
	"strings"
)

func invalidTransactions(transactions []string) []string {
	if nil == transactions || len(transactions) < 1 {
		return nil
	}
	inval := []string{}
	nameMap := make(map[string][]string)

	// for i := 0; i < len(transactions); i++ {
	// tx := strings.Split(transactions[i], ",")
	for _, transaction := range transactions {
		tx := strings.Split(transaction, ",")

		name := tx[0]
		time1, _ := strconv.Atoi(tx[1])
		amount, _ := strconv.Atoi(tx[2])
		city := tx[3]
		//时间比较
		if _, ok := nameMap[name]; ok {
			for _, nameTX := range nameMap[name] {
				nametx := strings.Split(nameTX, ",")
				time2, _ := strconv.Atoi(nametx[1])
				// amountMap, _ := strconv.Atoi(nametx[2])
				if abs(time1, time2) <= 60 && nametx[3] != city {
					// if amountMap <= 1000 {
					inval = append(inval, nameTX)
					// }
					if amount <= 1000 {
						inval = append(inval, transaction)
					}
				}
			}
		}
		nameMap[name] = append(nameMap[name], transaction)
		//价格比较
		if amount > 1000 {
			inval = append(inval, transaction)
		}

	}
	// 城市比较时总会重复，所以加个 过滤
	return RemoveRepByLoop(inval)
}

func abs(a, b int) int {
	d := a - b
	if d > 0 {
		return d
	} else {
		return 0 - d
	}
}
func RemoveRepByLoop(slc []string) []string {
	result := []string{} // 存放结果
	for i := range slc {
		flag := true
		for j := range result {
			if slc[i] == result[j] {
				flag = false // 存在重复元素，标识为false
				break
			}
		}
		if flag { // 标识为false，不添加进结果
			result = append(result, slc[i])
		}
	}
	return result
}
