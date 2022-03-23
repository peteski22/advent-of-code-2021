package main

import (
	"fmt"
	"os"
	"strconv"
)

func base2StringToInt(s string) int64 {
	var i int64
	var err error
	if i, err = strconv.ParseInt(s, 2, 64); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return i
}

func getDay03Part1Result(vs []string) int64 {
	const MaxInputLength = 12
	var gs string
	var es string

	for i := 0; i < MaxInputLength; i++ {
		var bits string
		for _, v := range vs {
			bits += string(v[i])
		}
		g, e := getGammaAndEpsilon(bits)
		gs += g
		es += e
	}

	gamma := base2StringToInt(gs)
	epsilon := base2StringToInt(es)

	return gamma * epsilon
}

func getGammaAndEpsilon(vs string) (gamma string, epsilon string) {
	var ones int
	var zeros int

	for i := 0; i < len(vs); i++ {
		v, err := strconv.Atoi(string(vs[i]))
		if err != nil {
			panic(fmt.Sprint("Unable to convert character to string -> int", vs[i]))
		}
		if v == 0 {
			zeros++
		} else if v == 1 {
			ones++
		} else {
			panic(fmt.Sprint("Unexpected non-binary value", v))
		}
	}

	if ones > zeros {
		return "1", "0"
	} else {
		return "0", "1"
	}
}
