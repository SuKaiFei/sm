<template>
    <div style="background: rgb(245, 245, 246) !important;border-bottom: 3px solid rgb(217, 220, 221);">
        <b-navbar toggleable="lg" type="light" variant="light"
                  style="background: #F5F5F6 !important;">
            <b-col md="6" offset-md="3" class="navbar" style="padding: 0 0;">
                <b-navbar-nav
                        v-for="tag in tags" :key="tag.code" class="ml-auto">
                    <router-link class="mx-auto" :to="{ path: '/articles?tag={{tag.code}}'}">{{tag.name}}
                    </router-link>
                </b-navbar-nav>
            </b-col>
        </b-navbar>
    </div>
</template>

<script>
	import {mapGetters} from 'vuex'
	import {tags} from '@/api/tag'
	import Breadcrumb from '@/components/Breadcrumb'
	import Hamburger from '@/components/Hamburger'

	export default {
		components: {
			Breadcrumb,
			Hamburger
		},
		data() {
			return {
				tags: []
			}
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
			getTags() {
				tags().then(res => {
					if (res && res.code === 0) {
						this.tags = res.data
					}
				})
			}
		},
		created() {
			this.getTags()
		}
	}
</script>

<style lang="scss" scoped>
    .ml-auto, .mx-auto {
        color: #6d7b80;
        margin-right: 12px !important;
    }

    .navbar .navbar-toggler {
        margin: auto 0 auto auto;
    }
</style>
