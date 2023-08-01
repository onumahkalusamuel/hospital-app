package main

import (
	"github.com/kardianos/service"
	"github.com/onumahkalusamuel/hospital-app/internal"
)

var logger service.Logger

func main() {

	svcConfig := &service.Config{
		Name:        "HcmsSvc",
		DisplayName: "HCMS",
		Description: "Hospital Card Management System",
	}
	runAsService(svcConfig, func() {

		go internal.WebServer()
	})
}

func runAsService(svcConfig *service.Config, run func()) error {
	s, err := service.New(&program{exec: run}, svcConfig)
	if err != nil {
		return err
	}
	logger, err = s.Logger(nil)
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
