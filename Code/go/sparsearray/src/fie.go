package utils

import (
	"io/ioutil"
	"encoding/binary"
	"bytes"
	"log"
	"os"
	"fmt"
)

func File(sparseArr []ValNode )  (vn []ValNode){
	for i, v := range sparseArr {
		fmt.Printf("sparseArr[%d] = %v \n",i,v)	
		// row := IntToBytes(v.Row)
		// fmt.Println(row)
		// writeFileByte(row)
		// writeFileByte([]byte("\r"))
		// cal := IntToBytes(v.Cal)
		// writeFileByte(cal)
		// writeFileByte([]byte("\r"))
		// value := IntToBytes(v.Value)
		// writeFileByte(value)
		// writeFileByte([]byte("\n"))

	}

 
	// writeFileByte([]byte(1))
	filePath := "sparses.txt"
	res, err := ioutil.ReadFile(filePath)
	CheckErrorNew(err,"读取数据失败")
	fmt.Println("读取数据:",string(res))

	
	return
}

func writeFileByte(data  []byte){
	filePath := "sparses.txt"
	
	file, err := os.OpenFile(filePath,  os.O_RDWR | os.O_APPEND |os. O_CREATE,066) 
	
	defer file.Close() 
	CheckErrorReturn(err)
	err = ioutil.WriteFile(filePath,data,066)
	if err != nil {
		log.Fatal(err)
	}
	
}


//整形转换成字节
func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian,&tmp)
	return bytesBuffer.Bytes()
}
 
//字节转换成整形
func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}