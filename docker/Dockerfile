FROM ubuntu:16.04
MAINTAINER Jason Lee <jawc@hotmail.com>

RUN apt-get update -y && apt-get install -y && mkdir /opt/test-server

ADD test-server /opt/test-server/test-server

WORKDIR /opt/test-server

CMD ./test-server