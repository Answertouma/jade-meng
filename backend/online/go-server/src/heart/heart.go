package main

import (
	"fmt"
	"heart/handler/paris"
	"net/http"
	"time"
)

const (
	pparis = "/paris"
)

const (
	huitailang = "/huitailang"
)

func main() {

	fmt.Println(time.Now())
	fmt.Println("--------------jade start--------------")
	http.Handle(pparis, paris.NewParisHandler())
	http.ListenAndServe("192.168.43.95:8050", nil)
	fmt.Println(time.Now())
	fmt.Println("--------------jade end----------------")

}
