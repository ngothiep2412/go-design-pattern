package logic_test

import (
	"context"
	"microservice/internal/database"
	"microservice/internal/logic"
	"os"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

// If environment variable TIME not found, use time.Nnow()
// Otherwise, use time.Parse(go.Env("TIME"))
var currentTIme time.Time

// black box testing
// white box testing
// test-driven development (TDD)

// Say Hello: Output: "Hello " + thing
func TestSayHello(t *testing.T) {
	t.Parallel()

	output := logic.SayHello("World")
	// if output != "Hello World" {
	// 	t.Fail()
	// }

	// output = logic.SayHello("")
	// if output != "Hello" {
	// 	t.Fail()
	// 	return

	// }
	assert.Equal(t, "Hello World", output, "incorrect output")
}

func TestCurrentTime(t *testing.T) {
	// t.Log(currentTIme)
	mockController := gomock.NewController(t)
	defer mockController.Finish()

	userDataAccessor := database.NewMockUserDataAccessor(mockController)
	userDataAccessor.EXPECT().GetUser(gomock.Any(), uint64(1)).Return(database.User{
		ID:   1,
		Name: "Thiep",
	}, nil).AnyTimes()

	_ = logic.Hello{
		UserDataAccessor: userDataAccessor,
	}

	user, err := userDataAccessor.GetUser(context.Background(), 1)
	assert.Nil(t, err)
	assert.Equal(t, database.User{
		ID:   1,
		Name: "Thiep",
	}, user)

	user, err = userDataAccessor.GetUser(context.Background(), 1)
	assert.Nil(t, err)
	assert.Equal(t, database.User{
		ID:   1,
		Name: "Thiep",
	}, user)
}

func TestMain(m *testing.M) {
	var err error
	timeEnvVar := os.Getenv("TIME")
	if timeEnvVar == "" {
		currentTIme = time.Now()
	} else {
		currentTIme, err = time.Parse(time.RFC3339, timeEnvVar)
		if err != nil {
			return
		}
	}

	m.Run()
}
