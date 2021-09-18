package main

import "errors"

type ShirtsCache struct{}

func (s *ShirtsCache) GetClone() (ItemInfoGetter, error) {
	return nil, errors.New("not yet implemented")
}
