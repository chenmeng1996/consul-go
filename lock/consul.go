package lock

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateSession(key string) (sessionId string) {
	body := map[string]string{
		"LockDelay": "15s",
		"Name":      key,
		"Node":      "consul-0",
		"Behavior":  "release",
		"TTL":       "30s",
	}
	bodyJson, err := json.Marshal(body)
	if err != nil {
		log.Println(err)
	}
	request, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:32500/v1/session/create", bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
	}
	client := http.Client{}
	response, err := client.Do(request)
	result, err := ParseResponse(response)
	sessionId = result["ID"].(string)
	return
}

func Lock(key, sessionId string) bool {
	body := "test"
	bodyJson, _ := json.Marshal(body)
	request, err := http.NewRequest(http.MethodPut, "http://localhost:32500/v1/kv/lock/"+key, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
	}

	query := request.URL.Query()
	query.Add("acquire", sessionId)
	request.URL.RawQuery = query.Encode()
	//fmt.Println(request.URL.String())

	client := http.Client{}
	response, err := client.Do(request)

	if err != nil {
		log.Println(err)
		return false
	}
	result, err := ParseResponse2Interface(response)
	if err != nil {
		log.Println(err)
		return false
	}
	resultBool, ok := result.(bool)
	return resultBool && ok
}

func UnLock(key, sessionId string) bool {
	body := "test"
	bodyJson, _ := json.Marshal(body)
	request, err := http.NewRequest(http.MethodPut, "http://localhost:32500/v1/kv/lock/"+key, bytes.NewBuffer(bodyJson))
	if err != nil {
		log.Println(err)
	}

	query := request.URL.Query()
	query.Add("release", sessionId)
	request.URL.RawQuery = query.Encode()

	client := http.Client{}
	response, err := client.Do(request)

	result, err := ParseResponse2Interface(response)
	if err != nil {
		log.Println(err)
		return false
	}
	resultBool, ok := result.(bool)
	return resultBool && ok
}

func DeleteSession(sessionId string) {
	request, err := http.NewRequest(http.MethodPut, "http://127.0.0.1:32500/v1/session/destroy/"+sessionId, nil)
	if err != nil {
		log.Println(err)
	}
	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	if response == nil {
		log.Println("response空")
	}
	result, err := ParseResponse2Interface(response)
	if err != nil {
		log.Println(err)
	}
	resultBool, ok := result.(bool)
	if !(ok && resultBool) {
		log.Println("删除session失败")
	}
}

func ParseResponse(response *http.Response) (map[string]interface{}, error) {
	result, err := ParseResponse2Interface(response)
	if err != nil {
		return nil, err
	}
	result1, ok := result.(map[string]interface{})
	if !ok {
		return nil, errors.New("类型转换错误")
	}
	return result1, nil
}

func ParseResponse2Interface(response *http.Response) (interface{}, error) {
	var result interface{}
	body, err := ioutil.ReadAll(response.Body)
	if err == nil {
		err = json.Unmarshal(body, &result)
	}
	return result, err
}
