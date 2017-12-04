# 一个基于grpc的Myapp web程序

---

## **现有功能：**
### **1. echo服务**
`实现服务器回显客户端输入的功能`
```sequence
Client->Server: Dial:serverip:4959
Note right of Server:  listen port 4959
Client->Server: Send string
Note right of Server:  packaging messange
Server->Client: Send echo
```

### **2. Time服务**
`实现服务器定时发送当前时间到客户端`
```sequence
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
![login](https://github.com/jonyhy96/train1/blob/master/login.jpg)
输入用户名，密码以登陆
`登陆成功过后会自动跳转到profile界面，否则则跳转到注册界面`
### **5. 用户个人信息显示**
![profile](https://github.com/jonyhy96/train1/blob/master/profile.jpg)
## **Useage：** ##
### **1.下载源码** 
```
    git clone https://github.com/jonyhy96/myapp.git
```
`您可能要修改源码中对mysql的操作指令，示例中使用的mysql是root:12345，database为myapp，table为customer`
`如果您不做修改可能会报错，出错原因大都是找不到 数据库或表单的问题，修改源码main.go即可解决`
### **2. 进入myapp目录下** 
```
    cd myapp
```
### **3. 编译出myapp镜像** 
```
    make
```
`你可以通过修改Dockerfile来修改myapp版本和名称等信息`
### **4.运行myapp服务** 
```
    docker run -p accessport:8088 -p echoport:4959 -p timeport:4950 myapp:latest ./myapp
```
**作者** [@jonyhy](https://weibo.com/u/5991880963)
