// This file was generated by counterfeiter
package dbngfakes

import (
	"sync"
	"time"

	"github.com/concourse/atc"
	"github.com/concourse/atc/dbng"
)

type FakeWorkerFactory struct {
	GetWorkerStub        func(name string) (*dbng.Worker, bool, error)
	getWorkerMutex       sync.RWMutex
	getWorkerArgsForCall []struct {
		name string
	}
	getWorkerReturns struct {
		result1 *dbng.Worker
		result2 bool
		result3 error
	}
	WorkersStub        func() ([]*dbng.Worker, error)
	workersMutex       sync.RWMutex
	workersArgsForCall []struct{}
	workersReturns     struct {
		result1 []*dbng.Worker
		result2 error
	}
	WorkersForTeamStub        func(teamName string) ([]*dbng.Worker, error)
	workersForTeamMutex       sync.RWMutex
	workersForTeamArgsForCall []struct {
		teamName string
	}
	workersForTeamReturns struct {
		result1 []*dbng.Worker
		result2 error
	}
	StallWorkerStub        func(name string) (*dbng.Worker, error)
	stallWorkerMutex       sync.RWMutex
	stallWorkerArgsForCall []struct {
		name string
	}
	stallWorkerReturns struct {
		result1 *dbng.Worker
		result2 error
	}
	StallUnresponsiveWorkersStub        func() ([]*dbng.Worker, error)
	stallUnresponsiveWorkersMutex       sync.RWMutex
	stallUnresponsiveWorkersArgsForCall []struct{}
	stallUnresponsiveWorkersReturns     struct {
		result1 []*dbng.Worker
		result2 error
	}
	DeleteFinishedRetiringWorkersStub        func() error
	deleteFinishedRetiringWorkersMutex       sync.RWMutex
	deleteFinishedRetiringWorkersArgsForCall []struct{}
	deleteFinishedRetiringWorkersReturns     struct {
		result1 error
	}
	LandFinishedLandingWorkersStub        func() error
	landFinishedLandingWorkersMutex       sync.RWMutex
	landFinishedLandingWorkersArgsForCall []struct{}
	landFinishedLandingWorkersReturns     struct {
		result1 error
	}
	SaveWorkerStub        func(worker atc.Worker, ttl time.Duration) (*dbng.Worker, error)
	saveWorkerMutex       sync.RWMutex
	saveWorkerArgsForCall []struct {
		worker atc.Worker
		ttl    time.Duration
	}
	saveWorkerReturns struct {
		result1 *dbng.Worker
		result2 error
	}
	LandWorkerStub        func(name string) (*dbng.Worker, error)
	landWorkerMutex       sync.RWMutex
	landWorkerArgsForCall []struct {
		name string
	}
	landWorkerReturns struct {
		result1 *dbng.Worker
		result2 error
	}
	RetireWorkerStub        func(name string) (*dbng.Worker, error)
	retireWorkerMutex       sync.RWMutex
	retireWorkerArgsForCall []struct {
		name string
	}
	retireWorkerReturns struct {
		result1 *dbng.Worker
		result2 error
	}
	PruneWorkerStub        func(name string) error
	pruneWorkerMutex       sync.RWMutex
	pruneWorkerArgsForCall []struct {
		name string
	}
	pruneWorkerReturns struct {
		result1 error
	}
	DeleteWorkerStub        func(name string) error
	deleteWorkerMutex       sync.RWMutex
	deleteWorkerArgsForCall []struct {
		name string
	}
	deleteWorkerReturns struct {
		result1 error
	}
	HeartbeatWorkerStub        func(worker atc.Worker, ttl time.Duration) (*dbng.Worker, error)
	heartbeatWorkerMutex       sync.RWMutex
	heartbeatWorkerArgsForCall []struct {
		worker atc.Worker
		ttl    time.Duration
	}
	heartbeatWorkerReturns struct {
		result1 *dbng.Worker
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeWorkerFactory) GetWorker(name string) (*dbng.Worker, bool, error) {
	fake.getWorkerMutex.Lock()
	fake.getWorkerArgsForCall = append(fake.getWorkerArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("GetWorker", []interface{}{name})
	fake.getWorkerMutex.Unlock()
	if fake.GetWorkerStub != nil {
		return fake.GetWorkerStub(name)
	} else {
		return fake.getWorkerReturns.result1, fake.getWorkerReturns.result2, fake.getWorkerReturns.result3
	}
}

func (fake *FakeWorkerFactory) GetWorkerCallCount() int {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return len(fake.getWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) GetWorkerArgsForCall(i int) string {
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	return fake.getWorkerArgsForCall[i].name
}

func (fake *FakeWorkerFactory) GetWorkerReturns(result1 *dbng.Worker, result2 bool, result3 error) {
	fake.GetWorkerStub = nil
	fake.getWorkerReturns = struct {
		result1 *dbng.Worker
		result2 bool
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeWorkerFactory) Workers() ([]*dbng.Worker, error) {
	fake.workersMutex.Lock()
	fake.workersArgsForCall = append(fake.workersArgsForCall, struct{}{})
	fake.recordInvocation("Workers", []interface{}{})
	fake.workersMutex.Unlock()
	if fake.WorkersStub != nil {
		return fake.WorkersStub()
	} else {
		return fake.workersReturns.result1, fake.workersReturns.result2
	}
}

func (fake *FakeWorkerFactory) WorkersCallCount() int {
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	return len(fake.workersArgsForCall)
}

func (fake *FakeWorkerFactory) WorkersReturns(result1 []*dbng.Worker, result2 error) {
	fake.WorkersStub = nil
	fake.workersReturns = struct {
		result1 []*dbng.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) WorkersForTeam(teamName string) ([]*dbng.Worker, error) {
	fake.workersForTeamMutex.Lock()
	fake.workersForTeamArgsForCall = append(fake.workersForTeamArgsForCall, struct {
		teamName string
	}{teamName})
	fake.recordInvocation("WorkersForTeam", []interface{}{teamName})
	fake.workersForTeamMutex.Unlock()
	if fake.WorkersForTeamStub != nil {
		return fake.WorkersForTeamStub(teamName)
	} else {
		return fake.workersForTeamReturns.result1, fake.workersForTeamReturns.result2
	}
}

func (fake *FakeWorkerFactory) WorkersForTeamCallCount() int {
	fake.workersForTeamMutex.RLock()
	defer fake.workersForTeamMutex.RUnlock()
	return len(fake.workersForTeamArgsForCall)
}

func (fake *FakeWorkerFactory) WorkersForTeamArgsForCall(i int) string {
	fake.workersForTeamMutex.RLock()
	defer fake.workersForTeamMutex.RUnlock()
	return fake.workersForTeamArgsForCall[i].teamName
}

func (fake *FakeWorkerFactory) WorkersForTeamReturns(result1 []*dbng.Worker, result2 error) {
	fake.WorkersForTeamStub = nil
	fake.workersForTeamReturns = struct {
		result1 []*dbng.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) StallWorker(name string) (*dbng.Worker, error) {
	fake.stallWorkerMutex.Lock()
	fake.stallWorkerArgsForCall = append(fake.stallWorkerArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("StallWorker", []interface{}{name})
	fake.stallWorkerMutex.Unlock()
	if fake.StallWorkerStub != nil {
		return fake.StallWorkerStub(name)
	} else {
		return fake.stallWorkerReturns.result1, fake.stallWorkerReturns.result2
	}
}

func (fake *FakeWorkerFactory) StallWorkerCallCount() int {
	fake.stallWorkerMutex.RLock()
	defer fake.stallWorkerMutex.RUnlock()
	return len(fake.stallWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) StallWorkerArgsForCall(i int) string {
	fake.stallWorkerMutex.RLock()
	defer fake.stallWorkerMutex.RUnlock()
	return fake.stallWorkerArgsForCall[i].name
}

func (fake *FakeWorkerFactory) StallWorkerReturns(result1 *dbng.Worker, result2 error) {
	fake.StallWorkerStub = nil
	fake.stallWorkerReturns = struct {
		result1 *dbng.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) StallUnresponsiveWorkers() ([]*dbng.Worker, error) {
	fake.stallUnresponsiveWorkersMutex.Lock()
	fake.stallUnresponsiveWorkersArgsForCall = append(fake.stallUnresponsiveWorkersArgsForCall, struct{}{})
	fake.recordInvocation("StallUnresponsiveWorkers", []interface{}{})
	fake.stallUnresponsiveWorkersMutex.Unlock()
	if fake.StallUnresponsiveWorkersStub != nil {
		return fake.StallUnresponsiveWorkersStub()
	} else {
		return fake.stallUnresponsiveWorkersReturns.result1, fake.stallUnresponsiveWorkersReturns.result2
	}
}

func (fake *FakeWorkerFactory) StallUnresponsiveWorkersCallCount() int {
	fake.stallUnresponsiveWorkersMutex.RLock()
	defer fake.stallUnresponsiveWorkersMutex.RUnlock()
	return len(fake.stallUnresponsiveWorkersArgsForCall)
}

func (fake *FakeWorkerFactory) StallUnresponsiveWorkersReturns(result1 []*dbng.Worker, result2 error) {
	fake.StallUnresponsiveWorkersStub = nil
	fake.stallUnresponsiveWorkersReturns = struct {
		result1 []*dbng.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) DeleteFinishedRetiringWorkers() error {
	fake.deleteFinishedRetiringWorkersMutex.Lock()
	fake.deleteFinishedRetiringWorkersArgsForCall = append(fake.deleteFinishedRetiringWorkersArgsForCall, struct{}{})
	fake.recordInvocation("DeleteFinishedRetiringWorkers", []interface{}{})
	fake.deleteFinishedRetiringWorkersMutex.Unlock()
	if fake.DeleteFinishedRetiringWorkersStub != nil {
		return fake.DeleteFinishedRetiringWorkersStub()
	} else {
		return fake.deleteFinishedRetiringWorkersReturns.result1
	}
}

func (fake *FakeWorkerFactory) DeleteFinishedRetiringWorkersCallCount() int {
	fake.deleteFinishedRetiringWorkersMutex.RLock()
	defer fake.deleteFinishedRetiringWorkersMutex.RUnlock()
	return len(fake.deleteFinishedRetiringWorkersArgsForCall)
}

func (fake *FakeWorkerFactory) DeleteFinishedRetiringWorkersReturns(result1 error) {
	fake.DeleteFinishedRetiringWorkersStub = nil
	fake.deleteFinishedRetiringWorkersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerFactory) LandFinishedLandingWorkers() error {
	fake.landFinishedLandingWorkersMutex.Lock()
	fake.landFinishedLandingWorkersArgsForCall = append(fake.landFinishedLandingWorkersArgsForCall, struct{}{})
	fake.recordInvocation("LandFinishedLandingWorkers", []interface{}{})
	fake.landFinishedLandingWorkersMutex.Unlock()
	if fake.LandFinishedLandingWorkersStub != nil {
		return fake.LandFinishedLandingWorkersStub()
	} else {
		return fake.landFinishedLandingWorkersReturns.result1
	}
}

func (fake *FakeWorkerFactory) LandFinishedLandingWorkersCallCount() int {
	fake.landFinishedLandingWorkersMutex.RLock()
	defer fake.landFinishedLandingWorkersMutex.RUnlock()
	return len(fake.landFinishedLandingWorkersArgsForCall)
}

func (fake *FakeWorkerFactory) LandFinishedLandingWorkersReturns(result1 error) {
	fake.LandFinishedLandingWorkersStub = nil
	fake.landFinishedLandingWorkersReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerFactory) SaveWorker(worker atc.Worker, ttl time.Duration) (*dbng.Worker, error) {
	fake.saveWorkerMutex.Lock()
	fake.saveWorkerArgsForCall = append(fake.saveWorkerArgsForCall, struct {
		worker atc.Worker
		ttl    time.Duration
	}{worker, ttl})
	fake.recordInvocation("SaveWorker", []interface{}{worker, ttl})
	fake.saveWorkerMutex.Unlock()
	if fake.SaveWorkerStub != nil {
		return fake.SaveWorkerStub(worker, ttl)
	} else {
		return fake.saveWorkerReturns.result1, fake.saveWorkerReturns.result2
	}
}

func (fake *FakeWorkerFactory) SaveWorkerCallCount() int {
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	return len(fake.saveWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) SaveWorkerArgsForCall(i int) (atc.Worker, time.Duration) {
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	return fake.saveWorkerArgsForCall[i].worker, fake.saveWorkerArgsForCall[i].ttl
}

func (fake *FakeWorkerFactory) SaveWorkerReturns(result1 *dbng.Worker, result2 error) {
	fake.SaveWorkerStub = nil
	fake.saveWorkerReturns = struct {
		result1 *dbng.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) LandWorker(name string) (*dbng.Worker, error) {
	fake.landWorkerMutex.Lock()
	fake.landWorkerArgsForCall = append(fake.landWorkerArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("LandWorker", []interface{}{name})
	fake.landWorkerMutex.Unlock()
	if fake.LandWorkerStub != nil {
		return fake.LandWorkerStub(name)
	} else {
		return fake.landWorkerReturns.result1, fake.landWorkerReturns.result2
	}
}

func (fake *FakeWorkerFactory) LandWorkerCallCount() int {
	fake.landWorkerMutex.RLock()
	defer fake.landWorkerMutex.RUnlock()
	return len(fake.landWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) LandWorkerArgsForCall(i int) string {
	fake.landWorkerMutex.RLock()
	defer fake.landWorkerMutex.RUnlock()
	return fake.landWorkerArgsForCall[i].name
}

func (fake *FakeWorkerFactory) LandWorkerReturns(result1 *dbng.Worker, result2 error) {
	fake.LandWorkerStub = nil
	fake.landWorkerReturns = struct {
		result1 *dbng.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) RetireWorker(name string) (*dbng.Worker, error) {
	fake.retireWorkerMutex.Lock()
	fake.retireWorkerArgsForCall = append(fake.retireWorkerArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("RetireWorker", []interface{}{name})
	fake.retireWorkerMutex.Unlock()
	if fake.RetireWorkerStub != nil {
		return fake.RetireWorkerStub(name)
	} else {
		return fake.retireWorkerReturns.result1, fake.retireWorkerReturns.result2
	}
}

func (fake *FakeWorkerFactory) RetireWorkerCallCount() int {
	fake.retireWorkerMutex.RLock()
	defer fake.retireWorkerMutex.RUnlock()
	return len(fake.retireWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) RetireWorkerArgsForCall(i int) string {
	fake.retireWorkerMutex.RLock()
	defer fake.retireWorkerMutex.RUnlock()
	return fake.retireWorkerArgsForCall[i].name
}

func (fake *FakeWorkerFactory) RetireWorkerReturns(result1 *dbng.Worker, result2 error) {
	fake.RetireWorkerStub = nil
	fake.retireWorkerReturns = struct {
		result1 *dbng.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) PruneWorker(name string) error {
	fake.pruneWorkerMutex.Lock()
	fake.pruneWorkerArgsForCall = append(fake.pruneWorkerArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("PruneWorker", []interface{}{name})
	fake.pruneWorkerMutex.Unlock()
	if fake.PruneWorkerStub != nil {
		return fake.PruneWorkerStub(name)
	} else {
		return fake.pruneWorkerReturns.result1
	}
}

func (fake *FakeWorkerFactory) PruneWorkerCallCount() int {
	fake.pruneWorkerMutex.RLock()
	defer fake.pruneWorkerMutex.RUnlock()
	return len(fake.pruneWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) PruneWorkerArgsForCall(i int) string {
	fake.pruneWorkerMutex.RLock()
	defer fake.pruneWorkerMutex.RUnlock()
	return fake.pruneWorkerArgsForCall[i].name
}

func (fake *FakeWorkerFactory) PruneWorkerReturns(result1 error) {
	fake.PruneWorkerStub = nil
	fake.pruneWorkerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerFactory) DeleteWorker(name string) error {
	fake.deleteWorkerMutex.Lock()
	fake.deleteWorkerArgsForCall = append(fake.deleteWorkerArgsForCall, struct {
		name string
	}{name})
	fake.recordInvocation("DeleteWorker", []interface{}{name})
	fake.deleteWorkerMutex.Unlock()
	if fake.DeleteWorkerStub != nil {
		return fake.DeleteWorkerStub(name)
	} else {
		return fake.deleteWorkerReturns.result1
	}
}

func (fake *FakeWorkerFactory) DeleteWorkerCallCount() int {
	fake.deleteWorkerMutex.RLock()
	defer fake.deleteWorkerMutex.RUnlock()
	return len(fake.deleteWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) DeleteWorkerArgsForCall(i int) string {
	fake.deleteWorkerMutex.RLock()
	defer fake.deleteWorkerMutex.RUnlock()
	return fake.deleteWorkerArgsForCall[i].name
}

func (fake *FakeWorkerFactory) DeleteWorkerReturns(result1 error) {
	fake.DeleteWorkerStub = nil
	fake.deleteWorkerReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeWorkerFactory) HeartbeatWorker(worker atc.Worker, ttl time.Duration) (*dbng.Worker, error) {
	fake.heartbeatWorkerMutex.Lock()
	fake.heartbeatWorkerArgsForCall = append(fake.heartbeatWorkerArgsForCall, struct {
		worker atc.Worker
		ttl    time.Duration
	}{worker, ttl})
	fake.recordInvocation("HeartbeatWorker", []interface{}{worker, ttl})
	fake.heartbeatWorkerMutex.Unlock()
	if fake.HeartbeatWorkerStub != nil {
		return fake.HeartbeatWorkerStub(worker, ttl)
	} else {
		return fake.heartbeatWorkerReturns.result1, fake.heartbeatWorkerReturns.result2
	}
}

func (fake *FakeWorkerFactory) HeartbeatWorkerCallCount() int {
	fake.heartbeatWorkerMutex.RLock()
	defer fake.heartbeatWorkerMutex.RUnlock()
	return len(fake.heartbeatWorkerArgsForCall)
}

func (fake *FakeWorkerFactory) HeartbeatWorkerArgsForCall(i int) (atc.Worker, time.Duration) {
	fake.heartbeatWorkerMutex.RLock()
	defer fake.heartbeatWorkerMutex.RUnlock()
	return fake.heartbeatWorkerArgsForCall[i].worker, fake.heartbeatWorkerArgsForCall[i].ttl
}

func (fake *FakeWorkerFactory) HeartbeatWorkerReturns(result1 *dbng.Worker, result2 error) {
	fake.HeartbeatWorkerStub = nil
	fake.heartbeatWorkerReturns = struct {
		result1 *dbng.Worker
		result2 error
	}{result1, result2}
}

func (fake *FakeWorkerFactory) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getWorkerMutex.RLock()
	defer fake.getWorkerMutex.RUnlock()
	fake.workersMutex.RLock()
	defer fake.workersMutex.RUnlock()
	fake.workersForTeamMutex.RLock()
	defer fake.workersForTeamMutex.RUnlock()
	fake.stallWorkerMutex.RLock()
	defer fake.stallWorkerMutex.RUnlock()
	fake.stallUnresponsiveWorkersMutex.RLock()
	defer fake.stallUnresponsiveWorkersMutex.RUnlock()
	fake.deleteFinishedRetiringWorkersMutex.RLock()
	defer fake.deleteFinishedRetiringWorkersMutex.RUnlock()
	fake.landFinishedLandingWorkersMutex.RLock()
	defer fake.landFinishedLandingWorkersMutex.RUnlock()
	fake.saveWorkerMutex.RLock()
	defer fake.saveWorkerMutex.RUnlock()
	fake.landWorkerMutex.RLock()
	defer fake.landWorkerMutex.RUnlock()
	fake.retireWorkerMutex.RLock()
	defer fake.retireWorkerMutex.RUnlock()
	fake.pruneWorkerMutex.RLock()
	defer fake.pruneWorkerMutex.RUnlock()
	fake.deleteWorkerMutex.RLock()
	defer fake.deleteWorkerMutex.RUnlock()
	fake.heartbeatWorkerMutex.RLock()
	defer fake.heartbeatWorkerMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeWorkerFactory) recordInvocation(key string, args []interface{}) {
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

var _ dbng.WorkerFactory = new(FakeWorkerFactory)
