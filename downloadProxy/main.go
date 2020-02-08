package main

import (
	"fmt"
	"io"

	"net/http"
	"strconv"

	"strings"
	_ "tools/downloadProxy/log"

	log "github.com/donnie4w/go-logger/logger"
)

var download_cnt int

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	key := ""
	for k, v := range r.Form {
		key = k
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	url := key + "/archive/master.zip"

	re, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Fprintf(w, url+"下载出错")
		return
	}
	io.Copy(w, re.Body)
	download_cnt++
	log.Debug("url:", url)
	log.Debug("下载完成，网站总访问次数：" + strconv.Itoa(download_cnt))
	fmt.Fprintf(w, "下载完成，网站总访问次数："+strconv.Itoa(download_cnt)) //这个写入到w的是输出到客户端的
}

func main() {

	http.HandleFunc("/", sayhelloName)       //设置访问的路由
	err := http.ListenAndServe(":8080", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
