package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	resp, err := http.Get("http://www.baidu.com/s?wd=www&pn=30&ie=utf-8")
	if err != nil {
		fmt.Println("aiyou")
	}

	cc, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(cc))
}


