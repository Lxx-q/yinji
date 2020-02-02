
window.SERVER = "yinji"

//js 引入 其他的  js 和 css 文件
window.import = {
	css: function(path) {
		if(!path || path.length === 0) {
			throw new Error('参数"path"错误');
		}
		var head = document.getElementsByTagName('head')[0];
		var link = document.createElement('link');
		link.href = path;
		link.rel = 'stylesheet';
		link.type = 'text/css';
		head.appendChild(link);
		return this;
	},
	js: function(path) {
		if(!path || path.length === 0) {
			throw new Error('参数"path"错误');
		}
		var head = document.getElementsByTagName('head')[0];
		var script = document.createElement('script');
		script.src = path;
		script.type = 'text/javascript';
		head.appendChild(script);
		return this;
	}
}
