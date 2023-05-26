package cron

import (
	"log"

	"github.com/robfig/cron/v3"
)

const spec = "0 0 0 * * ?" // 基于 Quartz Cron 表达式的时间规则，表示每天 0 点执行一次
//const spec = "*/10 * * * * ?" //10秒执行一次，用于测试

// Init 初始化定时任务
func Init() {
	// 创建一个支持到秒级别的定时任务管理器  使用 cron.New 方法进行初始化
	c := cron.New(cron.WithSeconds()) //支持到秒级别
	_, err := c.AddFunc(spec, GoRun)
	if err != nil {
		log.Printf("定时任务开启失败 %s", err)
	}
	c.Start() // 启动定时任务管理器
	log.Printf("定时任务已开启")
	select {} // 让程序保持运行状态，以便可以一直执行定时任务（阻塞进程）
}

func GoRun() {
	log.Printf("定时任务已启动")
	// 查询数据
	//resp, err := bc.ChannelQuery("queryPrescription", [][]byte{}) //调用智能合约
	//if err != nil {
	//	log.Printf("定时任务-querySellingList失败%s", err.Error())
	//	return
	//}

	// 反序列化json
	//var data []model.
	//if err = json.Unmarshal(bytes.NewBuffer(resp.Payload).Bytes(), &data); err != nil {
	//	log.Printf("定时任务-反序列化json失败%s", err.Error())
	//	return
	//}

	// for循环处理data
	//for _, v := range data {
	//	fmt.Println(v)
	//}
}
