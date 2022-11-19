


```go

type Locker interface {
	Lock()
	Unlock()
}


// 互斥锁
type Mutex struct {
	// contains filtered or unexported fields
}
func (m *Mutex) Lock()
func (m *Mutex) TryLock() bool
func (m *Mutex) Unlock()


// 读写锁
type RWMutex struct {
	// contains filtered or unexported fields
}
func (rw *RWMutex) Lock()
func (rw *RWMutex) RLock()
func (rw *RWMutex) RLocker() Locker
func (rw *RWMutex) RUnlock()
func (rw *RWMutex) TryLock() bool
func (rw *RWMutex) TryRLock() bool
func (rw *RWMutex) Unlock()


// 条件变量
type Cond struct {

	// L is held while observing or changing the condition
	L Locker
	// contains filtered or unexported fields
}
func NewCond(l Locker) *Cond
func (c *Cond) Broadcast()
func (c *Cond) Signal()
func (c *Cond) Wait()


// 只会执行一次
type Once struct {
	// contains filtered or unexported fields
}
func (o *Once) Do(f func())


//
type WaitGroup struct {
	// contains filtered or unexported fields
}
func (wg *WaitGroup) Add(delta int)
func (wg *WaitGroup) Done()
func (wg *WaitGroup) Wait()


//
type Pool struct {

	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() any
	// contains filtered or unexported fields
}
func (p *Pool) Get() any
func (p *Pool) Put(x any)


type Map struct {
	// contains filtered or unexported fields
}
func (m *Map) Delete(key any)
func (m *Map) Load(key any) (value any, ok bool)
func (m *Map) LoadAndDelete(key any) (value any, loaded bool)
func (m *Map) LoadOrStore(key, value any) (actual any, loaded bool)
func (m *Map) Range(f func(key, value any) bool)
func (m *Map) Store(key, value any)

```
