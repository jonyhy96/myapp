<html>
<head>
<meta charset="utf-8">
<style>
*{
	margin:0;
	padding:0;}
li{
	list-style-type:none;
	margin:10px auto 0 auto;
	width:400px;
	padding:0;}
a{
	text-decoration:none;
	color:#000;}
.login_putin ul li input{
	margin: 0;
	width:70%;
	height:20px;
	padding: 1em 2em 1em ;
	-webkit-border-radius:.3em;
	-moz-border-radius: .3em;
	border: 1px solid #999;
}
.login_btn{
	width:200px;
	margin:10px auto 0 auto;
}
.regist_btn input{
	width:200px;
	height:30px;
	-webkit-border-radius:.3em;
	background:#1b85fd;
	font-weight:bolder;
	letter-spacing:1em;
	border:#1263be solid 3px;
}
.login_btn input{
	width:200px;
	height:30px;
	-webkit-border-radius:.3em;
	background:#1b85fd;
	font-weight:bolder;
	letter-spacing:1em;
	border:#1263be solid 3px;
}
</style>
<title>Welcom to myapp</title>
</head>
<body bgcolor="black">
	<div align="center" class="login_img">
		<p style="color:red" align="left">OEM:{{.OEM}}</p> 
		<p style="color:red" align="left">VER:{{.VER}}</p> 
		</br>
	   <img src="/static/img/login.jpg" width=300px>	
	</div>
	<form method="post">
    <div align="center" class="login_putin">
    	<ul>
        	<li><input name="username" type="text"></li>
            <li><input name="password" type="password" ></li>
        </ul>
    </div>
    <div align="center" class="login_btn">
    	<input name="in_btn" type="submit" value="登录">
    </div>
	<div align="center" class="regist_btn">
    	<input name="re_btn" type="submit" value="注册">
    </div>
	</form>
</body>
</html>
