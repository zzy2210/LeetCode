package simplifypath

import "strings"

/*
	71。简化路径
	把每个分割符都入栈
	遇到../ 就出栈 其他时候就入栈

	好像不需要单独想怎么分割，用split
*/

func simplifyPath(path string) string {
	pathStack := []string{}
	tmp := strings.Split(path, "/")
	for _, v := range tmp {

		if v == ".." {
			if len(pathStack) > 0 {
				pathStack = pathStack[:len(pathStack)-1]
			}
		} else if v != "" && v != "." {
			pathStack = append(pathStack, v)
		}
	}

	return "/" + strings.Join(pathStack, "/")
}
