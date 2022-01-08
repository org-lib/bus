#config args
Svc := Bus
Name := bus
Cnf  := -conf [INSTALLDIR]assets/config.yaml
Type := json
WixJson := wix.json
Version := 0.1
Arch := amd64
ProjectPath := $(GOPATH)/$(Name)
BuildDir := build/amd64
MainDoGo := example/makeapp/main.go
MakeWixMainDoGo := $(ProjectPath)/example/makewix/main.go

#gowix.json
BuildGoWixJson := $(MakeWixMainDoGo)

ExeFile := $(ProjectPath)/$(BuildDir)/$(Name).exe
BuildGoMainFile := $(ProjectPath)/$(MainDoGo)
MakeMsi := makemsi.exe make --msi
MakeDotMsi := $(ProjectPath)/$(BuildDir)/$(Svc).msi
#go env config
GoBuild := go build -o
GoRun :=  go run
GoENV := g use 1.16
GoTIDY := go mod tidy
GOVENDOR := go mod vendor
#go build and make msi
GoBUILD = $(GoBuild) $(ExeFile) $(BuildGoMainFile)
GoBuildWix = $(GoRun) $(BuildGoWixJson) "$(Svc)" "$(Name)" "$(Type)" "$(Cnf)" "$(WixJson)"
DoMakeMSI = $(MakeMsi) $(MakeDotMsi) --version $(Version) --arch $(Arch)

normal: makebus

makebus:
	@ECHO "Go to building $(Svc).msi"
	@ECHO $(ProjectPath)
	cd $(ProjectPath)
	$(GoENV)
	$(GoTIDY)
	$(GOVENDOR)
	$(GoBUILD)
	$(GoBuildWix)
	$(DoMakeMSI)
	-del $(WixJson)
#@ping 127.0.0.1 -n 20
#无法删除....
#-del $(ExeFile)
