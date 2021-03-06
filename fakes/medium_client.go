// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	sync "sync"

	medium "github.com/Medium/medium-sdk-go"
	resource "github.com/cappyzawa/medium-resource"
)

type FakeMediumClient struct {
	CreatePostStub        func(medium.CreatePostOptions) (*medium.Post, error)
	createPostMutex       sync.RWMutex
	createPostArgsForCall []struct {
		arg1 medium.CreatePostOptions
	}
	createPostReturns struct {
		result1 *medium.Post
		result2 error
	}
	createPostReturnsOnCall map[int]struct {
		result1 *medium.Post
		result2 error
	}
	GetUserStub        func(string) (*medium.User, error)
	getUserMutex       sync.RWMutex
	getUserArgsForCall []struct {
		arg1 string
	}
	getUserReturns struct {
		result1 *medium.User
		result2 error
	}
	getUserReturnsOnCall map[int]struct {
		result1 *medium.User
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeMediumClient) CreatePost(arg1 medium.CreatePostOptions) (*medium.Post, error) {
	fake.createPostMutex.Lock()
	ret, specificReturn := fake.createPostReturnsOnCall[len(fake.createPostArgsForCall)]
	fake.createPostArgsForCall = append(fake.createPostArgsForCall, struct {
		arg1 medium.CreatePostOptions
	}{arg1})
	fake.recordInvocation("CreatePost", []interface{}{arg1})
	fake.createPostMutex.Unlock()
	if fake.CreatePostStub != nil {
		return fake.CreatePostStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createPostReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMediumClient) CreatePostCallCount() int {
	fake.createPostMutex.RLock()
	defer fake.createPostMutex.RUnlock()
	return len(fake.createPostArgsForCall)
}

func (fake *FakeMediumClient) CreatePostCalls(stub func(medium.CreatePostOptions) (*medium.Post, error)) {
	fake.createPostMutex.Lock()
	defer fake.createPostMutex.Unlock()
	fake.CreatePostStub = stub
}

func (fake *FakeMediumClient) CreatePostArgsForCall(i int) medium.CreatePostOptions {
	fake.createPostMutex.RLock()
	defer fake.createPostMutex.RUnlock()
	argsForCall := fake.createPostArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMediumClient) CreatePostReturns(result1 *medium.Post, result2 error) {
	fake.createPostMutex.Lock()
	defer fake.createPostMutex.Unlock()
	fake.CreatePostStub = nil
	fake.createPostReturns = struct {
		result1 *medium.Post
		result2 error
	}{result1, result2}
}

func (fake *FakeMediumClient) CreatePostReturnsOnCall(i int, result1 *medium.Post, result2 error) {
	fake.createPostMutex.Lock()
	defer fake.createPostMutex.Unlock()
	fake.CreatePostStub = nil
	if fake.createPostReturnsOnCall == nil {
		fake.createPostReturnsOnCall = make(map[int]struct {
			result1 *medium.Post
			result2 error
		})
	}
	fake.createPostReturnsOnCall[i] = struct {
		result1 *medium.Post
		result2 error
	}{result1, result2}
}

func (fake *FakeMediumClient) GetUser(arg1 string) (*medium.User, error) {
	fake.getUserMutex.Lock()
	ret, specificReturn := fake.getUserReturnsOnCall[len(fake.getUserArgsForCall)]
	fake.getUserArgsForCall = append(fake.getUserArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetUser", []interface{}{arg1})
	fake.getUserMutex.Unlock()
	if fake.GetUserStub != nil {
		return fake.GetUserStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getUserReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeMediumClient) GetUserCallCount() int {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	return len(fake.getUserArgsForCall)
}

func (fake *FakeMediumClient) GetUserCalls(stub func(string) (*medium.User, error)) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = stub
}

func (fake *FakeMediumClient) GetUserArgsForCall(i int) string {
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	argsForCall := fake.getUserArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeMediumClient) GetUserReturns(result1 *medium.User, result2 error) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	fake.getUserReturns = struct {
		result1 *medium.User
		result2 error
	}{result1, result2}
}

func (fake *FakeMediumClient) GetUserReturnsOnCall(i int, result1 *medium.User, result2 error) {
	fake.getUserMutex.Lock()
	defer fake.getUserMutex.Unlock()
	fake.GetUserStub = nil
	if fake.getUserReturnsOnCall == nil {
		fake.getUserReturnsOnCall = make(map[int]struct {
			result1 *medium.User
			result2 error
		})
	}
	fake.getUserReturnsOnCall[i] = struct {
		result1 *medium.User
		result2 error
	}{result1, result2}
}

func (fake *FakeMediumClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createPostMutex.RLock()
	defer fake.createPostMutex.RUnlock()
	fake.getUserMutex.RLock()
	defer fake.getUserMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeMediumClient) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ resource.MediumClient = new(FakeMediumClient)
