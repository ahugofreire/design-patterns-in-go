package main

import "fmt"

// Interface Handler
type Handler interface {
	SetNext(handler Handler) Handler
	Handle(request string) string
}

// Implementação base da Handler
type BaseHandler struct {
	nextHandler Handler
}

func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.nextHandler = handler
	return handler
}

// Implementação concreta da Handler
type ConcreteHandler1 struct {
	BaseHandler
}

func (h *ConcreteHandler1) Handle(request string) string {
	if request == "request1" {
		return "ConcreteHandler1 handled the request"
	} else if h.nextHandler != nil {
		return h.nextHandler.Handle(request)
	} else {
		return "No handler found for the request"
	}
}

// Implementação concreta da Handler
type ConcreteHandler2 struct {
	BaseHandler
}

func (h *ConcreteHandler2) Handle(request string) string {
	if request == "request2" {
		return "ConcreteHandler2 handled the request"
	} else if h.nextHandler != nil {
		return h.nextHandler.Handle(request)
	} else {
		return "No handler found for the request"
	}
}

// Implementação concreta da Handler
type ConcreteHandler3 struct {
	BaseHandler
}

func (h *ConcreteHandler3) Handle(request string) string {
	if request == "request3" {
		return "ConcreteHandler3 handled the request"
	} else if h.nextHandler != nil {
		return h.nextHandler.Handle(request)
	} else {
		return "No handler found for the request"
	}
}

func main() {
	// Criando a cadeia de Handlers
	handler1 := &ConcreteHandler1{}
	handler2 := &ConcreteHandler2{}
	handler3 := &ConcreteHandler3{}
	handler1.SetNext(handler2).SetNext(handler3)

	// Enviando solicitações
	fmt.Println(handler1.Handle("request1"))
	fmt.Println(handler1.Handle("request2"))
	fmt.Println(handler1.Handle("request3"))
	fmt.Println(handler1.Handle("request4"))
}
