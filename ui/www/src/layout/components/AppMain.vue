<template>
    <section class="app-main" style="background-color: #F5F5F6;">
        <transition name="fade-transform" mode="out-in">
            <div style="height: 100%;width: 80%;margin: 0 auto;">
                <b-container style="padding-right: 0;padding-left: 0;margin-right: 0;margin-left: 0;max-width: 100%;">
                    <b-row>
                        <b-col class="col-lg-9 col-md-8 col-sm-7">
                            <router-view :key="key"/>
                        </b-col>
                        <b-col class="col-lg-3 col-md-4 col-sm-5">
                            <sidebar></sidebar>
                        </b-col>
                    </b-row>
                </b-container>
            </div>
        </transition>
    </section>
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
