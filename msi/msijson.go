package msi

import (
	"bytes"
	"github.com/spf13/viper"
	"strings"
)

//定义 wix.json 模板

var (
	FileJosn = `{
  "product": "xxxxxx",
  "company": "shangguannihao",
  "license": "LICENSE",
  "upgrade-code": "",
  "files": {
    "guid": "",
    "items": [
      "build/amd64/xxxxxx.exe"
    ]
  },
  "directories": [
    "assets"
  ],
  "env": {
    "guid": "",
    "vars": [
      {
        "name": "some",
        "value": "value",
        "permanent": "no",
        "system": "no",
        "action": "set",
        "part": "last"
      }
    ]
  },
  "shortcuts": {
    "guid": "",
    "items": [
      {
        "name": "xxxxxx",
        "description": "xxxxxx web server",
        "target": "[INSTALLDIR]\\xxxxxx.exe",
        "wdir": "INSTALLDIR",
        "icon":"assets/ico.ico"
      }
    ]
  },
  "hooks": [
    {"when": "install", "command": "sc.exe create XxxxxxSvc binPath=\"[INSTALLDIR]xxxxxx.exe -conf [INSTALLDIR]assets\\config.yaml\" type=share start=auto DisplayName=\"xxxxxx\""},
    {"when": "install", "command": "sc.exe start XxxxxxSvc"},
    {"when": "uninstall", "command": "sc.exe delete XxxxxxSvc"}
  ],
  "choco": {
    "description": "xxxxxx program",
    "project-url": "xxxxxx",
    "tags": "xxxxxx",
    "license-url": "xxxxxx"
  }
}`
)

//定义MSI信息

type Msi struct {
	Task     int64    `json:"task"`
	Svc      string   `json:"svc"`
	Display  string   `json:"display"`
	Commands []string `json:"commands"`
}

//获取json文件

func SetJson(svc string, name string, filetype string, filename string) {
	v := viper.New()
	v.SetConfigType(filetype) // 设置配置文件的类型
	fj := strings.ReplaceAll(strings.ReplaceAll(FileJosn, "xxxxxx", name), "Xxxxxx", svc)
	v.ReadConfig(bytes.NewBuffer([]byte(fj)))
	v.WriteConfigAs(filename)
}
