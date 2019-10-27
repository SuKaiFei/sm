<template>
    <b-row>
        <b-col md="8" lg="9" sm="7">
            <b-list-group>
                <b-list-group-item v-for="article in articles" :key="article.id">
                    <b-card :img-src="article.avatar" img-alt="Card image" img-left
                            class="mb-3 card-group">
                        <div class="card-title">
                            <!--                    <a href="#">网站成立喽</a>-->
                            <b-button variant="link" @click="getInfo(article.id)">{{article.title}}</b-button>
                        </div>
                        <ul class="card-info">
                            <li>
                                <a href="#">
                                    <b-badge>{{article.tag}}</b-badge>
                                </a>
                            </li>
                            &nbsp; • &nbsp;
                            <li>
                                <a href="#">{{article.author}}</a>
                            </li>
                            &nbsp; • &nbsp;
                            <li>
                                {{article.create_time}}
                            </li>
                            &nbsp; • &nbsp;
                            <li>
                                <b-button pill variant="outline-primary" size="sm">
                                    <svg-icon icon-class="like"/>&nbsp;
                                    <b-badge variant="primary">{{article.pv}}</b-badge>
                                </b-button>
                            </li>
                        </ul>
                    </b-card>
                </b-list-group-item>
            </b-list-group>
            <b-pagination per-page="10" v-model="page.num" @input="getList" :total-rows="page.total"
                          align="fill"></b-pagination>
        </b-col>
        <b-col md="4" lg="3" sm="5">
            <sidebar></sidebar>
        </b-col>
    </b-row>
</template>

<script>
	import {list} from '@/api/article'
	import Sidebar from '@/views/sidebar/index'

	export default {
		name: "Index",
		components: {
			Sidebar
		},
		data() {
			return {
				articles: [],
				page: {
					num: 1,
					total: 0,
				},
				form: {
					pn: '',
				},
			}
		},
		watch: {
			'$route': function (route) {
				this.tag = route.query.tag
				this.items[1].text = this.tag
				this.getList()
			},
		},
		methods: {
			getInfo(aid) {
				this.$router.push({path: '/article/info', query: {aid: aid}})
			},
			getList() {
				this.form.pn = this.page.num
				list(this.form).then((res) => {
					if (res && res.code === 0) {
						this.articles = res.data.result
						this.page.total = res.data.page.total
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
		},

	}
</script>

<style type="scss" scoped>
    .list-group-item {
        padding: 0 0;
        border: 0;
    }

    .list-group-item .card-body {
        padding: 1rem;
    }

    .list-group-item .card-img-left {
        width: 64px;
        height: 64px;
        margin: auto auto auto 10px;
    }

    .card-group {
        width: 100%;
        height: 100%;
        margin-bottom: 0 !important;
    }

    .card-title {
        margin-bottom: 0 !important;
    }

    .card-title button {
        margin: 0 0;
        padding: 0 0;
    }

    .card-info {
        margin: 0 0;
        padding-left: 0;
    }

    .card-info ul {
        display: inline;
        white-space: nowrap;
    }

    .card-info li {
        display: inline-block;
        white-space: nowrap;
    }

    .card-info li:first-child {
        padding: 0 0;
    }
</style>
