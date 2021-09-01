// Package gosorts
// Created by RTT.
// Author: teocci@yandex.com on 2021-Sep-01
package gosorts

import "sort"

func Heapsort(data sort.Interface) {
	heapSort(data, 0, data.Len())
}

func GuessIntShift(data Int64Interface, l int) uint {
	return guessIntShift(intWrapper{data}, l)
}

func SetQSortCutoff(i int) int {
	orig := qSortCutoff
	qSortCutoff = i
	return orig
}

func SetMinOffload(i int) int {
	orig := minOffload
	minOffload = i
	return orig
}

func Checking() bool {
	return true
}
