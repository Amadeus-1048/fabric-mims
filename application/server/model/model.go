package model

// ----------------------         Account 用户   ----------------------------------

type AccountIdBody struct {
	AccountId string `json:"account_id"`
}

type AccountRequestBody struct {
	Args []AccountIdBody `json:"args"`
}

type CreateAccountBody struct {
	AccountName string `json:"account_name"`
	Operator    string `json:"operator"`
}

// ----------------------         Prescription 病历   ----------------------------------

type PrescriptionRequestBody struct {
	Doctor    string `json:"doctor"`    // 医生ID
	Patient   string `json:"patient"`   // 患者Id
	Diagnosis string `json:"diagnosis"` // 诊断结果
	//Drug      []Drug `json:"drug"`      // 药品列表及用量
	DrugName   string `json:"drug_name"`   // 药品名
	DrugAmount string `json:"drug_amount"` // 药品用量
	Hospital   string `json:"hospital"`    // 医院 ID
	Comment    string `json:"comment"`     // 备注
}

type PrescriptionQueryRequestBody struct {
	Patient string `json:"patient"` // 患者AccountId
}

// ----------------------         DrugOrder 药品订单   ----------------------------------

type DrugOrderRequestBody struct {
	//Drug      []Drug `json:"drug"`      // 药品列表及用量
	DrugName     string `json:"drug_name"`    // 药品名
	DrugAmount   string `json:"drug_amount"`  // 药品用量
	Prescription string `json:"prescription"` // 处方ID
	Patient      string `json:"patient"`      // 患者Id
	DrugStore    string `json:"drug_store"`   // 药店Id
}

type DrugOrderQueryRequestBody struct {
	Patient   string `json:"patient"` // 患者AccountId
	DrugStore string `json:"drug_store"`
}

// ----------------------         InsuranceCover 保险报销   ----------------------------------

type InsuranceCoverRequestBody struct {
	Prescription string `json:"prescription"` // 处方ID
	Patient      string `json:"patient"`      // 患者Id
	Status       string `json:"status"`       // 订单状态
}

type InsuranceCoverQueryRequestBody struct {
	Patient        string `json:"patient"`         // 患者Id
	InsuranceCover string `json:"insurance_cover"` // 报销订单ID
}

type UpdateInsuranceCoverRequestBody struct {
	InsuranceCover string `json:"insurance_cover"` // 报销订单ID
	Patient        string `json:"patient"`         // 病人ID
	InsuranceID    string `json:"insurance_id"`    // 保险机构ID
	Status         string `json:"status"`          // 订单状态
}
