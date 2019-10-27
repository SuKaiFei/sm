'use strict'
const path = require('path')
const CompressionPlugin = require("compression-webpack-plugin")

const resolve = (dir) => path.join(__dirname, dir);


const name = '杂货铺社区' // page title

// If your port is set to 80,
// use administrator privileges to execute the command line.
// For example, Mac: sudo npm run
// You can change the port by the following methods:
// port = 9528 npm run dev OR npm run dev --port = 9528
const port = process.env.port || process.env.npm_config_port || 9528 // dev port

if (process.env.NODE_ENV === 'development') {
	process.env.VUE_APP_BASE_API = `http://localhost:8000/sm`
	process.env.VUE_APP_BASE_UPLOAD = `http://localhost:8001/`
} else {
	process.env.VUE_APP_BASE_API = `https://www.sukf.top/sm`
	process.env.VUE_APP_BASE_UPLOAD = `https://www.sukf.top/upload/`
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
	},
	configureWebpack: config => {
		if (process.env.NODE_ENV === 'production') {
			return {
				plugins: [
					new CompressionPlugin({
						test: /\.js$|\.html$|.\css/, //匹配文件名
						threshold: 10240,//对超过10k的数据压缩
						deleteOriginalAssets: false //不删除源文件
					})
				]
			}
		}
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
