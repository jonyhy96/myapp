<html>
  <head>
    <script src="/static/js/p5.js"></script>
    <script src="/static/js/sketch.js"></script>
    <script language="javascript"> 
    function geturl(){ 
        var url="http://api.baiyug.cn/vip/index.php?url="+document.getElementById("url").value; 
        document.location=url; 
        } </script> 
        
  </head>
  <header class="hero-unit">
    <div class="container">
        <div class="row">
            <div class="hero-text">
                <p style="color:red" align="left">OEM:{{.OEM}}</p> 
		        <p style="color:red" align="left">VER:{{.VER}}</p> 
                <h1>欢迎您！{{.Username}}</h1>
                <h2>This is My Test Version</h2>
                <p>VIP视频解析</p>
                <input name="url" type="text" id="url" value="" /> 
                <input type="submit" name="Submit" value="播放" onclick="geturl();"/>
                </br>
                <p>您的用户名称为： {{.Username}}</p>
                </br>
                <p>您的密码为：{{.Password}}</p>
                </br>
                <p>您的个人信息：{{.Userinfo}}</p>
                </br>
                <p>{{.Host}}</p>
                <p>{{.UserAgent}}</p>
                <p>{{.user_ip}}</p>
            </div>
        </div>
    </div>
</header>
  <body>
  </body>
</html>