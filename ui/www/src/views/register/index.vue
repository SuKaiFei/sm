<template>
    <div>
        <b-breadcrumb :items="items"></b-breadcrumb>
        <b-card>
            <b-form @submit.stop.prevent="onSubmit()" :novalidate="true">
                <b-form-group
                        label-cols-lg="3"
                        label="邮箱:"
                        label-for="input-email">
                    <b-form-input
                            id="input-email"
                            v-model="form.email"
                            type="email"
                            placeholder="请输入邮箱"
                            :state="this.$v.form.email.$dirty ? !this.$v.form.email.$error : null"
                            aria-describedby="input-email-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="input-email-live-feedback">
                        邮箱格式不正确.
                    </b-form-invalid-feedback>
                </b-form-group>
                <b-form-group
                        label-cols-lg="3"
                        label="头像:"
                        label-for="input-avatar">
                    <b-img v-if="avatar_url!==''" thumbnail style="margin: 0 auto;display: flex;" width="75px"
                           :src="avatar_url" rounded
                           alt="头像"></b-img>
                    <b-form-file v-model="form.avatar" placeholder="请上传头像" @change="onFileChange"
                                 :state="this.$v.form.avatar.$dirty ? !this.$v.form.avatar.$error : null"
                                 accept="image/*" browse-text="选择"></b-form-file>
                </b-form-group>
                <b-form-group
                        label-cols-lg="3"
                        label="昵称:"
                        label-for="input-nickname">
                    <b-form-input
                            v-model="form.nickname"
                            placeholder="请输入昵称"
                            :state="this.$v.form.nickname.$dirty ? !this.$v.form.nickname.$error : null"
                            aria-describedby="input-nickname-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="input-nickname-live-feedback">
                        昵称格式不正确.
                    </b-form-invalid-feedback>
                </b-form-group>
                <b-form-group
                        label-cols-lg="3"
                        label="性别:"
                        label-for="select-sex">
                    <b-form-select id="select-sex" v-model="form.sex" :options="sexType"
                                   :state="this.$v.form.sex.$dirty ? !this.$v.form.sex.$error : null"></b-form-select>
                    <b-form-invalid-feedback id="input-sex-live-feedback">
                        请选择性别.
                    </b-form-invalid-feedback>
                </b-form-group>

                <b-form-group
                        label-cols-lg="3" label="密码:" label-for="input-password">
                    <b-form-input
                            id="input-password"
                            type="password"
                            placeholder="请输入密码"
                            v-model="form.password"
                            :state="this.$v.form.password.$dirty ? !this.$v.form.password.$error : null"
                            aria-describedby="input-password-live-feedback"></b-form-input>
                    <b-form-invalid-feedback id="input-password-live-feedback">
                        密码长度大于或等于6.
                    </b-form-invalid-feedback>
                </b-form-group>
                <b-form-group
                        label-cols-lg="3" label="确认密码:" label-for="input-confirm_password">
                    <b-form-input
                            id="input-confirm_password"
                            type="password"
                            placeholder="请再次输入密码"
                            v-model="form.confirm_password"
                            :state="this.$v.form.confirm_password.$dirty ? !this.$v.form.confirm_password.$error : null"
                            aria-describedby="input-confirm_password-live-feedback"></b-form-input>
                    <b-form-invalid-feedback id="input-confirm_password-live-feedback">
                        您输入的密码不匹配
                    </b-form-invalid-feedback>
                </b-form-group>
                <b-form-group
                        label-cols-lg="3" label="验证码:" label-for="input-captcha">
                    <b-form-input
                            id="input-captcha"
                            type="text"
                            placeholder="请输入验证码"
                            v-model="form.captcha"
                            :state="this.$v.form.captcha.$dirty ? !this.$v.form.captcha.$error : null"
                            aria-describedby="input-captcha-live-feedback"></b-form-input>
                    <b-form-invalid-feedback id="input-captcha-live-feedback">
                        请输入四位验证码.
                    </b-form-invalid-feedback>
                </b-form-group>

                <b-button block type="submit" variant="primary">注册</b-button>
            </b-form>
        </b-card>
    </div>
</template>

<script>
	import {validationMixin} from 'vuelidate'
	import {required, minLength, maxLength, email, url, sameAs, integer} from 'vuelidate/lib/validators'

	export default {
		mixins: [validationMixin],
		name: "Register",
		data() {
			return {
				avatar_url: '',
				form: {
					email: '',
					nickname: '',
					avatar: '',
					sex: null,
					password: '',
					confirm_password: '',
					captcha: '',
				},
				sexType: [
					{value: null, text: '请选择'},
					{value: 'man', text: '男孩'},
					{value: 'girl', text: '女孩'},
				],
				items:
					[
						{
							text: '首页',
							href: '#'
						},
						{
							text: '注册',
							href: '#',
							active: true
						}
					]
			}
		},
		validations: {
			form: {
				email: {
					email,
					required
				},
				sex: {
					required
				},
				password: {
					required,
					minLength: minLength(6)
				},
				confirm_password: {
					sameAsPassword: sameAs('password')
				},
				avatar: {
					url,
					required
				},
				nickname: {
					required,
					minLength: minLength(4)
				},
				captcha: {
					required,
					integer,
					minLength: minLength(4),
					maxLength: maxLength(4)
				}
			}
		},
		methods: {
			onFileChange(e) {
				const file = e.target.files[0];
				this.avatar_url = URL.createObjectURL(file);
			},
			onSubmit() {
				this.$v.form.$touch()
				if (this.$v.form.$anyError) {
					return false
				}
				return false
			}
		}
	}
</script>

<style scoped>

</style>
