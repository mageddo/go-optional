package optional

import "errors"

type Optional interface {
	Map(func(o interface{}) interface{}) Optional
	OrElse(interface{}) interface{}
	IsPresent() bool
	Get() (interface{}, error)
}

type OptionalImpl struct {
	value interface{}
}

func (op *OptionalImpl) Map(fn func(o interface{}) interface{}) Optional {
	if op.IsPresent() {
		op.value = fn(op.value)
	}
	return op
}

func (op *OptionalImpl) OrElse(o interface{}) interface{} {
	if op.IsPresent() {
		return op.value
	}
	return o
}

func (op *OptionalImpl) Get() (interface{}, error) {
	if op.IsPresent() {
		return op.value, nil
	}
	return nil, errors.New("No such element")
}

func(op *OptionalImpl) IsPresent() bool {
	return op.value != nil
}

func OfNullable(o interface{}) Optional {
	return &OptionalImpl{o}
}
