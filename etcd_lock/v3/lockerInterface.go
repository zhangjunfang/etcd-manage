package v3

//分布式锁  获取锁接口方法
type Locker interface {
	Acquire(key string, ttl int64) (Lock, error)
	WaitAcquire(key string, ttl uint64) (Lock, error)
}

//分布式锁 锁释放接口
type Lock interface {
	Release() error
}
