# 安装方法
## 1.下载项目
```bash
git clone https://git.imooc.com/coding-339/chat.git
```


## 2.项目配置
建议采用goland IDE 
ADD Configuration->左上角+->go build类型
右侧填写
files :{你的项目路径}/hellox.x/main.go
workdir:{你的项目路径}/hellox.x/

将asset目录复制到对应项目下
如将asset复制到hellox.x下
## 3.依赖包安装

本项目依赖`x/net`,`x/time`包可能会被墙，使用如下指令可以安装
```bash
$mkdir -p $GOPATH/src/golang.org/x/
$cd $GOPATH/src/golang.org/x/
$git clone https://github.com/golang/net.git net
$git clone https://github.com/golang/time.git time
```
其他包依赖
```bash
go get github.com/go-xorm/xorm
go get github.com/gorilla/websocket
go get gopkg.in/fatih/set.v0
go get github.com/go-xorm/xorm
```
# 联系作者
如有疑问请联系作者
```
微信号jiepool-winlion,请备注来自imooc
微信公众号号智能生活优选,每天更新和技术相关的,有机会获得各种项目代码:betaidea 
qq群golang服务开发,吹牛装B真的大牛，里面都有：367573447
```
