package main

import "testing"

func TestClone(t *testing.T) {
	shirtsCache := GetShirtsCloner()
	if shirtsCache == nil {
		t.Fatal("Received cache was nil")
	}

	item1, err := shirtsCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	if item1 == whitePrototype {
		t.Error("item1 cannot be equal to the white prototype")
	}

	shirt1, ok := item1.(*Shirt)

	if !ok {
		t.Fatal("Type assertion for shirt1 couldn't be done successfully")
	}

	shirt1.SKU = "fdsgfgf"

	item2, err := shirtsCache.GetClone(White)
	if err != nil {
		t.Error(err)
	}

	shirt2, ok := item2.(*Shirt)

	if !ok {
		t.Fatal("Type assertion for shirt2 couldn't be done successfully")
	}

	if shirt2 == shirt1 {
		t.Error("shirt1 and shirt2 cannot be the same object")
	}

	if shirt1.SKU == shirt2.SKU {
		t.Error("SKU of shirt1 and shirt2 cannot be the same")
	}

	t.Logf("LOG: shirt1 info: %s", shirt1.GetInfo())
	t.Logf("LOG: shirt2 info: %s", shirt2.GetInfo())

	t.Logf("LOG: The memory positions of the shirts are different %p != %p", &shirt1, &shirt2)
}
