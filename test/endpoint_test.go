package test

import (
	"testing"

	"github.com/chicagobuss/userService/mock"
	"github.com/chicagobuss/userService/userService"
	"github.com/go-kit/kit/endpoint"
)

func TestUserServiceEndpoint(t *testing.T) {
	mockedUserService := &mock.MockUserService{
		GetFunc: func(id string) (*userService.User, error) {
			panic("TODO: mock out the Get method")
		},
	}
	// ep := endpoint.Endpoint ({UserService: mockedUserService})

	func MockedUserServiceEndpoint(s mockedUserService) endpoint.Endpoint {
		return func(ctx context.Context, request interface{}) (response interface{}, err error) {
			req := request.(postProfileRequest)
			e := s.PostProfile(ctx, req.Profile)
			return postProfileResponse{Err: e}, nil
		}
	}
	ep := MockerUserServiceEndpoints(mockedUserService)

	resp, err := ep.GetUser(endpoint.{ID: "travisjeffery"})
	// ...
	if !us.GetUserCalled() {
		t.Error("expected endpoint to call GetUser")
	}
}
