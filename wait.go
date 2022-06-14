package main

import (
	"fmt"
	"net/http"
	"sync"
)
var signals = []string{"test"}
var wg sync.WaitGroup

func main() {
	website :=[]string{
		"https://google.com",
		"https://fb.com",
		"https://whatsapp.com",
		"https://.chrome.com",

		}

    for _, web:=range website {
		go Statuscode(web)
		wg.Add(1)
	}
	wg.Wait()
}
func Statuscode(endpoint string){
	defer wg.Done()
	res,err := http.Get(endpoint)
	if err != nil{
		fmt.Println("oops is endpoint")
	}else {
		signals = append(signals,endpoint)
		fmt.Printf( "%d is status for %s",res.StatusCode, endpoint)

	}
	
}
//wait groups are advanced methods of time.packages
