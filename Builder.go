package main

import "fmt"

var i int
type Phone struct {
	id int
	name,country string
	OnMarketYear,storage int

}
type PhoneBuilder struct {
	Phone *Phone
}

func (p*PhoneBuilder)Build() *Phone  {
	i++
	p.Phone.id=i
	return p.Phone
}

type PhoneNameBuilder struct {
	PhoneBuilder
}

func (p*PhoneBuilder)  StringInfo() *PhoneNameBuilder  {
	return &PhoneNameBuilder{*p}
}

type PhoneInfoBuilder struct {
	PhoneBuilder
}

func (p*PhoneBuilder) IntInfo() *PhoneInfoBuilder  {
	return &PhoneInfoBuilder{*p}
}
func NewPhoneBuilder()*PhoneBuilder  {
	return &PhoneBuilder{&Phone{}}
}
func (pb*PhoneNameBuilder)Names(name string)*PhoneNameBuilder  {
	pb.Phone.name=name
	return pb
}
func (pb*PhoneNameBuilder) Countries(country string) *PhoneNameBuilder {
	pb.Phone.country=country
	return pb

}
func (ib*PhoneInfoBuilder) Storages(storage int)*PhoneInfoBuilder {
	ib.Phone.storage=storage
	return ib
}
func (ib*PhoneInfoBuilder) OnMarketYears(OnMarketYear int) *PhoneInfoBuilder {
	ib.Phone.OnMarketYear=OnMarketYear
	return ib

}
func main()  {
	ne:=NewPhoneBuilder()
	ne.StringInfo().Names("OPPO").Countries("China")
	ne.IntInfo().OnMarketYears(2017).Storages(128)
	var c =ne.Build()
	fmt.Println(c)

	aa:=NewPhoneBuilder()
	aa.StringInfo().Names("Iphone").Countries("America")
	aa.IntInfo().OnMarketYears(2008).Storages(32)
	var e =aa.Build()
	fmt.Println(e)

	bb:=NewPhoneBuilder()
	bb.StringInfo().Names("Samsung").Countries("Korea")
	bb.IntInfo().OnMarketYears(2010).Storages(64)
	var w =bb.Build()
	fmt.Println(w)
}
