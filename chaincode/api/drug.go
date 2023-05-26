package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"time"
)

// CreateDrugOrder 创建药品订单
func CreateDrugOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 5 {
		return shim.Error("参数个数不满足")
	}
	drugName := args[0]       // 药品名
	drugAmount := args[1]     // 药品数量
	prescriptionID := args[2] // 处方id
	patientID := args[3]      // 患者id
	DrugStoreID := args[4]

	if drugName == "" || drugAmount == "" || prescriptionID == "" || patientID == "" || DrugStoreID == "" {
		return shim.Error("参数存在空值")
	}

	// 判断是否为药店操作
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountV2Key, []string{DrugStoreID})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("操作人权限验证失败%s", err))
	}

	record := &model.DrugOrder{
		ID:           stub.GetTxID()[:16],
		Name:         drugName,
		Amount:       drugAmount,
		Prescription: prescriptionID,
		Patient:      patientID,
		DrugStore:    DrugStoreID,
		Created:      time.Now().Format("2006-01-02 15:04:05"),
	}

	// 写入账本
	if err := utils.WriteLedger(record, stub, model.DrugKey, []string{record.Patient, record.DrugStore, record.ID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	recordByte, err := json.Marshal(record)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(recordByte)
}

// QueryDrugOrder 查询药品订单(可查询所有，也可根据所有人查询名下处方)
func QueryDrugOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var drugOrderList []model.DrugOrder
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.DrugKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var p model.DrugOrder
			err := json.Unmarshal(v, &p)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryDrugOrder-反序列化出错: %s", err))
			}
			drugOrderList = append(drugOrderList, p)
		}
	}
	drugOrderByte, err := json.Marshal(drugOrderList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryDrugOrder-序列化出错: %s", err))
	}
	return shim.Success(drugOrderByte)
}
