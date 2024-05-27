package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

const comprimento = 100
const largura = 25

var tabuleiro = [largura][comprimento]string{}

type cobra struct {
	tamanho   int
	posicao_X int
	posicao_Y int
}

var tamanho_cobrinha = 2

func limpar_console() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func desenhar() {
	limpar_console()

	var output strings.Builder

	output.WriteString("┌")
	for i := 0; i < comprimento; i++ {
		output.WriteString("─")
	}
	output.WriteString("┐\n")

	for _, lista := range tabuleiro {
		output.WriteString("│")
		for _, elemento := range lista {
			output.WriteString(elemento)
		}
		output.WriteString("│\n")
	}

	output.WriteString("└")
	for i := 0; i < comprimento; i++ {
		output.WriteString("─")
	}
	output.WriteString("┘\n")

	fmt.Print(output.String())
}

func posicionar_cobrinha(cobrinha1 cobra) {
	for i := 0; i < tamanho_cobrinha; i++ {
		tabuleiro[cobrinha1.posicao_Y][cobrinha1.posicao_X+i] = "*"
	}
}

func mover_cobrinha(cobrinha2 *cobra) {
	if cobrinha2.posicao_X != comprimento {
		cobrinha2.posicao_X++
	}
}

func limpar_tabuleiro() {
	for i := range tabuleiro {
		for o := range tabuleiro[i] {
			tabuleiro[i][o] = " "
		}
	}
}

func main() {
	var cobrinha cobra
	cobrinha.tamanho = 2
	cobrinha.posicao_X = comprimento / 2
	cobrinha.posicao_Y = largura / 2

	var exit = false

	posicionar_cobrinha(cobrinha)
	//var teste_pisca = 0

	for !exit {
		limpar_tabuleiro()
		posicionar_cobrinha(cobrinha)
		desenhar()
		time.Sleep(500 * time.Millisecond)
		mover_cobrinha(&cobrinha)

		/*
			if teste_pisca < 5 {
				teste_pisca++
			} else {
				if tabuleiro[10][10] == " " {
					tabuleiro[10][10] = "*"
				} else {
					tabuleiro[10][10] = " "
				}
			}*/
	}
}
