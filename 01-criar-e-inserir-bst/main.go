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

	// Se v for igual a no.Valor, o valor duplicado e ignorado.
	return no
}

func emOrdem(no *No, valores *[]int) {
	if no == nil {
		return
	}

	emOrdem(no.Esquerdo, valores)
	*valores = append(*valores, no.Valor)
	emOrdem(no.Direito, valores)
}

func main() {
	arvore := NovaArvore()
	fmt.Println("Arvore inicialmente vazia:", arvore.Raiz == nil)

	valores := []int{50, 30, 70, 20, 40, 60, 80, 35, 65}
	for _, valor := range valores {
		arvore.Inserir(valor)
	}

	// Teste de valor duplicado: o segundo 35 nao sera inserido.
	arvore.Inserir(35)

	resultado := []int{}
	emOrdem(arvore.Raiz, &resultado)
	fmt.Println("Valores da BST em ordem:", resultado)
}
