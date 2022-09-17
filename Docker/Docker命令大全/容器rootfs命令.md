# Docker commit 命令

------

**docker commit :**从容器创建一个新的镜像。

### 语法

```
docker commit [OPTIONS] CONTAINER [REPOSITORY[:TAG]]
```

OPTIONS说明：

- **-a :**提交的镜像作者；

- **-c :**使用Dockerfile指令来创建镜像；

- **-m :**提交时的说明文字；

- **-p :**在commit时，将容器暂停。

  

### 实例

将容器a404c6c174a2 保存为新的镜像,并添加提交人信息和说明信息。

```
runoob@runoob:~$ docker commit -a "runoob.com" -m "my apache" a404c6c174a2  mymysql:v1 
sha256:37af1236adef1544e8886be23010b66577647a40bc02c0885a6600b33ee28057
runoob@runoob:~$ docker images mymysql:v1
REPOSITORY          TAG                 IMAGE ID            CREATED             SIZE
mymysql             v1                  37af1236adef        15 seconds ago      329 MB
```





# Docker cp 命令

------

**docker cp :**用于容器与主机之间的数据拷贝。

### 语法

```
docker cp [OPTIONS] CONTAINER:SRC_PATH DEST_PATH|-
docker cp [OPTIONS] SRC_PATH|- CONTAINER:DEST_PATH
```

OPTIONS说明：

- **-L :**保持源目标中的链接

### 实例

将主机/www/runoob目录拷贝到容器96f7f14e99ab的/www目录下。

```
docker cp /www/runoob 96f7f14e99ab:/www/
```

将主机/www/runoob目录拷贝到容器96f7f14e99ab中，目录重命名为www。

```
docker cp /www/runoob 96f7f14e99ab:/www
```

将容器96f7f14e99ab的/www目录拷贝到主机的/tmp目录中。

```
docker cp  96f7f14e99ab:/www /tmp/
```



# Docker diff 命令

------

**docker diff :** 检查容器里文件结构的更改。

### 语法

```
docker diff [OPTIONS] CONTAINER
```

### 实例

查看容器mymysql的文件结构更改。

```
runoob@runoob:~$ docker diff mymysql
A /logs
A /mysql_data
C /run
C /run/mysqld
A /run/mysqld/mysqld.pid
A /run/mysqld/mysqld.sock
C /tmp
```