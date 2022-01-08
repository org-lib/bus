package config

import (
	"errors"
	"flag"
	"fmt"
	"github.com/org-lib/bus/pool"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

// ViperConfig 定义一个viper的struct
type ViperConfig struct {
	V *viper.Viper
}

// Config 声明一个ViperConfig类型的变量
var (
	path   string
	ftype  string
	Config ViperConfig
	Work   *pool.WaitGroup
)

func init() {
	//构建一个命令行参数，指定配置文件位置
	printPaser()
	if err := LoadConfig(&Config); err != nil {
		panic(err.Error())
	}
}

func printPaser() {
	flag.StringVar(&path, "conf", "", "自定义自己的配置文件路径和名称（默认config.yaml）")
	flag.StringVar(&ftype, "type", "yaml", "自定义配置文件类型（默认yaml）")
	help := flag.Bool("help", false, "Display usage")
	flag.CommandLine.SetOutput(os.Stdout)

	// 必须有这一行
	flag.Parse()
	if *help {
		fmt.Fprintf(os.Stdout, "Usage of library:\n")
		flag.PrintDefaults()
		os.Exit(1)
	}
}

// LoadConfig 读取配置文件载入配置
func LoadConfig(vc *ViperConfig) error {
	vc.V = viper.New()

	//自定义配置文件路径
	if path != "" {
		dir, mfile := filepath.Split(path)

		// 设置配置文件的文件名

		vc.V.SetConfigName(mfile)

		// 添加配置文件所在路径

		vc.V.AddConfigPath(dir)

		fileExt := filepath.Ext(path)
		//文件后缀类型支持  JSON, TOML, YAML, HCL, INI, envfile
		tmp := "yaml"
		switch fileExt {
		case ".yaml":
			tmp = "yaml"
		case ".json":
			tmp = "json"
		case ".toml":
			tmp = "toml"
		case ".ini":
			tmp = "ini"
		case ".hcl":
			tmp = "hcl"
		default:
			return errors.New(fmt.Sprintf("该文件 %v 后缀类型暂时不支持.", path))
		}
		//...文件类型等待更新

		// 设置配置文件类型

		vc.V.SetConfigType(tmp)

	} else {
		vc.V.SetConfigName("config.yaml")
		// 添加配置文件所在路径

		vc.V.AddConfigPath("$GOPATH/src/")
		vc.V.AddConfigPath("./")

		// 设置配置文件类型

		vc.V.SetConfigType("yaml")
	}

	if err := vc.V.ReadInConfig(); err != nil {
		return fmt.Errorf("Failed to load configuration file, err: %s", err)
	}

	return nil
}
