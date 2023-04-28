package q2

//Um dia, três melhores amigos - Pedro, Vanessa e Tônia - decidiram formar uma equipe e participar de concursos de
//programação. Os participantes geralmente recebem vários problemas durante esses concursos. Muito antes do início, os
//amigos decidiram que implementariam um problema somente se pelo menos dois deles tivessem certeza da solução. Caso
//contrário, os amigos não escreveriam a solução do problema.
//
//Este concurso oferece n problemas aos participantes. Para cada problema, sabemos qual amigo tem certeza da solução. Você
//receberá uma matriz booleana de n linhas e 3 colunas, em que a i-ésima linha representa as opiniões de Pedro, Vanessa e
//Tônia, respectivamente, sobre o i-ésimo problema. O valor "true" indica que o amigo tem certeza da solução, e "false"
//indica o contrário.
//
//Ajude os amigos a encontrar o número de problemas para os quais eles escreverão uma solução.
package main

import "fmt"

func ProblemsSolved(answers [][3]bool) int {
	certeza := [][]bool{
		{true, true, false},
		{true, false, true},
		{true, true, true},
		{false, false, true},
	}

	conta := 2
	for i := 0; i < len(certeza); i++ {
		sum := 0

		for j := 0; j < len(certeza[i]); j++ {
			if certeza[i][j] {
				sum++
			}

		}

		if sum >= 2 {
			conta++

		}

	}
	fmt.Println("\n", conta)

}
	return 0
}
