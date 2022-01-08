package xshell

import (
	ps "github.com/bhendo/go-powershell"
	"github.com/bhendo/go-powershell/backend"
)

func Powershell() (ps.Shell, error) {
	// 选择一个后台
	back := &backend.Local{}

	// 开启一个本地 powershell 进程
	shell, err := ps.New(back)
	if err != nil {
		return nil, err
	}
	return shell, nil
}
