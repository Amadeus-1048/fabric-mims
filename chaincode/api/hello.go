package api

import (
	"chaincode/pkg/utils"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// Hello 测试 Hyperledger Fabric 智能合约的编写和部署
// stub：智能合约与区块链网络进行交互的接口，类型为 shim.ChaincodeStubInterface。
// args：用于创建复合主键的键值列表，一个字符串数组。
func Hello(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 调用 utils.WriteLedger 方法将一个包含 "msg": "hello" 的 map 对象写入到账本中，用于测试智能合约中写入账本的功能是否正常
	err := utils.WriteLedger(map[string]interface{}{"msg": "hello"}, stub, "hello", []string{})
	if err != nil {
		return shim.Error(err.Error())
	}
	// 使用 shim.Success 方法返回一个包含 "hello world" 字符串的字节数组，作为函数执行成功的响应结果
	return shim.Success([]byte("hello world"))
}
