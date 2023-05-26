package main

import (
	"chaincode/api"
	"chaincode/model"
	"chaincode/pkg/utils"
	"fmt"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type BlockChainMedicalInfoManageSystem struct {
}

// Init 链码部署到链上并进行初始化时会执行该方法
func (t *BlockChainMedicalInfoManageSystem) Init(stub shim.ChaincodeStubInterface) pb.Response { // stub 是智能合约中的一个对象，用于与区块链网络进行交互
	fmt.Println("链码初始化")
	//初始化默认数据

	var accountV2Ids = [7]string{"0feceb66ffc1", "1feceb66ffc1", "2b86b273ff31", "34735e3a261e", "4e17408561be", "5b227771d4dd", "6f2d121de37b"}
	var userNameV2s = [7]string{"管理员", "医生", "①号病人", "②号病人", "③号病人", "药店", "保险机构"}
	//初始化账号数据
	for i, val := range accountV2Ids {
		account := &model.AccountV2{
			AccountId:   val,
			AccountName: userNameV2s[i],
		}
		// 写入账本
		if err := utils.WriteLedger(account, stub, model.AccountV2Key, []string{val}); err != nil {
			return shim.Error(fmt.Sprintf("%s", err))
		}
	}

	return shim.Success(nil)
}

// Invoke 实现Invoke接口调用智能合约
func (t *BlockChainMedicalInfoManageSystem) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// 获取调用该智能合约函数时传入的参数和函数名
	// 这些信息可以用于智能合约内部的逻辑处理，例如根据不同的函数名和参数执行不同的操作或者检查参数的有效性等
	funcName, args := stub.GetFunctionAndParameters() // 返回两个值：funcName 表示函数名，args 表示函数参数列表，以字符串数组的形式返回
	switch funcName {
	case "hello":
		return api.Hello(stub, args)
	// api v2
	case "createAccountV2":
		return api.CreateAccountV2(stub, args)
	case "queryAccountV2List":
		return api.QueryAccountV2List(stub, args)
	case "createPrescription":
		return api.CreatePrescription(stub, args)
	case "queryPrescription":
		return api.QueryPrescription(stub, args)
	case "createInsuranceCover":
		return api.CreateInsuranceCover(stub, args)
	case "queryInsuranceCover":
		return api.QueryInsuranceCover(stub, args)
	case "updateInsuranceCover":
		return api.UpdateInsuranceCover(stub, args)
	case "deleteInsuranceCover":
		return api.DeleteInsuranceCover(stub, args)
	case "createDrugOrder":
		return api.CreateDrugOrder(stub, args)
	case "queryDrugOrder":
		return api.QueryDrugOrder(stub, args)
	default:
		return shim.Error(fmt.Sprintf("没有该功能: %s", funcName))
	}
}

func main() {
	timeLocal, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		panic(err)
	}
	time.Local = timeLocal
	err = shim.Start(new(BlockChainMedicalInfoManageSystem)) // 启动智能合约
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
