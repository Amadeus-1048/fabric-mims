package api

import (
	"chaincode/model"
	"chaincode/pkg/utils"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// CreatePrescription 创建处方(医生)
func CreatePrescription(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	// 验证参数
	if len(args) != 7 {
		return shim.Error("参数个数不满足")
	}
	doctorID := args[0]   // 医生id
	patientID := args[1]  // 患者id
	diagnosis := args[2]  // 诊断结果
	drugName := args[3]   // 药品名
	drugAmount := args[4] // 药品数量
	hospitalID := args[5] // 医院ID
	comment := args[6]    // 备注
	if doctorID == "" || patientID == "" || diagnosis == "" || drugName == "" || drugAmount == "" || hospitalID == "" {
		return shim.Error("参数存在空值")
	}

	// 参数数据格式转换
	var drugs []model.Drug
	drugNames := strings.Split(drugName, ",")
	drugAmounts := strings.Split(drugAmount, ",")
	for i, v := range drugNames {
		drug := model.Drug{
			Name:   v,
			Amount: drugAmounts[i],
		}
		drugs = append(drugs, drug)
	}

	// 判断是否为医生操作
	resultsAccount, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountV2Key, []string{doctorID})
	if err != nil || len(resultsAccount) != 1 {
		return shim.Error(fmt.Sprintf("操作人权限验证失败%s", err))
	}
	var account model.AccountV2
	if err = json.Unmarshal(resultsAccount[0], &account); err != nil {
		return shim.Error(fmt.Sprintf("查询操作人信息-反序列化出错: %s", err))
	}
	if account.AccountName != "医生" {
		return shim.Error(fmt.Sprintf("操作人权限不足%s", err))
	}

	// 判断患者是否存在
	resultsPatient, err := utils.GetStateByPartialCompositeKeys(stub, model.AccountV2Key, []string{patientID})
	if err != nil || len(resultsPatient) != 1 {
		return shim.Error(fmt.Sprintf("患者信息验证失败%s", err))
	}

	prescription := &model.Prescription{
		ID:        stub.GetTxID()[:16],
		Patient:   patientID,
		Diagnosis: diagnosis,
		Drug:      drugs,
		Doctor:    doctorID,
		Hospital:  hospitalID,
		Created:   time.Now().Format("2006-01-02 15:04:05"),
		Comment:   comment,
	}
	// 写入账本
	if err := utils.WriteLedger(prescription, stub, model.PrescriptionKey, []string{prescription.Patient, prescription.ID}); err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	//将成功创建的信息返回
	prescriptionByte, err := json.Marshal(prescription)
	if err != nil {
		return shim.Error(fmt.Sprintf("序列化成功创建的信息出错: %s", err))
	}
	// 成功返回
	return shim.Success(prescriptionByte)
}

// QueryPrescription 查询处方(可查询所有，也可根据所有人查询名下处方)
func QueryPrescription(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var prescriptionList []model.Prescription
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.PrescriptionKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var p model.Prescription
			err := json.Unmarshal(v, &p)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryPrescription-反序列化出错: %s", err))
			}
			prescriptionList = append(prescriptionList, p)
		}
	}
	prescriptionByte, err := json.Marshal(prescriptionList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryPrescription-序列化出错: %s", err))
	}
	return shim.Success(prescriptionByte)
}

// QueryPatient 查询患者
func QueryPatient(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var patientList []model.Patient
	results, err := utils.GetStateByPartialCompositeKeys2(stub, model.PatientKey, args)
	if err != nil {
		return shim.Error(fmt.Sprintf("%s", err))
	}
	for _, v := range results {
		if v != nil {
			var p model.Patient
			err := json.Unmarshal(v, &p)
			if err != nil {
				return shim.Error(fmt.Sprintf("QueryPatient-反序列化出错: %s", err))
			}
			patientList = append(patientList, p)
		}
	}
	patientByte, err := json.Marshal(patientList)
	if err != nil {
		return shim.Error(fmt.Sprintf("QueryPatient-序列化出错: %s", err))
	}
	return shim.Success(patientByte)
}
