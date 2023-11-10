package chat

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const API_KEY = "ofDcudBBpNmmClKLfF6XX9yH"
const SECRET_KEY = "4GItIHjzPEeCkXPmYcwB3KjrG7t3c3t1"

type ApiMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Api struct {
	Message []ApiMessage `json:"messages"`
}

func wenxiapi(msg string) string {

	url := "https://aip.baidubce.com/rpc/2.0/ai_custom/v1/wenxinworkshop/chat/completions_pro?access_token=" + GetAccessToken()
	var apiMsg ApiMessage
	apiMsg.Role = "user"
	apiMsg.Content = msg

	bejson, err := json.Marshal(&Api{Message: []ApiMessage{apiMsg}})
	if err != nil {
		panic(err)
	}

	payload := strings.NewReader(string(bejson))

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, payload)

	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return err.Error()
	}
	fmt.Println(string(body))
	resultObj := map[string]string{}
	json.Unmarshal([]byte(body), &resultObj)
	return resultObj["result"]
	//return string(body)
}

/**
 * 使用 AK，SK 生成鉴权签名（Access Token）
 * @return string 鉴权签名信息（Access Token）
 */
func GetAccessToken() string {
	url := "https://aip.baidubce.com/oauth/2.0/token"
	postData := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s", API_KEY, SECRET_KEY)
	resp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(postData))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	accessTokenObj := map[string]string{}
	json.Unmarshal([]byte(body), &accessTokenObj)
	return accessTokenObj["access_token"]
}
