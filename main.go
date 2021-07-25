package main

import (
	"calcPiWithThreads/utils"
	"fmt"
)

func main() {
	var termsQuantity, threadsQnt int

	fmt.Print("Digite a quantidade de termos: ")
	_, _ = fmt.Scan(&termsQuantity)

	fmt.Print("Digite a quantidade de threads: ")
	_, _ = fmt.Scan(&threadsQnt)

	var numberOfExecutions = 5
	statistics := utils.NewStatisticsCalculator()
	calculator := utils.NewPiCalculator(termsQuantity, threadsQnt)

	for i := 0; i < numberOfExecutions; i++ {
		statistics.AddResultSet(
			*calculator.Calculate(),
			)
	}

	fmt.Println("Valores - Valores - ( Pi / Tempo de execução ):  [")
	for _, r := range statistics.GetResultSet() {
		fmt.Println("    ", r)
	}
	fmt.Println("]\nDuração média (ms): ", statistics.GetAverageTimeInMillis())
	fmt.Printf("Desvio padrão: %.2f\n", statistics.GetStandardDeviation())
	fmt.Printf("Coeficiente de variação: %.4f %%\n", statistics.GetCoefficientOfVariation())
}
