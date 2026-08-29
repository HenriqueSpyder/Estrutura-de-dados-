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

func (a *Arvore) Buscar(v int) *No {
	return buscar(a.Raiz, v)
}

func buscar(no *No, v int) *No {
	if no == nil || no.Valor == v {
		return no
	}

	if v < no.Valor {
		return buscar(no.Esquerdo, v)
	}
	return buscar(no.Direito, v)
}

func (a *Arvore) BuscarIter(v int) *No {
	atual := a.Raiz
	for atual != nil {
		if atual.Valor == v {
			return atual
		}
		if v < atual.Valor {
			atual = atual.Esquerdo
		} else {
			atual = atual.Direito
		}
	}
	return nil
}

func mostrarResultado(tipo string, valor int, no *No) {
	if no == nil {
		fmt.Printf("Busca %s: valor %d nao encontrado.\n", tipo, valor)
		return
	}
	fmt.Printf("Busca %s: valor %d encontrado.\n", tipo, no.Valor)
}

func main() {
	arvore := NovaArvore()
	for _, valor := range []int{50, 30, 70, 20, 40, 60, 80, 35, 65} {
		arvore.Inserir(valor)
	}

	mostrarResultado("recursiva", 65, arvore.Buscar(65))
	mostrarResultado("recursiva", 45, arvore.Buscar(45))
	mostrarResultado("iterativa", 65, arvore.BuscarIter(65))
	mostrarResultado("iterativa", 45, arvore.BuscarIter(45))
}
