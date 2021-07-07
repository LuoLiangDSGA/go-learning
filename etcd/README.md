## 本地搭建etcd集群

> 本地环境为macOS，在macOS上通过goreman搭建etcd集群

### 安装etcd

可以通过此脚本在macOS上安装etcd，安装成功后的etcd二进制文件目录为`/tmp/etcd-download-test/etcd`

``` shell
ETCD_VER=v3.5.0

# choose either URL
GOOGLE_URL=https://storage.googleapis.com/etcd
GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
DOWNLOAD_URL=${GOOGLE_URL}

rm -f /tmp/etcd-${ETCD_VER}-darwin-amd64.zip
rm -rf /tmp/etcd-download-test && mkdir -p /tmp/etcd-download-test

curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-darwin-amd64.zip -o /tmp/etcd-${ETCD_VER}-darwin-amd64.zip
unzip /tmp/etcd-${ETCD_VER}-darwin-amd64.zip -d /tmp && rm -f /tmp/etcd-${ETCD_VER}-darwin-amd64.zip
mv /tmp/etcd-${ETCD_VER}-darwin-amd64/* /tmp/etcd-download-test && rm -rf mv /tmp/etcd-${ETCD_VER}-darwin-amd64

/tmp/etcd-download-test/etcd --version
/tmp/etcd-download-test/etcdctl version
/tmp/etcd-download-test/etcdutl version
```

其他环境的脚本可[在此查看](https://github.com/etcd-io/etcd/releases)

### 安装goreman

[goreman](https://github.com/mattn/goreman) 是一个Go语言编写的多进程管理工具，是对Ruby下使用广泛的的Foreman的重写。

> 确保本地已经安装go环境

```shell
go get github.com/mattn/goreman
```

使用上述命令即可安装

### 编写集群脚本`procfile`

```shell

etcd1: /tmp/etcd-download-test/etcd --name infra1 --listen-client-urls http://127.0.0.1:2379 --advertise-client-urls http://127.0.0.1:2379 --listen-peer-urls http://127.0.0.1:12380 --initial-advertise-peer-urls http://127.0.0.1:12380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
etcd2: /tmp/etcd-download-test/etcd --name infra2 --listen-client-urls http://127.0.0.1:22379 --advertise-client-urls http://127.0.0.1:22379 --listen-peer-urls http://127.0.0.1:22380 --initial-advertise-peer-urls http://127.0.0.1:22380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
etcd3: /tmp/etcd-download-test/etcd --name infra3 --listen-client-urls http://127.0.0.1:32379 --advertise-client-urls http://127.0.0.1:32379 --listen-peer-urls http://127.0.0.1:32380 --initial-advertise-peer-urls http://127.0.0.1:32380 --initial-cluster-token etcd-cluster-1 --initial-cluster 'infra1=http://127.0.0.1:12380,infra2=http://127.0.0.1:22380,infra3=http://127.0.0.1:32380' --initial-cluster-state new --enable-pprof --logger=zap --log-outputs=stderr
#proxy: /tmp/etcd-download-test/etcd grpc-proxy start --endpoints=127.0.0.1:2379,127.0.0.1:22379,127.0.0.1:32379 --listen-addr=127.0.0.1:23790 --advertise-client-url=127.0.0.1:23790 --enable-pprof

```

使用脚本启动集群：

```shell
goreman -f setup/procfile start
```

验证结果：

```shell
$ /tmp/etcd-download-test/etcdctl --endpoints="http://127.0.0.1:12380,http://127.0.0.1:22380,http://127.0.0.1:32380" endpoint  health
http://127.0.0.1:32380 is healthy: successfully committed proposal: took = 19.388258ms
http://127.0.0.1:22380 is healthy: successfully committed proposal: took = 18.467176ms
http://127.0.0.1:12380 is healthy: successfully committed proposal: took = 18.280266ms
```

