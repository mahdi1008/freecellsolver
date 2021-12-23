package freecellsolver

import "sync"

type valueMap struct {
	sync.RWMutex
	m map[string]int
}

func (vm *valueMap) Get(s string) int {
	vm.Lock()
	defer vm.Unlock()
	return vm.m[s]
}

var ValueMap valueMap

type inverseValueMap struct {
	sync.RWMutex
	m map[int]string
}

func (ivm *inverseValueMap) Get(i int) string {
	ivm.Lock()
	defer ivm.Unlock()
	return ivm.m[i]
}

var InverseValueMap inverseValueMap

type suitMap struct {
	sync.RWMutex
	m map[string]string
}

func (sm *suitMap) Get(s string) string {
	sm.Lock()
	defer sm.Unlock()
	return sm.m[s]
}

var SuitMap suitMap

type seenMap struct {
	sync.RWMutex
	m map[uint]bool
}

func (sm *seenMap) Set(u uint) {
	sm.Lock()
	defer sm.Unlock()
	sm.m[u] = true
}

func (sm *seenMap) Get(u uint) bool {
	sm.RLock()
	defer sm.RUnlock()
	return sm.m[u]
}

var SeenMap seenMap

func initMaps() {
	ValueMap = valueMap{
		m: map[string]int{
			"A": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "X": 10, "J": 11, "Q": 12, "K": 13,
		},
	}
	InverseValueMap = inverseValueMap{
		m: map[int]string{
			1: "A", 2: "2", 3: "3", 4: "4", 5: "5", 6: "6", 7: "7", 8: "8", 9: "9", 10: "X", 11: "J", 12: "Q", 13: "K",
		},
	}
	SuitMap = suitMap{
		m: map[string]string{
			"s": "spades", "d": "diamonds", "h": "hearts", "c": "clubs",
		},
	}
	SeenMap = seenMap{
		m: map[uint]bool{},
	}
}
