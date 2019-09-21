FROM xushikuan/alpine-build:1.0

ENV WORK_DIR=/app
ENV TIME_ZONE=Asia/Singapore
WORKDIR $WORK_DIR
RUN ln -snf /usr/share/zoneinfo/$TIME_ZONE /etc/localtime && echo $TIME_ZONE > /etc/timezone

COPY ./main $WORK_DIR
ENTRYPOINT ./main