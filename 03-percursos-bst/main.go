package main

import "fmt"

type No struct {
	Valor             int
	Esquerdo, Direito *No
}

type Arvore struct {
	Raiz *No
}

func NovaArvore() *Arvore {
	return &Arvore{}
}

func NovoNo(v int) *No {
	return &No{Valor: v}
}

func (a *Arvore) Inserir(v int) {
	a.Raiz = inserir(a.Raiz, v)
}

func inserir(no *No, v int) *No {
	if no == nil {
		return NovoNo(v)
	}

	if v < no.Valor {
		no.Esquerdo = inserir(no.Esquerdo, v)
	} else if v > no.Valor {
		no.Direito = inserir(no.Direito, v)
	}

	return no
}

func PreOrdem(no *No, valores *[]int) {
	if no == nil {
		return
	}

	*valores = append(*valores, no.Valor)
	PreOrdem(no.Esquerdo, valores)
	PreOrdem(no.Direito, valores)
}

func EmOrdem(no *No, valores *[]int) {
	if no == nil {
		return
	}

	EmOrdem(no.Esquerdo, valores)
	*valores = append(*valores, no.Valor)
	EmOrdem(no.Direito, valores)
}

func PosOrdem(no *No, valores *[]int) {
	if no == nil {
		return
	}

	PosOrdem(no.Esquerdo, valores)
	PosOrdem(no.Direito, valores)
	*valores = append(*valores, no.Valor)
}

func EmLargura(raiz *No) []int {
	if raiz == nil {
		return []int{}
	}

	valores := []int{}
	fila := []*No{raiz}
	for len(fila) > 0 {
		atual := fila[0]
		fila = fila[1:]
		valores = append(valores, atual.Valor)

		if atual.Esquerdo != nil {
			fila = append(fila, atual.Esquerdo)
		}
		if atual.Direito != nil {
			fila = append(fila, atual.Direito)
		}
	}
	return valores
}

func main() {
	arvore := NovaArvore()
	for _, valor := range []int{50, 30, 70, 20, 40, 60, 80, 35, 65} {
		arvore.Inserir(valor)
	}

	preOrdem := []int{}
	emOrdem := []int{}
	posOrdem := []int{}

	PreOrdem(arvore.Raiz, &preOrdem)
	EmOrdem(arvore.Raiz, &emOrdem)
	PosOrdem(arvore.Raiz, &posOrdem)

	// Em uma BST, todos os valores da subarvore esquerda sao menores que a raiz,
	// e todos os da direita sao maiores. Por isso, EmOrdem gera valores crescentes.
	fmt.Println("Pre-ordem:", preOrdem)
	fmt.Println("Em ordem:", emOrdem)
	fmt.Println("Pos-ordem:", posOrdem)
	fmt.Println("Em largura:", EmLargura(arvore.Raiz))
}
