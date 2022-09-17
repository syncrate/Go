# Docker
## Docker的基本组成
镜像(image):  
容器(container):  
仓库(repository):仓库是集中存放镜像的地方

## 基础命令知识：命令行说明中格式
[   ]：内的内容是可选的

{   }：一组必须的项目，必须要在{a|b|c}内给出的选择里选一个

< >：表示必须被替换的占位

|    ：用于分隔多个互斥参数，含义为“或”，使用时只能选择一个


## 镜像命令
```shell
docker images [options] 
#options说明  -a列出所有的镜像（含历史） -q显示所有镜像ID
#说明 Repository:表示镜像的仓库源  Tag:镜像的标签版本号 Image ID:镜像ID  Created:镜像的创建时间 Size:镜像大小 
docker search [options]
#options说明 --limit:列出n个镜像 默认25  docker search --limit 5 redis
#搜索结果 Name:名字 Description:镜像说明 Stars:点赞数量 Official:是否是官方的 Automated:是否自动构建的

docker pull
docker system df
docker rmi

```

## 容器命令

