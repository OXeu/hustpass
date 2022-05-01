package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"
)

//使用方法：在下列常量中填入自己的账号密码
//使用 go build -o hust-pass && ./hust-pass 构建并运行
//Docker运行：在当前目录：docker-compose up -d
const username = "U202012345"
const password = "Password"

func main() {
	fmt.Println("HUST Pass Service Started.")
	for {
		runtime.GC()
		loop()
	}
}

func loop() {
	resp, _ := http.Get("http://123.123.123.123")
	if resp == nil {
		time.Sleep(5 * time.Second)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Resp Close Error: ", err)
		}
	}(resp.Body)
	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	if strings.Contains(strBody, "eportal") {
		//需要认证
		fmt.Println("Need to Auth. Start Authing...")
		queryString := strBody[strings.Index(strBody, "?")+1 : strings.LastIndex(strBody, "'")]
		queryStringEnc := url.QueryEscape(queryString)
		params := fmt.Sprintf("userId=%s&password=%s&queryString=%s",
			username,
			password,
			queryStringEnc)
		resp, err := http.Post("http://172.18.18.60:8080/eportal/InterFace.do?method=login",
			"application/x-www-form-urlencoded",
			strings.NewReader(params))
		if err != nil {
			fmt.Println("Auth Error: ", err)
		}
		body, _ := ioutil.ReadAll(resp.Body)
		strBody = string(body)
		if strings.Contains(strBody, "success") {
			fmt.Println("Auth Success!")
		} else {
			fmt.Println("Auth Failed: ", strBody)
		}
		defer func(Body io.ReadCloser) {
			err := Body.Close()
			if err != nil {
				fmt.Println("Resp Close Error: ", err)
			}
		}(resp.Body)
	}
}
