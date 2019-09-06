<template>
    <div>
        <div :class="classObj" class="app-wrapper">
            <div style="background: #ffffff !important;z-index: 10;">
                <b-navbar toggleable="lg" type="dark" variant="light"
                          style="background: #ffffff !important;width: 80%;margin: 0 auto;">
                    <b-navbar-brand style="color: #000000;font-size: 26px;" href="#">杂货铺社区</b-navbar-brand>

                    <b-navbar-toggle target="nav-collapse"></b-navbar-toggle>

                    <b-collapse id="nav-collapse" is-nav>

                        <!-- Right aligned nav items -->
                        <b-navbar-nav class="ml-auto">
                            <b-nav-form>
                                <b-form-input size="sm" class="mr-sm-2" placeholder="Search"></b-form-input>
                                <b-button size="sm" class="my-2 my-sm-0" type="submit"> 搜索</b-button>
                            </b-nav-form>

                            <b-nav-item-dropdown right>
                                <!-- Using 'button-content' slot -->
                                <template slot="button-content"><em>张三</em></template>
                                <b-dropdown-item href="#">个人中心</b-dropdown-item>
                                <b-dropdown-item href="#">退出</b-dropdown-item>
                            </b-nav-item-dropdown>
                        </b-navbar-nav>
                    </b-collapse>
                </b-navbar>
            </div>
            <navbar/>
            <app-main/>
        </div>

        <!-- 底部 -->
        <!-- Footer -->
        <footer class="section footer-classic context-dark bg-image" style="background: #2d3246;z-index: 10;">
            <div class="container">
                <div class="row row-30">
                    <div class="col-md-4 col-xl-5">
                        <div class="pr-xl-4"><a class="brand" href="index.html"><img class="brand-logo-light"
                                                                                     src="images/agency/logo-inverse-140x37.png"
                                                                                     alt="" width="140" height="37"
                                                                                     srcset="images/agency/logo-retina-inverse-280x74.png 2x"></a>
                            <p>杂货铺社区为技术专业人士提供他们构建未来所需的知识，工具和策略。</p>
                            <!-- Rights-->
                            <p class="rights"><span>©  </span><span
                                    class="copyright-year">2019</span><span> </span><span>  杂货铺社区</span></p>
                            <p class="rights"><span><a
                                    href="http://beian.miit.gov.cn/state/outPortal/loginPortal.action"> 京ICP备18057307号</a></span>
                            </p>
                        </div>
                    </div>
                    <div class="col-md-4" style="margin-top: 20px;">
                        <h5>交流</h5>
                        <dl class="contact-list">
                            <dt>邮箱:</dt>
                            <dd><a href="mailto:#">kaifeisu@gmail.com</a></dd>
                        </dl>
                        <dl class="contact-list">
                            <dt>手机号:</dt>
                            <dd><a href="tel:#">+86 176-0038-4700</a></dd>
                        </dl>
                    </div>
                    <div class="col-md-4 col-xl-3" style="margin-top: 20px;">
                        <h5>链接</h5>
                        <ul class="nav-list">
                            <li><a href="#">关于</a></li>
                            <li><a href="https://github.com/SUKF/sm" target="view_window">GitHub</a></li>
                            <li><a href="https://Gitee.com/SUKF/sm" target="view_window">Gitee</a></li>
                        </ul>
                    </div>
                </div>
            </div>
        </footer>
    </div>
</template>

<script>
	import {Navbar, AppMain} from './components'
	import ResizeMixin from './mixin/ResizeHandler'

	export default {
		name: 'Layout',
		components: {
			Navbar,
			AppMain
		},
		mixins: [ResizeMixin],
		computed: {
			device() {
				return this.$store.state.app.device
			},
			fixedHeader() {
				return this.$store.state.settings.fixedHeader
			},
			classObj() {
				return {
					mobile: this.device === 'mobile'
				}
			}
		},
		methods: {},
		mounted() {
			particlesJS.load('particles', 'https://www.sukf.top/particles.json');
		}
	}
</script>

<style lang="scss" scoped>
    .app-wrapper {
        position: relative;
    }

    .context-dark, .bg-gray-dark, .bg-primary {
        color: rgba(255, 255, 255, 0.8);
    }

    .footer-classic a, .footer-classic a:focus, .footer-classic a:active {
        color: #ffffff;
    }

    .nav-list li {
        padding-top: 5px;
        padding-bottom: 5px;
    }

    .nav-list li a:hover:before {
        margin-left: 0;
        opacity: 1;
        visibility: visible;
    }

    ul, ol {
        list-style: none;
        padding: 0;
        margin: 0;
    }

    .social-inner {
        display: flex;
        flex-direction: column;
        align-items: center;
        width: 100%;
        padding: 23px;
        font: 900 13px/1 "Lato", -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
        text-transform: uppercase;
        color: rgba(255, 255, 255, 0.5);
    }

    .social-container .col {
        border: 1px solid rgba(255, 255, 255, 0.1);
    }

    .nav-list li a:before {
        content: "\f14f";
        font: 400 21px/1 "Material Design Icons";
        color: #4d6de6;
        display: inline-block;
        vertical-align: baseline;
        margin-left: -28px;
        margin-right: 7px;
        opacity: 0;
        visibility: hidden;
        transition: .22s ease;
    }
</style>
