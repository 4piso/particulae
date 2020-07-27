package cfg

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// EnvProvider provide configuration from the environment. All key will be
// made UPPERCASE
type EnvProvider struct {
	Namespace string
}

// Provide implements the Provider interface
// store the config in this emppy map
// get the list available from the variables
func (ep *EnvProvider) Provide() (map[string]string, error) {
	config := map[string]string{}
	envs := os.Environ()
	if len(envs) == 0 {
		return nil, errors.New("no enviroment varibles found")
	}
	uspace := fmt.Sprintf("%s_", strings.ToUpper(ep.Namespace))
	for _, val := range envs {
		if !strings.HasPrefix(val, uspace) {
			continue
		}
		idx := strings.Index(val, "=")
		config[strings.ToUpper(strings.TrimPrefix(val[0:idx], uspace))] = val[idx+1:]
	}
	if len(config) == 0 {
		return nil, fmt.Errorf("Namespace %q was not found ", ep.Namespace)
	}
	return config, nil
}
