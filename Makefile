normal: build-ios
#初始化 mobile 环境命令
install:
	cd $HOME/Documents/gomodworkspace/
	git clone https://github.com/golang/mobile.git
	g use 1.18.6
	go env -w GOPROXY=https://goproxy.cn,direct
	cd $HOME/Documents/gomodworkspace/mobile/cmd/gobind
	go build -o /usr/local/bin/gomobile
	gomobile --help
#构建 mobile 开发环境
gomobile:

build-ios:
