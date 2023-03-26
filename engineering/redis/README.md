## Redis Cluster in Docker

- 拉取镜像

```shell
$ docker pull bitnami/redis-cluster:latest
```

- docker compose

```shell
# 建议使用附录 2 配置 
$ curl -sSL https://raw.githubusercontent.com/bitnami/bitnami-docker-redis-cluster/master/docker-compose.yml > docker-compose.yml
$ docker-compose up -d
```

## Reference

1. https://github.com/bitnami/bitnami-docker-redis-cluster
2. https://github.com/bitnami/bitnami-docker-redis-cluster/issues/3#issuecomment-1117426443