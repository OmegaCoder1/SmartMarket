package model

type Produto struct {
	nome  string
	preco float32
}

func GetNome(p Produto) string {
	return p.nome
}

func (p *Produto) SetNome(nome string) {
	p.nome = nome
}

func GetPreco(p Produto) float32 {
	return p.preco
}

func (p *Produto) SetPreco(preco float32) {
	p.preco = preco
}
