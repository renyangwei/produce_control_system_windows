# 生管系统客户端程序使用说明 #

> 所有的文件都请使用notepad++打开，安装文件在打包文档中。举例：打开app.conf文件的方式为：右键app.conf文件，选择**Edit with notepad++** ,修改完成后保存退出。

## 文件目录 ##

![](https://i.imgur.com/GxwjRMK.png)

### app.conf文件 ###

生管系统客户端配置文件，配置如下：

| 配置参数        | 说明                      |
| -------------- | ------------------------ |
| host           | 数据库主机地址            |
| user           | 数据库用户名，默认sa用户   |
| pwd            | 数据库密码，密码为空则不填  |
| port           | 数据库监听端口,默认1433    |
| database       | 数据库，name表示数据库名称， group表示对应的产线|
| server_name    | 数据库实例名称，默认MSSQLSERVER |
| time_interval  | 轮询时间间隔，单位：秒，默认5秒  |
| debug          | 调试模式,是否开启日志,0开启,1关闭，默认0|
| rows_limit     | 客户端每次查询订单和完工资料的条数，显示在订单或者完工资料界面，默认10条 |
| local_test     | 本地测试,0开启,1关闭,默认1，请不要修改|
| udp_port		 | udp监听端口，用于监听生管系统发送的实时数据，注意发送的数据必须是**utf-8**格式 |

实例：

    {
	"server": [{
		"host": "192.168.15.128",
		"user": "sa",
		"pwd": "123456",
		"port": 1433,
		"server_name":"MSSQLSERVER",
		"database": [{
			"name": "scgl",
			"group": "一号线"
		}]
	},
	{
		"host": "localhost",
		"user": "sa",
		"pwd": "123456",
		"port": 1433,
		"server_name":"MSSQLSERVER",
		"database": [{
			"name": "scgl-1",
			"group": "二号线"
		}]
	}],
	"time_interval": 5,
	"debug": 0,
	"rows_limit": 10,
	"local_test": 1
    }

**说明**：请根据实际情况修改配置，可以配置多个主机地址，每个主机上可以配置多个数据库，注意**group字段不要重复**

### 生管系统.apk ###

生管系统 Android应用文件，在Android手机上打开文件即可安装。

### PaperManagementClientTest.exe ###

生管系统客户端程序，打开会日志窗口，一般用于测试。

![运行窗口](https://i.imgur.com/5qNNhdI.png)

### PaperManagementClientBkg.exe ###

生管系统客户端正式程序，用于后台运行

## 运行生管系统客户端 ##

### 测试 ###

1. 根据实际情况配置**app.conf**参数;
2. 保证有可以访问的数据，如何安装SQLserver并导入数据请自行百度；
3. 双击运行**PaperManagementClientTest.exe**文件;
4. 打开生管系统APP，提示输入厂家名称,比如：测试，点击确定，看到数据表示成功。

### 正式运行 ###

1. 根据实际情况配置**app.conf**参数;
2. 双击运行**PaperManagementClientTest.exe**文件;
3. 打开生管系统APP，提示输入厂家名称，点击确定，看到数据表示成功；
4. 关闭窗口，修改**app.conf**中的**debug**参数为1；
5. 双击**PaperManagementClientBkg.exe**文件（双击后没有任何反应）,然后打开任务管理器 -> 进程，看到 **PaperManagementClientBkg.exe\*32** 表示运行成功（如图）;
6. 设置开机启动:单击“开始→程序”，你会发现一个“启动”菜单，右击“启动”菜单选择“打开”即可将其打开，其中的程序和快捷方式都会在系统启动时自动运行,右键“新建”-“快捷方式”，选择要启动的文件。

![后台运行](https://i.imgur.com/HlFO18k.png)

## 资源地址 ##
- windows server 2008 操作系统 迅雷下载地址：
ed2k://|file|cn_windows_server_2008_r2_standard_enterprise_datacenter_and_web_with_sp1_x64_dvd_617598.iso|3368839168|D282F613A80C2F45FF23B79212A3CF67|/

- sql server 2008 数据库 迅雷下载地址：
ed2k://|file|cn_sql_server_2008_r2_enterprise_x86_x64_ia64_dvd_522233.iso|4662884352|1DB025218B01B48C6B76D6D88630F541|/

- windows sql server 2008 远程连接配置：
https://jingyan.baidu.com/article/6c67b1d6ca06f02787bb1ed1.html



