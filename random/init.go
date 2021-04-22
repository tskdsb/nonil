package random

import (
	"math/rand"
	"reflect"
	"time"
)

func init() {
	r = rand.New(rand.NewSource(time.Now().UnixNano()))

	registeredGenerator = make(map[reflect.Type]valueGenerator)

	RegisterGenerator(reflect.TypeOf(time.Time{}), func() interface{} {
		return time.Now()
	})
}
