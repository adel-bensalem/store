package store

type Message struct {
	Type string
	Data interface{}
}

type StateReducer interface {
	ReduceState(state interface{}, message Message) interface{}
}

type Subscriber interface {
	Notify(state interface{})
}

type AbstractStore interface {
	GetState() interface{}
	Subscribe(subscriber Subscriber)
	Dispatch(message Message)
}

type Store struct {
	State interface{}
	subscribers []Subscriber
	StateReducer StateReducer
}

func (store *Store) GetState() interface{}  {
	return store.State
}

func (store *Store) Subscribe(subscriber Subscriber) {
	store.subscribers = append(store.subscribers, subscriber)
}

func (store *Store) Dispatch(message Message) {
	store.State = store.StateReducer.ReduceState(store.GetState(), message)

	for _, subscriber := range store.subscribers {
		subscriber.Notify(store.State)
	}
}