# 压力测试

---

## **使用工具：**
### **webbench**
`webbench是Linux下的一个网站压力测试工具，最多可以模拟3万个并发连接去测试网站的负载能力`
### **安装过程：**
```
    git clone https://github.com/EZLippi/WebBench
    cd WebBench
    make && make install
```
### **使用方法：**
```
    webbench -c 并发数 -t 运行测试时间 URL
    webbench -c 5000 -t 120 http://www.baidu.com
```
### **测试结果：**
![presstest](https://github.com/jonyhy96/train1/blob/master/presstest.jpg)
**作者** [@jonyhy](https://weibo.com/u/5991880963)
