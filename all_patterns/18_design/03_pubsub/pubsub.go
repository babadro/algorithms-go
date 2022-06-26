package _3_pubsub

import "math/rand"

type Message struct {
	TopicName string
	Val       int
}

type Subscriber struct {
	ID string
	CH chan Message
}

type group struct {
	subscribers []Subscriber
}

type topic struct {
	publisherIDs map[string]bool
	groups       map[string]group
}

type Broker struct {
	input  chan Message
	topics map[string]topic
}

func (b *Broker) Subscribe(id, groupName, topicName string) chan Message {
	ch := make(chan Message)

	if _, ok := b.topics[topicName]; !ok {
		b.topics[topicName] = topic{
			publisherIDs: make(map[string]bool),
			groups:       make(map[string]group),
		}
	}

	newSubscriber := Subscriber{
		ID: id,
		CH: ch,
	}

	gr := b.topics[topicName].groups[groupName]

	gr.subscribers = append(gr.subscribers, newSubscriber)
	b.topics[topicName].groups[groupName] = gr

	return ch
}

func (b *Broker) Unsubscribe(id, topicName, groupName string) {
	gr := b.topics[topicName].groups[groupName]
	for i := range gr.subscribers {
		subs := gr.subscribers[i]
		if gr.subscribers[i].ID == id {
			last := len(gr.subscribers) - 1
			subs[i], subs[last] = subs[last], subs[i]
			subs = subs[:last]

			b.topics[topicName][groupName] = subs

			return
		}
	}
}

func (b *Broker) NewPublisher(id, topicName string) {

}

func (b *Broker) Run() {
	for msg := range b.input {
		if groups, ok := b.topics[msg.TopicName]; !ok {
			b.topics[msg.TopicName] = make(map[string][]Subscriber)
		} else {
			for _, group := range groups {
				sub := group[rand.Intn(len(group))]
				select {
				case sub.CH <- msg:
				default: // skip sending
				}
			}
		}
	}
}
