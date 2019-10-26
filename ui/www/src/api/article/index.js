import request from '@/utils/request'

export function create(data) {
	return request({
		url: '/article/add',
		method: 'post',
		data
	})
}

export function info(data) {
	return request({
		url: '/article/info',
		method: 'get',
		params: data
	})
}

export function list(data) {
	return request({
		url: '/articles',
		method: 'get',
		params: data
	})
}

export function my(data) {
	return request({
		url: '/my/articles',
		method: 'get',
		params: data
	})
}
