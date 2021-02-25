package customer

import (
	"fmt"
	"strings"
)

type product struct {
	oberverList []Observer
	name        string
	inStock     bool
}

func NewProduct(_name string) *product {
	return &product{name: _name}
}

func (p *product) NotifyAll() {
	for _, observer := range p.oberverList {
		observer.Update(p.name)
	}
}

func (p *product) UpdateAvalabilitiy() {
	fmt.Printf("Product %s artık stoklardadır\n", p.name)
	fmt.Printf(strings.Repeat("-", 100))
	fmt.Printf("\n")
	p.inStock = true
	p.NotifyAll()
}

func (p *product) Register(o Observer) {
	p.oberverList = append(p.oberverList, o)
}

func (p *product) RegisterList(olist []Observer) {
	for _, o := range olist {
		p.oberverList = append(p.oberverList, o)
	}
}

func RemoveFromList(observerList []Observer, removeObserver Observer) []Observer {
	for index, observer := range observerList {
		if removeObserver.GetID() == observer.GetID(){
			return append(observerList[:index], observerList[index+1:]...)
		}
	}
	return observerList
}

func (p *product) Unregister(o Observer) []Observer {
	p.oberverList = RemoveFromList(p.oberverList, o)
	return p.oberverList
}
