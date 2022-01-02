package main

import (
        "fmt"
)

/*
编写一个小程序：
给定一个字符串数组
[“I”,“am”,“stupid”,“and”,“weak”]
用 for 循环遍历该数组并修改为
[“I”,“am”,“smart”,“and”,“strong”]
 */

var testArray = [...]string{"“I”","“am”","“stupid”","“and”","“weak”"}
func main() {

fmt.Printf("初始数组内容:testArray为 %+v\n",testArray)
testArr()
modifyArr()
fmt.Printf("循环修改后的数组内容:testArray为 %+v\n",testArray)
}

func testArr(){
                for i := 0; i < (len(testArray)); i++ {
                        fmt.Println(testArray[i])        //循环遍历数组内容
                }
}

func modifyArr(){
        testArray[2] = "“smart”"
        testArray[4] = "“strong”"
        for i := 0; i < (len(testArray)); i++ {
                fmt.Println(testArray[i])        //循环修改遍历数组内容
        }

}
