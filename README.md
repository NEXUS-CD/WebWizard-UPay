# UPay
# 项目骨架说明
- **configs**：配置文件
  - `config.go`：读取配置文件的代码
  - `config.yaml`：默认的配置文件
  - `config_docker.yaml`：Docker环境下使用的配置文件
- **controllers**：控制器
  - `base_controller.go`：基础控制器
- **database**：数据库相关代码
  - `mongodb.go`：MongoDB相关代码
  - `redis.go`：Redis相关代码
- **docker**：Docker文件夹
  - **mongo**：MongoDB容器化配置文件夹
    - `docker-compose.yml`：MongoDB容器化配置文件
  - **redis**：Redis容器化配置文件夹
    - `docker-compose.yml`：Redis容器化配置文件
- **docs**：Swagger API文档
  - `docs.go`：生成Swagger API文档的代码
  - `swagger.json`：Swagger API文档的JSON格式
  - `swagger.yaml`：Swagger API文档的YAML格式
- **logger**：日志相关代码
  - `app.log`：应用的日志文件
  - `logger.go`：日志相关代码
- **models**：数据模型代码
  - `mongo_model.go`：MongoDB数据模型代码
  - `mongodb.go`：MongoDB访问代码
- **router**：路由配置代码
  - `router.go`：路由配置代码
- **services**：服务代码
  - `hello_service.go`：示例服务代码

# 项目启动说明
1.go mod tidy 下载依赖
2.vscode启动或者是go run main.go启动