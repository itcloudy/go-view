# Go-View


#### 参考源码: https://github.com/gmarik/go-erd


#### 出发点
在做golang应用的数据库表设计时，可以使用PowerDesigner或者Navicat，设计好表，然后生成sql语句或者直接连接数据库建表
当时没法直接生成golang的struct，很多的框架提供的scaffold可以直接从数据库逆向生成struct，如谢大的bee

很喜欢PD来设计，但是没有mac版的，故想换种思路，直接设计struct，然后通过工具生成展示各struct间关系的图，在github上搜索发现了[https://github.com/gmarik/go-erd](https://github.com/gmarik/go-erd),
但是没有field的说明，故在参考该源码，做了下修改，希望生成的图中带上备注

同时也想在这基础上进行扩展，直接生成关系型数据库的sql语句，故有了此项目

由于参考源码没有license，故在此特别说明出处


## 使用
```
go get github.com/itcloudy/go-view
```

go-view项目目录下执行
```
go-erd -path $(pwd)/examples |dot -Tsvg > $(pwd)/views/db-table-relation.svg
```