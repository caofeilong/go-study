package main

import (
	"fmt"
	"regexp"
)

func main() {
	//preversion := "v01.02.13"

	// value3`
	//	content := `
	//option2: value2
	//option3: value3
	//option4: value4
	//option5: value5
	//`
	//
	//reg := regexp.MustCompile(`(?m)(?P<key>\w+):\s+(?P<value>\w+)$`)
	//fmt.Println(reg.SubexpNames())
	//
	//for i, namei := range reg.SubexpNames() {
	//	fmt.Println(i, namei)
	//}

	//fmt.Println(re.s)
	//template := "nnnn${key}=$value\n"
	//result := []byte{}
	//
	//	//fmt.Println(reg.FindAllStringSubmatch(content, -1))
	//
	//	//
	//for _, subm := range reg.FindAllStringSubmatchIndex(content, -1) {
	//	// Apply the captured submatches to the template and append the output
	//	// to the result.
	//	fmt.Println(subm)
	//
	//	//result = reg.ExpandString(result, template, content, subm)
	//}

	//	//
	//	fmt.Println(string(result))
	//reg := regexp.MustCompile(`^v(?P<major>\d+).(?P<minor>\d)+.(?P<patch>\d)+$`)
	//fmt.Println(reg.ReplaceAllString("v1.1.1","v${major}.3.${patch}"))
	//fmt.Println(reg.MatchString("v1.1.1"))
	//fmt.Println(reg.ReplaceAllStringFunc("v1.1.1", func(s string) string {
	//	return "${major}"
	//}))

	//re := regexp.MustCompile(`a(x*)b`)
	//fmt.Println(re.ReplaceAllString("-ab-axxb-", "T"))
	//fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1"))
	//fmt.Println(re.ReplaceAllString("-ab-axxb-", "$1W"))
	//fmt.Println(re.ReplaceAllString("-ab-axxb-", "${1}W"))

	//fmt.Println(template[2:]...)
	//dst := []byte{}
	//var a = []interface{}
	//fmt.Println(string(template[1]))
	//for len(template) > 0 {
	//
	//	i := strings.Index(template, "$")
	//	fmt.Println(i)
	//	if i < 0 {
	//		break
	//	}
	//	template = template[i:]
	//	fmt.Println(string(template[1]))
	//	//dst = append(dst, )
	//
	//}

	//fmt.Println(utf8.DecodeRuneInString("cfl"))
	//fmt.Println(2*0 + 1)
	//fmt.Println(2*1 + 1)

	//updateName := "major"
	//preTag := "v0.0.0"

	//reg := regexp.MustCompile(`^v(?P<major>\d+)\.(?P<minor>\d+)\.(?P<patch>\d+)$`)
	//
	//
	//fmt.Println(string(reg.Find([]byte("v0.0.0"))))

	conetnt := []byte("cfl: cfl123 age: age")

	reg := regexp.MustCompile(`bdfasdfaaa\\\\w+:\s+\w+$`)


	for _, mt := range reg.FindAll(conetnt, -1) {
		fmt.Println(string(mt))
	}

	//var replaceIndex int
	//matchIndex := reg.FindStringSubmatchIndex(preTag)
	//fmt.Println(matchIndex)
	//newTag := []byte("v")
	//for i, iname := range reg.SubexpNames() {
	//	fmt.Println(iname, i)
	//	//if iname == "" {
	//	//	fmt.Println("nil")
	//	//}
	//	//if updateName == iname {
	//	//} else {
	//	newTag = append(newTag, preTag[matchIndex[2*i]:matchIndex[2*i+1]]...)
	//	//}
	//}
	////sort.SearchStrings()
	//
	//fmt.Println(newTag)
	//
	//fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
	//fmt.Println(strings.Title("her royal highness"))
	//fmt.Println(strings.Title("loud noises"))
	//fmt.Println(strings.Title("хлеб"))
	//
	//
	//fmt.Println(strings.ToTitle("her royal highness"))
	//fmt.Println(strings.ToTitle("loud noises"))
	//fmt.Println(strings.ToTitle("хлеб"))
	//
	//
	//fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}
