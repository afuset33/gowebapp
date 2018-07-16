package main

import (
	"gowebapp/checker"
	"gowebapp/validator"
	"html/template"
	"log"
	"net/http"
)

type Input struct {
	ErrMsg []string
}

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
	if err := t.ExecuteTemplate(w, "input.tpl", new(Input)); err != nil {
		log.Fatal(err)
	}
	log.Printf("call inputHandler")
}

/*
パスワード強度判定結果画面を表示するためのリクエストハンドラ
*/
func resultHandler(w http.ResponseWriter, r *http.Request) {

	// フォームからパスワードを取得
	r.ParseForm()
	password := r.Form.Get("password")
	// バリデーションチェック
	if valid, errMsgs := validation(password); !valid {
		input := Input{
			ErrMsg: errMsgs}
		// テンプレートをパース
		t := template.Must(template.ParseFiles("../templates/input.tpl"))
		// テンプレートを描画
		if err := t.ExecuteTemplate(w, "input.tpl", input); err != nil {
			log.Fatal(err)
		}
		log.Printf("bariche")
		log.Printf("call resultHandler")
		return
	}

	// テンプレートをパース
	t := template.Must(template.ParseFiles("../templates/result.tpl"))

	// パスワードの強度を判定
	var strength string
	satisfy, suggestions := checker.GetSatisfiedCondition(password)

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
		Password:    password,
		Suggestions: suggestions}

	// テンプレートを描画
	if err := t.ExecuteTemplate(w, "result.tpl", result); err != nil {
		log.Fatal(err)
	}
	log.Printf("call resultHandler")
	return
}

func validation(password string) (result bool, errMsgs []string) {
	result = true
	if valid, msg := validator.Required(password); !valid {
		errMsgs = append(errMsgs, msg)
		result = false
	}
	if valid, msg := validator.CharType(password); !valid {
		errMsgs = append(errMsgs, msg)
		result = false
	}
	if valid, msg := validator.MaxLength(password); !valid {
		errMsgs = append(errMsgs, msg)
		result = false
	}
	return
}
