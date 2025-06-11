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

# 检查依赖是否发生变化
check_dependencies() {
    echo -e "${GREEN}检查依赖变化...${NC}"
    
    # 如果 go.mod 或 go.sum 文件不存在，说明依赖发生变化
    if [ ! -f "go.mod" ] || [ ! -f "go.sum" ]; then
        return 1
    fi
    
    # 计算 go.mod 和 go.sum 的 MD5 值
    local old_go_mod_md5=$(cat go.mod.md5 2>/dev/null || echo "")
    local old_go_sum_md5=$(cat go.sum.md5 2>/dev/null || echo "")
    
    local new_go_mod_md5=$(md5sum go.mod | awk '{print $1}')
    local new_go_sum_md5=$(md5sum go.sum | awk '{print $1}')
    
    # 保存新的 MD5 值
    echo "$new_go_mod_md5" > go.mod.md5
    echo "$new_go_sum_md5" > go.sum.md5
    
    # 比较 MD5 值
    if [ "$old_go_mod_md5" != "$new_go_mod_md5" ] || [ "$old_go_sum_md5" != "$new_go_sum_md5" ]; then
        echo -e "${YELLOW}检测到依赖变化！${NC}"
        return 1
    fi
    
    echo -e "${GREEN}依赖未发生变化${NC}"
    return 0
}

echo -e "${GREEN}开始部署...${NC}"

# 获取当前代码的 commit hash
OLD_HASH=$(git rev-parse HEAD)

# 强制重置本地修改并拉取最新代码
echo -e "${GREEN}强制更新代码...${NC}"
git reset --hard
git clean -fd
git pull origin main

# 获取更新后的 commit hash
NEW_HASH=$(git rev-parse HEAD)

# 检查是否有代码更新
if [ "$OLD_HASH" != "$NEW_HASH" ]; then
    echo -e "${YELLOW}检测到代码更新${NC}"
    
    # 检查依赖是否发生变化
    if check_dependencies; then
        echo -e "${YELLOW}尝试重启容器...${NC}"
        # 依赖未变化，尝试重启容器
        if docker-compose restart; then
            echo -e "${GREEN}容器重启成功！${NC}"
        else
            echo -e "${YELLOW}容器重启失败，开始重新构建...${NC}"
            docker-compose up -d --build
            cleanup_docker
        fi
    else
        echo -e "${YELLOW}检测到依赖变化，开始重新构建...${NC}"
        # 依赖发生变化，直接重新构建
        docker-compose up -d --build
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