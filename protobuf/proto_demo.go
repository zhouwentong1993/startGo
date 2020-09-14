package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
)

func main() {
	s := &SearchResult{
		QueryResult: "aaa查询数据结果",
		Row:         10,
		Pages:       1110,
		Map: map[string]uint32{},
		Rep: nil,
	}
	bytes, err := proto.Marshal(s)
	if err != nil {
		return
	}
	newS := &SearchResult{}
	err = proto.Unmarshal(bytes, newS)
	if err != nil {
		return
	}
	fmt.Println(newS)

}
