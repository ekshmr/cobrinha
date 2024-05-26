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

func main() {

	var exit = false

	for i := range tabuleiro {
		for o := range tabuleiro[i] {
			tabuleiro[i][o] = " "
		}
	}

	var teste_pisca = 0

	for exit != true {
		desenhar()
		time.Sleep(10 * time.Millisecond)
		if teste_pisca < 5 {
			teste_pisca++
		} else {
			if tabuleiro[10][10] == " " {
				tabuleiro[10][10] = "*"
			} else {
				tabuleiro[10][10] = " "
			}
		}
	}
}
