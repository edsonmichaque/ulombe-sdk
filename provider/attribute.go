package provider

type Attribute struct {
    Name string
    Type string
    Required bool
}

func NewAttribute(name, typ string) Attribute {
        return Attribute{
            Name: name,
            Type: typ,
	    Required: false,
        }
}

func NewRequiredAttribute(name, typ string) Attribute {
	return Attribute{
		Name: name,
		Type: typ,
		Required: true,
	}
}

