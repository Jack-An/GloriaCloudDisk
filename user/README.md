## 用户服务 user-service

## Timeline

    -- add get_user & create user api   2022.5.8

### user-api 提供对外接口

运行命令

```bash
    cd user/api
    go run user.go -f etc/user-api.yaml
```

### user-rpc  提供内部调用

运行命令

```bash
    cd user/rpc
    go run user.go -f etc/user.yaml
```

