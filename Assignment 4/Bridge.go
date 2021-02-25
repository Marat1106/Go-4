package main

import "fmt"

type BrushPen interface {
	write()
	WriteByHand(Write)
}
type Write interface {
	WriteByMachine()
}

type BigBrushPen struct {
	writer Write
}

func ( b*BigBrushPen) write() {
	fmt.Println("Big brush is using,by machine")
	b.writer.WriteByMachine()
}

func (b *BigBrushPen) WriteByHand(w Write) {
	fmt.Println("Big brush is using,by hand")
	b.writer =w
}
type SmallBrushPen struct {
	writer Write
}

func (s *SmallBrushPen) write() {
	fmt.Println("Small brush is using ,by machine")
	s.writer.WriteByMachine()
}

func (s *SmallBrushPen) WriteByHand(w Write) {
	fmt.Println("small brush is using,by hand")
	s.writer = w
}
type american struct {
}

func (a *american) WriteByMachine() {
	fmt.Println("You're using american's brush pen")
}
type chinese struct {
}

func (c *chinese) WriteByMachine() {
	fmt.Println("You're using chinese's brush pen")
}
func main()  {
	aBrushPen:=&american{}
	cBrushPen:=&chinese{}
	BrushPen:=&BigBrushPen{}
	BrushPen.WriteByHand(aBrushPen)
	BrushPen.write()
	fmt.Println()

	secondBrushPen := &SmallBrushPen{}
	secondBrushPen.WriteByHand(cBrushPen)
	secondBrushPen.write()
	fmt.Println()
}
