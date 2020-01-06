FROM python:3.6.10-alpine3.11

MAINTAINER chandler.song chandler605@outlook.com

RUN echo "http://mirrors.aliyun.com/alpine/v3.11/main/" > /etc/apk/repositories && \
    apk update && \
    apk upgrade && \
    apk add --no-cache --virtual build-deps build-base gcc bash&& \
    pip install aws-sam-cli -i https://mirrors.aliyun.com/pypi/simple/ && \
    apk del build-deps

WORKDIR /opt/sam

ENTRYPOINT ["sh","samCli/sam_entrypoint.sh"]
CMD [""]
