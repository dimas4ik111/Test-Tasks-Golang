// Реализовать конкурентную запись данных в map.

package main

import (
	"fmt"
	"sync"
)

type myMap struct {
	mx sync.RWMutex
	mp map[string]int
}

func newMyMap() *myMap {
	return &myMap{
		mp: make(map[string]int),
	}
}

func (m *myMap) Load(key string) (int, bool) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	val, ok := m.mp[key]
	return val, ok
}

func (m *myMap) Store(key string, value int) {
	m.mx.Lock()
	defer m.mx.Unlock()
	m.mp[key] = value
}

func (m *myMap) Delete(key string) {
	m.mx.RLock()
	defer m.mx.RUnlock()
	delete(m.mp, key)
}

// func (obj *DefenseMap) Delete(key interface{}) {
// 	obj.Lock()
// 	defer obj.Unlock()
// 	delete(obj.m, key)
// }

func printCaseMap(v int, ok bool) {
	if ok {
		fmt.Println(v)
	} else {
		fmt.Println("XD :(")
	}
}

func main() {
	exMap := newMyMap()
	exMap.Store("ex1", 5)
	exMap.Store("ex2", 66)

	// for v := range exMap.mp {
	// 	fmt.Println(v)
	// }

	v, ok := exMap.Load("ex1")
	printCaseMap(v, ok)
	exMap.Delete("ex1")
	v, ok = exMap.Load("ex1")
	printCaseMap(v, ok)
	// printCaseMap(exMap.Load("school"))
	// printCaseMap(exMap.Load("ex1"))
	// printCaseMap(exMap.Load("ex2"))
}
