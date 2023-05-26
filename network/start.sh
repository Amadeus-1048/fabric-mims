#!/bin/bash


# 检查操作系统类型
echo -e "\n"
echo "检查操作系统类型"
if [[ `uname` == 'Darwin' ]]; then
  echo "当前操作系统是 Mac"
  export PATH=${PWD}/hyperledger-fabric-darwin-amd64-1.4.12/bin:$PATH
elif [[ `uname` == 'Linux' ]]; then
  echo "当前操作系统是 Linux"
  export PATH=${PWD}/hyperledger-fabric-linux-amd64-1.4.12/bin:$PATH
else
  echo "当前操作系统不是 Mac 或 Linux，脚本无法继续执行！"
  exit 1
fi

echo -e "\n"
echo "一、清理环境"
echo "清除链码容器、链码镜像、Hyperledger Fabric网络的配置和数据"
./stop.sh

echo -e "\n\n\n"
echo "二、生成Hyperledger Fabric网络的证书和秘钥（ MSP 材料），生成结果将保存在 crypto-config 文件夹中"
echo "这些证书是身份的代表，在实体相互通信和交易的时候，可以对其身份进行签名和验证。"
echo "crypto-config.yaml定义了一个排序节点组织（OrdererOrgs）和两个对等节点组织（PeerOrgs），其中每个组织都包括名称、域名、节点和用户等信息。"
#排序节点组织QQ只包括一个节点orderer.qq.com。而对等节点组织Taobao和JD都包括两个节点（peer0和peer1），并且每个组织都有一个Admin用户和一个名为User1的普通用户。
cryptogen generate --config=./crypto-config.yaml

echo -e "\n\n\n"
echo "三、创建Hyperledger Fabric网络的排序通道创世区块"
echo "在Hyperledger Fabric网络中，创世区块（Genesis Block）是一个特殊的区块，它是整个区块链的第一个区块，没有前一个区块，也没有前一个区块的哈希。创世区块中包含了整个网络的初始状态、配置参数、通道信息、组织信息等重要的信息。在部署和启动Hyperledger Fabric网络时，首先需要生成创世区块。"
configtxgen -profile TwoOrgsOrdererGenesis -outputBlock ./config/genesis.block -channelID first-channel

echo -e "\n\n\n"
echo "四、生成通道配置事务'appchannel.tx'"
echo "在Hyperledger Fabric网络中，通道的交易配置用于定义和配置通道中的交易策略和参数，是确保交易安全性和可信度的关键"
configtxgen -profile TwoOrgsChannel -outputCreateChannelTx ./config/appchannel.tx -channelID appchannel

echo -e "\n\n\n"
echo "五、为 Taobao 定义锚节点"
echo "锚节点用于提高通道内部的通信效率，以便节点能够更快地发现其他节点的状态变化。"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/TaobaoAnchor.tx -channelID appchannel -asOrg Taobao

echo -e "\n\n\n"
echo "六、为 JD 定义锚节点"
configtxgen -profile TwoOrgsChannel -outputAnchorPeersUpdate ./config/JDAnchor.tx -channelID appchannel -asOrg JD

echo -e "\n\n\n"
echo "七、区块链 ： 启动"
echo "使用docker-compose工具启动容器化的Hyperledger Fabric网络，并设置不同节点的环境变量CORE_PEER_ADDRESS、CORE_PEER_LOCALMSPID和CORE_PEER_MSPCONFIGPATH，以便在后续的操作中访问和操作不同组织和节点"
docker-compose up -d
echo "正在等待节点的启动完成，等待10秒"
sleep 10

TaobaoPeer0Cli="CORE_PEER_ADDRESS=peer0.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp"
TaobaoPeer1Cli="CORE_PEER_ADDRESS=peer1.taobao.com:7051 CORE_PEER_LOCALMSPID=TaobaoMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/taobao.com/users/Admin@taobao.com/msp"
JDPeer0Cli="CORE_PEER_ADDRESS=peer0.jd.com:7051 CORE_PEER_LOCALMSPID=JDMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/jd.com/users/Admin@jd.com/msp"
JDPeer1Cli="CORE_PEER_ADDRESS=peer1.jd.com:7051 CORE_PEER_LOCALMSPID=JDMSP CORE_PEER_MSPCONFIGPATH=/etc/hyperledger/peer/jd.com/users/Admin@jd.com/msp"

echo -e "\n\n\n"
echo "八、创建通道"
echo "通道是Hyperledger Fabric中用于将多个组织的节点连接在一起的一种机制。通道中的所有节点共享相同的账本，并可以相互通信和交换数据"
docker exec cli bash -c "$TaobaoPeer0Cli peer channel create -o orderer.qq.com:7050 -c appchannel -f /etc/hyperledger/config/appchannel.tx"

echo -e "\n\n\n"
echo "九、将所有节点加入通道"
docker exec cli bash -c "$TaobaoPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$TaobaoPeer1Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$JDPeer0Cli peer channel join -b appchannel.block"
docker exec cli bash -c "$JDPeer1Cli peer channel join -b appchannel.block"

echo -e "\n\n\n"
echo "十、更新锚节点"
echo "锚节点是在通道中用于向其他节点传播配置信息的一种特殊节点。"
docker exec cli bash -c "$TaobaoPeer0Cli peer channel update -o orderer.qq.com:7050 -c appchannel -f /etc/hyperledger/config/TaobaoAnchor.tx"
docker exec cli bash -c "$JDPeer0Cli peer channel update -o orderer.qq.com:7050 -c appchannel -f /etc/hyperledger/config/JDAnchor.tx"

# -n 链码名，可以自己随便设置
# -v 版本号
# -p 链码目录，在 /opt/gopath/src/ 目录下
echo -e "\n\n\n"
echo "十一、安装链码"
echo "链码是Hyperledger Fabric中的智能合约，用于实现业务逻辑和操作账本数据。"
docker exec cli bash -c "$TaobaoPeer0Cli peer chaincode install -n fabric-mims -v 1.0.0 -l golang -p chaincode"
docker exec cli bash -c "$JDPeer0Cli peer chaincode install -n fabric-mims -v 1.0.0 -l golang -p chaincode"

# 只需要其中一个节点实例化
# -n 对应上一步安装链码的名字
# -v 版本号
# -C 是通道，在fabric的世界，一个通道就是一条不同的链
# -c 为传参，传入init参数
echo -e "\n\n\n"
echo "十二、实例化链码"
echo "实例化链码是将链码部署到通道中并启动它的过程，它需要在所有的对等节点上进行。"
docker exec cli bash -c "$TaobaoPeer0Cli peer chaincode instantiate -o orderer.qq.com:7050 -C appchannel -n fabric-mims -l golang -v 1.0.0 -c '{\"Args\":[\"init\"]}' -P \"AND ('TaobaoMSP.member','JDMSP.member')\""

echo "正在等待链码实例化完成，等待5秒"
sleep 5

# 进行链码交互，验证链码是否正确安装及区块链网络能否正常工作
echo -e "\n\n\n"
echo "十三、验证链码。在cli容器中进行链码交互，验证链码是否正确安装及区块链网络能否正常工作"
echo "使用变量TaobaoPeer0Cli指定在peer0.taobao.com节点上执行调用链码的命令，并使用peer chaincode invoke命令调用链码。指定了链码的名称（fabric-mims）、通道的名称（appchannel）以及调用链码的函数和参数"
docker exec cli bash -c "$TaobaoPeer0Cli peer chaincode invoke -C appchannel -n fabric-mims -c '{\"Args\":[\"hello\"]}'"

if docker exec cli bash -c "$JDPeer0Cli peer chaincode invoke -C appchannel -n fabric-mims -c '{\"Args\":[\"hello\"]}'" 2>&1 | grep "Chaincode invoke successful"; then
  echo "[Successful] network 部署成功。后续如需暂时停止运行，可以执行 docker-compose stop 命令（数据已持久化保存在Docker Volume中，不会丢失数据）。"
  exit 0
fi

echo "[Failed] network 未部署成功，请根据log信息定位具体问题并解决。"