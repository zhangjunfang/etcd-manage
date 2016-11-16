package etcd_lock

import (
	"sort"

	"context"
)

func (locker *EtcdLocker) Wait(key string) error {
	key = addPrefix(key)
	for {
		res, err := locker.Client.Get(context.Background(), key, &etcd.GetOptions{Recursive: true, Sort: true})
		if err != nil {
			if etcd.IsKeyNotFound(err) {
				break
			}
			return err
		}

		if len(res.Node.Nodes) == 0 {
			break
		}

		sort.Sort(res.Node.Nodes)
		currentLock := res.Node.Nodes[0]

		watcher := locker.kapi.Watcher(currentLock.Key, &etcd.WatcherOptions{AfterIndex: currentLock.CreatedIndex})
		_, err = watcher.Next(context.Background())
		if err != nil {
			return err
		}
	}

	return nil
}
