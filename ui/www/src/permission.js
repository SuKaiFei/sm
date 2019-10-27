import router from './router'
import store from './store'
import NProgress from 'nprogress' // progress bar
import 'nprogress/nprogress.css' // progress bar style
import {getToken, setToken} from '@/utils/auth' // get token from cookie
NProgress.configure({showSpinner: false}) // NProgress Configuration

const whiteList = ['/login'] // no redirect whitelist

function guid() {
	/**
	 * @return {string}
	 */
	function S4() {
		return (((1 + Math.random()) * 0x10000) | 0).toString(16).substring(1)
	}

	return (S4() + S4() + "-" + S4() + "-" + S4() + "-" + S4() + "-" + S4() + S4() + S4());
}

router.beforeEach(async (to, from, next) => {
	/* 路由发生变化修改页面title */
	if (to.meta.title) {
		document.title = to.meta.title
	}

	// start progress bar
	NProgress.start()

	// determine whether the user has logged in
	let hasToken = getToken()
	if (!hasToken) {
		hasToken = guid()
		store.commit('user/SET_TOKEN', hasToken)
	}
	setToken(hasToken);

	if (to.path === '/login') {
		// if is logged in, redirect to the home page
		next({path: '/'})
		NProgress.done()
	} else {
		const email = store.getters.email
		if (email || email === -1) {
			next()
		} else {
			try {
				// get user info
				await store.dispatch('user/getInfo')
				next()
			} catch (error) {
				store.commit("user/SET_EMAIL", -1)
				next()
				NProgress.done()
			}
		}
	}
})

router.afterEach(() => {
	// finish progress bar
	NProgress.done()
})
