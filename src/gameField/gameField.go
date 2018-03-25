package gameField

import Bacterium "github.com/vehsamrak/genetics/src/bacterium"

const bacteriumInitialCount = 10

type gameField struct {
	microorganisms []microorganism
}

func New() gameField {
	return gameField{}
}

func (field *gameField) allBacterias() []microorganism {
	return field.microorganisms
}

func (field *gameField) addBacterium(microorganism microorganism) {
	field.microorganisms = append(field.microorganisms, microorganism)
}

func (field *gameField) removeBacterium(microorganism microorganism) {
	for bacteriumKey, bacterium := range field.microorganisms {
		if bacterium == microorganism {
			field.microorganisms[bacteriumKey] = field.microorganisms[len(field.microorganisms)-1]
			field.microorganisms = field.microorganisms[:len(field.microorganisms)-1]
			break
		}
	}
}

func (field *gameField) Populate() {
	for i := 0; i < bacteriumInitialCount; i++ {
		bacterium := Bacterium.New()
		field.addBacterium(bacterium)
	}
}
