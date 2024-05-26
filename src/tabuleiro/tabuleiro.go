package tabuleiro

type Tabuleiro struct {
	Linhas  []int
	Colunas []int
}

func main() {
	// Conceito inicial de linhas e colunas
	linhas := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12,
		13, 14, 15, 16, 17, 18, 19, 20, 21, 22,
	}
	colunas := []int{
		1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11,
	}
	// Criar um tabuleiro vazio
	t := Tabuleiro{linhas, colunas}
}
