Svc := Bus
Name := bus
Type := json
Cnf  := -conf [INSTALLDIR]assets/config.yaml
Version := 0.1
ProjectPath := $(GOPATH)/$(Name)

GoBuild := go build -o
GoRun :=  go run
ExeFile := $(ProjectPath)/$(Name).exe
BuildGoMainFile := $(ProjectPath)\example\makeapp\main.go
BuildGoWixJson := $(ProjectPath)\example\makewix\main.go

GOENV := g use 1.16
GOTIDY := go mod tidy
GOVENDOR := go mod vendor

MakeMsiExe := makemsi.exe make --msi
GOBUILD = $(GoBuild) $(ExeFile) $(BuildGoMainFile)
GOBUILDWIX = $(GoRun) $(BuildGoWixJson) "$(Svc)" "$(Name)" "$(Cnf)" "$(Type)"
MAKEMSI = $(MakeMsiExe) $(ProjectPath)\build\amd64\$(Svc).msi --version $(Version) --arch amd64

normal: makebus

makebus:
	@ECHO "Go to building $(Svc).msi"
	@ECHO $(ProjectPath)
	cd $(ProjectPath)
	$(GOENV)
	$(GOTIDY)
	$(GOVENDOR)
	$(GOBUILD)
	$(GOBUILDWIX)
	$(MAKEMSI)
	-del wix.json
