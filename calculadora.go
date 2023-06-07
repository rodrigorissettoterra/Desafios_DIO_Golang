package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Obter a operação desejada do usuário
	fmt.Println("Calculadora - Digite a operação desejada (+, -, *, /):")
	operacao := lerEntrada()

	// Obter os números do usuário
	numeros := obterNumeros()

	// Realizar a operação selecionada
	switch operacao {
	case "+":
		resultado := adicao(numeros...)
		fmt.Println("Resultado:", resultado)
	case "-":
		resultado := subtracao(numeros...)
		fmt.Println("Resultado:", resultado)
	case "*":
		resultado := multiplicacao(numeros...)
		fmt.Println("Resultado:", resultado)
	case "/":
		resultado, err := divisao(numeros...)
		if err != nil {
			fmt.Println("Erro:", err)
		} else {
			fmt.Println("Resultado:", resultado)
		}
	default:
		fmt.Println("Operação inválida.")
	}
}

// Função para ler a entrada do usuário
func lerEntrada() string {
	var entrada string
	fmt.Scanln(&entrada)
	return entrada
}

// Função para obter os números do usuário
func obterNumeros() []float64 {
	var numeros []float64

	fmt.Println("Digite os números (pressione Enter após cada número, digite 's' para parar):")

	for {
		entrada := lerEntrada()

		if entrada == "s" {
			break
		}

		numero, err := strconv.ParseFloat(entrada, 64)
		if err != nil {
			fmt.Println("Número inválido. Tente novamente.")
			continue
		}

		numeros = append(numeros, numero)
	}

	return numeros
}

// Função de adição
func adicao(numeros ...float64) float64 {
	soma := 0.0

	for _, numero := range numeros {
		soma += numero
	}

	return soma
}

// Função de subtração
func subtracao(numeros ...float64) float64 {
	resultado := numeros[0]

	for i := 1; i < len(numeros); i++ {
		resultado -= numeros[i]
	}

	return resultado
}

// Função de multiplicação
func multiplicacao(numeros ...float64) float64 {
	produto := 1.0

	for _, numero := range numeros {
		produto *= numero
	}

	return produto
}

// Função de divisão
func divisao(numeros ...float64) (float64, error) {
	if len(numeros) < 2 {
		return 0.0, fmt.Errorf("pelo menos dois números são necessários para a divisão")
	}

	resultado := numeros[0]

	for i := 1; i < len(numeros); i++ {
		if numeros[i] == 0 {
			return 0.0, fmt.Errorf("divisão por zero não é permitida")
		}
		resultado /= numeros[i]
	}

	return resultado, nil
}
