package main

import (
	"fmt"
	"sort"
)

/*
	编写一个函数来查找字符串数组中的最长公共前缀。
	如果不存在公共前缀，返回空字符串 ""。


	示例 1：
	输入：strs = ["flower","flow","flight"]
	输出："fl"

	示例 2：
	输入：strs = ["dog","racecar","car"]
	输出：""
	解释：输入不存在公共前缀。

	来源：力扣（LeetCode）
	链接：https://leetcode.cn/problems/longest-common-prefix
	著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/
func main()  {
	fmt.Println(longestCommonPrefix([]string{"flower","flow","flight"}))
}

func longestCommonPrefix(strs []string) string {
	sort.Slice(strs, func(i, j int) bool {
		return 	len(strs[i])<len(strs[j])
	})
	base:=strs[0]
	for i := 0; i < len(strs); i++ {
		for i := 0; i < len(strs[0]); i++ {
			for _, str := range strs[1:] {
				if i >= len(str) || str[i] != base[i] {
					return base[:i]
				}
			}

		}
	}
	return base
}