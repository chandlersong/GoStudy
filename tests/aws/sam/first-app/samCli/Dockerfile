FROM ubuntu:bionic

MAINTAINER chandler.song chandler605@outlook.com

# only need add this if you are in china
ADD sources.list /etc/apt/sources.list

RUN apt-get -y update && \
    apt-get -y upgrade

RUN apt-get install -y  gcc python3.6  python3-pip
    # only need add -i this if you are in china
RUN pip3 install aws-sam-cli -i https://mirrors.aliyun.com/pypi/simple/

ENV LC_ALL=C.UTF-8
ENV LANG=C.UTF-8


WORKDIR /opt/sam

ADD sam_entrypoint.sh /opt

#ENTRYPOINT ["sam","local","start-api"]
ENTRYPOINT ["samCli/sam_entrypoint.sh"]
CMD [""]
