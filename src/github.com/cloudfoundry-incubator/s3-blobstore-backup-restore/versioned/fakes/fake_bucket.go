// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"sync"

	"github.com/cloudfoundry-incubator/s3-blobstore-backup-restore/s3bucket"
	"github.com/cloudfoundry-incubator/s3-blobstore-backup-restore/versioned"
)

type FakeBucket struct {
	CopyVersionStub        func(string, string, string, string) error
	copyVersionMutex       sync.RWMutex
	copyVersionArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}
	copyVersionReturns struct {
		result1 error
	}
	copyVersionReturnsOnCall map[int]struct {
		result1 error
	}
	IsVersionedStub        func() (bool, error)
	isVersionedMutex       sync.RWMutex
	isVersionedArgsForCall []struct {
	}
	isVersionedReturns struct {
		result1 bool
		result2 error
	}
	isVersionedReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	ListVersionsStub        func() ([]s3bucket.Version, error)
	listVersionsMutex       sync.RWMutex
	listVersionsArgsForCall []struct {
	}
	listVersionsReturns struct {
		result1 []s3bucket.Version
		result2 error
	}
	listVersionsReturnsOnCall map[int]struct {
		result1 []s3bucket.Version
		result2 error
	}
	NameStub        func() string
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 string
	}
	nameReturnsOnCall map[int]struct {
		result1 string
	}
	RegionStub        func() string
	regionMutex       sync.RWMutex
	regionArgsForCall []struct {
	}
	regionReturns struct {
		result1 string
	}
	regionReturnsOnCall map[int]struct {
		result1 string
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBucket) CopyVersion(arg1 string, arg2 string, arg3 string, arg4 string) error {
	fake.copyVersionMutex.Lock()
	ret, specificReturn := fake.copyVersionReturnsOnCall[len(fake.copyVersionArgsForCall)]
	fake.copyVersionArgsForCall = append(fake.copyVersionArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 string
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("CopyVersion", []interface{}{arg1, arg2, arg3, arg4})
	fake.copyVersionMutex.Unlock()
	if fake.CopyVersionStub != nil {
		return fake.CopyVersionStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.copyVersionReturns
	return fakeReturns.result1
}

func (fake *FakeBucket) CopyVersionCallCount() int {
	fake.copyVersionMutex.RLock()
	defer fake.copyVersionMutex.RUnlock()
	return len(fake.copyVersionArgsForCall)
}

func (fake *FakeBucket) CopyVersionCalls(stub func(string, string, string, string) error) {
	fake.copyVersionMutex.Lock()
	defer fake.copyVersionMutex.Unlock()
	fake.CopyVersionStub = stub
}

func (fake *FakeBucket) CopyVersionArgsForCall(i int) (string, string, string, string) {
	fake.copyVersionMutex.RLock()
	defer fake.copyVersionMutex.RUnlock()
	argsForCall := fake.copyVersionArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeBucket) CopyVersionReturns(result1 error) {
	fake.copyVersionMutex.Lock()
	defer fake.copyVersionMutex.Unlock()
	fake.CopyVersionStub = nil
	fake.copyVersionReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) CopyVersionReturnsOnCall(i int, result1 error) {
	fake.copyVersionMutex.Lock()
	defer fake.copyVersionMutex.Unlock()
	fake.CopyVersionStub = nil
	if fake.copyVersionReturnsOnCall == nil {
		fake.copyVersionReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.copyVersionReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBucket) IsVersioned() (bool, error) {
	fake.isVersionedMutex.Lock()
	ret, specificReturn := fake.isVersionedReturnsOnCall[len(fake.isVersionedArgsForCall)]
	fake.isVersionedArgsForCall = append(fake.isVersionedArgsForCall, struct {
	}{})
	fake.recordInvocation("IsVersioned", []interface{}{})
	fake.isVersionedMutex.Unlock()
	if fake.IsVersionedStub != nil {
		return fake.IsVersionedStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.isVersionedReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBucket) IsVersionedCallCount() int {
	fake.isVersionedMutex.RLock()
	defer fake.isVersionedMutex.RUnlock()
	return len(fake.isVersionedArgsForCall)
}

func (fake *FakeBucket) IsVersionedCalls(stub func() (bool, error)) {
	fake.isVersionedMutex.Lock()
	defer fake.isVersionedMutex.Unlock()
	fake.IsVersionedStub = stub
}

func (fake *FakeBucket) IsVersionedReturns(result1 bool, result2 error) {
	fake.isVersionedMutex.Lock()
	defer fake.isVersionedMutex.Unlock()
	fake.IsVersionedStub = nil
	fake.isVersionedReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) IsVersionedReturnsOnCall(i int, result1 bool, result2 error) {
	fake.isVersionedMutex.Lock()
	defer fake.isVersionedMutex.Unlock()
	fake.IsVersionedStub = nil
	if fake.isVersionedReturnsOnCall == nil {
		fake.isVersionedReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.isVersionedReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) ListVersions() ([]s3bucket.Version, error) {
	fake.listVersionsMutex.Lock()
	ret, specificReturn := fake.listVersionsReturnsOnCall[len(fake.listVersionsArgsForCall)]
	fake.listVersionsArgsForCall = append(fake.listVersionsArgsForCall, struct {
	}{})
	fake.recordInvocation("ListVersions", []interface{}{})
	fake.listVersionsMutex.Unlock()
	if fake.ListVersionsStub != nil {
		return fake.ListVersionsStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.listVersionsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBucket) ListVersionsCallCount() int {
	fake.listVersionsMutex.RLock()
	defer fake.listVersionsMutex.RUnlock()
	return len(fake.listVersionsArgsForCall)
}

func (fake *FakeBucket) ListVersionsCalls(stub func() ([]s3bucket.Version, error)) {
	fake.listVersionsMutex.Lock()
	defer fake.listVersionsMutex.Unlock()
	fake.ListVersionsStub = stub
}

func (fake *FakeBucket) ListVersionsReturns(result1 []s3bucket.Version, result2 error) {
	fake.listVersionsMutex.Lock()
	defer fake.listVersionsMutex.Unlock()
	fake.ListVersionsStub = nil
	fake.listVersionsReturns = struct {
		result1 []s3bucket.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) ListVersionsReturnsOnCall(i int, result1 []s3bucket.Version, result2 error) {
	fake.listVersionsMutex.Lock()
	defer fake.listVersionsMutex.Unlock()
	fake.ListVersionsStub = nil
	if fake.listVersionsReturnsOnCall == nil {
		fake.listVersionsReturnsOnCall = make(map[int]struct {
			result1 []s3bucket.Version
			result2 error
		})
	}
	fake.listVersionsReturnsOnCall[i] = struct {
		result1 []s3bucket.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeBucket) Name() string {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *FakeBucket) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeBucket) NameCalls(stub func() string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeBucket) NameReturns(result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) NameReturnsOnCall(i int, result1 string) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) Region() string {
	fake.regionMutex.Lock()
	ret, specificReturn := fake.regionReturnsOnCall[len(fake.regionArgsForCall)]
	fake.regionArgsForCall = append(fake.regionArgsForCall, struct {
	}{})
	fake.recordInvocation("Region", []interface{}{})
	fake.regionMutex.Unlock()
	if fake.RegionStub != nil {
		return fake.RegionStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.regionReturns
	return fakeReturns.result1
}

func (fake *FakeBucket) RegionCallCount() int {
	fake.regionMutex.RLock()
	defer fake.regionMutex.RUnlock()
	return len(fake.regionArgsForCall)
}

func (fake *FakeBucket) RegionCalls(stub func() string) {
	fake.regionMutex.Lock()
	defer fake.regionMutex.Unlock()
	fake.RegionStub = stub
}

func (fake *FakeBucket) RegionReturns(result1 string) {
	fake.regionMutex.Lock()
	defer fake.regionMutex.Unlock()
	fake.RegionStub = nil
	fake.regionReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) RegionReturnsOnCall(i int, result1 string) {
	fake.regionMutex.Lock()
	defer fake.regionMutex.Unlock()
	fake.RegionStub = nil
	if fake.regionReturnsOnCall == nil {
		fake.regionReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.regionReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBucket) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.copyVersionMutex.RLock()
	defer fake.copyVersionMutex.RUnlock()
	fake.isVersionedMutex.RLock()
	defer fake.isVersionedMutex.RUnlock()
	fake.listVersionsMutex.RLock()
	defer fake.listVersionsMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.regionMutex.RLock()
	defer fake.regionMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBucket) recordInvocation(key string, args []interface{}) {
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

var _ versioned.Bucket = new(FakeBucket)
