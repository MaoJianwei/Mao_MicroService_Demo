FROM ubuntu:20.04
MAINTAINER Jianwei Mao <maojianwei2012@126.com>

WORKDIR /

COPY "microservice.kvstore" "/"
RUN chmod 777 /microservice.kvstore

EXPOSE 9876
CMD ["/microservice.kvstore"]
