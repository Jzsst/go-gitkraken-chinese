package translate

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func aliTranslateEn2Ch(msg string) (string, error) {
	//这里添加post的body内容
	data := make(url.Values)
	data["srcLanguage"] = []string{"en"}
	data["tgtLanguage"] = []string{"zh"}
	data["bizType"] = []string{"message"}
	data["srcText"] = []string{msg}

	tarUrl := "https://translate.alibaba.com/translationopenseviceapp/trans/TranslateTextAddAlignment.do"
	//把post表单发送给目标服务器
	res, err := http.PostForm(tarUrl, data)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	body, err := ioutil.ReadAll(res.Body)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(res.Body)
	return string(body), nil
}

func AliTranslate(text string) string {
	if text == "" {
		return ""
	}
	str, err := aliTranslateEn2Ch(text)
	if err != nil {
		fmt.Println(err)
	}
	//json转map
	zhAli := make(map[string]interface{})
	err = json.Unmarshal([]byte(str), &zhAli)
	if err != nil {
		fmt.Println("需要转化为的数据格式为json：", err)
	}
	zhStr := zhAli["listTargetText"].([]interface{})[0]
	return zhStr.(string)
}
