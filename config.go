package ablib

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"strings"
)

type ABConfig struct {
	configs map[string]Experiment
}

func (c *ABConfig) String() string {
	var sb strings.Builder
	sb.WriteString("experiments:\n")
	for key := range c.configs {
		sb.WriteString(fmt.Sprintf("- %s\n", key))
	}
	return sb.String()
}

// MakeConfig builds an experiment configuration given a file
// with a YAML spec.
//
// The YAML spec should specify each experiment with the following
// format:
// foo:
//   desc: "optional description of the experiment"
//   comp:
//   -
//     name: bar
//     dist: 40
//   -
//     name: baz
//     dist: 60
//
// This will create an experiment "foo" with 2 variations.
// "bar" will be chosen with 40/100 probability, while
// "baz" will be chosen with 60/100 probability.
func MakeConfig(filename string) (ABConfig, error) {
	conf := ABConfig{
		configs: make(map[string]Experiment),
	}
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		return conf, fmt.Errorf("failed to read Experiment: %w", err)
	}
	if err = yaml.Unmarshal(dat, &conf.configs); err != nil {
		return conf, fmt.Errorf("failed to unmarshal Experiment: %w", err)
	}
	return conf, nil
}

// Experiment returns the experiment with the provided name.
func (c *ABConfig) Experiment(name string) (Experiment, error) {
	conf, ok := c.configs[name]
	if !ok {
		return conf, fmt.Errorf("experiment %s not found", name)
	}
	conf.Name = name
	return conf, nil
}



