package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"go-gitkraken-chinese/cst"
	"go-gitkraken-chinese/translate"
	"go-gitkraken-chinese/utils"
	"io/ioutil"
	"os"
	"strings"
)

func readFile(filename string) map[string]string {
	f, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("文件路径错误：", err)
	}
	jsonEn := string(f)
	enMap := make(map[string]string)
	err = json.Unmarshal([]byte(jsonEn), &enMap)
	if err != nil {
		fmt.Println("需要转化为的数据格式为json：", err)
	}
	zhMap := make(map[string]string)
	for key, value := range enMap {

		str := ""
		if utils.ElementInArray(value, cst.GetLeftSide()) {
			utils.StringTranslationCutting(value)
			//值中位置定位
			//leftStr := value[:strings.Index(value, leftSide)]
			//rightStr := value[strings.Index(value, rightSide)+len(rightSide):]
			//noTran := value[strings.Index(value, leftSide) : strings.Index(value, rightSide)+len(rightSide)]
			//str = translate.AliTranslate(leftStr) + noTran + translate.AliTranslate(rightStr)
		} else {
			//全文翻译
			str = translate.AliTranslate(value)
		}
		fmt.Println(str)
		zhMap[key] = str

	}
	return zhMap
}

func writeFile(zhMap map[string]string, fileDeal string, transType bool) {

	if transType {
		//增量
		file, err := os.OpenFile(fileDeal, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("文件打开失败", err)
		}
		//及时关闭file句柄
		defer func(file *os.File) {
			err := file.Close()
			if err != nil {
			}
		}(file)
		//写入文件时，使用带缓存的 *Writer
		write := bufio.NewWriter(file)
		//map2json

		_, err = write.WriteString(cst.GetDefaultStr())
		if err != nil {
			return
		}
		//Flush将缓存的文件真正写入到文件中
		err = write.Flush()
		if err != nil {
			return
		}
	}

}
func main() {
	filePath := "D:\\Codes\\Go\\src\\go-gitkraken-chinese\\data\\fanyi\\zh - 副本.json"
	//计算处理后文件路径
	fileDeal := "D:\\Codes\\Go\\src\\go-gitkraken-chinese\\data\\fanyi\\zh.json"
	fmt.Println(filePath)
	fmt.Println(fileDeal)
	zhMap := readFile(filePath)
	//zhMap := make(map[string]string)
	transType := true
	writeFile(zhMap, fileDeal, transType)
	fmt.Println(zhMap)
}
