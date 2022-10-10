FROM alpine:latest AS final

ENV APP_PATH="/app"
WORKDIR "/app"

# 拷贝程序
COPY output ${APP_PATH}

RUN pwd
RUN chmod +x ./bootstrap.sh
CMD ["sh","./bootstrap.sh"]