package motorcycle

import "time"

type MotorType int

const (
	Kawasaki MotorType = iota
	Honda
	Suzuki
	Ducati
)

type ServiceData struct {
	Id           int
	PersonalName string
	Model        MotorType
	Price        float32
	Date         time.Time
}

var counter *int
var dataHistoryList map[int]ServiceData

func UpdateDataChange(model *ServiceData) {
	if dataHistoryList == nil {
		dataHistoryList = make(map[int]ServiceData)
		counter = new(int)
	}
	*counter++
	model.Id = *counter
	dataHistoryList[*counter] = *model
}

func Undo() ServiceData {
	index := *counter - 1
	if index >= 1 {
		*counter = index
		return dataHistoryList[index]
	}
	return dataHistoryList[*counter]
}

func Redo() ServiceData {
	index := *counter + 1
	if index <= len(dataHistoryList) {
		*counter = index
		return dataHistoryList[*counter]
	}
	return dataHistoryList[*counter]
}
