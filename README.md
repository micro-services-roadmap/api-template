## biz

1. 修改 go.mod 文件: 仓库地址
    - `wordpress-plus/api-app` --> `wordpress-plus/api-kol`
2. 依赖 `rpc-tracing` 服务

## develop

1. 设置私有仓库

   ```shell
   go env -w GOPRIVATE="github.com/wordpress-plus"
   ```

2. 使用 `.deploy/goctl` 目录进行相关业务逻辑的生成

   ```shell
   rm -rf internal/types && rm -rf internal/handler && goctl api go -api ./doc/api/app.api -dir ./ --home=./.deploy/goctl/1.6.5
   ```

3. 生成 swagger 文档

   ```shell
   # go install github.com/zeromicro/goctl-swagger@latest
   goctl api plugin -plugin goctl-swagger="swagger -filename ./doc/swagger/swagger.json" -api ./doc/api/app.api -dir .
   ```

4. 填充相关业务逻辑

   - /internal/logic/xxx
