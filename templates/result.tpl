<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>パスワード強度判定結果</title>
</head>
<body>
    <h2>パスワード強度判定結果</h2>
    <p>入力されたパスワード</p>
    <div>
    {{if .isErr}}
    {{.ErrMsg}}
    {{end}}
    {{.Password}}
    </div>
    <p>入力されたパスワードの強度</p>
    {{.ResultMsg}}
    <div>
        
    </div>
    <form method="POST" action="/input">
        <button name="back" type="submit">戻る</button>
    </form>
</body>
</html>