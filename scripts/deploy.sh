#!/bin/bash

# 确保脚本在错误时退出
set -e

# 定义颜色
GREEN='\033[0;32m'
NC='\033[0m'

echo -e "${GREEN}开始部署...${NC}"

# 拉取最新代码
echo -e "${GREEN}拉取最新代码...${NC}"
git pull origin main

# 构建并启动容器
echo -e "${GREEN}构建并启动容器...${NC}"
docker-compose up -d --build

# 清理未使用的镜像
echo -e "${GREEN}清理未使用的镜像...${NC}"
docker image prune -f

echo -e "${GREEN}部署完成！${NC}" 