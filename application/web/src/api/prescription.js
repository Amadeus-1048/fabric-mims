import request from '@/utils/request'

// 新建病历(医生)
export function createPrescription(data) {
    return request({
        url: '/createPrescription',
        method: 'post',
        data
    })
}

// 获取病历信息(空json{}可以查询所有，指定patient可以查询指定患者的所有病历)
export function queryPrescriptionList(data) {
    return request({
        url: '/queryPrescription',
        method: 'post',
        data
    })
}
