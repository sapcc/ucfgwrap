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

package ucfgwrap

import (
	"github.com/elastic/go-ucfg"
	"github.com/elastic/go-ucfg/json"
	"github.com/elastic/go-ucfg/yaml"
)

type Config struct {
	config  *ucfg.Config
	options []ucfg.Option
}

func (c *Config) Unpack(to interface{}) error {
	return c.config.Unpack(to, c.options...)
}

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
