package botsend

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"path"
	"strconv"
)

func GetHook(key string) string {
	return "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?debug=1&key=" + key
}

func Request(key string, body []byte) *http.Response {
	url := GetHook(key)
	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	return response
}

type UploadResponse struct {
	Code    int    `json:"errcode"`
	Message string `json:"errmsg"`
	Type    string `json:"type"`
	Id      string `json:"media_id"`
}

func Upload(key string, filepath string) string {
	url := "https://qyapi.weixin.qq.com/cgi-bin/webhook/upload_media?key=" + key + "&type=file"

	filename := path.Base(filepath)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Fatal(err)
	}
	disp := fmt.Sprintf("Content-Disposition: form-data; name=\"media\";filename=\"%s\"; filelength=%d\r\n", filename, len(data))

	boundary := "-------------------------botsend-1234567890"
	start := "--" + boundary + "\r\n"
	end := "--" + boundary + "--\r\n"
	body := make([]byte, 0)
	body = append(body, []byte(start)...)
	body = append(body, []byte(disp)...)
	body = append(body, []byte("Content-Type: application/octet-stream\r\n\r\n")...)
	body = append(body, data...)
	body = append(body, []byte(end)...)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(body))
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Content-Type", "multipart/form-data; boundary="+boundary)
	request.Header.Set("Content-Length", strconv.Itoa(len(body)))

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	data, err = ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var uploadResponse UploadResponse
	err = json.Unmarshal(data, &uploadResponse)
	if err != nil {
		log.Fatal(err)
	}

	return uploadResponse.Id
}
