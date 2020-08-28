package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

// Go 跟字符串相关的操作都在 strings 类里，跟 Java 不一样，Java 对字符串的相关操作都在 String 类里面
// 面向对象和非面向对象的区别
func main() {
	// Go 字符串像数字一样比较就可以，不需要像 Java 一样通过 equals
	s1 := "abc"
	s2 := "cba"
	fmt.Println(s1 == s2)
	fmt.Println(s1 >= s2)
	fmt.Println(s1 <= s2)
	fmt.Println(strings.Compare(s1, s2))
	fmt.Println(strings.EqualFold(s1, s2))

	s3 := " Java   Python Go"
	// 去空格的
	fmt.Println(strings.Fields(s3))
	// 是否分隔，类比于 Java 的 split 方法，这里的参数是一个 bool 类型返回值的函数，标识着何时应该分隔
	fmt.Println(strings.FieldsFunc(s3, unicode.IsSpace))

	// 纯正的 split 方法
	s4 := "Java,Python,Golang"
	fmt.Println(strings.SplitAfter(s4, ","))

	// 前后缀匹配
	s5 := "http://www.baidu.com"
	fmt.Println(strings.HasPrefix(s5, "http"))

	s6 := strings.Split("abcdefg", "")
	fmt.Println(join(s6, ","))

	s7 := "hello "
	fmt.Println(strings.Repeat(s7, 10))
}

func join(strs []string, seq string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	buffer := bytes.NewBufferString(strs[0])
	for _, str := range strs[1:] {
		buffer.WriteString(seq)
		buffer.WriteString(str)
	}
	return buffer.String()
}
