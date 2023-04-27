package serializer

import (
	"fmt"

	"gopkg.in/yaml.v2"
)

type yamlSerializer struct{}

func NewYamlSerializer() *yamlSerializer {
	return &yamlSerializer{}
}

func (s *yamlSerializer) Serialize(a TestStruct, _ int) []byte {
	res, err := yaml.Marshal(a)
	if err != nil {
		panic(err)
	}

	return res
}

func (s *yamlSerializer) Deserialize(data []byte, _ int) TestStruct {
	var res TestStruct
	err := yaml.Unmarshal(data, &res)
	if err != nil {
		panic(fmt.Errorf("could not deserialize YAML data: %s", err))
	}

	return res
}
