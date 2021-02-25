package main

import "fmt"

type food interface {
	getPrice() int
}
type CornRice struct {

}

func (r*CornRice) getPrice() int  {
	return 150
}
type riceWhitPotato struct {
	CornRice food
}

func (p*riceWhitPotato) getPrice() int  {
	ricePrice:=p.CornRice.getPrice()
	return ricePrice+400
}
type riceWithChicken struct {
	CornRice food
}

func (c*riceWithChicken)getPrice() int  {
	ricePrice:=c.CornRice.getPrice()
	return ricePrice+500
}

type riceWithFish struct {
	Rice food
}

func (f*riceWithFish)getPrice() int {
	ricePrice:=f.Rice.getPrice()
	return ricePrice+600
}
type noodles struct {
}

func (n*noodles)getPrice() int {
	return 300
}

type noodlesWithPotato struct {
	noodles food
}

func (n*noodlesWithPotato)getPrice() int {
	noodlesPrice:=n.noodles.getPrice()
	return noodlesPrice+500
}
type noodlesWithChicken struct {
	noodles food
}

func (n*noodlesWithChicken)getPrice() int  {
	noodlesPrice:=n.noodles.getPrice()
	return noodlesPrice+600
}

type noodlesWithFish struct {
	noodles food
}

func (n*noodlesWithFish)getPrice() int {
	noodlesPrice:=n.noodles.getPrice()
	return noodlesPrice+600
}
func main(){
	rice:= &CornRice{}
	riceWhitPotato:=&riceWhitPotato{
		rice,
	}
	riceWithChickenAndPotato:=&riceWithChicken{
		riceWhitPotato,
	}
	riceWithFish:=riceWithFish{
		rice,
	}
	fmt.Printf("Price of cornRice with pomato and chicken is %d\n",riceWithChickenAndPotato.getPrice())
	fmt.Printf("Price of cornRice with fish is %d\n",riceWithFish.getPrice())
	noodles:=&noodles{}
	noodlesWithPotato:=&noodlesWithPotato{
		noodles: noodles,
	}
	noodlesWithChickenAndPotato:=&noodlesWithChicken{
		noodles: noodlesWithPotato,
	}
	noodlesWithFish:=noodlesWithFish{
		noodles,
	}
	fmt.Printf("Price of noodles with potato and chicken is %d\n",noodlesWithChickenAndPotato.getPrice())
	fmt.Printf("Price of noodles with fish is %d\n",noodlesWithFish.getPrice())
}
