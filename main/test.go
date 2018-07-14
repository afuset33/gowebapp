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
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("../templates/result.tpl"))

	// フォームからパスワードを取得
	r.ParseForm()

	// パスワードの強度を判定
	var strength string = "弱"
	if checker.LengthCheck(r.Form.Get("password"), 8) {
		strength = "中"
	}
	if checker.ComboUpperLowerCase(r.Form.Get("password")) {
		strength = "強"
	}
	result := Result{
		ResultMsg: "パスワードの強度は " + strength + " です",
		Password:  r.Form.Get("password")}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "result.tpl", result); err != nil {
		log.Fatal(err)
	}
}
