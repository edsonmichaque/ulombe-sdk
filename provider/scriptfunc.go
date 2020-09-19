package provider

type ScriptFunction func(op string, data map[string]interface{}) string

