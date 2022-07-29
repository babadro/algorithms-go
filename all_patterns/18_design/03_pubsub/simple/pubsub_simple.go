package simple

type Message struct {
	subscriberID int
	value        string
}

type Broker struct {
	capacity    int
	input       chan Message
	subscribers map[int]chan Message
}

func (b *Broker) Publish(m Message) {
	b.input <- m
}

func (b *Broker) Subscribe() (int, chan Message) {
	ch := make(chan Message, b.capacity)
	id := len(b.subscribers) + 1
	b.subscribers[id] = ch

	return id, ch
}

func (b *Broker) Run() {
	for message := range b.input {
		ch, ok := b.subscribers[message.subscriberID]
		if !ok {
			continue
		}

		if len(ch) == b.capacity {
			continue
		}

		ch <- message
	}
}
