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
		txflag := false // 当前交易添加标志
		name := tx[0]
		time1, _ := strconv.Atoi(tx[1])
		amount, _ := strconv.Atoi(tx[2])
		city := tx[3]
		//时间比较
		if _, ok := nameMap[name]; ok {
			for _, nameTX := range nameMap[name] {

				nametx := strings.Split(nameTX, ",")
				time2, _ := strconv.Atoi(nametx[1])
				if abs(time1, time2) <= 60 && nametx[3] != city {
					if txflag {
						inval = append(inval, nameTX)
						txflag = false
					} else {
						inval = append(inval, nameTX, transaction)
						txflag = true
					}

				}
			}

		}
		nameMap[name] = append(nameMap[name], transaction)
		//价格比较
		if !txflag {
			if amount > 1000 {
				inval = append(inval, transaction)
			}

		}

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
