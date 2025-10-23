// This code is a sample benchmarking template picked up from
// https://brandur.org/fragments/bytes-buffer-vs-strings-builder
//
// Usage:
//
//	go test -bench=. -benchmem
package main

import (
	"bytes"
	"strings"
	"testing"
)

var fragments = []string{
	"This",
	"is a series of",
	"string fragments",
	"that will be concatenated together",
	"into a single larger string",
	"so that we can",
	"determine which of Go's various",
	"tools for doing this",
	"is most efficient.",
	"I found a few articles",
	"online",
	"but most were poorly cited",
	"or",
	"behind a Medium login wall",
	"or otherwise",
	"not of admirable quality.",
}

func BenchmarkBytesBuffer(b *testing.B) {
	for b.Loop() {
		var buf bytes.Buffer

		for _, fragment := range fragments {
			buf.WriteString(fragment)
			buf.WriteString(" ")
		}

		_ = buf.String()
	}
}

func BenchmarkConcatenateStrings(b *testing.B) {
	for b.Loop() {
		var str string

		for _, fragment := range fragments {
			str += fragment
			str += " "
		}
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for b.Loop() {
		var sb strings.Builder

		for _, fragment := range fragments {
			sb.WriteString(fragment)
			sb.WriteString(" ")
		}

		_ = sb.String()
	}
}
