## run step
stp1:进入go_path 目录新建github.com/yarm 如mkdir -p github.com/yarm   

stp2:git clone git@github.com:15902124763/go-scp.git  

stp3:下载第三方包   
进度条：  
https://github.com/qianlnk/pgbar  
gopath目录github.com新建qianlnk，然后git clone源码

   
st4:生成可执行exe
go build scp.go


## 命令行
scp -R /usr/local/upload.txt root@youId:/opt/src/  
或   
scp -R ./upload.txt root@youId:/opt/src/


## 可执行的exe  
可执行的exe在项目路径bin下的scp.exe文件  
Windows在控制台执行该命令：  
只需将下载好的scp.exe文件拷贝到C盘的路径C:\Windows\System32即可；  


