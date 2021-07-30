package petridish

import "image"

func (pd *PetriDish) RegisterNewObserver(observer chan image.Image) {
	pd.observers = append(pd.observers, observer)
}
