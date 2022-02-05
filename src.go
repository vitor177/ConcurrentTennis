// Parte da concorrência baseada em https://gobyexample.com/waitgroups e https://go.dev/tour/concurrency/2
// https://flaviocopes.com/go-random/
// João Vítor Fonseca de Mendonça

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Variáveis utilizadas para controle dos jogadores 1 e 2
var wg sync.WaitGroup
var matchpoint, pontos1, pontos2, pontosPartida, jogador1, jogador2 = 0, 0, 0, 0, "Jogador 1", "Jogador 2"
var set = 1

// Função para exibir o placar do jogo
func exibePlacar(p1 int, p2 int) {
	fmt.Println("-----------------------------------------------------------")
	fmt.Println("PLACAR - ", p1, " x ", p2)
	fmt.Println("-----------------------------------------------------------")
}

// Função para cálculo do valor absoluto (fim de partida com vantagem de 2 pontos)
func abs(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

// Função para controle simultâneo dos dois jogadores
func player(nome string, vez chan bool) {

	// Serve para executar o Done() após a função retornar!
	defer wg.Done()

	for {
		// canalAberto recebe um valor booleano para controle do canal
		canalAberto := <-vez
		if !canalAberto {
			fmt.Println("A partida acabou, e o Vencedor foi: ", nome)
			fmt.Println("-----------------------------------------------------------")
			fmt.Println("PLACAR FINAL - ", pontos1, " x ", pontos2)
			return
		}
		// Gera número para decidir se vai acertar a bola ou não
		aux := rand.Intn(4)

		// 75% de chance de acertar a bola
		if aux == 0 || aux == 1 || aux == 2 {
			fmt.Println(nome, " acertou a bola no campo adversário.")
			// Sleep para simular o tempo da jogada do Jogador
			//time.Sleep(1 * time.Second)
		} else {
			// 50% de chance de perder o ponto por não ter acertado a bola, ou jogar a bola para fora!
			aux2 := rand.Intn(2)
			if aux2 == 0 {
				fmt.Println(nome, " não acertou a bola")
			}
			if aux2 == 1 {
				fmt.Println(nome, " jogou para fora")
			}

			fmt.Println("RODADA DE NÚMERO: ", set)

			switch nome {
			// Incrementa pontuação do Jogador 2
			case jogador1:
				pontos2++
			// Incrementa pontuação do Jogador 1
			case jogador2:
				pontos1++
			}
			exibePlacar(pontos1, pontos2)

			// Incrementa o número da rodada
			set++
		}
		// Chegar se está na primeira rodada de Matchpoint
		if (pontos1 == pontosPartida-1 || pontos2 == pontosPartida-1) && matchpoint == 0 {
			fmt.Println("ATENÇÃO: PRIMEIRA RODADA DE MATCHPOINT")
			matchpoint++
		}
		// Para tratar se há a vantagem de dois pontos assim que chega o valor de pontos do Sett.
		if (abs(pontos1, pontos2) >= 2) && (pontos1 >= pontosPartida || pontos2 >= pontosPartida) {
			// Partida acabou e o canal é fechado
			close(vez)
			return
		}
		vez <- canalAberto
	}
}

func main() {

	// Para garantir aleatoriedade à cada execução
	rand.Seed(time.Now().UnixNano())

	// Canal criado passando valores booleanos
	vez := make(chan bool)

	// Quantidade de pontos estabelecida
	var quantidade int
	fmt.Println("Digite a quantidade de pontos: ")
	fmt.Scanln(&quantidade)

	// Tratamento de erro para valores inadequados da quantidade de pontos máxima
	for quantidade <= 0 {
		fmt.Println("Digite uma quantidade válida, valor positivo e não nulo: ")
		fmt.Scanln(&quantidade)
	}
	pontosPartida = quantidade

	// Por padrão, adicionar a quantidade de Go Rotinas
	wg.Add(2)

	fmt.Println()
	fmt.Println("Início de partida - ", jogador1, " vs ", jogador2)
	fmt.Println()

	// Go rotina para player 1
	go player(jogador1, vez)
	vez <- true
	// Go rotina para player 2
	go player(jogador2, vez)

	// Wait para esperar a finalização da execução das Go rotinas
	wg.Wait()
}
