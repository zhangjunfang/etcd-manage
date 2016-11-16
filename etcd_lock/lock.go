package etcd_lock

import "fmt"

func T() {

	l, err := etcd_lock.Acquire(client, "/name", 60)
	if lockErr, ok := err.(*etcd_lock.Error); ok {
		// Key already locked
		fmt.Println(lockErr)
		return
	} else if err != nil {
		// Communication with etcd has failed or other error
		panic(err)
	}

	// It's ok, lock is granted for 60 secondes

	// When the opration is done we release the lock
	err = l.Release()
	if err != nil {
		// Something wrong can happen during release: connection problem with etcd
		panic(err)
	}

}
