package single

import "sync"

var lock = sync.Mutex{}

type Single struct{}

var _single *Single

func GetInstance() *Single {
	if _single == nil {
		lock.Lock()
		defer lock.Unlock()
		_single = &Single{}
	}
	return _single
}
