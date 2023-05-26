package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// QueryAccountList 查询账户列表
// stub：智能合约与区块链网络进行交互的接口，类型为 shim.ChaincodeStubInterface。
// args：用于创建复合主键的键值列表，一个字符串数组。
func QueryAccountList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var accountList []model.Account
	// 根据 model.AccountKey 和 args 的前缀从账本中获取所有的符合条件的账户信息
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	// 遍历账户信息，将其转换为 model.Account 对象添加到 accountList 列表中
	for _, v := range results {
		if v != nil {
			var account model.Account
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			accountList = append(accountList, account)
		}
	}
	// 将 accountList 转换为字节数组，然后使用 shim.Success 方法将其作为查询结果返回给调用方
	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(accountListByte)
}

// QueryAccountV2List 查询账户列表
// stub：智能合约与区块链网络进行交互的接口，类型为 shim.ChaincodeStubInterface。
// args：用于创建复合主键的键值列表，一个字符串数组。
func QueryAccountV2List(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var accountList []model.AccountV2
	// 根据 model.AccountKey 和 args 的前缀从账本中获取所有的符合条件的账户信息
	results, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountV2Key, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	// 遍历账户信息，将其转换为 model.Account 对象添加到 accountList 列表中
	for _, v := range results {
		if v != nil {
			var account model.AccountV2
			err := json.Unmarshal(v, &account)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryAccountList-反序列化出错: %s", err))
			}
			accountList = append(accountList, account)
		}
	}
	// 将 accountList 转换为字节数组，然后使用 shim.Success 方法将其作为查询结果返回给调用方
	accountListByte, err := json.Marshal(accountList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryAccountList-序列化出错: %s", err))
	}
	return shim.Success(accountListByte)
}

// CreateAccountV2 创建角色
func CreateAccountV2(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 2 {
		return shim.Error("参数个数不满足")
	}
	userName := args[0] // 用户名
	operator := args[1] // 操作人ID

	if operator == "" || userName == "" {
		return shim.Error("参数存在空值")
	}

	// 判断是否为管理员操作
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountV2Key, []string{operator})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("操作人权限验证失败%s", err))
	}
	var account model.AccountV2
	if err = json.Unmarshal(resultsAccount[0], &account); err != nil {
		return shim.Error(fmt.Sprintf("查询操作人信息-反序列化出错: %s", err))
	}
	if account.AccountName != "管理员" {
		return shim.Error(fmt.Sprintf("操作人权限不足%s", err))
	}

	newAccount := &model.AccountV2{
		AccountId:   stub.GetTxID()[:12],
		AccountName: userName,
	}
	// 写入账本
	if err := utils.WriteLedger(newAccount, stub, model.AccountV2Key, []string{newAccount.AccountId}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	prescriptionByte, err := json.Marshal(newAccount)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(prescriptionByte)
}
