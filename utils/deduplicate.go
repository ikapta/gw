package utils

import "sort"

/**
 * 字符串数组去重
 * inspired from https://github.com/golang/go/wiki/SliceTricks#in-place-deduplicate-comparable
 */
func DeduplicateStr(arr []string) []string {
	sort.Strings(arr)

	length := len(arr)
	if length == 0 {
		return arr
	}

	j := 0
	for i := 1; i < length; i++ {
		if arr[i] != arr[j] {
			j++
			if j < i {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}

	return arr[:j+1]
}

/**
 * []int去重
 */
func DeduplicateInt(in []int) []int {
	sort.Ints(in)

	j := 0
	for i := 1; i < len(in); i++ {
		if in[j] == in[i] {
			continue
		}
		j++
		// preserve the original data
		// in[i], in[j] = in[j], in[i]
		// only set what is required
		in[j] = in[i]
	}

	return in[:j+1]
}