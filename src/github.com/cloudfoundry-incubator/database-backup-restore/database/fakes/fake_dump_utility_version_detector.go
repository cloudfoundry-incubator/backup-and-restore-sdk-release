// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/database-backup-restore/database"
	"github.com/cloudfoundry-incubator/database-backup-restore/version"
)

type FakeDumpUtilityVersionDetector struct {
	GetVersionStub        func() (version.SemanticVersion, error)
	getVersionMutex       sync.RWMutex
	getVersionArgsForCall []struct {
	}
	getVersionReturns struct {
		result1 version.SemanticVersion
		result2 error
	}
	getVersionReturnsOnCall map[int]struct {
		result1 version.SemanticVersion
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDumpUtilityVersionDetector) GetVersion() (version.SemanticVersion, error) {
	fake.getVersionMutex.Lock()
	ret, specificReturn := fake.getVersionReturnsOnCall[len(fake.getVersionArgsForCall)]
	fake.getVersionArgsForCall = append(fake.getVersionArgsForCall, struct {
	}{})
	fake.recordInvocation("GetVersion", []interface{}{})
	fake.getVersionMutex.Unlock()
	if fake.GetVersionStub != nil {
		return fake.GetVersionStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.getVersionReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeDumpUtilityVersionDetector) GetVersionCallCount() int {
	fake.getVersionMutex.RLock()
	defer fake.getVersionMutex.RUnlock()
	return len(fake.getVersionArgsForCall)
}

func (fake *FakeDumpUtilityVersionDetector) GetVersionCalls(stub func() (version.SemanticVersion, error)) {
	fake.getVersionMutex.Lock()
	defer fake.getVersionMutex.Unlock()
	fake.GetVersionStub = stub
}

func (fake *FakeDumpUtilityVersionDetector) GetVersionReturns(result1 version.SemanticVersion, result2 error) {
	fake.getVersionMutex.Lock()
	defer fake.getVersionMutex.Unlock()
	fake.GetVersionStub = nil
	fake.getVersionReturns = struct {
		result1 version.SemanticVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeDumpUtilityVersionDetector) GetVersionReturnsOnCall(i int, result1 version.SemanticVersion, result2 error) {
	fake.getVersionMutex.Lock()
	defer fake.getVersionMutex.Unlock()
	fake.GetVersionStub = nil
	if fake.getVersionReturnsOnCall == nil {
		fake.getVersionReturnsOnCall = make(map[int]struct {
			result1 version.SemanticVersion
			result2 error
		})
	}
	fake.getVersionReturnsOnCall[i] = struct {
		result1 version.SemanticVersion
		result2 error
	}{result1, result2}
}

func (fake *FakeDumpUtilityVersionDetector) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getVersionMutex.RLock()
	defer fake.getVersionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDumpUtilityVersionDetector) recordInvocation(key string, args []interface{}) {
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

var _ database.DumpUtilityVersionDetector = new(FakeDumpUtilityVersionDetector)
