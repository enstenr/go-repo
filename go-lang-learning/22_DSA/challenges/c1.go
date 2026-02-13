package main

import (
	"fmt"
	"strings"
)

func DNAStrand(dna string) string {
	// replace A with T andC with G vice version in each char in string
	dnamap := map[string]string{}
	dnamap["A"] = "T"
	dnamap["T"] = "A"
	dnamap["G"] = "C"
	dnamap["C"] = "G"

	str := []byte(dna)
	var builder strings.Builder
	for _, c := range str {
		builder.WriteString(dnamap[string(c)])
	}
	return builder.String()
}

var dnareplacer *strings.Replacer = strings.NewReplacer(
	"A", "T",
	"T", "A",
	"G", "C",
	"C", "G")

func main() {
	fmt.Println(DNAStrand("TTGC"))
	fmt.Println(dnareplacer.Replace("TGC"))
}
