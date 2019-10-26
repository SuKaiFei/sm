<template>
    <div>
        <b-breadcrumb :items="items"></b-breadcrumb>
        <b-card :title="form.title">
            <div id="article_id" v-html='form.text'></div>
        </b-card>
    </div>
</template>

<script>
	import {info} from '@/api/article'

	export default {
		name: "Article",

		data() {
			return {
				tags: [],
				form: {
					title: '',
					text: '',
					tag: null,
				},
				items:
					[
						{
							text: '首页',
							href: '#'
						},
						{
							text: '帖子',
							href: '#',
						},
						{
							text: '详情',
							href: '#',
							active: true
						}
					]
			}
		},
		methods: {
			getInfo(aid) {
				this.$router.push({path: '/article/info', query: {aid: aid}})
			},
			getList() {
				info({id: this.$route.query.aid}).then((res) => {
					if (res && res.code === 0) {
						this.form = res.data
					} else {
						this.$bvToast.toast(`获取帖子详情失败，${res.message}`, {
							title: '操作提示',
							variant: 'danger',
							solid: true
						})
					}
				})
			}
		},
		mounted() {
			this.getList()
		}
	}
</script>

<style>
    #article_id img {
        max-width: 100%;
    }
</style>
