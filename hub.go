package main

import "sync"

type Message struct {
	OP      string
	Content string
}

type Consumer struct {
	Queue chan Message
}

func NewConsumer() *Consumer {
	return &Consumer{
		Queue: make(chan Message),
	}
}

type Hub struct {
	sync.RWMutex
	consumers map[*Consumer]bool
}

func NewHub() *Hub {
	return &Hub{
		consumers: map[*Consumer]bool{},
	}
}

func (h *Hub) Register(c *Consumer) {
	h.Lock()
	defer h.Unlock()

	h.consumers[c] = true
}

func (h *Hub) Unregister(c *Consumer) {
	h.Lock()
	defer h.Unlock()

	if _, ok := h.consumers[c]; ok {
		close(c.Queue)
		delete(h.consumers, c)
	}
}

func (h *Hub) Broadcast(msg Message) {
	for consumer := range h.consumers {
		select {
		case consumer.Queue <- msg:
		default:
			h.Unregister(consumer)
		}
	}
}
