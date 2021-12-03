# 发布文章
release:
	hexo generate
	hexo deploy
	hexo clean

#本地运行
run:
	hexo server