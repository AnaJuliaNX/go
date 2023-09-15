package modelo

type Livro struct {
	ID      int    `json:"id"`
	Titulo  string `json:"titulo"`
	Autor   string `json:"autor"`
	Estoque int    `json:"estoque"`
}

type DadosDoRequestDiminuirEstoque struct {
	Quantidade int `json:"quantidade"`
}

type DadosDoRequestAdicionarEstoque struct {
	Quantidade int `json:"quantidade"`
}
