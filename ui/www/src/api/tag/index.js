import request from '@/utils/request'

export function tags(data) {
	return request({
		url: '/tags',
		method: 'get',
		data
	})
}
