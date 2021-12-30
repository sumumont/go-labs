# 1.增加国际化配置
# 2.go底层的依赖库
FROM alpine:3.11
RUN apk add --no-cache libc6-compat
RUN apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone