package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"naive_test2/global"
)

func pushCommandToBafCloud(apiUrl string, command interface{}, apiKey string) error {
	// 构造请求体
	requestBody, _ := json.Marshal(command)
	global.SugarLogger.Debug("req Body:", string(requestBody))
	// 创建请求
	req, err := http.NewRequest("POST", apiUrl, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}

	// 设置请求头
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", apiKey))

	// 发送请求
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// 读取响应
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// 打印响应结果
	global.SugarLogger.Debug("Response Status:", resp.Status)
	global.SugarLogger.Debug("Response Body:", string(body))

	return nil
}

type Command struct {
	Uid         string `json:"uid"`
	Topic       string `json:"topic"`
	CommandType int    `json:"type"`
	Msg         string `json:"msg"`
}

func Turn(i int) {
	apiUrl := "https://apis.bemfa.com/va/postJsonMsg" // 这是一个示例 URL，替换为实际的 API URL
	var msg string
	if i == 0 {
		msg = "off"
	} else {
		msg = "on"
	}
	command := Command{
		CommandType: 1,
		Msg:         msg,
		Uid:         global.CONFIG.BafaConfig.Uid,
		Topic:       global.CONFIG.BafaConfig.Topic,
	} // 替换为你要推送的指令
	apiKey := "your_api_key_here" // 替换为你的 API 密钥

	err := pushCommandToBafCloud(apiUrl, command, apiKey)
	if err != nil {
		log.Fatal("Error sending command to BafCloud:", err)
	}
}
