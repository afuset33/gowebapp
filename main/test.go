package main

import (
	"html/template"
	"log"
	"net/http"
	"gowebapp/validator"
)

type Result struct {
	ResultMsg string
	Password  string
	Err
}

type Err struct {
	ErrMsg string
	isErr  bool
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
	result := Result{
		ResultMsg: "パスワードの強度は 強 です",
		Password:  r.Form.Get("password")}

	if validator.Required(result.Password) {
		err := Err{
			ErrMsg: "パスワードが入力されていません",
			isErr:  true}
		result.Err = err
	}
	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "result.tpl", result); err != nil {
		log.Fatal(err)
	}
}
