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

func menor(no *No) *No {
	for no.Esquerdo != nil {
		no = no.Esquerdo
	}
	return no
}

func Remover(no *No, v int) *No {
	if no == nil {
		return nil
	}

	if v < no.Valor {
		no.Esquerdo = Remover(no.Esquerdo, v)
		return no
	}
	if v > no.Valor {
		no.Direito = Remover(no.Direito, v)
		return no
	}

	// Caso 1: no folha ou no com apenas o filho direito.
	if no.Esquerdo == nil {
		return no.Direito
	}
	// Caso 2: no com apenas o filho esquerdo.
	if no.Direito == nil {
		return no.Esquerdo
	}

	// Caso 3: no com dois filhos. O sucessor e o menor da subarvore direita.
	sucessor := menor(no.Direito)
	no.Valor = sucessor.Valor
	no.Direito = Remover(no.Direito, sucessor.Valor)
	return no
}

func (a *Arvore) Remover(v int) {
	// A raiz precisa receber o retorno, pois ela tambem pode ser removida.
	a.Raiz = Remover(a.Raiz, v)
}

func EmOrdem(no *No, valores *[]int) {
	if no == nil {
		return
	}
	EmOrdem(no.Esquerdo, valores)
	*valores = append(*valores, no.Valor)
	EmOrdem(no.Direito, valores)
}

func novaArvoreDeTeste() *Arvore {
	arvore := NovaArvore()
	for _, valor := range []int{50, 30, 70, 20, 40, 60, 80, 35, 65} {
		arvore.Inserir(valor)
	}
	return arvore
}

func testarRemocao(descricao string, valor int) {
	// Cada teste usa uma arvore separada, como solicitado no enunciado.
	arvore := novaArvoreDeTeste()
	arvore.Remover(valor)

	resultado := []int{}
	EmOrdem(arvore.Raiz, &resultado)
	fmt.Printf("%s - remover %d: %v\n", descricao, valor, resultado)
}

func main() {
	testarRemocao("No folha", 20)
	testarRemocao("No com um filho", 60)
	testarRemocao("No com dois filhos", 50)
}
