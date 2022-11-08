package main

import (
	"fmt"
)

type Race int

const (
	Mouse Race = iota
	Lynx
	Otter
	Squirrel
	// FlyingSquirrel
	Hedgehog
	Rabbit
	// Flying
	Crow
	Gull
	Owl
	Swan
	// Medium
	Fox
	Wolf
	// Large
	Elk
	Deer
	Moose
	Bear
)

type Size int

const (
	SizeTiny Size = iota
	SizeSmall
	SizeMedium
	SizeLarge
)

type Attributes int

const (
	Herbivore Attributes = 0b00_01
	Carnivore Attributes = 0b00_10
	Omnivore  Attributes = Herbivore | Carnivore
	Flying    Attributes = 0b01_00
	Swimming  Attributes = 0b10_00
)

type RaceAttributes struct {
	Name string
	Size
	Speed      int
	Attributes Attributes
}

var Races = []RaceAttributes{
	{"Mouse", SizeTiny, 10, Herbivore},
	{"Lynx", SizeSmall, 20, Carnivore},
	{"Otter", SizeSmall, 20, Herbivore},
	{"Squirrel", SizeTiny, 20, Herbivore},
	{"Hedgehog", SizeTiny, 20, Herbivore},
	{"Rabbit", SizeSmall, 20, Herbivore},
	{"Crow", SizeTiny, 20, Herbivore},
	{"Gull", SizeSmall, 20, Omnivore | Flying},
	{"Owl", SizeSmall, 20, Carnivore | Flying},
	{"Swan", SizeMedium, 20, Herbivore | Flying},
	{"Fox", SizeSmall, 20, Carnivore},
	{"Wolf", SizeMedium, 20, Carnivore},
	{"Elk", SizeLarge, 20, Herbivore},
	{"Deer", SizeLarge, 20, Herbivore},
	{"Moose", SizeLarge, 20, Herbivore},
	{"Bear", SizeLarge, 20, Omnivore},
}

func GetRaceID(name string) (int, error) {
	for i, raceAttr := range Races {
		if raceAttr.Name == name {
			return i, nil
		}
	}
	return 0, fmt.Errorf("cannot find race with name %s", name)
}

type Currency int

const (
	Acorns Currency = iota
	Mushrooms
	Shinies
	Gems
)

type Wallet []float32

func NewWallet() Wallet {
	return Wallet{0, 0, 0, 0}
}

type Actor struct {
	ID     int
	Name   string
	Wallet Wallet
}
