package main

import "fmt"

type Visitor interface {
	VisitSquare(*square)
	VisitCircle(*circle)
	VisitTriangle(*triangle)
}

type shape interface {
	getType() string
	accept(Visitor)
}

func (s *square) getType() string {
	return "square"
}

func (c *circle) getType() string {
	return "circle"
}

func (t *triangle) getType() string {
	return "triangle"
}

type areaCalc struct {
	area int
}

func (a *areaCalc) VisitCircle(c *circle) {
	a.area = 1
	fmt.Println("area circle =", a.area)
}

func (a *areaCalc) VisitSquare(s *square) {
	a.area = 2
	fmt.Println("area square =", a.area)
}

func (a *areaCalc) VisitTriangle(t *triangle) {
	a.area = 3
	fmt.Println("area triangle =", a.area)
}

type middleCoordinats struct {
	x int
	y int
}

func (a *middleCoordinats) VisitCircle(c *circle) {
	a.x = 0
	a.y = 0
	fmt.Println("middle cicrle (x, y) = ", a.x, a.y)
}

func (a *middleCoordinats) VisitSquare(s *square) {
	a.x = 1
	a.y = 1
	fmt.Println("middle square (x, y) = ", a.x, a.y)
}

func (a *middleCoordinats) VisitTriangle(t *triangle) {
	a.x = 2
	a.y = 3
	fmt.Println("middle triangle (x, y) = ", a.x, a.y)
}

type square struct {
	side int
}

func (s *square) accept(v Visitor) {
	v.VisitSquare(s)
} 

type circle struct {
	r int
}

func (c *circle) accept(v Visitor) {
	v.VisitCircle(c)
}

type triangle struct {
	side int
}

func (t *triangle) accept(v Visitor) {
	v.VisitTriangle(t)
}

func main() {
	sq := &square{side: 1}
	cr := &circle{r: 1}
	tr := &triangle{side: 1}

	fmt.Println("sq = ", sq.getType())
	fmt.Println("cr = ", cr.getType())
	fmt.Println("tr = ", tr.getType())

	fmt.Println()

	mid := &middleCoordinats{}
	sq.accept(mid)
	cr.accept(mid)
	tr.accept(mid)

	fmt.Println()

	ar := &areaCalc{}
	sq.accept(ar)
	cr.accept(ar)
	tr.accept(ar)
}
