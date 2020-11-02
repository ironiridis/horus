package main

type Issuer struct {
	Name    string
	Email   string
	Phone   Phone
	Address []string
	Site    Website
}

type Recipient struct {
	Name    string
	Address []string
}

type Website struct {
	URL  string
	Name string
}

type Phone string

func (p Phone) Fmt() string {
	//TODO(cjh): render +1aaaeeennnn phone into aaa.eee.nnnn
	return string(p)
}
