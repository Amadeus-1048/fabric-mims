import request from '@/utils/request'

// 新建药品订单(医生)
export function createDrugOrder(data) {
    return request({
        url: '/createDrugOrder',
        method: 'post',
        data
    })
}

// 获取病历信息(空json{}可以查询所有，指定patient可以查询指定患者的所有病历)
export function queryDrugOrderList(data) {
    return request({
        url: '/queryDrugOrderList',
        method: 'post',
        data
    })
}
