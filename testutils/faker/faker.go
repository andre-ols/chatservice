package faker

import (
	"reflect"

	"github.com/andre-ols/chatservice/testutils/faker/providers"
	"github.com/go-faker/faker/v4"
	"github.com/go-faker/faker/v4/pkg/options"
)

func InitializeProviders() {
	_ = faker.AddProvider("time_now", func(v reflect.Value) (any, error) {
		return providers.TimeNow(), nil
	})

	_ = faker.AddProvider("time_zero", func(v reflect.Value) (any, error) {
		return providers.TimeZero(), nil
	})
}

func Setup() {
	options.SetRandomMapAndSliceMinSize(2)
	options.SetRandomMapAndSliceMaxSize(5)
}
