FROM alpine

MAINTAINER 苏凯飞

WORKDIR /sm

ADD /tms/data-parser-service.yml /sm/
ADD /tms/cmd /sm/

RUN apk add -U tzdata
RUN cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
RUN apk add --no-cache ca-certificates apache2-utils

CMD ["./cmd"]
