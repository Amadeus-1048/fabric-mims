import request from '@/utils/request'

// 获取登录界面角色选择列表
export function queryAccountList() {
    return request({
        url: '/queryAccountV2List',
        method: 'post'
    })
}

// 登录
export function login(data) {
    return request({
        url: '/queryAccountV2List',
        method: 'post',
        data
    })
}

// 创建角色
export function createAccount(data) {
    return request({
        url: '/createAccountV2',
        method: 'post',
        data
    })
}