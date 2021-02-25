### 从输入URL到页面加载发生了什么？

1. DNS 解析
2. TCP 连接
3. 发送 HTTP/HTTPS 请求
4. 服务器处理请求并返回 HTTP/HTTPS  报文
5. 浏览器解析渲染界面
6. 连接结束

#### DNS 解析过程

浏览器搜索自己的 DNS 缓存  → 搜索操作系统中的 DNS 缓存 → 搜索操作系统的 hosts 文件（windows 环境下）→ 操作系统将域名发送至 LDNS (本地域名服务器，本校或者运营商的DNS )，LDNS 查询自己的 DNS → LDNS 向Root Name Server （根域名服务器，存储用于每个域（com、net）的解析的顶级域名服务器的地址）请求，根域名服务器向 LDNS 返回具体域的 DNS 服务器地址→LDNS 向 com 域的域名服务器发起请求，得到具体的 IP 地址→LDNS 将得到的 IP 地址返回给操作系统，同时自己也将IP地址缓存。

#### 负载均衡

在请求真正进入到服务器前，可能还会先经过负载均衡服务器，它的作用是将请求合理地分配到多个服务器上，同时具备防御攻击等功能。比如在应用层实现的反向代理服务器 Nginx。

负载均衡服务器将请求定向到对应的web 服务器上，比如tomcat。

#### tomcat 处理过程

![](https://raw.githubusercontent.com/objcoding/objcoding.github.io/master/images/tomcat2.png)

Server：代表整个Tomcat，它包含所有的容器；

Service：相当于一个集合，包含多个Connector（连接）、一个Engine（引擎），它还负责处理所有Connector（连接）获取的客户请求；

Connector：一个Connector（连接）在指定的接口上侦听客户的请求，并将客户的请求交给Engine（引擎）来进行处理并获得回应返回给客户请求；

Engine：一个Engine（引擎）下可以配置多个虚拟主机Host，每个主机都有一个域名，当Engine获得一个请求时，会把这个请求发送的相应的Host上，Engine有一个默认的虚拟主机，如果没有虚拟主机能够匹配这个请求，那就由这个默认的虚拟主机来进行处理请求；

Host：代表一个Virtual host，每个虚拟主机都和某个网络域名想匹配，每个虚拟主机下面可以部署一个或者多个web app，每个web对应一个context，有一个context path，当一个host获取请求时，就把该请求匹配到某个context上；

Context：一个context对应一个web aplication，一个web由一个或多个servlet组成，Context在创建的时候将根据配置文件CATALINA_HOME/conf/web.xml和WEBAPP_HOME/WEB-INF/web.xml载入servlet类，当context获取请求时，将在自己的映射表中需找相匹配的servlet类，如果找到，则执行该类，获得请求的回应，并返回。

#### 处理请求过程

现在我来模拟Tomcat处理一个Http请求的过程：

设置一个来自客户端URL：http://localhost:8080/test/index

1. 服务器8080端口接收到客户发来的请求，被一个在那里监听的叫HTTP1.1的Connector获取了这个链接请求；
2. Connector把请求交给同在Service下的Engine去处理，并等待Engine的响应；
3. Engine把url解析，并把请求传给相对应的Host处理，如果没有相对应的Host，则用默认名叫localhost的Host来处理；
4. Host再把url解析为/webgateway/index.html，匹配context-path为/webgateway的Context去处理（如果匹配不到就把该请求交给路径名为””的Context去处理）；
5. context-path为/webgateway的Context会匹配Servlet Mapping为/index的Servlet处理；
6. 构造HttpServletRequest对象和HttpServletResponse对象，作为参数调用Servlet的doGet或doPost方法；
7. Context把处理完的HttpServletResponse对象返回给Host；
8. Host把HttpServletResponse对象返回给Engine；
9. Engine把HttpServletResponse对象返回给Connector；
10. Connector把HttpServletResponse对象返回给客户浏览器。
