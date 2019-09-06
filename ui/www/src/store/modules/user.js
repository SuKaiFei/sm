const state = {
  name: '',
  avatar: ''
}

const mutations = {
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  }
}

const actions = {
  // get user info
  getInfo({ commit, state }) {
    return new Promise((resolve, reject) => {
    })
  },
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}

