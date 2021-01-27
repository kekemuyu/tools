package web

import (
	"html/template"
	"log"
	"net/http"
)

//网站首页
func index(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./web/static/html/index.html") //加载模板文件
	t.Execute(w, nil)                                           //模板文件显示在页面上
}

func test(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./web/static/html/test.html") //加载模板文件
	t.Execute(w, nil)                                          //模板文件显示在页面上
}
func Start() {
	http.Handle("web/static", http.StripPrefix("web/static", http.FileServer(http.Dir(""))))
	http.HandleFunc("/", index)
	http.HandleFunc("/test", test)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
