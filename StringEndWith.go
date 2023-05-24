/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   StringEndWith.go                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/05/24 15:38:14 by npiya-is          #+#    #+#             */
/*   Updated: 2023/05/24 16:19:12 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func solution(str, ending string) bool {
	var i, j int
	if len(str) == 0 && len(ending) != 0 {
		return false
	} else if len(str) == 0 && len(ending) == 0 {
		return true
	}
	i = len(str) - 1
	fmt.Println("len str ", i)
	for j = len(ending) - 1; j >= 0; j-- {
		fmt.Printf("str[%d] : %c, ending[%d] : %c\n", i, str[i], j, ending[j])
		if str[i] != ending[j] {
			break
		}
		i--
	}
	if j < 0 {
		return true
	}
	return false
}

func main() {
	// fmt.Println("test with \"\"", "\"\"", solution("", ""))
	// fmt.Println("test with \" \"", "\"\"", solution(" ", ""))
	fmt.Println("test with \"abc\"", "\"c\"", solution("abc", "c"))
	// fmt.Println("test with \"a\"", "\"z\"", solution("a", "z"))
	// fmt.Println("test with \"anaana\"", "\"ana\"", solution("anaana", "ana"))
}
