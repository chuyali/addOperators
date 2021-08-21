package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	str := "202"
	var target int64 = 2
	if !isNum(str) {
		return
	}
	addOperators(str, target)
}

func addOperators(str string, target int64) {
	res := []string{}
	backtrack(str, "", target, 0, 0, 0, &res)
	fmt.Println(res)
}
func backtrack(str, path string, target, sum, start, pre int64, res *[]string) {
	strLen := int64(len(str))
	if start == strLen && sum == target {
		*res = append(*res, path)
		return
	} else if start >= strLen {
		return
	} else {
		for i := start; i < strLen; i++ {
			if string(str[start]) == "0" && i > start {
				break
			}
			curStr := string(str[start : i+1])
			cur, _ := strconv.ParseInt(curStr, 10, 64)
			if start == 0 {
				backtrack(str, path+curStr, target, sum+cur, i+1, cur, res)
			} else {
				//做加法 +cur
				backtrack(str, path+"+"+curStr, target, sum+cur, i+1, cur, res)

				//做减法 -cur
				backtrack(str, path+"-"+curStr, target, sum-cur, i+1, -cur, res)

				//做乘法 *cur
				backtrack(str, path+"*"+curStr, target, sum-pre+pre*cur, i+1, pre*cur, res)

				//做除法 /cur
				if cur == 0 {
					break
				}
				backtrack(str, path+"/"+curStr, target, sum-pre+pre/cur, i+1, pre/cur, res)
			}
		}
	}
}
func isNum(str string) bool {
	pattern := "^\\d+$"
	result, _ := regexp.MatchString(pattern, str)
	return result
}
