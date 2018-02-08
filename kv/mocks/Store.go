// Code generated by mockery v1.0.0
package mocks

import kv "github.com/getamis/sirius/kv"
import mock "github.com/stretchr/testify/mock"

// Store is an autogenerated mock type for the Store type
type Store struct {
	mock.Mock
}

// AtomicDelete provides a mock function with given fields: key, expected
func (_m *Store) AtomicDelete(key string, expected *kv.KeyValue) (bool, error) {
	ret := _m.Called(key, expected)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string, *kv.KeyValue) bool); ok {
		r0 = rf(key, expected)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, *kv.KeyValue) error); ok {
		r1 = rf(key, expected)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// AtomicPut provides a mock function with given fields: key, value, expected, opts
func (_m *Store) AtomicPut(key string, value []byte, expected *kv.KeyValue, opts ...kv.PutOption) (*kv.KeyValue, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key, value, expected)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 *kv.KeyValue
	if rf, ok := ret.Get(0).(func(string, []byte, *kv.KeyValue, ...kv.PutOption) *kv.KeyValue); ok {
		r0 = rf(key, value, expected, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kv.KeyValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, []byte, *kv.KeyValue, ...kv.PutOption) error); ok {
		r1 = rf(key, value, expected, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *Store) Close() {
	_m.Called()
}

// Delete provides a mock function with given fields: key
func (_m *Store) Delete(key string) error {
	ret := _m.Called(key)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteTree provides a mock function with given fields: prefix
func (_m *Store) DeleteTree(prefix string) error {
	ret := _m.Called(prefix)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(prefix)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Exists provides a mock function with given fields: key
func (_m *Store) Exists(key string) (bool, error) {
	ret := _m.Called(key)

	var r0 bool
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(key)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Get provides a mock function with given fields: key
func (_m *Store) Get(key string) (*kv.KeyValue, error) {
	ret := _m.Called(key)

	var r0 *kv.KeyValue
	if rf, ok := ret.Get(0).(func(string) *kv.KeyValue); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*kv.KeyValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// List provides a mock function with given fields: prefix
func (_m *Store) List(prefix string) ([]*kv.KeyValue, error) {
	ret := _m.Called(prefix)

	var r0 []*kv.KeyValue
	if rf, ok := ret.Get(0).(func(string) []*kv.KeyValue); ok {
		r0 = rf(prefix)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*kv.KeyValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(prefix)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Lock provides a mock function with given fields: key, opts
func (_m *Store) Lock(key string, opts ...kv.LockOption) (kv.Locker, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 kv.Locker
	if rf, ok := ret.Get(0).(func(string, ...kv.LockOption) kv.Locker); ok {
		r0 = rf(key, opts...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(kv.Locker)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, ...kv.LockOption) error); ok {
		r1 = rf(key, opts...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Put provides a mock function with given fields: key, value, opts
func (_m *Store) Put(key string, value []byte, opts ...kv.PutOption) error {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, key, value)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, []byte, ...kv.PutOption) error); ok {
		r0 = rf(key, value, opts...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Watch provides a mock function with given fields: key, stopCh
func (_m *Store) Watch(key string, stopCh <-chan struct{}) (<-chan *kv.KeyValue, error) {
	ret := _m.Called(key, stopCh)

	var r0 <-chan *kv.KeyValue
	if rf, ok := ret.Get(0).(func(string, <-chan struct{}) <-chan *kv.KeyValue); ok {
		r0 = rf(key, stopCh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan *kv.KeyValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, <-chan struct{}) error); ok {
		r1 = rf(key, stopCh)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// WatchTree provides a mock function with given fields: prefix, stopCh
func (_m *Store) WatchTree(prefix string, stopCh <-chan struct{}) (<-chan []*kv.KeyValue, error) {
	ret := _m.Called(prefix, stopCh)

	var r0 <-chan []*kv.KeyValue
	if rf, ok := ret.Get(0).(func(string, <-chan struct{}) <-chan []*kv.KeyValue); ok {
		r0 = rf(prefix, stopCh)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(<-chan []*kv.KeyValue)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, <-chan struct{}) error); ok {
		r1 = rf(prefix, stopCh)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
