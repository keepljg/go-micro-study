FROM alpine

# 在容器的根目录下创建 app 目录
RUN mkdir /app

# 将工作目录切换到 /app 下
WORKDIR /app

ADD fileinfo-srv-srv /app/fileinfo-srv-srv

CMD [ "./fileinfo-srv-srv" ]
