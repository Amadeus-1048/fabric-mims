package model

// Account 账户，虚拟管理员和若干业主账号
type Account struct {
	AccountId string  `json:"accountId"` //账号ID
	UserName  string  `json:"userName"`  //账号名
	Balance   float64 `json:"balance"`   //余额
}

// objectType  对象类型，用于创建复合主键
const (
	AccountKey = "account-key"

	AccountV2Key    = "account-v2-key"
	PrescriptionKey = "prescription-key"
	PatientKey      = "patient-key"
	InsuranceKey    = "insurance-key"
	DrugKey         = "drug-key"
)

// --------------------------------------------------------------------

// AccountV2 账号
type AccountV2 struct {
	AccountId   string `json:"account_id"`   // 账号ID
	AccountName string `json:"account_name"` // 账号名
}

// Hospital 医院
type Hospital struct {
	ID      string          `json:"id"`
	Name    string          `json:"name"`
	Admins  []HospitalAdmin `json:"admins"`
	Doctors []Doctor        `json:"doctors"`
}

// HospitalAdmin 医院管理员
type HospitalAdmin struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Doctor 医生
type Doctor struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Prescription 医疗处方
type Prescription struct {
	ID        string `json:"id"`        // 医疗处方ID
	Patient   string `json:"patient"`   // 患者ID
	Diagnosis string `json:"diagnosis"` // 诊断结果
	Drug      []Drug `json:"drug"`      // 药品列表及用量
	Doctor    string `json:"doctor"`    // 开方医师 AccountV2Id
	Hospital  string `json:"hospital"`  // 医院 ID
	Created   string `json:"created"`   // 创建时间
	Comment   string `json:"comment"`   // 备注
}

// Patient 患者
type Patient struct {
	ID     string `json:"id"`     // 患者 AccountV2Id
	Name   string `json:"name"`   // 患者姓名
	Age    int    `json:"age"`    // 患者年龄
	Gender string `json:"gender"` // 患者性别
}

// Drug 药品
type Drug struct {
	//ID      string `json:"id"`
	Name   string `json:"Name"`   // 药品名
	Amount string `json:"amount"` // 药品数量
}

// DrugOrder 药品订单
type DrugOrder struct {
	ID           string `json:"id"`           // 订单ID
	Name         string `json:"Name"`         // 药品名
	Amount       string `json:"amount"`       // 药品数量
	Prescription string `json:"prescription"` // 处方ID
	Patient      string `json:"patient"`      // 患者ID
	DrugStore    string `json:"drug_store"`   // 药店id
	Created      string `json:"created"`      // 创建时间
}

// DrugStore 药店
type DrugStore struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// Insurance 保险机构
type Insurance struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// InsuranceCover 保险报销订单
type InsuranceCover struct {
	ID           string `json:"id"`           // 订单ID
	Prescription string `json:"prescription"` // 处方ID
	Patient      string `json:"patient"`      // 患者ID
	Status       string `json:"status"`       // 订单状态
	Created      string `json:"created"`      // 创建时间
}

// InsuranceStatusConstant 保险状态
var InsuranceStatusConstant = func() map[string]string {
	return map[string]string{
		"processing": "处理中", // 患者发起保险报销申请，等待保险公司确认报销
		"cancelled":  "已取消", // 患者在保险公司确认报销之前取消保险报销申请
		"refused":    "已拒绝", // 保险公司拒绝确认报销
		"approved":   "已通过", // 保险公司确认报销，保险报销完成
	}
}

// DrugStatusConstant 药品状态
//var DrugStatusConstant = func() map[string]string {
//	return map[string]string{
//		"processing": "处理中", //
//		"done":       "完成",   //
//	}
//}
