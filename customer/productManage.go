package customer

type ProductManage interface {
	Register(Observer Observer)
	RegisterList(Observer []Observer)
	Unregister(Observer Observer) []Observer
	NotifyAll()
}
