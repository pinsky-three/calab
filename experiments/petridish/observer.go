package petridish

import "image"

func (pd *PetriDish) RegisterNewObserver(observer chan image.Image) {
	pd.observers = append(pd.observers, observer)
}

func (pd *PetriDish) GetMeanTPS() float64 {
	return pd.meanTPS
}
