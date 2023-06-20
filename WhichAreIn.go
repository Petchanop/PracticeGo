/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   WhichAreIn.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/27 00:45:01 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/21 00:31:14 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

func InArray(array1 []string, array2 []string) []string {
	for i := 0; i < len(array1)-1; i++ {
		if (array1[i] > array1[i+1] && array1[i+1] != "") || (array1[i] == "" && array1[i+1] != "") {
			tmp := array1[i]
			array1[i] = array1[i+1]
			array1[i+1] = tmp
			i = -1
		}
	}
	count := 0
	CopyArray := array1
	var num int
	for i, dup1 := range CopyArray {
		for j, dup2 := range CopyArray {
			if i != j && dup1 == dup2 && len(array1) > 1 {
				count++
				if count == 1 {
					num = j
				}
			}
		}
		if count > 0 {
			array1 = append(CopyArray[:num], CopyArray[num+count:]...)
		}
		count = 0
	}
	track := false
	rArray := []string{}
	for _, words := range array1 {
		for _, subwords := range array2 {
			if strings.Contains(subwords, words) {
				if words != "" {
					track = true
					break
				}
			}
		}
		if track {
			rArray = append(rArray, words)
		}
		track = false
	}
	array1 = rArray
	if len(rArray) != 0 {
		return array1
	} else {
		return []string{}
	}
}

func compareSlice(s1, s2 []string) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, str := range s1 {
		if i < len(s2) && str != s2[i] {
			return false
		}
	}
	return true
}

func printSlice(s []string, t bool) {
	if t {
		green := color.New(color.FgGreen)
		green.Print("[")
		for i, val := range s {
			green.Print(val)
			if i < len(s)-1 {
				fmt.Print(" ")
			}
		}
		green.Print("]")
	} else {
		red := color.New(color.FgRed)
		red.Print("[")
		for i, val := range s {
			red.Print(val)
			if i < len(s)-1 {
				fmt.Print(" ")
			}
		}
		red.Print("]")
	}
}

type addTest struct {
	arg1, arg2, expected []string
}

func main() {
	var a1 = []string{"live", "arp", "strong"}
	var a2 = []string{"lively", "alive", "harp", "sharp", "armstrong"}

	var a3 = []string{"cod", "code", "wars", "ewar", "ar"}
	var a4 = []string{"coding", "vscode", "warslock", "anaewar", "jarar"}

	var a5 = []string{"arp"}
	var a6 = []string{"h", "sharp", "a", "b"}

	var a7 = []string{""}
	var a8 = []string{"h", "sharp", "a", "b"}

	var a9 = []string{"cod", "code", "wars", "ewar", "ar"}
	var a10 = []string{}

	var a11 = []string{"", "", "arp"}
	var a12 = []string{"", "harp", "arp"}

	var a13 = []string{"de", "ab", "de"}
	var a14 = []string{"dest", "delicious"}

	var a15 = []string{"ew", "rp", "te", "test", "ice", "ice", "ice", "ice", "ii", "wh", "wh", "wh"}
	var a16 = []string{"net", "nii", "", "mhiew", "harp", "wh", "mice"}

	var addTests = []addTest{
		addTest{a1, a2, []string{"arp", "live", "strong"}},
		addTest{a3, a4, []string{"ar", "cod", "code", "ewar", "wars"}},
		addTest{a5, a6, []string{"arp"}},
		addTest{a7, a8, []string{}},
		addTest{a9, a10, []string{}},
		addTest{a11, a12, []string{"arp"}},
		addTest{a13, a14, []string{"de"}},
		addTest{a15, a16, []string{"ew", "ice", "ii", "rp", "wh"}},
	}
	expected := color.New(color.FgMagenta, color.Bold)
	for i, test := range addTests {
		output := InArray(test.arg1, test.arg2)
		t := compareSlice(output, test.expected)
		fmt.Print(i)
		printSlice(output, t)
		expected.Print(" expected to equal ")
		printSlice(test.expected, t)
		fmt.Println()
	}
}
