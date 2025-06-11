#!/bin/bash

# 确保脚本在错误时退出
set -e

# 定义颜色
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m'

# 清理未使用的镜像和构建缓存
cleanup_docker() {
    echo -e "${YELLOW}清理未使用的镜像和构建缓存...${NC}"
    # 删除所有未使用的镜像
    docker image prune -af
    # 删除所有未使用的构建缓存
    docker builder prune -af
    # 删除所有未使用的数据卷
    docker volume prune -f
    # 删除所有未使用的网络
    docker network prune -f
    echo -e "${GREEN}清理完成！${NC}"
}

# 显示 Docker 空间使用情况
show_docker_space() {
    echo -e "${GREEN}Docker 空间使用情况：${NC}"
    docker system df
}

echo -e "${GREEN}开始部署...${NC}"

# 获取当前代码的 commit hash
OLD_HASH=$(git rev-parse HEAD)

# 拉取最新代码
echo -e "${GREEN}拉取最新代码...${NC}"
git pull origin main

# 获取更新后的 commit hash
NEW_HASH=$(git rev-parse HEAD)

# 检查是否有代码更新
if [ "$OLD_HASH" != "$NEW_HASH" ]; then
    echo -e "${YELLOW}检测到代码更新，尝试重启容器...${NC}"
    
    # 先尝试重启容器
    if docker-compose restart; then
        echo -e "${GREEN}容器重启成功！${NC}"
    else
        echo -e "${YELLOW}容器重启失败，开始重新构建...${NC}"
        # 构建并启动容器
        docker-compose up -d --build
        # 清理未使用的镜像和构建缓存
        cleanup_docker
    fi
else
    echo -e "${GREEN}代码未更新，跳过部署步骤${NC}"
fi

echo -e "${GREEN}部署完成！${NC}"

# 显示容器状态
echo -e "${GREEN}容器状态：${NC}"
docker-compose ps

# 显示 Docker 空间使用情况
show_docker_space 