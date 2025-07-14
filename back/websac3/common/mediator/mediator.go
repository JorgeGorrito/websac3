package mediator

import (
	"errors"
	"fmt"
	"reflect"
	"websac3/common/dependencies/container"
)

var (
	ErrRequestTypeNil        = errors.New("mediator: request type cannot be nil")
	ErrHandlerNil            = errors.New("mediator: request handler cannot be nil")
	ErrTypeAlreadyRegistered = errors.New("mediator: request type is already registered")
	ErrTypeNotRegistered     = errors.New("mediator: request type is not registered")
	ErrInvalidHandler        = errors.New("mediator: handler must be a function or implement Handle()")
	ErrMediatorNotConfigured = errors.New("mediator: IMediator is not configured")
)

type RequestType reflect.Type

type RequestHandler any

type Handler[I any, O any] interface {
	Handle(request I, lang string) (O, error)
}

type IMediator interface {
	Register(t RequestType, h RequestHandler) error
	GetHandler(t RequestType) (RequestHandler, error)
}

type Mediator struct {
	handlers map[RequestType]RequestHandler
}

func New() *Mediator {
	return &Mediator{
		handlers: make(map[RequestType]RequestHandler),
	}
}

func (m *Mediator) Register(t RequestType, h RequestHandler) error {
	if t == nil {
		return ErrRequestTypeNil
	}
	if h == nil {
		return ErrHandlerNil
	}
	if _, exists := m.handlers[t]; exists {
		return ErrTypeAlreadyRegistered
	}
	m.handlers[t] = h
	return nil
}

func (m *Mediator) GetHandler(t RequestType) (RequestHandler, error) {
	if t == nil {
		return nil, ErrRequestTypeNil
	}
	h, exists := m.handlers[t]
	if !exists {
		return nil, ErrTypeNotRegistered
	}
	return h, nil
}

func Send[I any, O any](request I, lang string) (O, error) {
	var zero O
	t := reflect.TypeOf(request)

	iMediator := container.Inject[IMediator]()
	if iMediator == nil {
		return zero, ErrMediatorNotConfigured
	}

	h, err := iMediator.GetHandler(t)
	if err != nil {
		return zero, fmt.Errorf("mediator: %w", err)
	}

	switch handler := h.(type) {
	case Handler[I, O]:
		return handler.Handle(request, lang)
	case func(I) (O, error):
		return handler(request)
	default:
		return zero, ErrInvalidHandler
	}
}
