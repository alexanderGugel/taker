package taker

import "sync"

// Race runs the supplied tasks at the same time.
// If any of the tasks returns an error, the error will be returned and
// subsequent errors of the remaining tasks will be ignored.
// Equivalent to Promise.race
func Race(tasks ...Task) error {
	var finErr error

	// Used for blocking final return.
	var finMutex sync.Mutex

	// Used for synchronizing access to done counter.
	var updateMutex sync.Mutex
	done := 0

	if len(tasks) > 0 {
		finMutex.Lock()
	}

	for _, t := range tasks {
		go func(t Task) {
			err := t.Run()
			updateMutex.Lock()
			defer updateMutex.Unlock()
			done++
			if finErr == nil && (done == len(tasks) || err != nil) {
				finErr = err
				finMutex.Unlock()
			}
		}(t)
	}

	finMutex.Lock()
	defer finMutex.Unlock()

	return finErr
}
