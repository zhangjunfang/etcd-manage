package etcd_lock

import "context"

func (l *EtcdLock) Release() error {
	if l == nil {
		return errgo.New("nil lock")
	}
	_, err := l.Client..Delete(context.Background(), l.key)
	//_, err := l.kapi.Delete(&etcd.DeleteOptions{})
	if err != nil {
		return err
	}
	return nil
}
