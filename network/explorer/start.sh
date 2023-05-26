#!/bin/bash

priv_sk_path=$(ls ../crypto-config/peerOrganizations/taobao.com/users/Admin\@taobao.com/msp/keystore/)

cp -rf ./connection-profile/network_temp.json ./connection-profile/network.json

sed -i "" "s/priv_sk/$priv_sk_path/" ./connection-profile/network.json

docker-compose down -v
docker-compose up -d

echo -e "\n"
echo "已启动区块链浏览器，访问 http://localhost:8080/#/，用户名 admin，密码 123456"
echo -e "\n"