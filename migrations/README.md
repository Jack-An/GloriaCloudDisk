Generic single-database configuration.


### 添加db migration

```bash
    alembic revision -m 'some info'
```

### 更新db
```bash
    alembic -x dburl="mysql+pymysql://root@127.0.0.1/net_disk_dev?charset=utf8" upgrade head
```

