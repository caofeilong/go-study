package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	//jsonStr := `["cfl","sx"]`
	//var name []string
	//json.Unmarshal([]byte(jsonStr), &name)
	//
	//fmt.Printf("%+v\n", name)
	//
	//jsonObj := []string{"1", "2"}
	//data, _ := json.Marshal(jsonObj)
	//fmt.Println(string(data))
	//strconv.
	//fmt.Printf("%+v\n", )

	//	const jsonStream = `
	//[
	//	{"name": "Ed", "Text": "Knock knock."},
	//	{"Name": "Sam", "Text": "Who's there?"},
	//	{"Name": "Ed", "Text": "Go fmt."},
	//	{"Name": "Sam", "Text": "Go fmt who?"},
	//	{"Name": "Ed", "Text": "Go fmt yourself!"}
	//]
	//`
	//	type Message struct {
	//		Name, Text string
	//	}
	//	dec := json.NewDecoder(strings.NewReader(jsonStream))
	//
	//	t, _ := dec.Token()
	//
	//	fmt.Printf("%T:%v\n", t, t)
	//
	//	for dec.More() {
	//		var m Message
	//		dec.Decode(&m)
	//		fmt.Printf("%+v\n", m)
	//	}
	//
	//	t, _ = dec.Token()
	//
	//	fmt.Printf("%T:%v\n", t, t)
	//for {
	//	var m Message
	//	if err := dec.Decode(&m); err == io.EOF {
	//		break
	//	} else if err != nil {
	//		//log.Fatal(err)
	//	}
	//	fmt.Printf("%s: %s\n", m.Name, m.Text)
	//}



	//	goodJSON := `{"example": 1,
	//  "name":2
	//}`
	//	//badJSON := `{"example":2:]}}`
	//
	//	var out bytes.Buffer
	//	json.Compact(&out, []byte(goodJSON))
	//	var s string
	//	out.WriteString(s)
	//	fmt.Println(s)
	//fmt.Println(json.Valid([]byte(goodJSON)), json.Valid([]byte(badJSON)))

	//os.Stdout.Write(b)

	//type Road struct {
	//	Name   string
	//	Number int
	//}
	//roads := []Road{
	//	{"Diamond Fork", 29},
	//	{"Sheep Creek", 51},
	//}
	//
	//b, _ := json.MarshalIndent(roads, "", "  ")
	//fmt.Println(string(b))
	//if err != nil {
	//	//log.Fatal(err)
	//}
	//
	////fmt.Printf("%T,%v\n", string(b), string(b))
	//var out bytes.Buffer
	//json.Indent(&out, b, "", "  ")
	//out.WriteTo(os.Stdout)

	//	const jsonStream = `
	//	{"Message": "Hello", "Array": [1, 2, 3], "Null": null, "Number": 1.234}
	//`
	//	dec := json.NewDecoder(strings.NewReader(jsonStream))
	//	for {
	//		t, err := dec.Token()
	//		if err == io.EOF {
	//			//fmt.Printf(dec.More())
	//			break
	//		}
	//		if err != nil {
	//			//log.Fatal(err)
	//		}
	//		fmt.Printf("%T: %v", t, t)
	//		if dec.More() {
	//			fmt.Printf(" (more)")
	//		}
	//		fmt.Printf("\n")
	//	}

	//jsonStr := `{
	// 	"name":"cfl",
	//	"age":10,
	//	"friend":[{"name":"sx",
	//		"age":10
	//	}]
	//}`
	type User struct {
		Name   string
		Age    int64
		Friend []User
	}
	cfl := User{
		Name: "cfl",
		Age:  10,
		Friend: []User{
			{Name: "sx", Age: 10},
		},
	}

	jsonStr, _ := json.Marshal(cfl)
	fmt.Printf("%v\n", string(jsonStr))

	//fmt.Println(json.Valid([]byte(jsonStr)))
	//json.Unmarshal([]byte(jsonStr), &cfl)
	//fmt.Printf("%+v", cfl)

}

//部分接口总结
//NewDecoder(ioreader)  新建一个decode 对象
//dec.Decode(指针) 讲json string 转换为 指针对象
//dec.More() 是否读取完 Todo
//dec.token() 获取json 内容
//io.EOF TODO 判断io是否读完
//json.Indent  缩进格式之后的代码 TODO 流相关知识
//json.Marshal  将对象转换为字符串
//json.MarshalIndent 转为带缩进的字符串
//json.RawMessage(msg) rawMessage 类型
//json.Unmarshal  将字符串转换为对象
//json.Valid 校验json字符串是否合法
