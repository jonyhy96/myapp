<html>
  <head>
    <script src="/static/js/p5.js"></script>
    <script src="/static/js/sketch.js"></script>
    <script language="javascript"> 
    function geturl(){ 
        var url="http://api.baiyug.cn/vip/index.php?url="+document.getElementById("url").value; 
        document.location=url; 
        } </script> 
    <style type="text/css">
    .btn2{  
  border:none;  
  height:3.5em;  
  background-color:#3E8CD0;  
  color:white;  
  font-size:1.2em;  
  margin-top:0.5em;  
  width:20%;  
  border-radius:1em;  
}  
 </style>
  </head>
  <header class="hero-unit">
    <div class="container">
        <div class="row">
            <div class="hero-text">
                <h1>Welcome to the Sitepoint / Beego App!</h1>
                <p style="color:red" align="left">OEM:{{.OEM}}</p> 
		        <p style="color:red" align="left">VER:{{.VER}}</p> 
                <h2>This is My Test Version</h2>
                <input name="url" type="text" style="width:50%;height:10%"  id="url" value="https://v.qq.com/x/cover/v2bohhli64p784o.html" /> </br>
                <input class="btn2" type="submit" name="Submit" value="播放" onclick="geturl();"/>
                </br>
                <p>{{.Host}}</p>
                <p>{{.UserAgent}}</p>
                <p>{{.user_ip}}</p>
                </br>
                <p>{{.Website}} {{.Email}} {{.EmailName}} {{.Id}}</p>
            </div>
        </div>
    </div>
</header>
  <body>
  </body>
</html>