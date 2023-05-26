import request from '@/utils/request'

// 新建保险报销订单(医生)
export function createInsuranceCover(data) {
    return request({
        url: '/createInsuranceCover',
        method: 'post',
        data
    })
}

// 获取报销记录(空json{}可以查询所有，指定patient可以查询指定患者的所有记录)
export function queryInsuranceCoverList(data) {
    return request({
        url: '/queryInsuranceCoverList',
        method: 'post',
        data
    })
}

// 更新报销记录
export function updateInsuranceCoverList(data) {
    return request({
        url: '/updateInsuranceCover',
        method: 'post',
        data
    })
}