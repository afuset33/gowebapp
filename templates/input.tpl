<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>パスワード強度判定</title>
    <!--<link rel="stylesheet" type="text/css" href="../css/common.css">-->
</head>
<body>
    <h2>パスワード強度判定</h2>
    <p>パスワードを入力してください</p>
    <form method="POST" action="/result">
        {{if ne (len .ErrMsg) 0}}
        <ul class="errMsg">
            {{range $i, $v := .ErrMsg}}
            <li>
                {{$v}}
            </li>
            {{end}}
        </ul>
        {{end}}
        <div>
            <input name="password" type="text"></input>
        </div>
        <div>
            <button name="submit" type="submit">判定</button>
        </div>
    </form>
</body>
</html>