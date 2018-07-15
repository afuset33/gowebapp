package main

import (
	"gowebapp/checker"
	"html/template"
	"log"
	"net/http"
)

type Result struct {
	ResultMsg string
	Password  string
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/result", resultHandler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("../templates/input.tpl"))

	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
	log.Printf("call handler")
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("../templates/result.tpl"))

	// フォームからパスワードを取得
	r.ParseForm()

	// パスワードの強度を判定
	var strength string
	satisfy := checker.GetSatisfiedNum(r.Form.Get("password"))

	switch satisfy {
	case 0, 1, 2:
		strength = "弱"
	case 3, 4:
		strength = "中"
	case 5:
		strength = "強"
	}

	result := Result{
		ResultMsg: "パスワードの強度は " + strength + " です",
		Password:  r.Form.Get("password")}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "result.tpl", result); err != nil {
		log.Fatal(err)
	}
	log.Printf("call resultHandler")
}
