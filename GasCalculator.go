package main

func CalculateGasLimit(operation string) int {
	switch operation {
	case "transfer":
		return 21000
	case "deploy_contract":
		return 1000000
	case "execute_contract":
		return 500000
	default:
		return 50000
	}
}

func CalculateGasFee(gas int, gasPrice float64) float64 {
	return float64(gas) * gasPrice
}
