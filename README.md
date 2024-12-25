## biz

1. GO-ZERO: 1.7.3
2. 修改 go.mod 文件: 仓库地址
    - `wordpress-plus/app-api` --> `wordpress-plus/kol-api`
3. 依赖 `rpc-tracing` 服务
   ```shell
   # 设置私有仓库
   go env -w GOPRIVATE="github.com/wordpress-plus"
   ```

## develop

1. 使用 `.deploy/goctl` 目录进行相关业务逻辑的生成

   ```shell
   rm -rf internal/types && rm -rf internal/handler && goctl api go -api ./doc/api/app.api -dir ./ --home=./.deploy/1.7.3
   ```

2. 生成 swagger 文档

   ```shell
   # go install github.com/zeromicro/goctl-swagger@latest
   goctl api plugin -plugin goctl-swagger="swagger -filename ./doc/swagger/swagger.json" -api ./doc/api/app.api -dir .
   ```

3. 填充相关业务逻辑

   - /internal/logic/xxx
