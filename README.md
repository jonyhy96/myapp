# 一个基于grpc的Myapp web程序

---

## **现有功能：**
### **1. echo服务**
`实现服务器回显客户端输入的功能`
```seq
Client->Server: Dial:serverip:4959
Note right of Server:  listen port 4959
Client->Server: Send string
Note right of Server:  packaging messange
Server->Client: Send echo
```

### **2. Time服务**
`实现服务器定时发送当前时间到客户端`
```seq
Client->Server: Dial:serverip:4950
Note right of Server:  listen port 4950
Note right of Server:  packaging messange
Server->Client: Send Time
```
### **3. 用户注册服务**
`访问界面:http://serverip:8088/user/signup`
![signup](https://github.com/jonyhy96/train1/blob/master/signup.jpg)
输入用户名，密码，个人信息完成注册
`注册成功过后会自动跳转到登录界面`
### **4. 用户登录服务**
`访问界面:http://serverip:8088/user/login`
![signup](https://github.com/jonyhy96/train1/blob/master/login.jpg)
输入用户名，密码以登陆
`登陆成功过后会自动跳转到profile界面，否则则跳转到注册界面`

浏览器返回界面：
![view](https://github.com/jonyhy96/train1/blob/master/imag.jpg)
作者 [@jonyhy](https://weibo.com/u/5991880963)
