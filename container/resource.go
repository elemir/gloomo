package container

type Resource[T any] struct {
	value *T
}

func (r *Resource[T]) Get() (T, bool) {
	if r.value == nil {
		return Zero[T](), false
	}

	return *r.value, true
}

func (r *Resource[T]) Set(val T) {
	if r.value == nil {
		r.value = &val
	}

	*r.value = val
}
