'use strict'
const path = require('path')

const resolve = (dir) => path.join(__dirname, dir);


const name = '杂货铺社区' // page title

// If your port is set to 80,
// use administrator privileges to execute the command line.
// For example, Mac: sudo npm run
// You can change the port by the following methods:
// port = 9528 npm run dev OR npm run dev --port = 9528
const port = process.env.port || process.env.npm_config_port || 9528 // dev port

if (process.env.NODE_ENV === 'development') {
	process.env.VUE_APP_BASE_API = `http://192.168.31.98:8000/sm`
} else {
	process.env.VUE_APP_BASE_API = `https://api.sukf.top/sm`
}

// All configuration item explanations can be find in https://cli.vuejs.org/config/
module.exports = {
	/**
	 * You will need to set publicPath if you plan to deploy your site under a sub path,
	 * for example GitHub Pages. If you plan to deploy your site to https://foo.github.io/bar/,
	 * then publicPath should be set to "/bar/".
	 * In most cases please use '/' !!!
	 * Detail: https://cli.vuejs.org/config/#publicpath
	 */
	publicPath: '/',
	outputDir: 'dist',
	assetsDir: 'static',
	lintOnSave: process.env.NODE_ENV === 'development',
	productionSourceMap: false,
	devServer: {
		port: port,
		overlay: {
			warnings: false,
			errors: true
		},
		proxy: {
			// change xxx-api/login => mock/login
			// detail: https://cli.vuejs.org/config/#devserver-proxy
			[process.env.VUE_APP_BASE_API]: {
				target: `http://127.0.0.1:${port}/mock`,
				changeOrigin: true,
				pathRewrite: {
					['^' + process.env.VUE_APP_BASE_API]: ''
				}
			}
		},
	},
	chainWebpack: (config) => {
		config.module
			.rule('svg')
			.exclude
			.add(resolve('src/icons'))
			.end()
		config.module
			.rule('icons')
			.test(/\.svg$/)
			.include.add(resolve('src/icons'))
			.end()
			.use('svg-sprite-loader')
			.loader('svg-sprite-loader')
			.options({
				symbolId: 'icon-[name]'
			})
			.end()
		config.resolve.alias
			.set('@', resolve('src'))
			.set('assets', resolve('src/assets'))
			.set('components', resolve('src/components'))
			.set('layout', resolve('src/layout'))
			.set('base', resolve('src/base'))
			.set('static', resolve('src/static'))
	}
};
