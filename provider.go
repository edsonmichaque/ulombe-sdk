package sdk

import (
    provider "gitlab.com/ulombe/sdk/provider"
)

type Provider struct {
    Name string
    Resource string
    Operations map[string]provider.Operation
    ScriptGenerator provider.ScriptFunction
}

func NewProvider(name string, generator provider.ScriptFunction) Provider {
	return Provider{
		Name: name,
		Operations: make(map[string]provider.Operation),
		ScriptGenerator: generator,
	}
}

func (p Provider) AddOperation(o provider.Operation) {
	p.Operations[o.Name] = o
}

func (p Provider) GenerateScript(op string, data map[string]interface{}) string {
	return  p.ScriptGenerator(op, data)
}

func (p Provider) AddScriptGenerator(function func(pro *Provider) func(op string, data map[string]interface{}) string) {
	p.ScriptGenerator = function(&p)
}

func (p Provider) GetOperation(operation string) string {
	for k, v := range(p.Operations) {
		if k == operation {
			return k
		}

		for k2, v2 := range v.Aliases {
			if v2 == operation {
				return k
			}
		}
	}

	return nil
}
