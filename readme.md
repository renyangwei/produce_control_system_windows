## 使用说明 ##

### apk文件 ###

将 **PaperManagement.apk** 文件复制到Android手机，找到文件并安装，首次运行APP，弹出对话框提示“请输入名称”，输入“纸箱厂”，点击确定可以看到数据，app会自动刷新，右上角有倒计时，点击按钮可以停止/启动自动刷新功能

### PaperManagementClient.exe文件 ###

将**infor.txt**文件和**PaperManagementClient.exe**放到一个目录下，双击exe文件运行

### infor文件 ###

**infor.txt**文件中的数据表示实时数据，修改**infor.txt**文件中**Factory**和**Other**对应的数据（数据格式不能变），保存，看到**PaperManagementClient.exe**的窗口中提示**post success**后，等待APP自动刷新，就能看到修改后的数据，**infor.txt** 代表 **一号线**， **infor1.txt** 代表**二号线**，以此类推

### data文件 ###

**data.txt**文件中的数据表示历史数据，修改**data.txt**中的数据，保存，看到 **PaperManagementClient.exe**的窗口中提示**post success**后，进入APP中**历史数据**界面，可以查询对应时间、产线和班组的数据。**data.txt** 代表 **一号线**， **data1.txt** 代表 **二号线**， 以此类推

### 手动刷新功能 ###

在APP的**历史数据**界面中，长按**查询**，弹出对话框提示 **系统将重新读取数据，耗时较长**，点击确认即可手动刷新，如果在5秒内刷新失败，用户可以稍后再次查询。用户确定手动刷新后PaperManagementClient.exe程序目录下会生成location.txt文件，生管系统读取文件内容，将数据写入对应的**data.txt** 中

### path文件 ###

PaperManagementClient.exe程序读取**data.txt**和**infor.txt**的路径，每个路径之间用逗号隔开，最后一个路径后加上回车符号