import axios from 'axios'
axios.defaults.withCredentials = true

// create an axios instance
const service = axios.create({
	baseURL: process.env.VUE_APP_BASE_API, // url = base url + request url
	withCredentials: true, // send cookies when cross-domain requests
	crossDomain: true,
	timeout: 5000 // request timeout
})

service.defaults.withCredentials = true
// request interceptor
service.interceptors.request.use(
	config => {
		// do something before request is sent
		return config
	},
	error => {
		// do something with request error
		console.log(error) // for debug
		return Promise.reject(error)
	}
)

// response interceptor
service.interceptors.response.use(
	/**
	 * If you want to get http information such as headers or status
	 * Please return  response => response
	 */

	/**
	 * Determine the request status by custom code
	 * Here is just an example
	 * You can also judge the status by HTTP Status Code
	 */
	response => {
		return response.data
	},
	error => {
		console.log('err' + error) // for debug
		return Promise.reject(error)
	}
)

export default service
