package provider

type Attribute struct {
    Name string
    Type string
    Description string
}

func NewAttribute(name, typ, desc string) Attribute {
        return Attribute{
            Name: name,
            Type: typ,
            Description: desc,
        }
}

