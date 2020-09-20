package provider

type Operation struct {
    Name string
    Aliases []string
    Attributes map[string]Attribute
}

func NewOperation(name string, aliases []string, attributes map[string]Attribute) Operation {
	return Operation{
		Name: name,
		Aliases: aliases,
		Attributes: attributes,
	}
}

func (o Operation) AddAttribute(a Attribute) {
        o.Attributes[a.Name] = a
}
