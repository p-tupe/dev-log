package main

import (
	"fmt"
	"math/rand"
	"strings"
)

const SEED_TEXT = `In probability theory and statistics, a Markov chain or Markov process is a stochastic process describing a sequence of possible events in which the probability of each event depends only on the state attained in the previous event. Informally, this may be thought of as, "What happens next depends only on the state of affairs now." A countably infinite sequence, in which the chain moves state at discrete time steps, gives a discrete-time Markov chain (DTMC). A continuous-time process is called a continuous-time Markov chain (CTMC). Markov processes are named in honor of the Russian mathematician Andrey Markov.

Markov chains have many applications as statistical models of real-world processes. They provide the basis for general stochastic simulation methods known as Markov chain Monte Carlo, which are used for simulating sampling from complex probability distributions, and have found application in areas including Bayesian statistics, biology, chemistry, economics, finance, information theory, physics, signal processing, and speech processing.

The adjectives Markovian and Markov are used to describe something that is related to a Markov process.`

type Chain = map[string]map[string]int

func main() {
	chain := input(SEED_TEXT)
	output(chain, "A", 500)
}

func input(seedTxt string) Chain {
	chain := Chain{}
	words := strings.Fields(seedTxt)

	for i, currWord := range words {
		nextWord := ""
		if i < len(words)-1 {
			nextWord = words[i+1]
		}

		if chain[currWord] == nil {
			chain[currWord] = map[string]int{}
		}

		chain[currWord][nextWord]++
	}

	return chain
}

func output(chain Chain, word string, count int) {
	if count == 0 {
		return
	}

	fmt.Print(word + " ")

	node := chain[word]

	if node == nil {
		return
	}

	max := 0
	for _, count := range node {
		max += count
	}

	rand := rand.Intn(max)

	sum := 0
	for nextWord, count := range node {
		sum += count

		if sum > rand {
			word = nextWord
			break
		}
	}

	output(chain, word, count-1)
}
