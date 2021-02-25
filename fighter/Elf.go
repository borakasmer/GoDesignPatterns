package fighter

import "fmt"

type Elf struct {
	Id       FighterType
	ArrowHit int
	Life     int
	Mana     int
}

func NewElf(arrowHit, life, mana int) Elf {
	return Elf{ElfType, arrowHit, life, mana}
}

func (e Elf) GetId() FighterType {
	return e.Id
}

func (e Elf) Clone() Fighter {
	return NewElf(e.ArrowHit, e.Life, e.Mana)
}

func (e Elf) Fight() {
	fmt.Println("Elf Arrow Hit-Point:", e.ArrowHit)
}
