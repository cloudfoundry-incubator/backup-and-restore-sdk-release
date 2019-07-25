// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	gcs "gcs-blobstore-backup-restore"
	"sync"
)

type FakeBackupArtifact struct {
	ReadStub        func() (map[string]gcs.BucketBackup, error)
	readMutex       sync.RWMutex
	readArgsForCall []struct {
	}
	readReturns struct {
		result1 map[string]gcs.BucketBackup
		result2 error
	}
	readReturnsOnCall map[int]struct {
		result1 map[string]gcs.BucketBackup
		result2 error
	}
	WriteStub        func(map[string]gcs.BucketBackup) error
	writeMutex       sync.RWMutex
	writeArgsForCall []struct {
		arg1 map[string]gcs.BucketBackup
	}
	writeReturns struct {
		result1 error
	}
	writeReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBackupArtifact) Read() (map[string]gcs.BucketBackup, error) {
	fake.readMutex.Lock()
	ret, specificReturn := fake.readReturnsOnCall[len(fake.readArgsForCall)]
	fake.readArgsForCall = append(fake.readArgsForCall, struct {
	}{})
	fake.recordInvocation("Read", []interface{}{})
	fake.readMutex.Unlock()
	if fake.ReadStub != nil {
		return fake.ReadStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.readReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBackupArtifact) ReadCallCount() int {
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	return len(fake.readArgsForCall)
}

func (fake *FakeBackupArtifact) ReadCalls(stub func() (map[string]gcs.BucketBackup, error)) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = stub
}

func (fake *FakeBackupArtifact) ReadReturns(result1 map[string]gcs.BucketBackup, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	fake.readReturns = struct {
		result1 map[string]gcs.BucketBackup
		result2 error
	}{result1, result2}
}

func (fake *FakeBackupArtifact) ReadReturnsOnCall(i int, result1 map[string]gcs.BucketBackup, result2 error) {
	fake.readMutex.Lock()
	defer fake.readMutex.Unlock()
	fake.ReadStub = nil
	if fake.readReturnsOnCall == nil {
		fake.readReturnsOnCall = make(map[int]struct {
			result1 map[string]gcs.BucketBackup
			result2 error
		})
	}
	fake.readReturnsOnCall[i] = struct {
		result1 map[string]gcs.BucketBackup
		result2 error
	}{result1, result2}
}

func (fake *FakeBackupArtifact) Write(arg1 map[string]gcs.BucketBackup) error {
	fake.writeMutex.Lock()
	ret, specificReturn := fake.writeReturnsOnCall[len(fake.writeArgsForCall)]
	fake.writeArgsForCall = append(fake.writeArgsForCall, struct {
		arg1 map[string]gcs.BucketBackup
	}{arg1})
	fake.recordInvocation("Write", []interface{}{arg1})
	fake.writeMutex.Unlock()
	if fake.WriteStub != nil {
		return fake.WriteStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.writeReturns
	return fakeReturns.result1
}

func (fake *FakeBackupArtifact) WriteCallCount() int {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	return len(fake.writeArgsForCall)
}

func (fake *FakeBackupArtifact) WriteCalls(stub func(map[string]gcs.BucketBackup) error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = stub
}

func (fake *FakeBackupArtifact) WriteArgsForCall(i int) map[string]gcs.BucketBackup {
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	argsForCall := fake.writeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBackupArtifact) WriteReturns(result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	fake.writeReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackupArtifact) WriteReturnsOnCall(i int, result1 error) {
	fake.writeMutex.Lock()
	defer fake.writeMutex.Unlock()
	fake.WriteStub = nil
	if fake.writeReturnsOnCall == nil {
		fake.writeReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeBackupArtifact) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.readMutex.RLock()
	defer fake.readMutex.RUnlock()
	fake.writeMutex.RLock()
	defer fake.writeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBackupArtifact) recordInvocation(key string, args []interface{}) {
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

var _ gcs.BackupArtifact = new(FakeBackupArtifact)