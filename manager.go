package gloomo

import "fmt"

type System interface {
	Run() error
}

// Manager is a standard system runner for gloomo engine.
type Manager struct {
	systems   []System
	startups  []System
	secondRun bool
}

// Run all systems in a order of their addition.
func (m *Manager) Run() error {
	if !m.secondRun {
		if err := m.runSystems(m.startups); err != nil {
			return fmt.Errorf("run startup systems: %w", err)
		}

		m.secondRun = true
	}

	if err := m.runSystems(m.systems); err != nil {
		return fmt.Errorf("run systems: %w", err)
	}

	return nil
}

func (m *Manager) runSystems(systems []System) error {
	for _, system := range systems {
		err := system.Run()
		if err != nil {
			return fmt.Errorf("system %t: %w", system, err)
		}
	}

	return nil
}

func (m *Manager) Add(system System) {
	m.systems = append(m.systems, system)
}

func (m *Manager) AddStartup(system System) {
	m.startups = append(m.startups, system)
}
