// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/sjkyspa/stacks/controller/api/api"
)

type FakeRouteRepository struct {
	CreateStub        func(params api.RouteParams) (apiErr error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		params api.RouteParams
	}
	createReturns struct {
		result1 error
	}
	GetRoutesStub        func() (routes api.Routes, apiErr error)
	getRoutesMutex       sync.RWMutex
	getRoutesArgsForCall []struct{}
	getRoutesReturns     struct {
		result1 api.Routes
		result2 error
	}
	GetAppsForRouteStub        func(routeId string) (apps api.Apps, apiErr error)
	getAppsForRouteMutex       sync.RWMutex
	getAppsForRouteArgsForCall []struct {
		routeId string
	}
	getAppsForRouteReturns struct {
		result1 api.Apps
		result2 error
	}
}

func (fake *FakeRouteRepository) Create(params api.RouteParams) (apiErr error) {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		params api.RouteParams
	}{params})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(params)
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeRouteRepository) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeRouteRepository) CreateArgsForCall(i int) api.RouteParams {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].params
}

func (fake *FakeRouteRepository) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeRouteRepository) GetRoutes() (routes api.Routes, apiErr error) {
	fake.getRoutesMutex.Lock()
	fake.getRoutesArgsForCall = append(fake.getRoutesArgsForCall, struct{}{})
	fake.getRoutesMutex.Unlock()
	if fake.GetRoutesStub != nil {
		return fake.GetRoutesStub()
	} else {
		return fake.getRoutesReturns.result1, fake.getRoutesReturns.result2
	}
}

func (fake *FakeRouteRepository) GetRoutesCallCount() int {
	fake.getRoutesMutex.RLock()
	defer fake.getRoutesMutex.RUnlock()
	return len(fake.getRoutesArgsForCall)
}

func (fake *FakeRouteRepository) GetRoutesReturns(result1 api.Routes, result2 error) {
	fake.GetRoutesStub = nil
	fake.getRoutesReturns = struct {
		result1 api.Routes
		result2 error
	}{result1, result2}
}

func (fake *FakeRouteRepository) GetAppsForRoute(routeId string) (apps api.Apps, apiErr error) {
	fake.getAppsForRouteMutex.Lock()
	fake.getAppsForRouteArgsForCall = append(fake.getAppsForRouteArgsForCall, struct {
		routeId string
	}{routeId})
	fake.getAppsForRouteMutex.Unlock()
	if fake.GetAppsForRouteStub != nil {
		return fake.GetAppsForRouteStub(routeId)
	} else {
		return fake.getAppsForRouteReturns.result1, fake.getAppsForRouteReturns.result2
	}
}

func (fake *FakeRouteRepository) GetAppsForRouteCallCount() int {
	fake.getAppsForRouteMutex.RLock()
	defer fake.getAppsForRouteMutex.RUnlock()
	return len(fake.getAppsForRouteArgsForCall)
}

func (fake *FakeRouteRepository) GetAppsForRouteArgsForCall(i int) string {
	fake.getAppsForRouteMutex.RLock()
	defer fake.getAppsForRouteMutex.RUnlock()
	return fake.getAppsForRouteArgsForCall[i].routeId
}

func (fake *FakeRouteRepository) GetAppsForRouteReturns(result1 api.Apps, result2 error) {
	fake.GetAppsForRouteStub = nil
	fake.getAppsForRouteReturns = struct {
		result1 api.Apps
		result2 error
	}{result1, result2}
}

var _ api.RouteRepository = new(FakeRouteRepository)
