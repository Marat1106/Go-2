package main

import "fmt"

type iShoe interface {
	setName(name string)
	getName()string
	setSize(size int)
	getSize()int
}
type shoe struct {
	name string
	size int
}

func (s*shoe) setName(name string)  {
	s.name=name
}
func (s*shoe) getName() string  {
	return s.name
}
func (s*shoe) setSize(size int)  {
	s.size=size
}
func (s*shoe)getSize()int  {
	return s.size

}

type nike struct {
	shoe
}

func newNike() iShoe {
	return &nike{
		shoe:shoe{
			name: "Nike",
			size: 40,
		},
	}
}


type adidas struct {
	shoe
}
func newAdidas() iShoe{
	return &adidas{
		shoe:shoe{
			name: "Adidas",
			size: 41,
		},
	}
}
func getShoe(shoeType string)(iShoe,error)  {
	if shoeType=="nike"{
		return newNike(),nil
	}
	if shoeType=="adidas" {
		return newAdidas(),nil
	}
	return nil,fmt.Errorf("There is not one")
}

func main()  {
	nike,_:=getShoe("nike")
	adidas,_:=getShoe("adidas")
	printDetails(nike)
	printDetails(adidas)
}
func printDetails(w iShoe)  {

	fmt.Printf("Name: %s",w.getName())
	fmt.Println()
	fmt.Printf("Size: %d",w.getSize())
	fmt.Println()
}
