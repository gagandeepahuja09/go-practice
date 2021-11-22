package main

import "errors"

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone(m int) (ItemInfoGetter, error) {
	switch m {
	case White:
		// since *whitePrototype will only give the value and not the actual object
		// hence &newItem is a clone.
		newItem := *whitePrototype
		return &newItem, nil
	default:
		return nil, errors.New("Shirt model not recognized")
	}
}
