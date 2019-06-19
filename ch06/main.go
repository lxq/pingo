// struct 自定义类型练习.
// 6.2 节
// @date 2019/6/19
// @author Allen Lin
// @email xqlin@qq.com

package main

import (
	"fmt"
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

func (sp StringPair) String() string {
	return fmt.Sprintf("%q + %q", sp.first, sp.second)
}

// Point 基本类型自定义.
type Point [2]int

// Exchange 实现了Exchanger接口.
func (p *Point) Exchange() {
	p[0], p[1] = p[1], p[0]
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

}
