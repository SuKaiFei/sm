<template>
    <b-container class="bv-example-row"
                 style="padding-right: 0;padding-left: 0;margin: 30px auto;">
        <router-view :key="key"/>
    </b-container>
</template>

<script>
	import Sidebar from '@/views/sidebar/index'

	export default {
		name: 'AppMain',
		components: {
			Sidebar
		},
		computed: {
			key() {
				return this.$route.path
			}
		}
	}
</script>

<style scoped>
    .app-main {
        /*50 = navbar  */
        /*min-height: calc(100vh - 296px);*/
        width: 100%;
        position: relative;
        overflow: hidden;
    }

    .fixed-header + .app-main {
        padding-top: 50px;
    }
</style>

<style lang="scss">
    // fix css style bug in open el-dialog
    .el-popup-parent--hidden {
        .fixed-header {
            padding-right: 15px;
        }
    }
</style>
