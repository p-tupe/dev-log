// This program shows how to generate text strings from a seed text using markov chain.
//
//	https://en.wikipedia.org/wiki/Markov_chain
package main

import (
	"fmt"
	"slices"
	"strings"
)

const SEED_TEXT = `In probability theory and statistics, a Markov chain or Markov process is a stochastic process describing a sequence of possible events in which the probability of each event depends only on the state attained in the previous event. Informally, this may be thought of as, "What happens next depends only on the state of affairs now." A countably infinite sequence, in which the chain moves state at discrete time steps, gives a discrete-time Markov chain (DTMC). A continuous-time process is called a continuous-time Markov chain (CTMC). Markov processes are named in honor of the Russian mathematician Andrey Markov.

Markov chains have many applications as statistical models of real-world processes. They provide the basis for general stochastic simulation methods known as Markov chain Monte Carlo, which are used for simulating sampling from complex probability distributions, and have found application in areas including Bayesian statistics, biology, chemistry, economics, finance, information theory, physics, signal processing, and speech processing.

The adjectives Markovian and Markov are used to describe something that is related to a Markov process.`

type Link struct {
	count int
	node  *Node
}

type Node struct {
	value string
	links []Link
}

func main() {
	fmt.Println("*** Markov Chain Text Generator ***")

	prevWord := ""
	mappedWords := map[string]*Node{}
	for currWord := range strings.FieldsSeq(SEED_TEXT) {
		if prevWord == "" {
			prevWord = currWord
			continue
		}

		var prevNode *Node

		if mappedWords[prevWord] == nil {
			prevNode = &Node{
				value: prevWord,
				links: []Link{},
			}
			mappedWords[prevWord] = prevNode
		} else {
			prevNode = mappedWords[prevWord]
		}

		idx := slices.IndexFunc(prevNode.links, func(l Link) bool {
			return l.node.value == currWord
		})

		if idx != -1 {
			prevNode.links[idx].count += 1
		} else {
			prevNode.links = append(prevNode.links, Link{count: 1, node: &Node{
				value: currWord,
				links: []Link{},
			}})
		}

		prevWord = currWord
	}

	word := "and"
	next := mappedWords[word]
	for range 200 {
		fmt.Print(next.value + " ")
		if len(next.links) > 0 {
			next = next.links[0].node
		} else {
			break
		}
	}
}
