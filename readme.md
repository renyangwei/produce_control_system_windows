# 生管系统客户端程序使用说明 #

> 所有的文件都请使用notepad++打开，安装文件在打包文档中。举例：打开app.conf文件的方式为：右键app.conf文件，选择**Edit with notepad++** ,修改完成后保存退出。

## 文件目录 ##

![目录文件](https://i.imgur.com/5ALAMMo.png)

### test目录 ###

该目录存放**infor.txt**和**data.txt**文件，前者表示实时数据，后者表示历史数据，正常情况下都由生管系统自动生成，这里用于测试。

> infor.txt表示一号线的实时数据，infor1.txt表示二号线的实时数据,依次类推。

### path.txt ###

指定infor.txt和data.txt文件的路径，每个路径表示一条产线，比如:

    test\,test\

表示现在有两条产线，每条产线路径之间用逗号隔开，请在最后一个路径后加上**回车符号**。

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

**说明**：请根据实际情况修改配置，可以配置多个主机地址，每个主机上可以配置多个数据库，注意**group**字段不要重复

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
3. 将**infor.txt**文件中**Factory**对应的数据改成数据库中的厂家名称（数据格式不能变），保存；
4. 将**data.txt**文件中**Factory**对应的数据改成数据库中的厂家名称（数据格式不能变），保存；
5. 双击运行**PaperManagementClientTest.exe**文件;
6. 打开生管系统APP，提示输入厂家名称，点击确定，看到数据表示成功。

### 正式运行 ###

1. 根据实际情况配置**app.conf**参数;
2. 配置**path.txt**文件中的路径；
3. 双击运行**PaperManagementClientTest.exe**文件;
4. 打开生管系统APP，提示输入厂家名称，点击确定，看到数据表示成功；
5. 关闭窗口，修改**app.conf**中的**debug**参数为1；
6. 双击**PaperManagementClientBkg.exe**文件（双击后没有任何反应）,然后打开任务管理器 -> 进程，看到 **PaperManagementClientBkg.exe\*32** 表示运行成功（如图）;
7. 设置开机启动:单击“开始→程序”，你会发现一个“启动”菜单，右击“启动”菜单选择“打开”即可将其打开，其中的程序和快捷方式都会在系统启动时自动运行,右键“新建”-“快捷方式”，选择要启动的文件。

![后台运行](https://i.imgur.com/HlFO18k.png)
