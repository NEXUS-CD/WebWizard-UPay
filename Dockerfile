# 使用官方的 Alpine 镜像作为运行环境
FROM alpine:3.14

# 设置工作目录
WORKDIR /app
RUN mkdir -p /app/logger
RUN touch /app/logger/app.log
# 将可执行文件和配置文件复制到工作目录
COPY UPay /app/
COPY ./configs/config_docker.yaml /app/configs/
RUN mv /app/configs/config_docker.yaml /app/configs/config.yaml
COPY ./configs/config.go /app/configs/

# 设置容器中的环境变量（如有需要）
# ENV APP_ENV=production

# 暴露应用程序使用的端口
EXPOSE 8080

# 启动应用程序
CMD ["/app/UPay"]
