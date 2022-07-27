// Code generated by counterfeiter. DO NOT EDIT.
package routesfakes

import (
	"example-project/routes"
	"sync"

	"github.com/gin-gonic/gin"
)

type FakeHandlerInterface struct {
	CreateEmployeeHandlerStub        func(*gin.Context)
	createEmployeeHandlerMutex       sync.RWMutex
	createEmployeeHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	DeleteByIdHandlerStub        func(*gin.Context)
	deleteByIdHandlerMutex       sync.RWMutex
	deleteByIdHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	GetAllEmployeesHandlerStub        func(*gin.Context)
	getAllEmployeesHandlerMutex       sync.RWMutex
	getAllEmployeesHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	GetEmployeeHandlerStub        func(*gin.Context)
	getEmployeeHandlerMutex       sync.RWMutex
	getEmployeeHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	LoginStub        func(*gin.Context)
	loginMutex       sync.RWMutex
	loginArgsForCall []struct {
		arg1 *gin.Context
	}
	LogoutStub        func(*gin.Context)
	logoutMutex       sync.RWMutex
	logoutArgsForCall []struct {
		arg1 *gin.Context
	}
	OAuthRedirectHandlerStub        func(*gin.Context)
	oAuthRedirectHandlerMutex       sync.RWMutex
	oAuthRedirectHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	OAuthStarterHandlerStub        func(*gin.Context)
	oAuthStarterHandlerMutex       sync.RWMutex
	oAuthStarterHandlerArgsForCall []struct {
		arg1 *gin.Context
	}
	UpdateByIdStub        func(*gin.Context)
	updateByIdMutex       sync.RWMutex
	updateByIdArgsForCall []struct {
		arg1 *gin.Context
	}
	ValidateTokenStub        func(*gin.Context)
	validateTokenMutex       sync.RWMutex
	validateTokenArgsForCall []struct {
		arg1 *gin.Context
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeHandlerInterface) CreateEmployeeHandler(arg1 *gin.Context) {
	fake.createEmployeeHandlerMutex.Lock()
	fake.createEmployeeHandlerArgsForCall = append(fake.createEmployeeHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.CreateEmployeeHandlerStub
	fake.recordInvocation("CreateEmployeeHandler", []interface{}{arg1})
	fake.createEmployeeHandlerMutex.Unlock()
	if stub != nil {
		fake.CreateEmployeeHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) CreateEmployeeHandlerCallCount() int {
	fake.createEmployeeHandlerMutex.RLock()
	defer fake.createEmployeeHandlerMutex.RUnlock()
	return len(fake.createEmployeeHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) CreateEmployeeHandlerCalls(stub func(*gin.Context)) {
	fake.createEmployeeHandlerMutex.Lock()
	defer fake.createEmployeeHandlerMutex.Unlock()
	fake.CreateEmployeeHandlerStub = stub
}

func (fake *FakeHandlerInterface) CreateEmployeeHandlerArgsForCall(i int) *gin.Context {
	fake.createEmployeeHandlerMutex.RLock()
	defer fake.createEmployeeHandlerMutex.RUnlock()
	argsForCall := fake.createEmployeeHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) DeleteByIdHandler(arg1 *gin.Context) {
	fake.deleteByIdHandlerMutex.Lock()
	fake.deleteByIdHandlerArgsForCall = append(fake.deleteByIdHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.DeleteByIdHandlerStub
	fake.recordInvocation("DeleteByIdHandler", []interface{}{arg1})
	fake.deleteByIdHandlerMutex.Unlock()
	if stub != nil {
		fake.DeleteByIdHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) DeleteByIdHandlerCallCount() int {
	fake.deleteByIdHandlerMutex.RLock()
	defer fake.deleteByIdHandlerMutex.RUnlock()
	return len(fake.deleteByIdHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) DeleteByIdHandlerCalls(stub func(*gin.Context)) {
	fake.deleteByIdHandlerMutex.Lock()
	defer fake.deleteByIdHandlerMutex.Unlock()
	fake.DeleteByIdHandlerStub = stub
}

func (fake *FakeHandlerInterface) DeleteByIdHandlerArgsForCall(i int) *gin.Context {
	fake.deleteByIdHandlerMutex.RLock()
	defer fake.deleteByIdHandlerMutex.RUnlock()
	argsForCall := fake.deleteByIdHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) GetAllEmployeesHandler(arg1 *gin.Context) {
	fake.getAllEmployeesHandlerMutex.Lock()
	fake.getAllEmployeesHandlerArgsForCall = append(fake.getAllEmployeesHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetAllEmployeesHandlerStub
	fake.recordInvocation("GetAllEmployeesHandler", []interface{}{arg1})
	fake.getAllEmployeesHandlerMutex.Unlock()
	if stub != nil {
		fake.GetAllEmployeesHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetAllEmployeesHandlerCallCount() int {
	fake.getAllEmployeesHandlerMutex.RLock()
	defer fake.getAllEmployeesHandlerMutex.RUnlock()
	return len(fake.getAllEmployeesHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) GetAllEmployeesHandlerCalls(stub func(*gin.Context)) {
	fake.getAllEmployeesHandlerMutex.Lock()
	defer fake.getAllEmployeesHandlerMutex.Unlock()
	fake.GetAllEmployeesHandlerStub = stub
}

func (fake *FakeHandlerInterface) GetAllEmployeesHandlerArgsForCall(i int) *gin.Context {
	fake.getAllEmployeesHandlerMutex.RLock()
	defer fake.getAllEmployeesHandlerMutex.RUnlock()
	argsForCall := fake.getAllEmployeesHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) GetEmployeeHandler(arg1 *gin.Context) {
	fake.getEmployeeHandlerMutex.Lock()
	fake.getEmployeeHandlerArgsForCall = append(fake.getEmployeeHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.GetEmployeeHandlerStub
	fake.recordInvocation("GetEmployeeHandler", []interface{}{arg1})
	fake.getEmployeeHandlerMutex.Unlock()
	if stub != nil {
		fake.GetEmployeeHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) GetEmployeeHandlerCallCount() int {
	fake.getEmployeeHandlerMutex.RLock()
	defer fake.getEmployeeHandlerMutex.RUnlock()
	return len(fake.getEmployeeHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) GetEmployeeHandlerCalls(stub func(*gin.Context)) {
	fake.getEmployeeHandlerMutex.Lock()
	defer fake.getEmployeeHandlerMutex.Unlock()
	fake.GetEmployeeHandlerStub = stub
}

func (fake *FakeHandlerInterface) GetEmployeeHandlerArgsForCall(i int) *gin.Context {
	fake.getEmployeeHandlerMutex.RLock()
	defer fake.getEmployeeHandlerMutex.RUnlock()
	argsForCall := fake.getEmployeeHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) Login(arg1 *gin.Context) {
	fake.loginMutex.Lock()
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.LoginStub
	fake.recordInvocation("Login", []interface{}{arg1})
	fake.loginMutex.Unlock()
	if stub != nil {
		fake.LoginStub(arg1)
	}
}

func (fake *FakeHandlerInterface) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeHandlerInterface) LoginCalls(stub func(*gin.Context)) {
	fake.loginMutex.Lock()
	defer fake.loginMutex.Unlock()
	fake.LoginStub = stub
}

func (fake *FakeHandlerInterface) LoginArgsForCall(i int) *gin.Context {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	argsForCall := fake.loginArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) Logout(arg1 *gin.Context) {
	fake.logoutMutex.Lock()
	fake.logoutArgsForCall = append(fake.logoutArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.LogoutStub
	fake.recordInvocation("Logout", []interface{}{arg1})
	fake.logoutMutex.Unlock()
	if stub != nil {
		fake.LogoutStub(arg1)
	}
}

func (fake *FakeHandlerInterface) LogoutCallCount() int {
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	return len(fake.logoutArgsForCall)
}

func (fake *FakeHandlerInterface) LogoutCalls(stub func(*gin.Context)) {
	fake.logoutMutex.Lock()
	defer fake.logoutMutex.Unlock()
	fake.LogoutStub = stub
}

func (fake *FakeHandlerInterface) LogoutArgsForCall(i int) *gin.Context {
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	argsForCall := fake.logoutArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) OAuthRedirectHandler(arg1 *gin.Context) {
	fake.oAuthRedirectHandlerMutex.Lock()
	fake.oAuthRedirectHandlerArgsForCall = append(fake.oAuthRedirectHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.OAuthRedirectHandlerStub
	fake.recordInvocation("OAuthRedirectHandler", []interface{}{arg1})
	fake.oAuthRedirectHandlerMutex.Unlock()
	if stub != nil {
		fake.OAuthRedirectHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) OAuthRedirectHandlerCallCount() int {
	fake.oAuthRedirectHandlerMutex.RLock()
	defer fake.oAuthRedirectHandlerMutex.RUnlock()
	return len(fake.oAuthRedirectHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) OAuthRedirectHandlerCalls(stub func(*gin.Context)) {
	fake.oAuthRedirectHandlerMutex.Lock()
	defer fake.oAuthRedirectHandlerMutex.Unlock()
	fake.OAuthRedirectHandlerStub = stub
}

func (fake *FakeHandlerInterface) OAuthRedirectHandlerArgsForCall(i int) *gin.Context {
	fake.oAuthRedirectHandlerMutex.RLock()
	defer fake.oAuthRedirectHandlerMutex.RUnlock()
	argsForCall := fake.oAuthRedirectHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) OAuthStarterHandler(arg1 *gin.Context) {
	fake.oAuthStarterHandlerMutex.Lock()
	fake.oAuthStarterHandlerArgsForCall = append(fake.oAuthStarterHandlerArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.OAuthStarterHandlerStub
	fake.recordInvocation("OAuthStarterHandler", []interface{}{arg1})
	fake.oAuthStarterHandlerMutex.Unlock()
	if stub != nil {
		fake.OAuthStarterHandlerStub(arg1)
	}
}

func (fake *FakeHandlerInterface) OAuthStarterHandlerCallCount() int {
	fake.oAuthStarterHandlerMutex.RLock()
	defer fake.oAuthStarterHandlerMutex.RUnlock()
	return len(fake.oAuthStarterHandlerArgsForCall)
}

func (fake *FakeHandlerInterface) OAuthStarterHandlerCalls(stub func(*gin.Context)) {
	fake.oAuthStarterHandlerMutex.Lock()
	defer fake.oAuthStarterHandlerMutex.Unlock()
	fake.OAuthStarterHandlerStub = stub
}

func (fake *FakeHandlerInterface) OAuthStarterHandlerArgsForCall(i int) *gin.Context {
	fake.oAuthStarterHandlerMutex.RLock()
	defer fake.oAuthStarterHandlerMutex.RUnlock()
	argsForCall := fake.oAuthStarterHandlerArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) UpdateById(arg1 *gin.Context) {
	fake.updateByIdMutex.Lock()
	fake.updateByIdArgsForCall = append(fake.updateByIdArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.UpdateByIdStub
	fake.recordInvocation("UpdateById", []interface{}{arg1})
	fake.updateByIdMutex.Unlock()
	if stub != nil {
		fake.UpdateByIdStub(arg1)
	}
}

func (fake *FakeHandlerInterface) UpdateByIdCallCount() int {
	fake.updateByIdMutex.RLock()
	defer fake.updateByIdMutex.RUnlock()
	return len(fake.updateByIdArgsForCall)
}

func (fake *FakeHandlerInterface) UpdateByIdCalls(stub func(*gin.Context)) {
	fake.updateByIdMutex.Lock()
	defer fake.updateByIdMutex.Unlock()
	fake.UpdateByIdStub = stub
}

func (fake *FakeHandlerInterface) UpdateByIdArgsForCall(i int) *gin.Context {
	fake.updateByIdMutex.RLock()
	defer fake.updateByIdMutex.RUnlock()
	argsForCall := fake.updateByIdArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) ValidateToken(arg1 *gin.Context) {
	fake.validateTokenMutex.Lock()
	fake.validateTokenArgsForCall = append(fake.validateTokenArgsForCall, struct {
		arg1 *gin.Context
	}{arg1})
	stub := fake.ValidateTokenStub
	fake.recordInvocation("ValidateToken", []interface{}{arg1})
	fake.validateTokenMutex.Unlock()
	if stub != nil {
		fake.ValidateTokenStub(arg1)
	}
}

func (fake *FakeHandlerInterface) ValidateTokenCallCount() int {
	fake.validateTokenMutex.RLock()
	defer fake.validateTokenMutex.RUnlock()
	return len(fake.validateTokenArgsForCall)
}

func (fake *FakeHandlerInterface) ValidateTokenCalls(stub func(*gin.Context)) {
	fake.validateTokenMutex.Lock()
	defer fake.validateTokenMutex.Unlock()
	fake.ValidateTokenStub = stub
}

func (fake *FakeHandlerInterface) ValidateTokenArgsForCall(i int) *gin.Context {
	fake.validateTokenMutex.RLock()
	defer fake.validateTokenMutex.RUnlock()
	argsForCall := fake.validateTokenArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeHandlerInterface) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createEmployeeHandlerMutex.RLock()
	defer fake.createEmployeeHandlerMutex.RUnlock()
	fake.deleteByIdHandlerMutex.RLock()
	defer fake.deleteByIdHandlerMutex.RUnlock()
	fake.getAllEmployeesHandlerMutex.RLock()
	defer fake.getAllEmployeesHandlerMutex.RUnlock()
	fake.getEmployeeHandlerMutex.RLock()
	defer fake.getEmployeeHandlerMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.logoutMutex.RLock()
	defer fake.logoutMutex.RUnlock()
	fake.oAuthRedirectHandlerMutex.RLock()
	defer fake.oAuthRedirectHandlerMutex.RUnlock()
	fake.oAuthStarterHandlerMutex.RLock()
	defer fake.oAuthStarterHandlerMutex.RUnlock()
	fake.updateByIdMutex.RLock()
	defer fake.updateByIdMutex.RUnlock()
	fake.validateTokenMutex.RLock()
	defer fake.validateTokenMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeHandlerInterface) recordInvocation(key string, args []interface{}) {
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

var _ routes.HandlerInterface = new(FakeHandlerInterface)
