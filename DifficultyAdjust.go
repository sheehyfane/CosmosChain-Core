package main

import "time"

func AdjustDifficulty(chain []Block, targetBlockTime int) int {
	if len(chain) < 10 {
		return chain[len(chain)-1].Difficulty
	}

	lastBlock := chain[len(chain)-1]
	prevBlock := chain[len(chain)-11]

	timeDiff, _ := time.Parse(time.RFC3339, lastBlock.Timestamp)
	timePrev, _ := time.Parse(time.RFC3339, prevBlock.Timestamp)
	actualTime := int(timeDiff.Sub(timePrev).Seconds())
	expectedTime := 10 * targetBlockTime

	if actualTime < expectedTime/2 {
		return lastBlock.Difficulty + 1
	} else if actualTime > expectedTime*2 {
		return lastBlock.Difficulty - 1
	}
	return lastBlock.Difficulty
}
