<template>
    <div>
        <div v-if="this.$store.getters.name !== ''" class="sidebar-group">
            <b-card
                    no-body
                    :img-src="this.avatar"
                    img-width="200px"
                    img-alt="Image"
                    img-top>
                <b-card-body>
                    <b-card-title>{{this.name}}</b-card-title>
                    <b-card-sub-title class="mb-2">精通卖萌，求铲屎官</b-card-sub-title>
                    <b-card-text>
                        头圆、颜面部短，前肢五指，后肢四趾，趾端具锐利而弯曲的爪，爪能伸缩。夜行性。
                    </b-card-text>
                </b-card-body>

                <b-card-body>
                    <a href="#" class="card-link">发布帖子</a>
                    <a href="#" class="card-link">我的帖子</a>
                </b-card-body>
            </b-card>
        </div>
        <div v-if="this.$store.getters.name === ''" class="sidebar-group box_white">
            <div class="top">
                <h3 class="title"><i class="glyphicon glyphicon-user"></i>&nbsp;<a href="/users" target="_blank"
                                                                                   title="点击更多">会员登录</a></h3>
            </div>
            <div class="sb-content">
                <b-form @submit.stop.prevent="onSubmit()" :novalidate="true">
                    <b-form-group
                            label-cols-lg="3"
                            label="邮箱:"
                            label-for="input-1"
                    >
                        <b-form-input
                                id="input-1"
                                name="username"
                                :value="form.username"
                                v-model="form.username"
                                type="email"
                                :state="this.$v.form.username.$dirty ? !this.$v.form.username.$error : null"
                                aria-describedby="input-1-live-feedback"
                        ></b-form-input>
                        <b-form-invalid-feedback id="input-1-live-feedback">
                            邮箱格式不正确.
                        </b-form-invalid-feedback>
                    </b-form-group>

                    <b-form-group
                            label-cols-lg="3" label="密码:" label-for="input-2">
                        <b-form-input
                                id="input-2"
                                name="password"
                                :value="form.password"
                                type="password"
                                v-model="form.password"
                                :state="this.$v.form.password.$dirty ? !this.$v.form.password.$error : null"
                                aria-describedby="input-2-live-feedback"
                        ></b-form-input>
                        <b-form-invalid-feedback id="input-2-live-feedback">
                            密码长度大于或等于6.
                        </b-form-invalid-feedback>
                    </b-form-group>

                    <b-button block type="submit" variant="primary">登录</b-button>
                    <b-button block variant="outline-primary" @click="signInWithShow">第三方登录
                    </b-button>
                </b-form>
            </div>
        </div>
        <div class="sidebar-group box_white">
            <div class="top">
                <h3 class="title"><i class="glyphicon glyphicon-user"></i>&nbsp;<a href="/users" target="_blank"
                                                                                   title="点击更多">新会员</a></h3>
            </div>
            <div class="sb-content user-list" style="clear: both;" data-limit="9">
                <ul style="list-style: none;margin: 0 0 0 0;padding: 0 0 0 0;">
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                    <li class="user">
                        <a>
                            <div class="avatar">
                                <img src="https://placekitten.com/g/400/450" alt="Image">
                            </div>
                            <div>
                                <span class="nick-name">小花猫</span>
                            </div>
                        </a>
                    </li>
                </ul>
            </div>
        </div>
        <div class="sidebar-group box_white">
            <div class="top">
                <h3 class="title">
                    <i class="glyphicon glyphicon-user"></i>&nbsp;
                    <a href="#" target="_blank" title="点击更多">加入我们</a></h3>
            </div>
            <div class="sb-content">
                QQ群：331497557
            </div>
        </div>
        <div>
            <b-modal id="sign-in-with" hide-footer ref="sign-in-with">
                <template v-slot:modal-title>
                    第三方登录
                </template>
                <div class="d-block text-center">
                    <b-button class="mt-3" variant="success" block>
                        <svg-icon icon-class="wechat"/>&nbsp;微信
                    </b-button>
                    <b-button class="mt-2" variant="dark" block>
                        <svg-icon icon-class="github"/>&nbsp;GitHub
                    </b-button>
                </div>
            </b-modal>
        </div>
    </div>
</template>

<script>
	import {validationMixin} from 'vuelidate'
	import {required, minLength, email} from 'vuelidate/lib/validators'
	import {login} from '@/api/user'
	import {mapGetters} from 'vuex'

	export default {
		mixins: [validationMixin],
		name: 'Sidebar',
		data() {
			return {
				form: {
					username: null,
					password: null,
				}
			}
		},
		computed: {
			...mapGetters({
				name: 'name',
				email: 'email',
				avatar: 'avatar',
			})
		},
		validations: {
			form: {
				username: {
					required,
					email
				},
				password: {
					required,
					minLength: minLength(6)
				}
			}
		},
		methods: {
			signInWithShow() {
				this.$bvModal.show('sign-in-with')
			},
			signInWithHide() {
				this.$bvModal.hide('sign-in-with')
			},
			onSubmit() {
				this.$v.form.$touch()
				if (this.$v.form.$anyError) {
					return
				}
				login(this.form).then(res => {
					if (res && res.code === 0) {
						const userInfo = res.data;
						this.$store.commit('user/SET_NAME', userInfo.nickname)
						this.$store.commit('user/SET_EMAIL', userInfo.email)
						this.$store.commit('user/SET_AVATAR', userInfo.avatar)
					}
				})
				return false
			}
		}
	}
</script>

<style type="scss" scoped>
    .sidebar-group {
        margin-bottom: 12px;
    }

    .box_white {
        background: #FFF;
        overflow: hidden;
        border-radius: 0.25rem;
        border: 1px solid rgba(0, 0, 0, 0.125);
    }

    .user-list {
        margin: 0 0 0 0 !important;
    }

    .user-list .user {
        text-align: center;
        margin-bottom: 8px;
        width: 33%;
        float: left;
    }

    .user-list .user .avatar img {
        border-radius: 20%;
        width: 90%;
    }

    .user-list .card-body {
        padding: 0;
    }

    .sb-content {
        margin: auto 5px 5px 5px;
    }

    .sidebar-group .top {
        line-height: 24px;
        border-bottom: solid 1px #007bff;
        font-size: 14px;
        font-weight: 700;
        margin-bottom: 4px;
        margin-top: 10px;
        padding-left: 10px;
    }
</style>
