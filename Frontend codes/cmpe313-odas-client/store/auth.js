export const state = () => ({
  user: undefined,
})

export const mutations = {
  login(state, _user) {
    state.user = _user
  },
  logout(state) {
    state.user = undefined
  },
}
