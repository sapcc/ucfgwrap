/*******************************************************************************
*
* Copyright 2022 SAP SE
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You should have received a copy of the License along with this
* program. If not, you may obtain a copy of the License at
*
*     http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*
*******************************************************************************/

// Small wrapper for ucfg that passes a slice of ucfg.Options to every unpack function.
package ucfgwrap

import (
	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/json"
	"github.com/elastic/go-ucfg/yaml"
)

// Config bundles a ucfg.Config with a slice of ucfg.Options.
type Config struct {
	config  *ucfg.Config
	options []ucfg.Option
}

// Unpack unpacks to the given pointer.
func (c *Config) Unpack(to interface{}) error {
	return c.config.Unpack(to, c.options...)
}

// Wrap wraps a ucfg.Config together with the options contained
// in the parent ucfgwrap.Config. This handy when parsing into
// ucfg.Configs but passing around ucfgwrap.Configs.
func (c *Config) Wrap(config *ucfg.Config) Config {
	return Config{
		config:  config,
		options: c.options,
	}
}

// FromYAML parses the given yaml into a ucfgwrap.Config using the given options.
func FromYAML(data []byte, opts ...ucfg.Option) (Config, error) {
	yamlConf, err := yaml.NewConfig(data, opts...)
	if err != nil {
		return Config{}, err
	}
	return Config{
		config:  yamlConf,
		options: opts,
	}, nil
}

// FromYAMLFile parses the given file into a ucfgwrap.Config using the given options.
func FromYAMLFile(path string, opts ...ucfg.Option) (Config, error) {
	yamlConf, err := yaml.NewConfigWithFile(path, opts...)
	if err != nil {
		return Config{}, err
	}
	return Config{
		config:  yamlConf,
		options: opts,
	}, nil
}

// FromJSON parses the given json into a ucfgwrap.Config using the given options.
func FromJSON(data []byte, opts ...ucfg.Option) (Config, error) {
	jsonConf, err := json.NewConfig(data, opts...)
	if err != nil {
		return Config{}, err
	}
	return Config{
		config:  jsonConf,
		options: opts,
	}, nil
}

// FromJSONFile parses the given file into a ucfgwrap.Config using the given options.
func FromJSONFile(path string, opts ...ucfg.Option) (Config, error) {
	jsonConf, err := json.NewConfigWithFile(path, opts...)
	if err != nil {
		return Config{}, err
	}
	return Config{
		config:  jsonConf,
		options: opts,
	}, nil
}
