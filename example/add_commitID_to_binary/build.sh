#!/bin/sh

# 获取当前commit号
CommitID=$(git rev-parse HEAD)

# 获取当前分支名称
Branch=$(git rev-parse --abbrev-ref HEAD)

# 获取最近的tag
Tag=$(git describe --abbrev=0 --tags)

# 打包时间
DATE=$(date +'%Y-%m-%dT%H:%M:%m+08:00')

go build -ldflags "-X 'main.Tag=$Tag' -X 'main.Branch=$Branch' -X 'main.CommitID=$CommitID' -X 'main.DATE=$DATE'"