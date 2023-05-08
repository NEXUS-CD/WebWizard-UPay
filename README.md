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
3.建议vscode启动
# 开发说明
0.在main.go写入 swag的tag.name  比如: // @tag.name unit (unit是你的模块名)
1.先在router文件写入路由
1.1.userGroup := r.Group("/users")这样是可以理解为声明一个根路由
1.2	{
		userGroup.GET("/:id", controllers.GetUser)
	}这样是可以理解为声明了一个users/:id的路由，并调用controllers下的GetUser函数
2.进入controllers书写代码逻辑，并记得书写swagger注释，参照其他接口
3.进入services书写代码逻辑
4.进入models书写代码逻辑（这里主要是操作数据库）

# 代码开发规范
1.常量全大写
2.注释
3.命名小驼峰
# 分支名规范
前缀描述-WW-分支号，列：feat-WW-1
1.前缀描述:
  1.1 新功能：feat
  1.2 缺陷/bug：fix
  1.3 优化：perf
  1.4 样式优化：style
  1.5 文档完善：doc
# commit规范
git commit -m "前缀: ww-1@github名 任务描述xxxxx"
# 推送代码流程
https://www.yuque.com/yuqueyonghubka1if/piw85r/vbxdn16pnka5sf3l

# 部署流程