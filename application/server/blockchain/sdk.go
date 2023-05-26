package blockchain

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/core/config"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
)

// 配置信息
var (
	sdk *fabsdk.FabricSDK // Fabric SDK
	//configPath    = "config.yaml"                                // 配置文件路径
	channelName   = "appchannel"                                 // 通道名称
	user          = "Admin"                                      // 用户
	chainCodeName = "fabric-mims"                                // 链码名称
	endpoints     = []string{"peer0.jd.com", "peer0.taobao.com"} // 要发送交易的节点

	configPath = "./config-local-dev.yaml" // 配置文件路径(本地开发时使用)  工作目录是main.go所在的目录，所以不用加./..
)

// Init 初始化
func Init() {
	var err error
	// 通过配置文件初始化SDK
	sdk, err = fabsdk.New(config.FromFile(configPath)) // 创建一个 FabricSDK 实例
	if err != nil {
		panic(err)
	}
}

// ChannelExecute 区块链交互
// fcn: 要调用的链码函数名
// args: 传递给链码函数的参数列表
func ChannelExecute(fcn string, args [][]byte) (channel.Response, error) {
	// 创建客户端，表明在通道的身份
	ctx := sdk.ChannelContext(channelName, fabsdk.WithUser(user)) // 创建一个 Fabric SDK 的 Channel 上下文，用于与指定的通道进行交互
	cli, err := channel.New(ctx)                                  // 使用 Fabric SDK 创建了一个客户端 cli，并通过 channel.New 方法进行初始化
	if err != nil {
		return channel.Response{}, err
	}
	// 对区块链账本的写操作（调用了链码的invoke执行指定的 fcn 函数）
	resp, err := cli.Execute(channel.Request{
		ChaincodeID: chainCodeName, // 要调用的链码名称
		Fcn:         fcn,
		Args:        args, // 传递给链码函数的参数列表
	}, channel.WithTargetEndpoints(endpoints...)) // 指定要调用的 Peer 节点的地址，以实现负载均衡和故障恢复等功能
	if err != nil {
		return channel.Response{}, err
	}
	//返回链码执行后的结果
	return resp, nil // 将链码执行后的结果 resp 返回给调用方
}

// ChannelQuery 区块链查询
// fcn: 要查询的链码函数名
// args: 传递给链码函数的参数列表
func ChannelQuery(fcn string, args [][]byte) (channel.Response, error) {
	// 创建客户端，表明在通道的身份
	ctx := sdk.ChannelContext(channelName, fabsdk.WithUser(user))
	cli, err := channel.New(ctx)
	if err != nil {
		return channel.Response{}, err
	}
	// 对区块链账本查询的操作（调用了链码的invoke），只返回结果
	resp, err := cli.Query(channel.Request{
		ChaincodeID: chainCodeName, // 要查询的链码名称
		Fcn:         fcn,
		Args:        args, // 传递给链码函数的参数列表
	}, channel.WithTargetEndpoints(endpoints...))
	if err != nil {
		return channel.Response{}, err
	}
	//返回链码执行后的结果
	return resp, nil
}
