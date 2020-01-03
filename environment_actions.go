package penv

import "fmt"

// DAO defines the interface for loading and saving a set of environment
// variables
type DAO interface {
	Load() (*Environment, error)
	Save(*Environment) error
}

type (

	// Environment is a collection of appenders, setters and unsetters
	Environment struct {
		Appenders []NameValue
		Setters   []NameValue
		Unsetters []NameValue
	}

	// NameValue is a name value pair
	NameValue struct {
		Name  string
		Value string
	}

	// EnvironmentActions Performs actions on the env DAO's
	EnvironmentActions struct {
		dao DAO
	}
)

func filter(arr []NameValue, cond func(NameValue) bool) []NameValue {
	nvs := make([]NameValue, 0, len(arr))
	for _, nv := range arr {
		if cond(nv) {
			nvs = append(nvs, nv)
		}
	}
	return nvs
}

// AppendEnv permanently appends an environment variable
func (environmentActions *EnvironmentActions) AppendEnv(name, value string) error {
	env, err := environmentActions.dao.Load()
	if err != nil {
		return fmt.Errorf("failed to load environment: %v", err)
	}
	env.Setters = filter(env.Setters, func(nv NameValue) bool {
		return nv.Name != name || nv.Value != value
	})
	env.Appenders = filter(env.Appenders, func(nv NameValue) bool {
		return nv.Name != name || nv.Value != value
	})
	env.Appenders = append(env.Appenders, NameValue{name, value})
	// if it's being unset, remove it from the list
	env.Unsetters = filter(env.Unsetters, func(nv NameValue) bool {
		return nv.Name != name
	})
	err = environmentActions.dao.Save(env)
	if err != nil {
		return fmt.Errorf("failed to save environment: %v", err)
	}

	return nil
}

// SetEnv permanently sets an environment variable
func (environmentActions *EnvironmentActions) SetEnv(name, value string) error {
	env, err := environmentActions.dao.Load()
	if err != nil {
		return fmt.Errorf("failed to load environment: %v", err)
	}
	env.Setters = filter(env.Setters, func(nv NameValue) bool {
		return nv.Name != name
	})
	env.Setters = append(env.Setters, NameValue{name, value})
	env.Unsetters = filter(env.Unsetters, func(nv NameValue) bool {
		return nv.Name != name
	})
	err = environmentActions.dao.Save(env)
	if err != nil {
		return fmt.Errorf("failed to save environment: %v", err)
	}

	return nil
}

// UnsetEnv permanently unsets an environment variable
func (environmentActions *EnvironmentActions) UnsetEnv(name string) error {
	env, err := environmentActions.dao.Load()
	if err != nil {
		return fmt.Errorf("failed to load environment: %v", err)
	}
	env.Setters = filter(env.Setters, func(nv NameValue) bool {
		return nv.Name != name
	})
	env.Appenders = filter(env.Appenders, func(nv NameValue) bool {
		return nv.Name != name
	})
	env.Unsetters = filter(env.Unsetters, func(nv NameValue) bool {
		return nv.Name != name
	})
	env.Unsetters = append(env.Unsetters, NameValue{name, ""})

	err = environmentActions.dao.Save(env)
	if err != nil {
		return fmt.Errorf("failed to save environment: %v", err)
	}

	return nil
}
