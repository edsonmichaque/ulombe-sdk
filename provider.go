package sdk

import (
    provider "gitlab.com/ulombe/sdk/provider"
)

type Provider struct {
    Name string
    Resource string
    Operations map[string]provider.Operation
    ScriptGenerator provider.ScriptGeneratorFunc
}

func NewProvider(name string, generator provider.ScriptGeneratorFunc) Provider {
	return Provider{
		Name: name,
		Operations: []string{},
		ScriptGenerator: generator,
	}
}

func (p Provider) AddOperation(o provider.Operation) {
	p.Operations[o.Name] = o
}

func (p Provider) GenerateScript(op string, data map[string]interface{}) string {
	return  p.ScriptGenerator(op, data)
}
