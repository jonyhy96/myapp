<html>
  <head>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/p5.js/0.5.16/p5.js"></script>
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
                <h1>Welcome to the Sitepoint / Beego App!</h1>
                <h2>This is My Test Version</h2>
                <input name="url" type="text" id="url" value="" /> 
                <input type="submit" name="Submit" value="播放" onclick="geturl();"/>
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