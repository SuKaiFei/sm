<template>
    <div style="background: rgb(245, 245, 246) !important;border-bottom: 3px solid rgb(217, 220, 221);">
        <b-navbar toggleable="lg" type="dark" variant="light" style="margin: 0 auto;width: 60%;background: #F5F5F6 !important;">
          <b-navbar-nav class="ml-auto">Java</b-navbar-nav>
          <b-navbar-nav class="ml-auto">Golang</b-navbar-nav>
          <b-navbar-nav class="ml-auto">微服务</b-navbar-nav>
          <b-navbar-nav class="ml-auto">大数据</b-navbar-nav>
          <b-navbar-nav class="ml-auto">物联网</b-navbar-nav>
          <b-navbar-nav class="ml-auto">AI</b-navbar-nav>
          <b-navbar-nav class="ml-auto">安全</b-navbar-nav>
          <b-navbar-nav class="ml-auto">性能</b-navbar-nav>
        </b-navbar>
    </div>
</template>

<script>
	import {mapGetters} from 'vuex'
	import Breadcrumb from '@/components/Breadcrumb'
	import Hamburger from '@/components/Hamburger'

	export default {
		components: {
			Breadcrumb,
			Hamburger
		},
		computed: {
			...mapGetters([
				'sidebar',
				'avatar'
			])
		},
		methods: {
			toggleSideBar() {
				this.$store.dispatch('app/toggleSideBar')
			},
			async logout() {
				await this.$store.dispatch('user/logout')
				this.$router.push(`/login?redirect=${this.$route.fullPath}`)
			}
		}
	}
</script>

<style lang="scss" scoped>
  .ml-auto, .mx-auto{
    color: #6d7b80;
    margin-right: auto !important;
  }
</style>
