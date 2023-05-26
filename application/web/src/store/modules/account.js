import {
  login
} from '@/api/accountV2'
import {
  getToken,
  setToken,
  removeToken
} from '@/utils/auth'
import {
  resetRouter
} from '@/router'

const getDefaultState = () => {
  return {
    token: getToken(),
    account_id: '',
    account_name: '',
    roles: []
  }
}

const state = getDefaultState()

const mutations = {
  RESET_STATE: (state) => {
    Object.assign(state, getDefaultState())
  },
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_ACCOUNTID: (state, account_id) => {
    state.account_id = account_id
  },
  SET_USERNAME: (state, account_name) => {
    state.account_name = account_name
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  }
}

const actions = {
  login({
    commit
  }, account_id) {
    return new Promise((resolve, reject) => {
      login({
        args: [{
          account_id: account_id
        }]
      }).then(response => {
        commit('SET_TOKEN', response[0].account_id)
        setToken(response[0].account_id)
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
  // get user info
  getInfo({
    commit,
    state
  }) {
    return new Promise((resolve, reject) => {
      login({
        args: [{
          account_id: state.token
        }]
      }).then(response => {
        var roles
        if (/管理员/.test(response[0].account_name)) {
          roles = ['admin']
        } else if (/医生/.test(response[0].account_name)){
          roles = ['doctor']
        } else if (/病人/.test(response[0].account_name)){
          roles = ['patient']
        } else if (/药店/.test(response[0].account_name)){
          roles = ['drugstore']
        } else if (/保险机构/.test(response[0].account_name)){
          roles = ['insurance']
        }
        commit('SET_ROLES', roles)
        commit('SET_ACCOUNTID', response[0].account_id)
        commit('SET_USERNAME', response[0].account_name)
        resolve(roles)
      }).catch(error => {
        reject(error)
      })
    })
  },
  logout({
    commit
  }) {
    return new Promise(resolve => {
      removeToken()
      resetRouter()
      commit('RESET_STATE')
      resolve()
    })
  },

  resetToken({
    commit
  }) {
    return new Promise(resolve => {
      removeToken()
      commit('RESET_STATE')
      resolve()
    })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
