package _3_pubsub

import (
	"math/rand"
	"time"
)

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

	b.createTopicIfNotExists(topicName)

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
		sub := gr.subscribers[i]
		if sub.ID == id {
			last := len(gr.subscribers) - 1
			gr.subscribers[i], gr.subscribers[last] = gr.subscribers[last], gr.subscribers[i]
			gr.subscribers = gr.subscribers[:last]

			if len(gr.subscribers) == 0 {
				delete(b.topics[topicName].groups, groupName)
			}

			b.deleteTopicIfEmpty(topicName)

			return
		}
	}
}

func (b *Broker) RegisterPublisher(id, topicName string) chan<- Message {
	b.createTopicIfNotExists(topicName)
	t := b.topics[topicName]
	t.publisherIDs[id] = true

	return b.input
}

func (b *Broker) UnregisterPublisher(id, topicName string) {
	delete(b.topics[topicName].publisherIDs, id)

	b.deleteTopicIfEmpty(topicName)
}

func (b *Broker) createTopicIfNotExists(topicName string) {
	if _, ok := b.topics[topicName]; !ok {
		b.topics[topicName] = topic{
			publisherIDs: make(map[string]bool),
			groups:       make(map[string]group),
		}
	}
}

func (b *Broker) deleteTopicIfEmpty(topicName string) {
	t := b.topics[topicName]
	if len(t.groups) == 0 && len(t.publisherIDs) == 0 {
		delete(b.topics, topicName)
	}
}

func (b *Broker) Run() {
	for msg := range b.input {
		t, ok := b.topics[msg.TopicName]
		if !ok {
			continue // skip non-existing topic
		}

	Loop:
		for _, gr := range t.groups {
			for _ = range gr.subscribers {
				subIDx := rand.Intn(len(gr.subscribers))
				select {
				case gr.subscribers[subIDx].CH <- msg:
					continue Loop
				case <-time.After(10 * time.Millisecond):
				}
			}
		}
	}
}
