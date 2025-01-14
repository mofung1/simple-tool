#!/bin/bash

# 设置错误时退出
set -e

# 设置编译环境变量
export CGO_ENABLED=0
export GOOS=linux
export GOARCH=amd64

# 获取当前时间作为版本号
VERSION=$(date "+%Y%m%d%H%M%S")

# 编译输出目录
OUTPUT_DIR="build"
rm -rf ${OUTPUT_DIR}
mkdir -p ${OUTPUT_DIR}

echo "====== 开始编译 ======"
echo "目标平台: Linux"
echo "目标架构: ${GOARCH}"
echo "版本号: ${VERSION}"

# 编译主程序
echo "正在编译..."
go build -ldflags "-X main.Version=${VERSION} -w -s" -o ${OUTPUT_DIR}/server ../cmd/main.go

# 创建必要的目录结构
#echo "正在创建目录结构..."
#mkdir -p ${OUTPUT_DIR}/config
#mkdir -p ${OUTPUT_DIR}/storage

# 创建版本信息文件
echo "版本号: ${VERSION}" > ${OUTPUT_DIR}/VERSION

# 打包
#echo "正在打包..."
#cd ${OUTPUT_DIR}
#tar -czf ../server-linux-amd64.tar.gz ./*
#cd ..

echo "====== 编译完成 ======"
echo "可执行文件: ${OUTPUT_DIR}/server"
echo "可执行文件大小: $(ls -lh ${OUTPUT_DIR}/server | awk '{print $5}')"
#echo "打包文件: server-linux-amd64.tar.gz"
#echo "压缩包大小: $(ls -lh server-linux-amd64.tar.gz | awk '{print $5}')"
