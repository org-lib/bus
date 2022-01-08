Svc := "Bus"
Name := "bus"
Type := "json"
GOENV = g use 1.16
SETPROXY = GOPROXY=https://goproxy.cn
GOTIDY = go mod tidy
GOVENDOR = go mod vendor
GOBUILD = go build -o .\build\amd64\bus.exe .\example\makeapp\main.go
GOBUILDWIX = go run .\example\makewix\main.go $(Svc) $(Name) $(Type)
MAKEMSI = makemsi.exe make --msi .\build\amd64\Bus.msi --version 0.1 --arch amd64

normal: makebus

makebus:
	@ECHO OFF
	$(GOENV)
	$(GOTIDY)
	$(GOVENDOR)
	$(GOBUILD)
	$(GOBUILDWIX)
	$(MAKEMSI)
