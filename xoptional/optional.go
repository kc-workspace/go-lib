package xoptional

type Optional[V any] struct {
	value *V
}

func (o *Optional[V]) Empty() bool {
	return o.value == nil
}

func (o *Optional[V]) Present() bool {
	return o.value != nil
}

func (o *Optional[V]) GetFn(fn func(*V) V) V {
	return fn(o.value)
}

func (o *Optional[V]) Get() V {
	return o.GetFn(func(v *V) V { return *v })
}

func (o *Optional[V]) Raw() *V {
	return o.value
}

func (o *Optional[V]) Or(fn func() V) V {
	if o.value == nil {
		return fn()
	}
	return *o.value
}

func (o *Optional[V]) OrElse(def V) V {
	return o.Or(func() V { return def })
}
