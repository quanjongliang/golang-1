package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func test() {

	url := "http://localhost:3000/videos"
	method := "GET"

	client := &http.Client{}
	req, _ := http.NewRequest(method, url, nil)

	req.Header.Add("Authorization", "Basic cHJhZ21hdGljOnJldmlld3M=")
	req.Header.Add("Content-Type", "application/json")

	res, _ := client.Do(req)

	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))
}
