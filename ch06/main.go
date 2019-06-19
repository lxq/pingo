// struct 自定义类型练习.
// 6.2 节
// @date 2019/6/19
// @author Allen Lin
// @email xqlin@qq.com

package main

import (
	"fmt"
	"io"
	"strings"
)

// Part 部门信息.
type Part struct {
	Id   int // 具名字段（聚合）
	Name string
}

// Lower 部门名称转小写.
// 指针接收器一般用于复杂自定义类型或需要更改数据的情况.
func (p *Part) Lower() {
	p.Name = strings.ToLower(p.Name)
}

// String 实现字符串转化.
// 值接收器一般用于简单自定义类型或无需更改类型值的情况.
func (p Part) String() string {
	return fmt.Sprintf("<<编号：%d, 姓名：%q>>", p.Id, p.Name)
}

// Exchanger 接口练习.
type Exchanger interface {
	Exchange()
}

// StringPair 字符串对.
type StringPair struct {
	first, second string
}

// Exchange 实现了Exchanger接口.
func (sp *StringPair) Exchange() {
	sp.first, sp.second = sp.second, sp.first
}

// String 实现了fmt.Stringer接口
func (sp StringPair) String() string {
	return fmt.Sprintf("%q + %q", sp.first, sp.second)
}

// Read 实现io.Reader接口.
// 把StringPair读取到data中并删除原来内存，返回已读取的字节数和错误。
func (sp *StringPair) Read(data []byte) (n int, err error) {
	if sp.first == "" && sp.second == "" {
		return 0, io.EOF
	}
	if sp.first != "" {
		n = copy(data, sp.first)
		sp.first = sp.first[n:]
	}
	if sp.second != "" && n < len(data) {
		m := copy(data[n:], sp.second)
		sp.second = sp.second[m:]
		n += m
	}

	return n, err
}

// 读取size个字节.
// 不关注r的具体类型，只需要满足接口即可.
func toBytes(r io.Reader, size int) ([]byte, error) {
	data := make([]byte, size)
	n, err := r.Read(data)
	if nil != err {
		return data, err
	}
	return data[:n], nil // 清除无用字节.
}

// Point 基本类型自定义.
type Point [2]int

// Exchange 实现了Exchanger接口.
func (p *Point) Exchange() {
	p[0], p[1] = p[1], p[0]
}

// 展示了接口的鸭子类型.
func exchage(ex ...Exchanger) {
	for _, e := range ex {
		e.Exchange()
	}
}

func main() {
	p := Part{1, "林秀全"}
	fmt.Println(p)
	p.Lower()
	fmt.Printf("%s\n", p) // 这里会直接调用Part.String()方法转化

	sp := StringPair{"weetgo", "COM"}
	pt := Point{123, 321}
	fmt.Println(sp)
	fmt.Println(pt)
	sp.Exchange()
	pt.Exchange()
	fmt.Println(sp)
	fmt.Println(pt)
	exchage(&sp, &pt) // 这里不能传值，必须传指针.
	fmt.Println(sp)
	fmt.Println(pt)

	const size = 16
	tj := &StringPair{"Tom", "Jerry"}
	md := StringPair{"Mouse", "Duck"}
	for _, r := range []io.Reader{tj, &md} {
		data, err := toBytes(r, size)
		if nil != err {
			fmt.Println(err)
		}
		fmt.Printf("%q\n", data)
	}
}
