// Code generated by counterfeiter. DO NOT EDIT.
package servicefakes

import (
	"example-project/model"
	"example-project/service"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
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
	GetEmployeesByDepartmentStub        func(string) ([]model.EmployeeReturn, error)
	getEmployeesByDepartmentMutex       sync.RWMutex
	getEmployeesByDepartmentArgsForCall []struct {
		arg1 string
	}
	getEmployeesByDepartmentReturns struct {
		result1 []model.EmployeeReturn
		result2 error
	}
	getEmployeesByDepartmentReturnsOnCall map[int]struct {
		result1 []model.EmployeeReturn
		result2 error
	}
	GetPaginatedStub        func(int, int) (model.PaginatedPayload, error)
	getPaginatedMutex       sync.RWMutex
	getPaginatedArgsForCall []struct {
		arg1 int
		arg2 int
	}
	getPaginatedReturns struct {
		result1 model.PaginatedPayload
		result2 error
	}
	getPaginatedReturnsOnCall map[int]struct {
		result1 model.PaginatedPayload
		result2 error
	}
	UpdateEmpStub        func(model.EmployeeReturn) (*mongo.UpdateResult, error)
	updateEmpMutex       sync.RWMutex
	updateEmpArgsForCall []struct {
		arg1 model.EmployeeReturn
	}
	updateEmpReturns struct {
		result1 *mongo.UpdateResult
		result2 error
	}
	updateEmpReturnsOnCall map[int]struct {
		result1 *mongo.UpdateResult
		result2 error
	}
	UpdateEmpShiftStub        func(model.Shift, string) (model.Employee, error)
	updateEmpShiftMutex       sync.RWMutex
	updateEmpShiftArgsForCall []struct {
		arg1 model.Shift
		arg2 string
	}
	updateEmpShiftReturns struct {
		result1 model.Employee
		result2 error
	}
	updateEmpShiftReturnsOnCall map[int]struct {
		result1 model.Employee
		result2 error
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

func (fake *FakeDatabaseInterface) GetEmployeesByDepartment(arg1 string) ([]model.EmployeeReturn, error) {
	fake.getEmployeesByDepartmentMutex.Lock()
	ret, specificReturn := fake.getEmployeesByDepartmentReturnsOnCall[len(fake.getEmployeesByDepartmentArgsForCall)]
	fake.getEmployeesByDepartmentArgsForCall = append(fake.getEmployeesByDepartmentArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetEmployeesByDepartmentStub
	fakeReturns := fake.getEmployeesByDepartmentReturns
	fake.recordInvocation("GetEmployeesByDepartment", []interface{}{arg1})
	fake.getEmployeesByDepartmentMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDatabaseInterface) GetEmployeesByDepartmentCallCount() int {
	fake.getEmployeesByDepartmentMutex.RLock()
	defer fake.getEmployeesByDepartmentMutex.RUnlock()
	return len(fake.getEmployeesByDepartmentArgsForCall)
}

func (fake *FakeDatabaseInterface) GetEmployeesByDepartmentCalls(stub func(string) ([]model.EmployeeReturn, error)) {
	fake.getEmployeesByDepartmentMutex.Lock()
	defer fake.getEmployeesByDepartmentMutex.Unlock()
	fake.GetEmployeesByDepartmentStub = stub
}

func (fake *FakeDatabaseInterface) GetEmployeesByDepartmentArgsForCall(i int) string {
	fake.getEmployeesByDepartmentMutex.RLock()
	defer fake.getEmployeesByDepartmentMutex.RUnlock()
	argsForCall := fake.getEmployeesByDepartmentArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDatabaseInterface) GetEmployeesByDepartmentReturns(result1 []model.EmployeeReturn, result2 error) {
	fake.getEmployeesByDepartmentMutex.Lock()
	defer fake.getEmployeesByDepartmentMutex.Unlock()
	fake.GetEmployeesByDepartmentStub = nil
	fake.getEmployeesByDepartmentReturns = struct {
		result1 []model.EmployeeReturn
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) GetEmployeesByDepartmentReturnsOnCall(i int, result1 []model.EmployeeReturn, result2 error) {
	fake.getEmployeesByDepartmentMutex.Lock()
	defer fake.getEmployeesByDepartmentMutex.Unlock()
	fake.GetEmployeesByDepartmentStub = nil
	if fake.getEmployeesByDepartmentReturnsOnCall == nil {
		fake.getEmployeesByDepartmentReturnsOnCall = make(map[int]struct {
			result1 []model.EmployeeReturn
			result2 error
		})
	}
	fake.getEmployeesByDepartmentReturnsOnCall[i] = struct {
		result1 []model.EmployeeReturn
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) GetPaginated(arg1 int, arg2 int) (model.PaginatedPayload, error) {
	fake.getPaginatedMutex.Lock()
	ret, specificReturn := fake.getPaginatedReturnsOnCall[len(fake.getPaginatedArgsForCall)]
	fake.getPaginatedArgsForCall = append(fake.getPaginatedArgsForCall, struct {
		arg1 int
		arg2 int
	}{arg1, arg2})
	stub := fake.GetPaginatedStub
	fakeReturns := fake.getPaginatedReturns
	fake.recordInvocation("GetPaginated", []interface{}{arg1, arg2})
	fake.getPaginatedMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDatabaseInterface) GetPaginatedCallCount() int {
	fake.getPaginatedMutex.RLock()
	defer fake.getPaginatedMutex.RUnlock()
	return len(fake.getPaginatedArgsForCall)
}

func (fake *FakeDatabaseInterface) GetPaginatedCalls(stub func(int, int) (model.PaginatedPayload, error)) {
	fake.getPaginatedMutex.Lock()
	defer fake.getPaginatedMutex.Unlock()
	fake.GetPaginatedStub = stub
}

func (fake *FakeDatabaseInterface) GetPaginatedArgsForCall(i int) (int, int) {
	fake.getPaginatedMutex.RLock()
	defer fake.getPaginatedMutex.RUnlock()
	argsForCall := fake.getPaginatedArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDatabaseInterface) GetPaginatedReturns(result1 model.PaginatedPayload, result2 error) {
	fake.getPaginatedMutex.Lock()
	defer fake.getPaginatedMutex.Unlock()
	fake.GetPaginatedStub = nil
	fake.getPaginatedReturns = struct {
		result1 model.PaginatedPayload
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) GetPaginatedReturnsOnCall(i int, result1 model.PaginatedPayload, result2 error) {
	fake.getPaginatedMutex.Lock()
	defer fake.getPaginatedMutex.Unlock()
	fake.GetPaginatedStub = nil
	if fake.getPaginatedReturnsOnCall == nil {
		fake.getPaginatedReturnsOnCall = make(map[int]struct {
			result1 model.PaginatedPayload
			result2 error
		})
	}
	fake.getPaginatedReturnsOnCall[i] = struct {
		result1 model.PaginatedPayload
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) UpdateEmp(arg1 model.EmployeeReturn) (*mongo.UpdateResult, error) {
	fake.updateEmpMutex.Lock()
	ret, specificReturn := fake.updateEmpReturnsOnCall[len(fake.updateEmpArgsForCall)]
	fake.updateEmpArgsForCall = append(fake.updateEmpArgsForCall, struct {
		arg1 model.EmployeeReturn
	}{arg1})
	stub := fake.UpdateEmpStub
	fakeReturns := fake.updateEmpReturns
	fake.recordInvocation("UpdateEmp", []interface{}{arg1})
	fake.updateEmpMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDatabaseInterface) UpdateEmpCallCount() int {
	fake.updateEmpMutex.RLock()
	defer fake.updateEmpMutex.RUnlock()
	return len(fake.updateEmpArgsForCall)
}

func (fake *FakeDatabaseInterface) UpdateEmpCalls(stub func(model.EmployeeReturn) (*mongo.UpdateResult, error)) {
	fake.updateEmpMutex.Lock()
	defer fake.updateEmpMutex.Unlock()
	fake.UpdateEmpStub = stub
}

func (fake *FakeDatabaseInterface) UpdateEmpArgsForCall(i int) model.EmployeeReturn {
	fake.updateEmpMutex.RLock()
	defer fake.updateEmpMutex.RUnlock()
	argsForCall := fake.updateEmpArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeDatabaseInterface) UpdateEmpReturns(result1 *mongo.UpdateResult, result2 error) {
	fake.updateEmpMutex.Lock()
	defer fake.updateEmpMutex.Unlock()
	fake.UpdateEmpStub = nil
	fake.updateEmpReturns = struct {
		result1 *mongo.UpdateResult
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) UpdateEmpReturnsOnCall(i int, result1 *mongo.UpdateResult, result2 error) {
	fake.updateEmpMutex.Lock()
	defer fake.updateEmpMutex.Unlock()
	fake.UpdateEmpStub = nil
	if fake.updateEmpReturnsOnCall == nil {
		fake.updateEmpReturnsOnCall = make(map[int]struct {
			result1 *mongo.UpdateResult
			result2 error
		})
	}
	fake.updateEmpReturnsOnCall[i] = struct {
		result1 *mongo.UpdateResult
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) UpdateEmpShift(arg1 model.Shift, arg2 string) (model.Employee, error) {
	fake.updateEmpShiftMutex.Lock()
	ret, specificReturn := fake.updateEmpShiftReturnsOnCall[len(fake.updateEmpShiftArgsForCall)]
	fake.updateEmpShiftArgsForCall = append(fake.updateEmpShiftArgsForCall, struct {
		arg1 model.Shift
		arg2 string
	}{arg1, arg2})
	stub := fake.UpdateEmpShiftStub
	fakeReturns := fake.updateEmpShiftReturns
	fake.recordInvocation("UpdateEmpShift", []interface{}{arg1, arg2})
	fake.updateEmpShiftMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDatabaseInterface) UpdateEmpShiftCallCount() int {
	fake.updateEmpShiftMutex.RLock()
	defer fake.updateEmpShiftMutex.RUnlock()
	return len(fake.updateEmpShiftArgsForCall)
}

func (fake *FakeDatabaseInterface) UpdateEmpShiftCalls(stub func(model.Shift, string) (model.Employee, error)) {
	fake.updateEmpShiftMutex.Lock()
	defer fake.updateEmpShiftMutex.Unlock()
	fake.UpdateEmpShiftStub = stub
}

func (fake *FakeDatabaseInterface) UpdateEmpShiftArgsForCall(i int) (model.Shift, string) {
	fake.updateEmpShiftMutex.RLock()
	defer fake.updateEmpShiftMutex.RUnlock()
	argsForCall := fake.updateEmpShiftArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeDatabaseInterface) UpdateEmpShiftReturns(result1 model.Employee, result2 error) {
	fake.updateEmpShiftMutex.Lock()
	defer fake.updateEmpShiftMutex.Unlock()
	fake.UpdateEmpShiftStub = nil
	fake.updateEmpShiftReturns = struct {
		result1 model.Employee
		result2 error
	}{result1, result2}
}

func (fake *FakeDatabaseInterface) UpdateEmpShiftReturnsOnCall(i int, result1 model.Employee, result2 error) {
	fake.updateEmpShiftMutex.Lock()
	defer fake.updateEmpShiftMutex.Unlock()
	fake.UpdateEmpShiftStub = nil
	if fake.updateEmpShiftReturnsOnCall == nil {
		fake.updateEmpShiftReturnsOnCall = make(map[int]struct {
			result1 model.Employee
			result2 error
		})
	}
	fake.updateEmpShiftReturnsOnCall[i] = struct {
		result1 model.Employee
		result2 error
	}{result1, result2}
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
	fake.getEmployeesByDepartmentMutex.RLock()
	defer fake.getEmployeesByDepartmentMutex.RUnlock()
	fake.getPaginatedMutex.RLock()
	defer fake.getPaginatedMutex.RUnlock()
	fake.updateEmpMutex.RLock()
	defer fake.updateEmpMutex.RUnlock()
	fake.updateEmpShiftMutex.RLock()
	defer fake.updateEmpShiftMutex.RUnlock()
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
