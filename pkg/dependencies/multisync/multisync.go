package multisync

type Mutex struct {
	allowedThreads int

	usedThreads int

	awaiting []chan struct{}
}

func NewMutex(allowedThreads int) *Mutex {
	return &Mutex{
		allowedThreads: allowedThreads,
		usedThreads:    0,
		awaiting:       []chan struct{}{},
	}
}

func (m *Mutex) Acquire() <-chan struct{} {
	ch := make(chan struct{})
	if m.usedThreads < m.allowedThreads {
		m.usedThreads++
		go func() {
			ch <- struct{}{}
			close(ch)
		}()
		return ch
	}
	m.awaiting = append(m.awaiting, ch)
	return ch
}

func (m *Mutex) Release() {
	if m.usedThreads == 0 {
		panic("Release called on a mutex with no acquired threads")
	}
	m.usedThreads--

	if len(m.awaiting) > 0 {
		next := m.awaiting[0]
		m.awaiting = m.awaiting[1:]

		m.usedThreads++
		go func() {
			next <- struct{}{}
			close(next)
		}()
	}
}
