package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/runzhen/golang/protobuf/example"
	"log"
)

func main() {
	// 为 AllPerson 填充数据
	p1 := example.Person{
		Id:   *proto.Int32(1),
		Name: *proto.String("hello, world"),
	}

	p2 := example.Person{
		Id:   2,
		Name: "gopher",
	}

	all_p := example.AllPerson{
		Per: []*example.Person{&p1, &p2},
	}

	// 对数据进行序列化
	data, err := proto.Marshal(&all_p)
	if err != nil {
		log.Fatalln("Mashal data error:", err)
	}
	fmt.Printf("[+] Marshal done.\n")

	// 对已经序列化的数据进行反序列化
	var target example.AllPerson
	err = proto.Unmarshal(data, &target)
	if err != nil {
		log.Fatalln("UnMashal data error:", err)
	}
	fmt.Printf("[+] UnMarshal done.\n")

	for i := 0; i < len(target.Per); i++ {
		println(target.Per[i].Name) // 打印第一个 person Name 的值进行反序列化验证
	}
}
