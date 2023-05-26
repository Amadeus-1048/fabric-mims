const getters = {
  sidebar: state => state.app.sidebar,
  device: state => state.app.device,
  token: state => state.account.token,
  account_id: state => state.account.account_id,
  account_name: state => state.account.account_name,
  balance: state => state.account.balance,
  roles: state => state.account.roles,
  permission_routes: state => state.permission.routes
}
export default getters
