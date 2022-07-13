FROM alpine
#最小的linux系统
ADD category /category
# 将user可执行文件放到根目录下
ENTRYPOINT [ "/category" ]
# 启动user
