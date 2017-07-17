// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"policy-server/api"
	"sync"
)

type CCClient struct {
	GetAppSpacesStub        func(token string, appGUIDs []string) (map[string]string, error)
	getAppSpacesMutex       sync.RWMutex
	getAppSpacesArgsForCall []struct {
		token    string
		appGUIDs []string
	}
	getAppSpacesReturns struct {
		result1 map[string]string
		result2 error
	}
	getAppSpacesReturnsOnCall map[int]struct {
		result1 map[string]string
		result2 error
	}
	GetSpaceStub        func(token, spaceGUID string) (*api.Space, error)
	getSpaceMutex       sync.RWMutex
	getSpaceArgsForCall []struct {
		token     string
		spaceGUID string
	}
	getSpaceReturns struct {
		result1 *api.Space
		result2 error
	}
	getSpaceReturnsOnCall map[int]struct {
		result1 *api.Space
		result2 error
	}
	GetSpaceGUIDsStub        func(token string, appGUIDs []string) ([]string, error)
	getSpaceGUIDsMutex       sync.RWMutex
	getSpaceGUIDsArgsForCall []struct {
		token    string
		appGUIDs []string
	}
	getSpaceGUIDsReturns struct {
		result1 []string
		result2 error
	}
	getSpaceGUIDsReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	GetUserSpaceStub        func(token, userGUID string, spaces api.Space) (*api.Space, error)
	getUserSpaceMutex       sync.RWMutex
	getUserSpaceArgsForCall []struct {
		token    string
		userGUID string
		spaces   api.Space
	}
	getUserSpaceReturns struct {
		result1 *api.Space
		result2 error
	}
	getUserSpaceReturnsOnCall map[int]struct {
		result1 *api.Space
		result2 error
	}
	GetUserSpacesStub        func(token, userGUID string) (map[string]struct{}, error)
	getUserSpacesMutex       sync.RWMutex
	getUserSpacesArgsForCall []struct {
		token    string
		userGUID string
	}
	getUserSpacesReturns struct {
		result1 map[string]struct{}
		result2 error
	}
	getUserSpacesReturnsOnCall map[int]struct {
		result1 map[string]struct{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *CCClient) GetAppSpaces(token string, appGUIDs []string) (map[string]string, error) {
	var appGUIDsCopy []string
	if appGUIDs != nil {
		appGUIDsCopy = make([]string, len(appGUIDs))
		copy(appGUIDsCopy, appGUIDs)
	}
	fake.getAppSpacesMutex.Lock()
	ret, specificReturn := fake.getAppSpacesReturnsOnCall[len(fake.getAppSpacesArgsForCall)]
	fake.getAppSpacesArgsForCall = append(fake.getAppSpacesArgsForCall, struct {
		token    string
		appGUIDs []string
	}{token, appGUIDsCopy})
	fake.recordInvocation("GetAppSpaces", []interface{}{token, appGUIDsCopy})
	fake.getAppSpacesMutex.Unlock()
	if fake.GetAppSpacesStub != nil {
		return fake.GetAppSpacesStub(token, appGUIDs)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getAppSpacesReturns.result1, fake.getAppSpacesReturns.result2
}

func (fake *CCClient) GetAppSpacesCallCount() int {
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	return len(fake.getAppSpacesArgsForCall)
}

func (fake *CCClient) GetAppSpacesArgsForCall(i int) (string, []string) {
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	return fake.getAppSpacesArgsForCall[i].token, fake.getAppSpacesArgsForCall[i].appGUIDs
}

func (fake *CCClient) GetAppSpacesReturns(result1 map[string]string, result2 error) {
	fake.GetAppSpacesStub = nil
	fake.getAppSpacesReturns = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetAppSpacesReturnsOnCall(i int, result1 map[string]string, result2 error) {
	fake.GetAppSpacesStub = nil
	if fake.getAppSpacesReturnsOnCall == nil {
		fake.getAppSpacesReturnsOnCall = make(map[int]struct {
			result1 map[string]string
			result2 error
		})
	}
	fake.getAppSpacesReturnsOnCall[i] = struct {
		result1 map[string]string
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpace(token string, spaceGUID string) (*api.Space, error) {
	fake.getSpaceMutex.Lock()
	ret, specificReturn := fake.getSpaceReturnsOnCall[len(fake.getSpaceArgsForCall)]
	fake.getSpaceArgsForCall = append(fake.getSpaceArgsForCall, struct {
		token     string
		spaceGUID string
	}{token, spaceGUID})
	fake.recordInvocation("GetSpace", []interface{}{token, spaceGUID})
	fake.getSpaceMutex.Unlock()
	if fake.GetSpaceStub != nil {
		return fake.GetSpaceStub(token, spaceGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSpaceReturns.result1, fake.getSpaceReturns.result2
}

func (fake *CCClient) GetSpaceCallCount() int {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return len(fake.getSpaceArgsForCall)
}

func (fake *CCClient) GetSpaceArgsForCall(i int) (string, string) {
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	return fake.getSpaceArgsForCall[i].token, fake.getSpaceArgsForCall[i].spaceGUID
}

func (fake *CCClient) GetSpaceReturns(result1 *api.Space, result2 error) {
	fake.GetSpaceStub = nil
	fake.getSpaceReturns = struct {
		result1 *api.Space
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpaceReturnsOnCall(i int, result1 *api.Space, result2 error) {
	fake.GetSpaceStub = nil
	if fake.getSpaceReturnsOnCall == nil {
		fake.getSpaceReturnsOnCall = make(map[int]struct {
			result1 *api.Space
			result2 error
		})
	}
	fake.getSpaceReturnsOnCall[i] = struct {
		result1 *api.Space
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpaceGUIDs(token string, appGUIDs []string) ([]string, error) {
	var appGUIDsCopy []string
	if appGUIDs != nil {
		appGUIDsCopy = make([]string, len(appGUIDs))
		copy(appGUIDsCopy, appGUIDs)
	}
	fake.getSpaceGUIDsMutex.Lock()
	ret, specificReturn := fake.getSpaceGUIDsReturnsOnCall[len(fake.getSpaceGUIDsArgsForCall)]
	fake.getSpaceGUIDsArgsForCall = append(fake.getSpaceGUIDsArgsForCall, struct {
		token    string
		appGUIDs []string
	}{token, appGUIDsCopy})
	fake.recordInvocation("GetSpaceGUIDs", []interface{}{token, appGUIDsCopy})
	fake.getSpaceGUIDsMutex.Unlock()
	if fake.GetSpaceGUIDsStub != nil {
		return fake.GetSpaceGUIDsStub(token, appGUIDs)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getSpaceGUIDsReturns.result1, fake.getSpaceGUIDsReturns.result2
}

func (fake *CCClient) GetSpaceGUIDsCallCount() int {
	fake.getSpaceGUIDsMutex.RLock()
	defer fake.getSpaceGUIDsMutex.RUnlock()
	return len(fake.getSpaceGUIDsArgsForCall)
}

func (fake *CCClient) GetSpaceGUIDsArgsForCall(i int) (string, []string) {
	fake.getSpaceGUIDsMutex.RLock()
	defer fake.getSpaceGUIDsMutex.RUnlock()
	return fake.getSpaceGUIDsArgsForCall[i].token, fake.getSpaceGUIDsArgsForCall[i].appGUIDs
}

func (fake *CCClient) GetSpaceGUIDsReturns(result1 []string, result2 error) {
	fake.GetSpaceGUIDsStub = nil
	fake.getSpaceGUIDsReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetSpaceGUIDsReturnsOnCall(i int, result1 []string, result2 error) {
	fake.GetSpaceGUIDsStub = nil
	if fake.getSpaceGUIDsReturnsOnCall == nil {
		fake.getSpaceGUIDsReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.getSpaceGUIDsReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetUserSpace(token string, userGUID string, spaces api.Space) (*api.Space, error) {
	fake.getUserSpaceMutex.Lock()
	ret, specificReturn := fake.getUserSpaceReturnsOnCall[len(fake.getUserSpaceArgsForCall)]
	fake.getUserSpaceArgsForCall = append(fake.getUserSpaceArgsForCall, struct {
		token    string
		userGUID string
		spaces   api.Space
	}{token, userGUID, spaces})
	fake.recordInvocation("GetUserSpace", []interface{}{token, userGUID, spaces})
	fake.getUserSpaceMutex.Unlock()
	if fake.GetUserSpaceStub != nil {
		return fake.GetUserSpaceStub(token, userGUID, spaces)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getUserSpaceReturns.result1, fake.getUserSpaceReturns.result2
}

func (fake *CCClient) GetUserSpaceCallCount() int {
	fake.getUserSpaceMutex.RLock()
	defer fake.getUserSpaceMutex.RUnlock()
	return len(fake.getUserSpaceArgsForCall)
}

func (fake *CCClient) GetUserSpaceArgsForCall(i int) (string, string, api.Space) {
	fake.getUserSpaceMutex.RLock()
	defer fake.getUserSpaceMutex.RUnlock()
	return fake.getUserSpaceArgsForCall[i].token, fake.getUserSpaceArgsForCall[i].userGUID, fake.getUserSpaceArgsForCall[i].spaces
}

func (fake *CCClient) GetUserSpaceReturns(result1 *api.Space, result2 error) {
	fake.GetUserSpaceStub = nil
	fake.getUserSpaceReturns = struct {
		result1 *api.Space
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetUserSpaceReturnsOnCall(i int, result1 *api.Space, result2 error) {
	fake.GetUserSpaceStub = nil
	if fake.getUserSpaceReturnsOnCall == nil {
		fake.getUserSpaceReturnsOnCall = make(map[int]struct {
			result1 *api.Space
			result2 error
		})
	}
	fake.getUserSpaceReturnsOnCall[i] = struct {
		result1 *api.Space
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetUserSpaces(token string, userGUID string) (map[string]struct{}, error) {
	fake.getUserSpacesMutex.Lock()
	ret, specificReturn := fake.getUserSpacesReturnsOnCall[len(fake.getUserSpacesArgsForCall)]
	fake.getUserSpacesArgsForCall = append(fake.getUserSpacesArgsForCall, struct {
		token    string
		userGUID string
	}{token, userGUID})
	fake.recordInvocation("GetUserSpaces", []interface{}{token, userGUID})
	fake.getUserSpacesMutex.Unlock()
	if fake.GetUserSpacesStub != nil {
		return fake.GetUserSpacesStub(token, userGUID)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.getUserSpacesReturns.result1, fake.getUserSpacesReturns.result2
}

func (fake *CCClient) GetUserSpacesCallCount() int {
	fake.getUserSpacesMutex.RLock()
	defer fake.getUserSpacesMutex.RUnlock()
	return len(fake.getUserSpacesArgsForCall)
}

func (fake *CCClient) GetUserSpacesArgsForCall(i int) (string, string) {
	fake.getUserSpacesMutex.RLock()
	defer fake.getUserSpacesMutex.RUnlock()
	return fake.getUserSpacesArgsForCall[i].token, fake.getUserSpacesArgsForCall[i].userGUID
}

func (fake *CCClient) GetUserSpacesReturns(result1 map[string]struct{}, result2 error) {
	fake.GetUserSpacesStub = nil
	fake.getUserSpacesReturns = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) GetUserSpacesReturnsOnCall(i int, result1 map[string]struct{}, result2 error) {
	fake.GetUserSpacesStub = nil
	if fake.getUserSpacesReturnsOnCall == nil {
		fake.getUserSpacesReturnsOnCall = make(map[int]struct {
			result1 map[string]struct{}
			result2 error
		})
	}
	fake.getUserSpacesReturnsOnCall[i] = struct {
		result1 map[string]struct{}
		result2 error
	}{result1, result2}
}

func (fake *CCClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getAppSpacesMutex.RLock()
	defer fake.getAppSpacesMutex.RUnlock()
	fake.getSpaceMutex.RLock()
	defer fake.getSpaceMutex.RUnlock()
	fake.getSpaceGUIDsMutex.RLock()
	defer fake.getSpaceGUIDsMutex.RUnlock()
	fake.getUserSpaceMutex.RLock()
	defer fake.getUserSpaceMutex.RUnlock()
	fake.getUserSpacesMutex.RLock()
	defer fake.getUserSpacesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *CCClient) recordInvocation(key string, args []interface{}) {
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
