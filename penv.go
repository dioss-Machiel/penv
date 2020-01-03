package penv

import (
	"fmt"
	"runtime"
)

type environmentActionsList []EnvironmentActions

var registered = make(environmentActionsList, 0)

// RegisterDAO registers a new data access object
func RegisterDAO(dao DAO) {
	ea := EnvironmentActions{dao: dao}
	registered = append(registered, ea)
}

// AppendEnv permanently appends an environment variable
func AppendEnv(name, value string) error {
	for _, r := range registered {
		err := r.AppendEnv(name, value)
		if err != nil {
			return fmt.Errorf("failed to append environment: %v", err)
		}
	}

	return nil
}

func init() {
	registerDAOsForPlatform()
}

func registerDAOsForPlatform() {

	//see darwin_dao and windows_dao for those platforms
	//Cannot be here since they are conditionally compiled

	if runtime.GOOS == "linux" {
		RegisterDAO(ProfileDAOInstance)
	}
}

// SetEnv permanently sets an environment variable
func SetEnv(name, value string) error {
	for _, r := range registered {
		err := r.SetEnv(name, value)
		if err != nil {
			return fmt.Errorf("failed to set environment: %v", err)
		}
	}

	return nil
}

// UnsetEnv permanently unsets an environment variable
func UnsetEnv(name string) error {
	for _, r := range registered {
		err := r.UnsetEnv(name)
		if err != nil {
			return fmt.Errorf("failed to unset environment: %v", err)
		}
	}

	return nil
}
