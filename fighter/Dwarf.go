package fighter
import "fmt"

type Dwarf struct {
	Id       FighterType
	AxePoint int
	Life     int
	Mana     int
}

func NewDwarf(arrowHit, life, mana int) Dwarf {
	return Dwarf{DwarfType, arrowHit, life, mana}
}

func (e Dwarf) GetId() FighterType {
	return e.Id
}

func (e Dwarf) Clone() Fighter {
	return NewDwarf(e.AxePoint, e.Life, e.Mana)
}

func (e Dwarf) Fight() {
	fmt.Println("Dwarf Axe Hit-Point:", e.AxePoint)
}
