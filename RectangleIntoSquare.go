/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   RectangleIntoSquare.go                             :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: npiya-is <marvin@42.fr>                    +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2023/06/24 19:42:24 by npiya-is          #+#    #+#             */
/*   Updated: 2023/06/24 21:15:12 by npiya-is         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package main

import "fmt"

func SquaresInRect(lng int, wdth int) []int {
	result := []int{}
	if lng == wdth {
		return nil
	}
	var high, low int
	if lng > wdth {
		high, low = lng, wdth
	} else {
		high, low = wdth, lng
	}
	for high > 0 {
		sq := low
		if sq > 0 {
			result = append(result, sq)
		}
		if low != 1 {
			high -= low
		}
		if high < low {
			tmp := high
			high = low
			low = tmp
		} else if low == 1 || high == 1 {
			high--
		}
		if low == 0 {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(SquaresInRect(5, 3))
}
