package suite

import (
	"github.com/andre-ols/chatservice/testutils/faker"
)

func init() {
	faker.InitializeProviders()
	faker.Setup()
}
