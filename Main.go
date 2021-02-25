package main

import (
	motor "DesignPatterns/motorcycle"
	"fmt"
	"strings"
	"time"
)

//STRATEGY DESIGN PATTERN
/*
type SqlConnection struct {
	connectionString string
}

func (con SqlConnection) Connect() {
	fmt.Println(("Sql " + con.connectionString))
}

type OracleConnection struct {
	connectionString string
}

func (con OracleConnection) Connect() {
	fmt.Println("Oracle " + con.connectionString)
}

type IDBconnection interface {
	Connect()
}

type DBConnection struct {
	db IDBconnection
}

func (con DBConnection) DBConnect() {
	con.db.Connect()
}
*/

// PROTOTYPE DESIGN PATTERN
// go mod init DesignPatterns
/*
var dataHistoryList map[int]fig.Fighter

func loadCache() {
	soldier := fig.NewSoldier(30, 20, 5)
	dwarf := fig.NewDwarf(50, 40, 15)
	dataHistoryList = make(map[int]fig.Fighter)
	dataHistoryList[int(soldier.GetId())] = soldier
	dataHistoryList[int(dwarf.GetId())] = dwarf
}
func main() {
	loadCache()
	fighter := dataHistoryList[int(fig.SoldierType)]
	f, ok := fighter.(fig.Soldier)
	if ok {
		newFighter := f.Clone().(fig.Soldier)
		newFighter.GunHit = 45
		newFighter.Fight()
	}

	elf := dataHistoryList[int(fig.ElfType)]
	e, ok := elf.(fig.Elf)
	if ok {
		newElf := e.Clone()
		newElf.Fight()
	} else {
		elf := fig.NewElf(15, 30, 3)
		dataHistoryList[int(elf.GetId())] = elf
		elf.Fight()
	}
	elf2 := dataHistoryList[int(fig.ElfType)]
	e2, ok := elf2.(fig.Elf)
	if ok {
		newElf2 := e2.Clone().(fig.Elf)
		newElf2.ArrowHit = 35
		newElf2.Fight()
	}

	dwarf := dataHistoryList[int(fig.DwarfType)]
	d, ok := dwarf.(fig.Dwarf)
	if ok {
		newDwarf := d.Clone()
		newDwarf.Fight()
	} else {
		dwarf := fig.NewDwarf(50, 40, 15)
		dataHistoryList[int(fig.DwarfType)] = dwarf
		dwarf.Fight()
	}
}
*/
//OBSERVER DESIGN PATTERN
/*
	func main() {
		ps5 := cus.NewProduct("Sony Playstation 5")
		observer1 := &cus.Customer{Id: "bora@borakasmer.com"}
		observer2 := &cus.Customer{Id: "martin@martinfowler.com"}
		ps5.RegisterList([]cus.Observer{observer1, observer2})
		ps5.UpdateAvalabilitiy()
	}
*/

//STRATEGY DESIGN PATTERN
/*
func main() {
	sqlConnection := SqlConnection{connectionString: "Connection is connected"}
	con := DBConnection{db: sqlConnection}
	con.DBConnect()

	orcConnection := OracleConnection{connectionString: "Connection is connected"}
	con2:=DBConnection{db:orcConnection}
	con2.DBConnect()
}*/

//MEMENTO DESIGN PATTERN
func main() {
	mainData := motor.ServiceData{PersonalName: "Bora Kasmer", Model: motor.Honda, Price: 105000, Date: time.Now()}
	motor.UpdateDataChange(&mainData)
	fmt.Println(mainData)

	mainData = motor.ServiceData{PersonalName: "Kent Back", Model: motor.Kawasaki, Price: 170000, Date: time.Now()}
	motor.UpdateDataChange(&mainData)
	fmt.Println(mainData)

	mainData = motor.ServiceData{PersonalName: "Martin Fowler", Model: motor.Ducati, Price: 330000, Date: time.Now()}
	motor.UpdateDataChange(&mainData)
	fmt.Println(mainData)

	mainData = motor.Undo()
	fmt.Println("1 Geri Git: ", mainData)
	fmt.Println(strings.Repeat("-", 100))

	mainData = motor.Undo()
	fmt.Println("1 Geri Git: ", mainData)
	fmt.Println(strings.Repeat("-", 100))

	mainData = motor.Redo()
	fmt.Println("1 İleri Git: ", mainData)
	fmt.Println(strings.Repeat("-", 100))

	mainData = motor.Redo()
	fmt.Println("1 İleri Git: ", mainData)
}
