<template>
    <div>
        <b-breadcrumb :items="items"></b-breadcrumb>
        <b-card>
            <b-form @submit.stop.prevent="onSubmit()" :novalidate="true">
                <b-form-group
                        label-cols-lg="3"
                        label="标题:"
                        label-for="input-title">
                    <b-form-input
                            id="input-title"
                            v-model="form.title"
                            placeholder="请输入标题"
                            :state="this.$v.form.title.$dirty ? !this.$v.form.title.$error : null"
                            aria-describedby="input-title-live-feedback"
                    ></b-form-input>
                    <b-form-invalid-feedback id="input-title-live-feedback">
                        标题格式不正确.
                    </b-form-invalid-feedback>
                </b-form-group>
                <b-form-group
                        label="正文:"
                        label-for="input-text">
                    <mavon-editor ref=md @imgAdd="imgAdd" @change="onMeChange" :boxShadow="false"
                                  v-model="form.md_text"/>
                    <b-form-invalid-feedback
                            :state="this.$v.form.md_text.$dirty ? !this.$v.form.md_text.$error : null">
                        请输入正文.
                    </b-form-invalid-feedback>
                </b-form-group>
                <b-form-group
                        label-cols-lg="3"
                        label="标签:"
                        label-for="select-tag">
                    <b-form-select id="select-tag" v-model="form.tag" :options="tags"
                                   :state="this.$v.form.tag.$dirty ? !this.$v.form.tag.$error : null"></b-form-select>
                    <b-form-invalid-feedback id="input-tag-live-feedback">
                        请选择标签.
                    </b-form-invalid-feedback>
                </b-form-group>

                <b-button block type="submit" variant="primary">发布</b-button>
            </b-form>
        </b-card>
    </div>
</template>

<script>
	import {validationMixin} from 'vuelidate'
	import {tags} from '@/api/tag'
	import {create} from '@/api/article'
	import {upload} from '@/api/file'
	import {required, maxLength} from 'vuelidate/lib/validators'

	export default {
		mixins: [validationMixin],
		name: "Article",
		data() {
			return {
				avatar_url: '',
				text_temp: '',
				tags: [],
				form: {
					title: '',
					md_text: '',
					html_text: '',
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
							text: '发布',
							href: '#',
							active: true
						}
					]
			}
		},
		validations: {
			form: {
				title: {
					maxLength: maxLength(50),
					required
				},
				tag: {
					required
				},
				md_text: {
					required
				},
			}
		},
		methods: {
			imgAdd(pos, $file) {
				// 第一步.将图片上传到服务器.
				let formData = new FormData();
				formData.append('file', $file);
				upload(formData).then((res) => {
					// 第二步.将返回的url替换到文本原位置![...](0) -> ![...](url)
					/**
					 * $vm 指为mavonEditor实例，可以通过如下两种方式获取
					 * 1. 通过引入对象获取: `import {mavonEditor} from ...` 等方式引入后，`$vm`为`mavonEditor`
					 * 2. 通过$refs获取: html声明ref : `<mavon-editor ref=md ></mavon-editor>，`$vm`为 `this.$refs.md`
					 */
					if (res && res.code === 0) {
						this.$refs.md.$img2Url(pos, process.env.VUE_APP_BASE_UPLOAD + res.data);
					}
				})
			},
			onMeChange(value, render) {
				this.form.html_text = render
			},
			onFileChange(e) {
				const file = e.target.files[0];
				this.avatar_url = URL.createObjectURL(file);
			},
			onSubmit() {
				this.$v.form.$touch()
				if (this.$v.form.$anyError) {
					return false
				}
				create(this.form).then((res) => {
					if (res && res.code === 0) {
						this.$bvToast.toast(`发布帖子 ${this.form.title} 成功`, {
							title: '操作提示',
							variant: 'success',
							solid: true
						})
						this.$router.push('/article/index')
					} else {
						this.$bvToast.toast(`发布帖子 ${this.form.title} 失败，${res.message}`, {
							title: '操作提示',
							variant: 'danger',
							solid: true
						})
					}
				})
				return false
			}
		},
		async mounted() {
			await tags().then(res => {
				if (res && res.code === 0) {
					let tags = [{value: null, text: '请选择'}]
					for (let i = 0; i < res.data.length; i++) {
						tags.push({value: res.data[i].code, text: res.data[i].name})
					}
					this.tags = tags
				}
			})
		}
	}
</script>

<style scoped>

</style>
