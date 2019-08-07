package utils

import (	
	"errors"
	"os"
	"log"
	"fmt"
)

type ValNode struct{
	Row int //行
	Cal int //列
	Value int //值
}

/*稀疏数组
使用稀疏数组，来保留类似前面的二维数组(棋盘、地图等等)
把稀疏数组存盘，并且可以从新恢复原来的二维数组数
*/

func Demo()  {
	//1、创建数组
	var arrys [11][11] int
	arrys[1][2] = 1
	arrys[2][3] = 2
	//2、遍历数组
	for _, value := range arrys {
		for _, v := range value {
			fmt.Printf("%d\t",v)
		}
		fmt.Println()
	}
	//3、原数组——》稀疏数组
	/*分析：
	1、遍历数组，查找非0值，创建节点结构体
	2、存放在切片稀疏数组
	*/

	var sparseArr []ValNode
	valNode := ValNode{}
	//标准稀疏数组
	valNode = ValNode{
		Row : 11,
		Cal : 11,
		Value : 0,
	}
	sparseArr = append(sparseArr,valNode)
	for i, value := range arrys {
		for j, v := range value {
			if v == 0 {
				continue
			}
			valNode = ValNode{
				Row : i,
				Cal : j,
				Value : v,
			}
			sparseArr = append(sparseArr,valNode)
		}
		
	}
	//4、输出稀疏数组
	for i, v := range sparseArr {
		fmt.Printf("sparseArr[%d] = %v \n",i,v)	
	}

	sparseArr = File(sparseArr)

	/*
	//5、保存文件
	data, err := json.Marshal(&sparseArr)
	CheckErrorNew(err,"序列化切片失败")
	fmt.Println(string(data))
	writeFile(data)
	//6、读取文件
	filePath := "sparseArr.txt"
	res, err := ioutil.ReadFile(filePath)
	CheckErrorNew(err,"读取数据失败")
	fmt.Println(string(res))
	err = json.Unmarshal(res, &sparseArr)
	CheckErrorNew(err,"反序列化失败")
	*/
	//7、恢复为原数组
	/*分析：
	1、按行读取
	2、遍历每一行
	*/

	var recoverArr [11][11]int

	for i, valNode := range sparseArr {
		if i == 0 {
			continue
		}
		recoverArr[valNode.Row][valNode.Cal] = valNode.Value
	}
	TravesingTwoArray(recoverArr)
}


//错误检查Return-普通
func CheckErrorReturn(err error) {
    if err != nil {
        log.Fatal(err)
        return  
    }
}

//遍历二维数组
func TravesingTwoArray(arrys [11][11]int)  {
	for _, value := range arrys {
		for _, v := range value {
			fmt.Printf("%d\t",v)
		}
		fmt.Println()
	}
}

func writeFile(data  []byte){
	filePath := "sparseArr.txt"
	
	file, err := os.OpenFile(filePath,  os.O_RDWR |os. O_CREATE,066) 
	
	defer file.Close() 
	CheckErrorReturn(err)
	count, err := file.Write(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Write %d \n", count,)
}

func CheckErrorNew(err error, errStr string) {
    if err != nil {
		log.Fatal(err,errors.New(errStr))
       return
    }
}