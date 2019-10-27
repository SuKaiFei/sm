<template>
    <div>
        <b-breadcrumb :items="items"></b-breadcrumb>
        <b-card style="padding: 1.125rem 1.5rem;">
            <h1>{{form.title}}</h1>
            <b-nav pills>
                <b-img rounded="circle" left :src="form.avatar"
                       width="40" height="40"></b-img>
                <b-nav-item disabled style="font-weight: 600;">{{form.author}}</b-nav-item>
                <b-nav-item disabled style="font-weight: 600;">{{form.tag}}</b-nav-item>
                <b-nav-item disabled style="font-weight: 600;">{{form.create_time}}</b-nav-item>
            </b-nav>
            <div style="background: #ebecef;margin-top: 1%;">
                <b-nav pills nobody>
                    <b-nav-item>
                        <b-button pill variant="outline-primary" size="sm">
                            <svg-icon icon-class="like"/>&nbsp;
                            <b-badge variant="primary">12</b-badge>
                        </b-button>
                    </b-nav-item>
                    <b-nav-item>
                        <b-button pill variant="outline-primary" size="sm">
                            <svg-icon icon-class="wechat"/>
                        </b-button>
                    </b-nav-item>
                </b-nav>
                <span style="float: right;margin-top: -35px;margin-right: 5%;">
                    <svg-icon icon-class="eye-open"/>&nbsp;
                    <b-badge>26.6K</b-badge>
                </span>
            </div>
            <div id="article_id" class="markdown-body" v-html='form.text'></div>
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
			getArticle() {
				info({id: this.$route.query.aid}).then((res) => {
					if (res && res.code === 0) {
						this.form = res.data
						document.title = this.form.tag + "-" + this.form.title + "-杂货铺社区"
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
			this.getArticle()
		}
	}
</script>

<style>
    #article_id img {
        max-width: 100%;
        margin-top: 2%;
    }

    #article_id {
        margin-top: 2%;
    }
</style>
