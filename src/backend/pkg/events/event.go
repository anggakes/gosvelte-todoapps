package events

import (
	"context"
)

type Msg interface{}

type HandlerFunc func(ctx context.Context, msg Msg) error

type IEvent interface {
	Publish(ctx context.Context, eventName string, msg Msg) error
	RegisterListener(eventName string, hf HandlerFunc)
}

func New() *EventBus {
	return &EventBus{
		listeners: map[string][]HandlerFunc{},
	}
}

type EventBus struct {
	listeners map[string][]HandlerFunc
}

func (c *EventBus) RegisterListener(eventName string, hf HandlerFunc) {

	_, exists := c.listeners[eventName]
	if !exists {
		c.listeners[eventName] = make([]HandlerFunc, 0)
	}

	c.listeners[eventName] = append(c.listeners[eventName], hf)
}

func (c *EventBus) Publish(ctx context.Context, eventName string, msg Msg) error {

	listeners := c.listeners[eventName]

	for _, listenerHandler := range listeners {
		err := listenerHandler(ctx, msg)
		if err != nil {
			return err
		}
	}

	return nil
}
