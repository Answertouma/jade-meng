package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type AnswerResponse interface {
}

type Answer struct {
	Love  string
	Sheep int32
}

func (*Answer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	message := r.URL.String()[7:]
	url := "http://www.baidu.com/s?wd=" +
		message +
		"&pn=30&ie=utf-8"
	fmt.Println(url)
	resp,err := http.Get(url)
	if err != nil {
		fmt.Println("aiyou")
	}
	baidu, err := ioutil.ReadAll(resp.Body)
	w.Write(baidu)
	return
}
