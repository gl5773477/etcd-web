# etcd-web
参考 https://github.com/silenceper/dcmp 修改的基于etcdv3的后台管理系统。

目前还很简陋，仅支持键值的操作：增删改查。

后续完善：
- 用户权限操作.
- 集群操作.

配置：conf/config.toml
修改etcd服务地址Endpoints，默认为本机2379.

启动
```sh
> ./build.sh run web
```
