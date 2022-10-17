package main

import (
	"errors"
	"fmt"
)

type Collector interface {
	SetCore()
	SetMemory()
	SetGPU()
	SetBrand()
	GetComputer() ComputerB
}

type ComputerB struct {
	core   int32
	memory int32
	GPU    string
	brand  string
}

type HP struct {
	core   int32
	memory int32
	GPU    string
	brand  string
}

func (pc *HP) SetCore() {
	pc.core = 6
}

func (pc *HP) SetMemory() {
	pc.memory = 16
}

func (pc *HP) SetGPU() {
	pc.GPU = "AMD"
}

func (pc *HP) SetBrand() {
	pc.brand = "HP"
}

func (pc *HP) GetComputer() ComputerB {
	return ComputerB{
		core:   pc.core,
		memory: pc.memory,
		GPU:    pc.GPU,
		brand:  pc.brand,
	}
}

type ASUS struct {
	core   int32
	memory int32
	GPU    string
	brand  string
}

func (pc *ASUS) SetCore() {
	pc.core = 12
}

func (pc *ASUS) SetMemory() {
	pc.memory = 16
}

func (pc *ASUS) SetGPU() {
	pc.GPU = "NVIDEA"
}

func (pc *ASUS) SetBrand() {
	pc.brand = "ASUS"
}

func (pc *ASUS) GetComputer() ComputerB {
	return ComputerB{
		core:   pc.core,
		memory: pc.memory,
		GPU:    pc.GPU,
		brand:  pc.brand,
	}
}

func GetCollector(coltype string) (Collector, error) {
	switch coltype {
	case "HP":
		return &HP{}, nil
	case "ASUS":
		return &ASUS{}, nil
	default:
		fmt.Println("Unnown brand:", coltype)
		return nil, errors.New("Unnown brand")
	}
}

func (pc ComputerB) PrintInfo() {
	fmt.Println(
		" brand : ", pc.brand, "\n",
		"code  : ", pc.core, "\n",
		"GPU   : ", pc.GPU, "\n",
		"memory: ", pc.memory)
	fmt.Println()
}

type Factory struct {
	Collector Collector
}

func NewFactory(pc Collector) *Factory {
	return &Factory{Collector: pc}
}

func (factory *Factory) CreateComputer() ComputerB {
	factory.Collector.SetCore()
	factory.Collector.SetGPU()
	factory.Collector.SetBrand()
	factory.Collector.SetMemory()
	return factory.Collector.GetComputer()
}

func (factory *Factory) SetCollector(pc Collector) ComputerB {
	factory.Collector.SetCore()
	factory.Collector.SetGPU()
	factory.Collector.SetBrand()
	factory.Collector.SetMemory()
	return factory.Collector.GetComputer()
}

func main() {
	HpCollector, ok1 := GetCollector("HP")
	AsusCollector, ok2 := GetCollector("ASUS")

	if ok1 != nil || ok2 != nil {
		return
	}

	factory := NewFactory(HpCollector)
	HpComp := factory.CreateComputer()
	HpComp.PrintInfo()

	factory.SetCollector(AsusCollector)
	AsusComp := factory.CreateComputer()
	AsusComp.PrintInfo()

	ErrPc, ok3 := GetCollector("Apple")

	if ok3 != nil {
		return
	}

	factory.SetCollector(ErrPc)
}

//Паттерн Builder относится к порождающим паттернам уровня объекта.
//Паттерн Builder определяет процесс поэтапного построения сложного продукта.
//После того как будет построена последняя его часть, продукт можно использовать.
