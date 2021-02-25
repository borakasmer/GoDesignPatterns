package customer

import "fmt"

type Customer struct{
	Id string
}

func(c *Customer) Update(productName string){
	fmt.Printf("%s müşterisine %s ürünü için email gönderiliyor.\n", c.Id, productName)
}

func(c *Customer) GetID() string{
	return c.Id
}