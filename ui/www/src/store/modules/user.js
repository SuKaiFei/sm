import {login, logout, getInfo} from '@/api/user'

const state = {
	token: '',
	name: '',
	email: '',
	avatar: ''
}

const mutations = {
	SET_NAME: (state, name) => {
		state.name = name
	},
	SET_TOKEN: (state, token) => {
		state.token = token
	},
	SET_EMAIL: (state, email) => {
		state.email = email
	},
	SET_AVATAR: (state, avatar) => {
		state.avatar = avatar
	}
}

const actions = {
	// get user info
	getInfo({commit, state}) {
		return new Promise((resolve, reject) => {
			getInfo(state.token).then(response => {
				if (response && response.code === 0) {
					const userInfo = response.data
					commit('SET_NAME', userInfo.nickname)
					commit('SET_EMAIL', userInfo.email)
					commit('SET_AVATAR', userInfo.avatar)
					resolve(userInfo)
				} else {
					reject('Verification failed, please Login again.')
				}
			}).catch(error => {
				reject(error)
			})
		})
	},
}

export default {
	namespaced: true,
	state: state,
	mutations: mutations,
	actions: actions
}

