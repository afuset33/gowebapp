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
        {{.Password}}
    </div>
    <p>入力されたパスワードの強度</p>
    <div>
        {{.ResultMsg}}
    </div>
    {{if len .Suggestions}}
    <ul>
        {{range $i, $v := .Suggestions}}
        <li>
            {{$v}}
        </li>
        {{end}}
    </ul>
    {{end}}
    <form method="POST" action="/input">
        <button name="back" type="submit">戻る</button>
    </form>
</body>
</html>