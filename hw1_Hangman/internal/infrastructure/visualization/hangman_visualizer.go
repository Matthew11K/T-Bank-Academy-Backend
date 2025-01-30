package visualization

import (
	"fmt"
	"math"
)

type HangmanVisualizer interface {
	DisplayHangman(attemptsLeft int, maxAttempts int)
}

type ConsoleHangmanVisualizer struct {
	stages []string
}

func NewConsoleHangmanVisualizer() *ConsoleHangmanVisualizer {
	return &ConsoleHangmanVisualizer{
		stages: generateHangmanStages(),
	}
}

func generateHangmanStages() []string {
	var stages []string
	stages = append(stages, generateInitialStages()...)
	stages = append(stages, generateMiddleStages()...)
	stages = append(stages, generateFinalStages()...)

	return stages
}

func generateInitialStages() []string {
	return []string{
		`
       ____
      |    |
      |    
      |   
      |    
      |   
    __|__
    `,
		`
       ____
      |    |
      |    O
      |   
      |    
      |   
    __|__
    `,
		`
       ____
      |    |
      |    O
      |    |
      |    
      |   
    __|__
    `,
	}
}

func generateMiddleStages() []string {
	return []string{
		`
       ____
      |    |
      |    O
      |   /|
      |    
      |   
    __|__
    `,
		`
       ____
      |    |
      |    O
      |   /|\
      |    
      |   
    __|__
    `,
		`
       ____
      |    |
      |    O
      |   /|\
      |   / 
      |   
    __|__
    `,
		`
       ____
      |    |
      |    O
      |   /|\
      |   / \
      |   
    __|__
    `,
	}
}

func generateFinalStages() []string {
	return []string{
		`
       ____
      |    |
      |    O
      |   /|\
      |   / \
      |    |
    __|__
    `,
		`
       ____
      |    |
      |   _O_
      |   /|\
      |   / \
      |    |
    __|__
    `,
		`
       ____
      |    |
      |   _O_
      |  //|\\
      |   / \\
      |    |
    __|__
    `,
	}
}

func (v *ConsoleHangmanVisualizer) DisplayHangman(attemptsLeft, maxAttempts int) {
	totalStages := len(v.stages)
	stageIndex := int(math.Round(float64(totalStages-1) * float64(maxAttempts-attemptsLeft) / float64(maxAttempts)))

	if stageIndex < 0 {
		stageIndex = 0
	}

	if stageIndex >= totalStages {
		stageIndex = totalStages - 1
	}

	fmt.Println(v.stages[stageIndex])
}
