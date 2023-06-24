/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ConsonantValue.go                                  :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/21 15:42:25 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/23 19:03:51 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import (
	"bytes"
	"fmt"
)

func solve(str string) uint {
	vowels := "aeiou"
	var sum uint
	var presum uint
	for _, val := range str {
		if !(bytes.IndexByte([]byte(vowels), byte(val)) >= 0) && val != ' ' {
			sum += uint(val - ('a' - 1))
			if sum > presum {
				presum = sum
			}
		} else {
			sum = 0
		}
	}
	return presum
}

func main() {
	fmt.Println(solve("test solve"))
}
