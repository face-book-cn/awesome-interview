### LINUX

##### EOF

```shell
EOF是bash中定义的一个结束标识符，可以是任意符号
写法：
	/usr/bin/cat << -EOF (加上 - 代表不需要按照标准的缩进)
		444
		555
		666
	EOF
这段话代表把下面的一段代码交给指定的cat来执行，EOF非常有用，你也可以给mysql,python等等不是bash的程序在bash代码中执行
```

##### linux允许远程使用root账号

```shell
1.先用其他账号登录上去
2.sudo passwd root 为管理员设置密码
3.su 切换到管理员
4.vi /etc/ssh/sshd_config 进去ssh配置文件
5.复制 PermitRootLogin yes 加入到配置文件中
6.systemctl restart ssh 或者 service ssh restart 重启服务
```

##### 小技巧

```linux
遇到需要输入yes 或者 y 的命令，可以是用echo yes | rm -i aa.txt 
后台启动java -jar   nohup java -jar service-0.0.1-releases.jar >/dev/null  &  
查询软件是否有安装 rpm -qa | grep "软件或者包的名字"
-i 以交互模式运行
watch -n 1 "/sbin/ifconfig eth0 | grep bytes" 显示eth0这个网卡的流量
格式化时间  date +%Y-%m-%d--%T （%T直接显示时分秒）
ntsysv 看服务是否是开机启动
cat /var/log/messages 查看系统日志
cat /var/log/secure 用户登录日志
centos7 查看ip的几种方式 ip addr , ifconfig , hostname -I
yum list 插件名字（docker-ce）--showduplicates | sort -r 列出仓库所有的插件版本，并选择指定的安装并排序
yum update 更新centos版本
man 命令 查看命令所有的用法（比如 man ls，man date）
```

##### 定时任务

> [计算时间网站](https://tool.lu/crontab/)

```shell
crontab -e 创建执行定时任务的命令
crontab -l 列出有哪些任务在执行
cat /etc/crontab 查看官方指导的说明

# 买那个了使用如下  eg：* * * * * /usr/local/dhy.sh
# .---------------- minute (0 - 59)
# | .------------- hour (0 - 23)
# | | .---------- day of month (1 - 31)
# | | | .------- month (1 - 12) OR jan,feb,mar,apr ... 月份的缩写
# | | | | .---- day of week (0 - 6) (Sunday=0 or 7) OR sun,mon,tue,wed,thu,fri,sat 日期缩写
# | | | | | 空一格
# * * * * * user-name command to be executed 用户要执行的文件位置或者是命令
```

##### 查看磁盘/文件大小

```linux
df -h 磁盘大小
du -h filename 文件大小
```



##### 系统级别命令

> **systemctl是CentOS7的服务管理工具中主要的工具，它融合之前service和chkconfig的功能于一体。**

```shell
su 切换用户
server开头的是centos6版本  systemctl是centos7版本

服务命令：systemctl command name.service
服务启动：service name start –> systemctl start name.service
服务停止：service name stop –> systemctl stop name.service
服务重启：service name restart –> systemctl restart name.service
服务状态：service name status –> systemctl status name.service
服务开机启动：systemctl enable firewalld.service
服务开机禁用：systemctl disable firewalld.service
查看服务是否开机启动：systemctl is-enabled firewalld.service
查看已启动的服务列表：systemctl list-unit-files | grep enabled
查看启动失败的服务列表：systemctl --failed

条件式重启(已启动才重启，否则不做任何操作)
    systemctl try-restart name.service
重载或重启服务(先加载，然后再启动)
	systemctl reload-or-try-restart name.service
	
红帽版本看发行版：cat /etc/redhat-release 
linux版本信息：lsb_release -a  
查看内核：uname -r
查看所以linux信息：uname -a（Linux localdomain 3.10.0-957.el7.x86_64 x86_64 x86_64 GNU/Linux）
```

##### 防火墙命令

> **只需要看前三段即可，后面的命令基本用不着**
>
> **centos6的防火墙 iptables 命令一样，只需要把 firewalld 改为 iptables 即可**

```shell
防火墙基本命令（具体看系统级别命令即可）
 	启动： systemctl start firewalld
    查看状态： systemctl status firewalld 
    停止运行： systemctl stop firewalld
    重启：systemctl restart firewalld
    
配置firewalld-cmd
    查看版本： firewall-cmd --version
    查看帮助： firewall-cmd --help
    显示状态： firewall-cmd --state
    查看所有打开的端口： firewall-cmd --zone=public --list-ports
    更新防火墙规则： firewall-cmd --reload
    更新防火墙规则，重启服务： firewall-cmd --completely-reload
    查看已激活的Zone信息:  firewall-cmd --get-active-zones
    查看指定接口所属区域： firewall-cmd --get-zone-of-interface=eth0
    拒绝所有包：firewall-cmd --panic-on
    取消拒绝状态： firewall-cmd --panic-off
    查看是否拒绝： firewall-cmd --query-panic
    
防火墙开启和关闭端口（以下都是指在public的zone下的操作，不同的Zone只要改变Zone后面的值就可以）
    添加：firewall-cmd --zone=public --add-port=80/tcp --permanent（--permanent永久生效，没有此参数重		   启后失效） 
    重新载入：firewall-cmd --reload
    查看：firewall-cmd --zone=public --query-port=80/tcp
    删除：firewall-cmd --zone=public --remove-port=80/tcp --permanent
 
 								下面命令基本用不着
 
防火墙信任级别 通过Zone的值指定
    drop: 丢弃所有进入的包，而不给出任何响应 
    block: 拒绝所有外部发起的连接，允许内部发起的连接 
    public: 允许指定的进入连接 
    external: 同上，对伪装的进入连接，一般用于路由转发 
    dmz: 允许受限制的进入连接 
    work: 允许受信任的计算机被限制的进入连接，类似 workgroup 
    home: 同上，类似 homegroup 
    internal: 同上，范围针对所有互联网用户 
    trusted: 信任所有连接
    	
防火墙管理服务
    以smtp服务为例， 添加到work zone
    添加：
    firewall-cmd --zone=work --add-service=smtp
    查看：
    firewall-cmd --zone=work --query-service=smtp
    删除：
    firewall-cmd --zone=work --remove-service=smtp
  
防火墙配置IP地址伪装
    查看：
    firewall-cmd --zone=external --query-masquerade
    打开：
    firewall-cmd --zone=external --add-masquerade
    关闭：
    firewall-cmd --zone=external --remove-masquerade
    
防火墙端口转发
    打开端口转发，首先需要打开IP地址伪装 firewall-cmd --zone=external --add-masquerade
    转发 tcp 22 端口至 3753： 
    firewall-cmd --zone=external --add-forward-port=22:porto=tcp:toport=3753
    转发端口数据至另一个IP的相同端口：
    firewall-cmd --zone=external --add-forward-port=22:porto=tcp:toaddr=192.168.1.112
    转发端口数据至另一个IP的 3753 端口：
    firewall-cmd --zone=external --add-forward-port=22:porto=tcp:toport=3753:toaddr=192.168.1.112
```

##### tail cat mv

```shell
tail -f 100 filename 显示文件的最后一百行
tail -r -n 10 filename 逆序显示文件的最后10行

跟tail功能类似
cat [filename] | wc -l 查看行数 实际就跟 cat -n [filename] 一样
cat 从第一行开始显示档案内容。 都要配合管道流使用 例如：cat [filename] | head -n 1 只显示头部第一行
	-n 带行号显示文件行数
	-b 去掉空白行的行号
  	tac 从最后一行開始显示档案内容。
  	more 分页显示档案内容。
  	less 与 more 相似，但支持向前翻页
  	head 仅仅显示前面几行
  	tail 仅仅显示后面几行
  	od 以二进制方式显示档案内容
  	
cut 命令从文件的每一行剪切字节、字符和字段并将这些字节、字符和字段写至标准输出。
	-c 以字符为单位进行切分
	例子：cat aa.txt | cut -c 1 （文件aa.txt中按照第一个字符来切分）只会打印出一个字符
  	
删除任何.log文件；删除前逐一询问确认
	rm -i *.log
删除test子目录及子目录中所有档案删除,并且不用一一确认
	rm -rf test
删除以-f开头的文件
	rm -- -f*
将文件file1改名为file2，如果file2已经存在，则询问是否覆盖
	mv -i log1.txt log2.txt
```

##### vi和vim

> vim是vi的升级版，有vim不推荐用vi

相比vi的优势

1. 多级撤消 （在vi编辑器中，按u只能撤消上次命令，而在vim里可以无限制的撤消）
2. 易用性 （ vi编辑器只能运行于unix中，而vim不仅可以运行于unix，还可用于windows、mac等多操作平台。）
3. 语法加亮 （vim用不同颜色加亮代码）
4.  对vi完全兼容 （完全可以把vim当作vi来用）

快捷键

```vim
1.正常 --> 插入模式
    i：在当前光标所在字符的前面，转为输入模式；
    a：在当前光标所在字符的后面，转为输入模式；
    o：在当前光标所在行的下方，新建一行，并转为输入模式；
    I：在当前光标所在行的行首，转为输入模式；
    A：在当前光标所在行的行尾，转为输入模式；
    O：在当前光标所在行的上方，新建一行，并转为输入模式；
    
2.退出
	在正常模式下按组合键shift zz可以保存并退出
	
3.移动光标（正常模式）
    1）逐字符移动：
        h: 左；
        l: 右；
        j: 下；
        k: 上；
        #h: 移动#个字符
    2）以单词为单位移动
        w: 移至下一个单词的词首；
        e: 跳至当前或下一个单词的词尾；
        b: 跳至当前或前一个单词的词首；
        #w: 移动#个单词
    3）行内跳转：
        0: 绝对行首；
        ^: 行首的第一个非空白字符；
        $: 绝对行尾
    4）行间跳转
        #G：跳转至第#行；
        gg: 第一行；
        G：最后一行
    5）末行模式
        .: 表示当前行；
        $: 最后一行；
        #：第#行；
        +#: 向下的#行
        
4、翻屏（正常模式）
        Ctrl+f: 向下翻一屏；
        Ctrl+b: 向上翻一屏；
        Ctrl+d: 向下翻半屏；
        Ctrl+u: 向上翻半屏
        
5、复制字符
    1）正常模式
    
        复制：
            yy：复制当前行
            nyy：复制当前行至下面的n行
        
        粘贴：
            p：粘贴到光标的后面
            P：粘贴到光标的前面
    2）可视模式

        复制：
            y：复制当前行
            ny：复制当前行至下面的n行

        粘贴：
            p：粘贴到光标的后面
            P：粘贴到光标的前面
            
6、删除字符（正常模式）
        x: 删除光标所在处的单个字符；
        #x: 删除光标所在处及向后的共#个字符；
        d$或D:从当前光标处删除至行尾；
        d^:从当前光标处删除之行首；
        dd: 删除当前光标所在行；
        #dd: 删除包括当前光标所在行在内的#行；
注：dd相当于剪切操作，如果你dd之后按p或者P可以进行粘贴。

7、替换字符
        r：替换单个字符（按完r在按你要替换的字符即可）
        R：替换多个字符（从你要替换的位置开始替换，直至你退出正常模式）
8、撤销编辑操作：u
        u：撤消前一次的编辑操作；
        #u：直接撤消最近#次编辑操作；
        温馨提示：连续u命令可撤消此前的n次编辑操作；
        
9、将另外一个文件（/path/sunhui.txt）的内容填充在当前文件夹中
        ：r   /path/sunhui.txt ：填充到当前文件所在光标的后面

10、修改vim配置文件
        vim   ~/.vimrc：修改当前用户的vim配置文件
        vim    /etc/vimrc：修改所有用户的vim配置文件
        例：在当前用户的vim配置文件中添加显示行数的命令
        vim    ~/.vimrc：在末行添加 set nu 即可
        
11、拓展（末行模式）
    1）显示或取消显示行号
        ：set    nu            //显示
        ：set    number    //显示
        ：set    nonu        //取消
    2）设置语法高亮
        ：syntax    on    //开启
        ：syntax    off    //关闭
    3）分屏
        ：vsp xxx.x    //将两个文件垂直分屏
        ：ctrl+w w   //切屏
注：该特性当前有效，如果想要永久有效需修改配置文件
```



##### 三剑客

> grep 更适合单纯的查找或匹配文本
>
> sed  更适合编辑匹配到的文本
>
> awk  更适合格式化文本，对文本进行较复杂格式处理
>
> 他们都可以使用正则表达式来扩展

- sed

 ```shell
sed可以直接替换文件中的内容
sed 's/ //g' filename 去掉所有空格
sed -i 's/原字符串/新字符串/' /home/1.txt 不加就是匹配到就结束
sed -i 's/原字符串/新字符串/g' /home/1.txt g表示 匹配每一行有行首到行尾的所有字符
sed -i 's/\r$//' portainer.sh #win下结尾是\n\r 在linux下是\r 因此需要改下才可以在linux下执行
 ```

- awk

``` linux
awk支持逻辑判断，可以在里面加if else语句等等
awk 'NR==2{print $2'} NR指定第几行 $指定第几列 NF字段的数量（一行字段数量）length($n)打印出行的字数的个数

以多个分隔符分割
awk -F'[: |]' '{pritn $0}'（以 : 空格 | 分割）

AWK 包含两种特殊的模式：BEGIN 和 END。
BEGIN 模式指定了处理文本之前需要执行的操作：
END 模式指定了处理完所有行之后所需要执行的操作：

$NF表示最后一个字段 NF表示被分隔符切开后有多少个字段

NR 一行的元素
BEGIN{FS=":";OFS="---"}  FS指以什么分割。OFS是输出的时候以什么显示
```

- grep

```linux
单个文件搜索字符串
	grep "literal-string" filename

-A 显示匹配行之后的n行
    grep -A n "string" filename

-B 显示匹配行之前的n行
    grep -B n "string" filename
    
-C 显示匹配行前后的n行
    grep -C n "string" filename
   
递归搜索：-r 搜索当前目录以及子目录下含“this”的全部文件。
    grep -r "this" *

不匹配搜索：-v 显示不含搜索字符串“go”的行。
    grep -v "go" demo_text

统计匹配的行数：-c 统计文件中含“go”字符串的行数。
    grep -c "go" filename

只显示含有符串的文件的文件名：-l 显示含“this”字符串的文件的文件名。
    grep -l "this" filename

输出时显示行号：-n 显示含文件中含“this”字符串的行的行号。
    grep -n "this" filename
```

##### java环境变量

```shell
下载对应的jdk版本
编辑vim /etc/profile文件
把jdk所在地址加入到环境变量
export JAVA_HOME=/usr/local/jdk8
export PATH=$JAVA_HOME/bin:$PATH
export CLASSPATH=.:$JAVA_HOME/lib/dt.jar:$JAVA_HOME/lib/tools.jar
export JRE_HOME=$JAVA_HOME/jre
```

##### 时间转换

```shell
D1=$(date -d now +%s)
......
D2=$(date -d now +%s)
timex=$(($D2-$D1))
打印两个时间相差的时间戳

date +%s   可以得到UNIX的时间戳;
日期时间与时间戳互转：
date -d "2015-08-04 00:00:00" +%s     输出：1438617600

时间戳转换为字符串：
date -d @1438617600  "+%Y-%m-%d"    输出：2015-08-04

如果需要得到指定日期的前后几天：
seconds=`date -d "2015-08-04 00:00:00" +%s`       #得到时间戳
seconds_new=`expr $seconds + 86400`                   #加上一天的秒数86400
date_new=`date -d @$seconds_new "+%Y-%m-%d"`   #获得指定日前加上一天的日前
```

##### 图形和命令界面启动

```shell
1、命令行界面启动
systemctl set-default multi-user.target

2、图形界面模式
systemctl set-default graphical.target
```

##### 管道

> 实际就是一个流，把上一个命令的结果作为一个流交给下一个命令处理

##### 磁盘和内存

```shell
df -hl 查看磁盘空间
free -m以M的方式显示，同理-g就是以G的方式显示
```

##### 服务

```shell
chkconfig --list查看所有的服务
chkconfig和service命令的区别
chkconfig是当前不生效，Linux重启之后才生效的命令(开机自启动项)
service是即使生效，重启后失效的命令

find / -name "filename" 查找功能（不要引号也可以）

获取外网地址 curl icanhazip.com

top -n 1代表运行一次就停止，不再刷新，-n运行 数字代表运行多少次
```

##### shell

```shell
tty 看是那个终端
命令和文件自动补齐
命令具有记忆功能 
	上下键，!number,!string,!$,!!（用在脚本里面的，表示执行上一个命令）
别名功能 
	alias 取一个别名，重新进去失效
	unalias 取消一个别名
快捷键（这里的 ^ 是 ctrl 是意思）
	^V 块选择，删除前面是一大片空格时候很有用
	^R 重新链接服务器
	^A ^E 去到文件的第一行和最后一行（需要用的自己去查下，其他感觉没什么用）
前台后台
	在一个界面使用时，可以使用^Z把他调到后台，然后在用fg调用回来
命令排序
	； 不具有逻辑
	&& || 具有逻辑
	
##############################
$ 后台运行
$> 混合重定向 （既有标准输入，又有错误输入）
$$ 逻辑符号
##############################

通配符（元字符）
	* 任意多个字符
	？任意一个字符
	[] 匹配括号内的任意一个字符 [0-9] 可能时0-9中的任意一个
	() 子shell中执行
	{} 集合执行 
		mkdir /home/dhy/{aaa,bbb}
	\ 转义符
	
变量 
	变量写法 $变量名 或 ${变量名}
	$# 变量个数
	$1 位置变量
	$? 上一个命令的返回值 0表示成功
	$0 脚本名字
	$* 所有的参数
	$$ 当前进程的pid
	$! 上一个进程的pid
	export 环境变量
	`` 反引号 可以写为 $() 一样的
变量运算
 	整数	expr，$(())，$[]，let
 	小数	echo "2*4" |bc
变量替换
	$url=www.baidu.com
	echo ${url#*.} 从前往后，最短匹配 
	echo ${url##*.} 从前往后，最长匹配 贪婪匹配 
	echo ${url%.*} 从后往前，最短匹配
	echo ${url%%.*} 从后往前，最长匹配 贪婪匹配 

 	echo ${url/baidu/dhy} 把百度换成dhy
 	echo ${url//n/N} 把所有的n换成N 贪婪匹配 
 	
 	${变量名-新的变量值} 
 	echo ${var3-ccccc} 会把var3的内容变为ccccc
 
shell 条件测试 
	Shell 条件测试 
	格式 1： test 条件表达式
	格式 2： [ 条件表达式 ] 
	格式 3： [[ 条件表达式 ]] 
	
	文件测试
	[ -e dir|file ]
	[ -d dir ] 
	[ -f file ] 是否存在，而且是文件 
	[ -r file ] 当前用户对该文件是否有读权限 
	[ -x file ] 
	[ -w file ] 
	[ -L file ] 
	
	数值比较（也可以用c语言的 < > ! == ...）
	[ 1 -gt 10 ] 大于 
	[ 1 -lt 10 ] 小于 
	[ 1 -eq 10 ] 等于 
	[ 1 -ne 10 ] 不等于 
	[ 1 -ge 10 ] 大于等于 
	[ 1 -le 10 ] 小于等于 

expect 
 	非交互式的命令，具体请上网查
 	
if case 
	if 条件测试 ;then 命令序列 fi if和then在一行必须有分号
	if then elseif then else fi
	
for
	for ((初值;条件;步长)) do  循环体 done（c） 
	for 变量名 [ in 取值列表 ] do 循环体 done（shell）
while 
	while 条件测试 do 循环体 done ==当条件测试成立（条件测试为真），执行循环体 
	
until 
	until 条件测试 do 循环体 done ==当条件测试成立（条件测试为假），执行循环体 
break 结束本次
continue 跳过
exit 退出
shift 
array
	shell中也支持数据，具体上网查
function 
	函数，跟其他语言差不多
	
输出重定向 标识符号为1 (可以写成 1>) 例子：echo 1 > a.txt（会把 1 写进a.txt）
    - 标准输出重定向 " > "  数据沿箭头方向流动，原来文件内容会被丢弃
    - 标准输出追加重定向 " >> "  在原来文件结尾追加内容
    
输入重定向 标识符号为0 (可以写成 <0) 例子：cat < a.txt（会把a.txt的内容打印到控制台）
    - 标准输入重定向 " < "  数据沿箭头方向流动，原来文件内容会被丢弃
    - 标准输入追加重定向 " << "  在原来文件结尾追加内容

输出错误流 标识符号为2 (必须写成 2>) 例子：echo 1 2> a.txt 
    - 2> 标准错误重定向：把流向标准错误的数据重新定位到后边的文件中，文件原本内容会丢弃
    - 2>> 标准错误追加重定向：把流向标准错误的数据重新定位到后边的文件文件结尾处，在其尾部添加数据。文件原本内容不会被丢弃

反向单引号 " `` "
	会转义命令 例子：echo `date` （2019年 07月 10日 星期三 17:46:43 CST）
管道
	" | " 通过管道把前一个命令的输出交给后一个命令继续处理。相当于流 管道截流 tee -a（加a时代表追加）
单引号
	'' 就算有变量也被当做字符串处理
双引号
	"" 有变量就当做变量处理
```







