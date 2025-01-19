package gloomo

import (
	"errors"
	"fmt"
)

type startupSystem struct {
	system   System
	finished bool
}

type System interface {
	Run() error
}

// Manager is a standard system runner for gloomo engine.
type Manager struct {
	systems  []System
	startups []startupSystem
}

// Run all systems in a order of their addition.
func (m *Manager) Run() error {
	var errs []error

	for i, startup := range m.startups {
		if startup.finished {
			continue
		}

		err := startup.system.Run()
		if err != nil {
			errs = append(errs, fmt.Errorf("startup system %T: %w", startup.system, err))
		} else {
			m.startups[i].finished = true
		}
	}

	for _, system := range m.systems {
		err := system.Run()
		if err != nil {
			errs = append(errs, fmt.Errorf("system %T: %w", system, err))
		}
	}

	return errors.Join(errs...)
}

func (m *Manager) Add(system System) {
	m.systems = append(m.systems, system)
}

func (m *Manager) AddStartup(system System) {
	m.startups = append(m.startups, startupSystem{
		system: system,
	})
}
