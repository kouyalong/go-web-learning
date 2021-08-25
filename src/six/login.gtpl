<html>
<head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>Golang学习 login</title>
</head>
<body>
<form action="/login" method="post">
	用户名:<input type="text" name="username">
	密码:<input type="password" name="password">
	<input type="hidden" name="token" value="{{.}}">
	<input type="submit" value="登录">
</form>
</body>
</html>