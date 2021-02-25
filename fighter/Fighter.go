package fighter

import "fmt"

var dataHistoryList map[int]Fighter

type FighterType int

const (
	SoldierType FighterType = 1
	ElfType                 = 2
	DwarfType               = 3
)

type Fighter interface {
	GetId() FighterType
	Fight()
	Clone() Fighter
}

type Soldier struct {
	Id     FighterType
	GunHit int
	Life   int
	Mana   int
}

func NewSoldier(gunHit, life, mana int) Soldier {
	soldier := Soldier{SoldierType, gunHit, life, mana}
	return soldier
}

func (s Soldier) GetId() FighterType {
	return s.Id
}

func (s Soldier) Clone() Fighter {
	return NewSoldier(s.GunHit, s.Life, s.Mana)
}

func (s Soldier) Fight() {
	fmt.Println("Soldier Gun Hit-Points:", s.GunHit)
}
