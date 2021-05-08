package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

// 计算两个Point的距离
// p是方法的接收者, 当调用时,p是调用对象的一个副本, 名字不能用this or self
func (p Point) Distance(q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}


// 接收者为指针,使用指针传递调用者地址,可以更新接收者
// 习惯上,如果Point的任何一个方法使用指针接收者,所有Point方法都应该使用指针接收者
func (p *Point) ScaleBy(factor float64){
	p.X *= factor
	p.Y *= factor
}

// 接收者为命名类型,p为实参变量的副本,无法修改原实参
func (p Point) ScaleBy2(factor float64){
	p.X *= factor
	p.Y *= factor
}

// Path 是连接多个点的直线段, 即使是slice类型,也可以给它定义方法
type Path []Point

// Distance 方法返回路径的长度
func (p Path) Distance() float64 {
	sum := 0.0
	for i := 1; i < len(p); i++ {
		sum += p[i-1].Distance(p[i])
	}
	return sum
}

func main() {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println("distance: ", p.Distance(q))

	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println("pathSum: ", perim.Distance())

	p.ScaleBy(10)
	fmt.Printf("after scaleby: %f %f\n", p.X, p.Y)
	p.ScaleBy2(10)
	fmt.Printf("after scaleby2: %f %f\n", p.X, p.Y)
}
