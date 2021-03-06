// Code generated by mocker; DO NOT EDIT
// github.com/travisjeffery/mocker
package mock

import (
	"sync"

	"github.com/chicagobuss/userService/userService"
)

var (
	lockMockUserServiceGet sync.RWMutex
)

// MockUserService is a mock implementation of UserService.
//
//     func TestSomethingThatUsesUserService(t *testing.T) {
//
//         // make and configure a mocked UserService
//         mockedUserService := &MockUserService{
//             GetFunc: func(id string) (*userService.User, error) {
// 	               panic("TODO: mock out the Get method")
//             },
//         }
//
//         // TODO: use mockedUserService in code that requires UserService
//         //       and then make assertions.
//
//     }
type MockUserService struct {
	// GetFunc mocks the Get method.
	GetFunc func(id string) (*userService.User, error)

	// calls tracks calls to the methods.
	calls struct {
		// Get holds details about calls to the Get method.
		Get []struct {
			// Id is the id argument value.
			Id string
		}
	}
}

// Reset resets the calls made to the mocked APIs.
func (mock *MockUserService) Reset() {
	lockMockUserServiceGet.Lock()
	mock.calls.Get = nil
	lockMockUserServiceGet.Unlock()
}

// Get calls GetFunc.
func (mock *MockUserService) Get(id string) (*userService.User, error) {
	if mock.GetFunc == nil {
		panic("moq: MockUserService.GetFunc is nil but UserService.Get was just called")
	}
	callInfo := struct {
		Id string
	}{
		Id: id,
	}
	lockMockUserServiceGet.Lock()
	mock.calls.Get = append(mock.calls.Get, callInfo)
	lockMockUserServiceGet.Unlock()
	return mock.GetFunc(id)
}

// GetCalled returns true if at least one call was made to Get.
func (mock *MockUserService) GetCalled() bool {
	lockMockUserServiceGet.RLock()
	defer lockMockUserServiceGet.RUnlock()
	return len(mock.calls.Get) > 0
}

// GetCalls gets all the calls that were made to Get.
// Check the length with:
//     len(mockedUserService.GetCalls())
func (mock *MockUserService) GetCalls() []struct {
	Id string
} {
	var calls []struct {
		Id string
	}
	lockMockUserServiceGet.RLock()
	calls = mock.calls.Get
	lockMockUserServiceGet.RUnlock()
	return calls
}
