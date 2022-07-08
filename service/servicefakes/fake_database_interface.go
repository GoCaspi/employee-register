// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"example-project/model"
	"example-project/service"
	"sync"
)

type FakeDatabaseInterface struct {
	DeleteByIDStub        func(string) (interface{}, error)
	deleteByIDMutex       sync.RWMutex
	deleteByIDArgsForCall []struct {
		arg1 string
	}
	deleteByIDReturns struct {
		result1 interface{}
		result2 error
	}
	deleteByIDReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	GetByIDStub        func(string) model.Employee
	getByIDMutex       sync.RWMutex
	getByIDArgsForCall []struct {
		arg1 string
	}
	getByIDReturns struct {
		result1 model.Employee
	}
	getByIDReturnsOnCall map[int]struct {
		result1 model.Employee
	}
	UpdateManyStub        func([]interface{}) (interface{}, error)
	updateManyMutex       sync.RWMutex
	updateManyArgsForCall []struct {
		arg1 []interface{}
	}
	updateManyReturns struct {
		result1 interface{}
		result2 error
	}
	updateManyReturnsOnCall map[int]struct {
		result1 interface{}
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDatabaseInterface) DeleteByID(arg1 string) (interface{}, error) {
	fake.deleteByIDMutex.Lock()
	ret, specificReturn := fake.deleteByIDReturnsOnCall[len(fake.deleteByIDArgsForCall)]
	fake.deleteByIDArgsForCall = append(fake.deleteByIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.DeleteByIDStub
	fakeReturns := fake.deleteByIDReturns
	fake.recordInvocation("DeleteByID", []interface{}{arg1})
	fake.deleteByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDatabaseInterface) DeleteByIDCallCount() int {
	fake.deleteByIDMutex.RLock()
	defer fake.deleteByIDMutex.RUnlock()
	return len(fake.deleteByIDArgsForCall)
}

func (fake *FakeDatabaseInterface) DeleteByIDCalls(stub func(string) (interface{}, error)) {
	fake.deleteByIDMutex.Lock()
	defer fake.deleteByIDMutex.Unlock()
	fake.DeleteByIDStub = stub
}

func (fake *FakeDatabaseInterface) DeleteByIDArgsForCall(i int) string {
	fake.deleteByIDMutex.RLock()
	defer fake.deleteByIDMutex.RUnlock()
	argsForCall := fake.deleteByIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDatabaseInterface) DeleteByIDReturns(result1 interface{}, result2 error) {
	fake.deleteByIDMutex.Lock()
	defer fake.deleteByIDMutex.Unlock()
	fake.DeleteByIDStub = nil
	fake.deleteByIDReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) DeleteByIDReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.deleteByIDMutex.Lock()
	defer fake.deleteByIDMutex.Unlock()
	fake.DeleteByIDStub = nil
	if fake.deleteByIDReturnsOnCall == nil {
		fake.deleteByIDReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.deleteByIDReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) GetByID(arg1 string) model.Employee {
	fake.getByIDMutex.Lock()
	ret, specificReturn := fake.getByIDReturnsOnCall[len(fake.getByIDArgsForCall)]
	fake.getByIDArgsForCall = append(fake.getByIDArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetByIDStub
	fakeReturns := fake.getByIDReturns
	fake.recordInvocation("GetByID", []interface{}{arg1})
	fake.getByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeDatabaseInterface) GetByIDCallCount() int {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	return len(fake.getByIDArgsForCall)
}

func (fake *FakeDatabaseInterface) GetByIDCalls(stub func(string) model.Employee) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = stub
}

func (fake *FakeDatabaseInterface) GetByIDArgsForCall(i int) string {
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	argsForCall := fake.getByIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDatabaseInterface) GetByIDReturns(result1 model.Employee) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = nil
	fake.getByIDReturns = struct {
		result1 model.Employee
	}{result1}
}

func (fake *FakeDatabaseInterface) GetByIDReturnsOnCall(i int, result1 model.Employee) {
	fake.getByIDMutex.Lock()
	defer fake.getByIDMutex.Unlock()
	fake.GetByIDStub = nil
	if fake.getByIDReturnsOnCall == nil {
		fake.getByIDReturnsOnCall = make(map[int]struct {
			result1 model.Employee
		})
	}
	fake.getByIDReturnsOnCall[i] = struct {
		result1 model.Employee
	}{result1}
}

func (fake *FakeDatabaseInterface) UpdateMany(arg1 []interface{}) (interface{}, error) {
	var arg1Copy []interface{}
	if arg1 != nil {
		arg1Copy = make([]interface{}, len(arg1))
		copy(arg1Copy, arg1)
	}
	fake.updateManyMutex.Lock()
	ret, specificReturn := fake.updateManyReturnsOnCall[len(fake.updateManyArgsForCall)]
	fake.updateManyArgsForCall = append(fake.updateManyArgsForCall, struct {
		arg1 []interface{}
	}{arg1Copy})
	stub := fake.UpdateManyStub
	fakeReturns := fake.updateManyReturns
	fake.recordInvocation("UpdateMany", []interface{}{arg1Copy})
	fake.updateManyMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDatabaseInterface) UpdateManyCallCount() int {
	fake.updateManyMutex.RLock()
	defer fake.updateManyMutex.RUnlock()
	return len(fake.updateManyArgsForCall)
}

func (fake *FakeDatabaseInterface) UpdateManyCalls(stub func([]interface{}) (interface{}, error)) {
	fake.updateManyMutex.Lock()
	defer fake.updateManyMutex.Unlock()
	fake.UpdateManyStub = stub
}

func (fake *FakeDatabaseInterface) UpdateManyArgsForCall(i int) []interface{} {
	fake.updateManyMutex.RLock()
	defer fake.updateManyMutex.RUnlock()
	argsForCall := fake.updateManyArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDatabaseInterface) UpdateManyReturns(result1 interface{}, result2 error) {
	fake.updateManyMutex.Lock()
	defer fake.updateManyMutex.Unlock()
	fake.UpdateManyStub = nil
	fake.updateManyReturns = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) UpdateManyReturnsOnCall(i int, result1 interface{}, result2 error) {
	fake.updateManyMutex.Lock()
	defer fake.updateManyMutex.Unlock()
	fake.UpdateManyStub = nil
	if fake.updateManyReturnsOnCall == nil {
		fake.updateManyReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 error
		})
	}
	fake.updateManyReturnsOnCall[i] = struct {
		result1 interface{}
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.deleteByIDMutex.RLock()
	defer fake.deleteByIDMutex.RUnlock()
	fake.getByIDMutex.RLock()
	defer fake.getByIDMutex.RUnlock()
	fake.updateManyMutex.RLock()
	defer fake.updateManyMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDatabaseInterface) recordInvocation(key string, args []interface{}) {
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

var _ service.DatabaseInterface = new(FakeDatabaseInterface)
