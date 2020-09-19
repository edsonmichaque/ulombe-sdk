package provider

type Operation struct {
    Name string
    Aliases []string
    Attributes map[string]Attribute
}

func NewOperation(name string, aliases []string) Operation {
        return Operation{
                Name: name,
                Aliases: aliases,
        }
}

func (o Operation) AddAttribute(a Attribute) {
        o.Attributes[a.Name] = a
}
