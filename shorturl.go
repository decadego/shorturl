package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"os"
	"strings"
)

type Result struct {
	UrlShort string `json:"url_short"`
}

func main() {
	arg_num := len(os.Args)
	if arg_num != 2 {
		fmt.Println("Useage: shorturl http://www.baidu.com")
		return
	}

	url_long := os.Args[1]
	if !strings.HasPrefix(url_long, "http://") && !strings.HasPrefix(url_long, "https://") {
		fmt.Println("The url should start with http:// or https://")
		return
	}

	resp, err := http.Get("http://api.t.sina.com.cn/short_url/shorten.json?source=3271760578&url_long=" + url_long)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	res := &[]Result{}
	err = json.Unmarshal([]byte(body), &res)
	if err != nil {
		fmt.Println(err)
	} else {
		a := *res
		fmt.Println(a[0].UrlShort)
	}
	return
}
