package customer

type Observer interface {
	Update(string)
	GetID()string
}
