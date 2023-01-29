部署要求：
环境：golang
数据存储：redis+mysql
导入数据库文件：./model/anwser.sql

构建二进制文件，程序入口 ./answer/api/answer.go
程序配置文件模板：./answer/api/etc/answer-api-example.yaml
启动程序 ./answer  -f 配置文件目录