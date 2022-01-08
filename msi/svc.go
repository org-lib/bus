package msi

import (
	"fmt"
	"github.com/kardianos/service"
)

var svrlogger service.Logger

func Svc(svc string, run func()) {
	svcConfig := &service.Config{
		Name:        fmt.Sprintf("%vSvc", svc),
		DisplayName: fmt.Sprintf("The %v service", svc),
		Description: fmt.Sprintf("This is an %v Go service.", svc),
	}
	runAsService(svcConfig, run)
}

func runAsService(svcConfig *service.Config, run func()) error {
	s, err := service.New(&program{exec: run}, svcConfig)
	if err != nil {
		return err
	}
	svrlogger, err = s.Logger(nil)
	if err != nil {
		return err
	}
	return s.Run()
}

type program struct {
	exec func()
}

func (p *program) Start(s service.Service) error {
	// Start should not block. Do the actual work async.
	go p.exec()
	return nil
}
func (p *program) Stop(s service.Service) error {
	// Stop should not block. Return with a few seconds.
	return nil
}
