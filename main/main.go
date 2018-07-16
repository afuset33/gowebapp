package main

import (
	"gowebapp/checker"
	"html/template"
	"log"
	"net/http"
)

type Result struct {
	ResultMsg   string
	Password    string
	Suggestions []string
}

/*
ハンドラをリスナーに登録し、リスナーを起動します
*/
func main() {
	http.HandleFunc("/", inputHandler)
	http.HandleFunc("/result", resultHandler)
	http.ListenAndServe(":8080", nil)
}

/*
パスワード強度判定入力画面を表示するためのリクエストハンドラ
*/
func inputHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("../templates/input.tpl"))

	// テンプレートを描画
	if err := t.Execute(w, nil); err != nil {
		log.Fatal(err)
	}
	log.Printf("call inputHandler")
}

/*
パスワード強度判定結果画面を表示するためのリクエストハンドラ
*/
func resultHandler(w http.ResponseWriter, r *http.Request) {
	// テンプレートをパース
	t := template.Must(template.ParseFiles("../templates/result.tpl"))

	// フォームからパスワードを取得
	r.ParseForm()

	// パスワードの強度を判定
	var strength string
	satisfy, suggestions := checker.GetSatisfiedCondition(r.Form.Get("password"))

	switch satisfy {
	case 0, 1, 2:
		strength = "弱"
	case 3, 4:
		strength = "中"
	case 5:
		strength = "強"
	}

	result := Result{
		ResultMsg:   "パスワードの強度は " + strength + " です",
		Password:    r.Form.Get("password"),
		Suggestions: suggestions}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "result.tpl", result); err != nil {
		log.Fatal(err)
	}
	log.Printf("call resultHandler")
}
