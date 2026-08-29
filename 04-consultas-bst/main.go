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

func Altura(no *No) int {
	if no == nil {
		return -1
	}

	alturaEsquerda := Altura(no.Esquerdo)
	alturaDireita := Altura(no.Direito)
	if alturaEsquerda > alturaDireita {
		return alturaEsquerda + 1
	}
	return alturaDireita + 1
}

func Contar(no *No) int {
	if no == nil {
		return 0
	}
	return 1 + Contar(no.Esquerdo) + Contar(no.Direito)
}

func ContarFolhas(no *No) int {
	if no == nil {
		return 0
	}
	if no.Esquerdo == nil && no.Direito == nil {
		return 1
	}
	return ContarFolhas(no.Esquerdo) + ContarFolhas(no.Direito)
}

func Minimo(no *No) *No {
	if no == nil {
		return nil
	}
	for no.Esquerdo != nil {
		no = no.Esquerdo
	}
	return no
}

func Maximo(no *No) *No {
	if no == nil {
		return nil
	}
	for no.Direito != nil {
		no = no.Direito
	}
	return no
}

func mostrarConsultas(nome string, arvore *Arvore) {
	fmt.Println(nome)
	fmt.Println("Altura:", Altura(arvore.Raiz))
	fmt.Println("Quantidade de nos:", Contar(arvore.Raiz))
	fmt.Println("Quantidade de folhas:", ContarFolhas(arvore.Raiz))

	minimo := Minimo(arvore.Raiz)
	maximo := Maximo(arvore.Raiz)
	if minimo == nil {
		fmt.Println("Minimo: arvore vazia")
		fmt.Println("Maximo: arvore vazia")
	} else {
		fmt.Println("Minimo:", minimo.Valor)
		fmt.Println("Maximo:", maximo.Valor)
	}
	fmt.Println()
}

func main() {
	vazia := NovaArvore()
	mostrarConsultas("Arvore vazia", vazia)

	unicoNo := NovaArvore()
	unicoNo.Inserir(50)
	mostrarConsultas("Arvore com um unico no", unicoNo)

	arvore := NovaArvore()
	for _, valor := range []int{50, 30, 70, 20, 40, 60, 80, 35, 65} {
		arvore.Inserir(valor)
	}
	mostrarConsultas("Arvore completa", arvore)
}
