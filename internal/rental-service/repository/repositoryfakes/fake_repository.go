// Code generated by counterfeiter. DO NOT EDIT.
package repositoryfakes

import (
	"sync"

	"github.com/svilenkomitov/rental-service/internal/rental-service/repository"
)

type FakeRepository struct {
	FetchRentalByIdStub        func(int) (*repository.Entity, error)
	fetchRentalByIdMutex       sync.RWMutex
	fetchRentalByIdArgsForCall []struct {
		arg1 int
	}
	fetchRentalByIdReturns struct {
		result1 *repository.Entity
		result2 error
	}
	fetchRentalByIdReturnsOnCall map[int]struct {
		result1 *repository.Entity
		result2 error
	}
	FetchRentalsStub        func(map[repository.QueryKey]interface{}) ([]*repository.Entity, error)
	fetchRentalsMutex       sync.RWMutex
	fetchRentalsArgsForCall []struct {
		arg1 map[repository.QueryKey]interface{}
	}
	fetchRentalsReturns struct {
		result1 []*repository.Entity
		result2 error
	}
	fetchRentalsReturnsOnCall map[int]struct {
		result1 []*repository.Entity
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRepository) FetchRentalById(arg1 int) (*repository.Entity, error) {
	fake.fetchRentalByIdMutex.Lock()
	ret, specificReturn := fake.fetchRentalByIdReturnsOnCall[len(fake.fetchRentalByIdArgsForCall)]
	fake.fetchRentalByIdArgsForCall = append(fake.fetchRentalByIdArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.FetchRentalByIdStub
	fakeReturns := fake.fetchRentalByIdReturns
	fake.recordInvocation("FetchRentalById", []interface{}{arg1})
	fake.fetchRentalByIdMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) FetchRentalByIdCallCount() int {
	fake.fetchRentalByIdMutex.RLock()
	defer fake.fetchRentalByIdMutex.RUnlock()
	return len(fake.fetchRentalByIdArgsForCall)
}

func (fake *FakeRepository) FetchRentalByIdCalls(stub func(int) (*repository.Entity, error)) {
	fake.fetchRentalByIdMutex.Lock()
	defer fake.fetchRentalByIdMutex.Unlock()
	fake.FetchRentalByIdStub = stub
}

func (fake *FakeRepository) FetchRentalByIdArgsForCall(i int) int {
	fake.fetchRentalByIdMutex.RLock()
	defer fake.fetchRentalByIdMutex.RUnlock()
	argsForCall := fake.fetchRentalByIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) FetchRentalByIdReturns(result1 *repository.Entity, result2 error) {
	fake.fetchRentalByIdMutex.Lock()
	defer fake.fetchRentalByIdMutex.Unlock()
	fake.FetchRentalByIdStub = nil
	fake.fetchRentalByIdReturns = struct {
		result1 *repository.Entity
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) FetchRentalByIdReturnsOnCall(i int, result1 *repository.Entity, result2 error) {
	fake.fetchRentalByIdMutex.Lock()
	defer fake.fetchRentalByIdMutex.Unlock()
	fake.FetchRentalByIdStub = nil
	if fake.fetchRentalByIdReturnsOnCall == nil {
		fake.fetchRentalByIdReturnsOnCall = make(map[int]struct {
			result1 *repository.Entity
			result2 error
		})
	}
	fake.fetchRentalByIdReturnsOnCall[i] = struct {
		result1 *repository.Entity
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) FetchRentals(arg1 map[repository.QueryKey]interface{}) ([]*repository.Entity, error) {
	fake.fetchRentalsMutex.Lock()
	ret, specificReturn := fake.fetchRentalsReturnsOnCall[len(fake.fetchRentalsArgsForCall)]
	fake.fetchRentalsArgsForCall = append(fake.fetchRentalsArgsForCall, struct {
		arg1 map[repository.QueryKey]interface{}
	}{arg1})
	stub := fake.FetchRentalsStub
	fakeReturns := fake.fetchRentalsReturns
	fake.recordInvocation("FetchRentals", []interface{}{arg1})
	fake.fetchRentalsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRepository) FetchRentalsCallCount() int {
	fake.fetchRentalsMutex.RLock()
	defer fake.fetchRentalsMutex.RUnlock()
	return len(fake.fetchRentalsArgsForCall)
}

func (fake *FakeRepository) FetchRentalsCalls(stub func(map[repository.QueryKey]interface{}) ([]*repository.Entity, error)) {
	fake.fetchRentalsMutex.Lock()
	defer fake.fetchRentalsMutex.Unlock()
	fake.FetchRentalsStub = stub
}

func (fake *FakeRepository) FetchRentalsArgsForCall(i int) map[repository.QueryKey]interface{} {
	fake.fetchRentalsMutex.RLock()
	defer fake.fetchRentalsMutex.RUnlock()
	argsForCall := fake.fetchRentalsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRepository) FetchRentalsReturns(result1 []*repository.Entity, result2 error) {
	fake.fetchRentalsMutex.Lock()
	defer fake.fetchRentalsMutex.Unlock()
	fake.FetchRentalsStub = nil
	fake.fetchRentalsReturns = struct {
		result1 []*repository.Entity
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) FetchRentalsReturnsOnCall(i int, result1 []*repository.Entity, result2 error) {
	fake.fetchRentalsMutex.Lock()
	defer fake.fetchRentalsMutex.Unlock()
	fake.FetchRentalsStub = nil
	if fake.fetchRentalsReturnsOnCall == nil {
		fake.fetchRentalsReturnsOnCall = make(map[int]struct {
			result1 []*repository.Entity
			result2 error
		})
	}
	fake.fetchRentalsReturnsOnCall[i] = struct {
		result1 []*repository.Entity
		result2 error
	}{result1, result2}
}

func (fake *FakeRepository) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.fetchRentalByIdMutex.RLock()
	defer fake.fetchRentalByIdMutex.RUnlock()
	fake.fetchRentalsMutex.RLock()
	defer fake.fetchRentalsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRepository) recordInvocation(key string, args []interface{}) {
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

var _ repository.Repository = new(FakeRepository)
