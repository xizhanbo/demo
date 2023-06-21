package main

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		if !strings.HasPrefix(strs[i], prefix) { // 不断更新前缀字符串
			prefix = prefix[:len(prefix)-1]
			if prefix == "" { // 如果前缀字符串为空，则直接返回空字符串
				return ""
			}
		}
	}
	return prefix
}

//时间复杂度O(1)
