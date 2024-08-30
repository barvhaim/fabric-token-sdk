// Code generated by counterfeiter. DO NOT EDIT.
package mock

import (
	"sync"

	"github.com/hyperledger-labs/fabric-token-sdk/token/services/network/common/rws/translator"
	"github.com/hyperledger-labs/fabric-token-sdk/token/token"
)

type TransferAction struct {
	GetInputsStub        func() ([]*token.ID, error)
	getInputsMutex       sync.RWMutex
	getInputsArgsForCall []struct {
	}
	getInputsReturns struct {
		result1 []*token.ID
		result2 error
	}
	getInputsReturnsOnCall map[int]struct {
		result1 []*token.ID
		result2 error
	}
	GetMetadataStub        func() map[string][]byte
	getMetadataMutex       sync.RWMutex
	getMetadataArgsForCall []struct {
	}
	getMetadataReturns struct {
		result1 map[string][]byte
	}
	getMetadataReturnsOnCall map[int]struct {
		result1 map[string][]byte
	}
	GetSerialNumbersStub        func() []string
	getSerialNumbersMutex       sync.RWMutex
	getSerialNumbersArgsForCall []struct {
	}
	getSerialNumbersReturns struct {
		result1 []string
	}
	getSerialNumbersReturnsOnCall map[int]struct {
		result1 []string
	}
	GetSerializedInputsStub        func() ([][]byte, error)
	getSerializedInputsMutex       sync.RWMutex
	getSerializedInputsArgsForCall []struct {
	}
	getSerializedInputsReturns struct {
		result1 [][]byte
		result2 error
	}
	getSerializedInputsReturnsOnCall map[int]struct {
		result1 [][]byte
		result2 error
	}
	GetSerializedOutputsStub        func() ([][]byte, error)
	getSerializedOutputsMutex       sync.RWMutex
	getSerializedOutputsArgsForCall []struct {
	}
	getSerializedOutputsReturns struct {
		result1 [][]byte
		result2 error
	}
	getSerializedOutputsReturnsOnCall map[int]struct {
		result1 [][]byte
		result2 error
	}
	IsGraphHidingStub        func() bool
	isGraphHidingMutex       sync.RWMutex
	isGraphHidingArgsForCall []struct {
	}
	isGraphHidingReturns struct {
		result1 bool
	}
	isGraphHidingReturnsOnCall map[int]struct {
		result1 bool
	}
	IsRedeemAtStub        func(int) bool
	isRedeemAtMutex       sync.RWMutex
	isRedeemAtArgsForCall []struct {
		arg1 int
	}
	isRedeemAtReturns struct {
		result1 bool
	}
	isRedeemAtReturnsOnCall map[int]struct {
		result1 bool
	}
	NumOutputsStub        func() int
	numOutputsMutex       sync.RWMutex
	numOutputsArgsForCall []struct {
	}
	numOutputsReturns struct {
		result1 int
	}
	numOutputsReturnsOnCall map[int]struct {
		result1 int
	}
	SerializeStub        func() ([]byte, error)
	serializeMutex       sync.RWMutex
	serializeArgsForCall []struct {
	}
	serializeReturns struct {
		result1 []byte
		result2 error
	}
	serializeReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	SerializeOutputAtStub        func(int) ([]byte, error)
	serializeOutputAtMutex       sync.RWMutex
	serializeOutputAtArgsForCall []struct {
		arg1 int
	}
	serializeOutputAtReturns struct {
		result1 []byte
		result2 error
	}
	serializeOutputAtReturnsOnCall map[int]struct {
		result1 []byte
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TransferAction) GetInputs() ([]*token.ID, error) {
	fake.getInputsMutex.Lock()
	ret, specificReturn := fake.getInputsReturnsOnCall[len(fake.getInputsArgsForCall)]
	fake.getInputsArgsForCall = append(fake.getInputsArgsForCall, struct {
	}{})
	stub := fake.GetInputsStub
	fakeReturns := fake.getInputsReturns
	fake.recordInvocation("GetInputs", []interface{}{})
	fake.getInputsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TransferAction) GetInputsCallCount() int {
	fake.getInputsMutex.RLock()
	defer fake.getInputsMutex.RUnlock()
	return len(fake.getInputsArgsForCall)
}

func (fake *TransferAction) GetInputsCalls(stub func() ([]*token.ID, error)) {
	fake.getInputsMutex.Lock()
	defer fake.getInputsMutex.Unlock()
	fake.GetInputsStub = stub
}

func (fake *TransferAction) GetInputsReturns(result1 []*token.ID, result2 error) {
	fake.getInputsMutex.Lock()
	defer fake.getInputsMutex.Unlock()
	fake.GetInputsStub = nil
	fake.getInputsReturns = struct {
		result1 []*token.ID
		result2 error
	}{result1, result2}
}

func (fake *TransferAction) GetInputsReturnsOnCall(i int, result1 []*token.ID, result2 error) {
	fake.getInputsMutex.Lock()
	defer fake.getInputsMutex.Unlock()
	fake.GetInputsStub = nil
	if fake.getInputsReturnsOnCall == nil {
		fake.getInputsReturnsOnCall = make(map[int]struct {
			result1 []*token.ID
			result2 error
		})
	}
	fake.getInputsReturnsOnCall[i] = struct {
		result1 []*token.ID
		result2 error
	}{result1, result2}
}

func (fake *TransferAction) GetMetadata() map[string][]byte {
	fake.getMetadataMutex.Lock()
	ret, specificReturn := fake.getMetadataReturnsOnCall[len(fake.getMetadataArgsForCall)]
	fake.getMetadataArgsForCall = append(fake.getMetadataArgsForCall, struct {
	}{})
	stub := fake.GetMetadataStub
	fakeReturns := fake.getMetadataReturns
	fake.recordInvocation("GetMetadata", []interface{}{})
	fake.getMetadataMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TransferAction) GetMetadataCallCount() int {
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	return len(fake.getMetadataArgsForCall)
}

func (fake *TransferAction) GetMetadataCalls(stub func() map[string][]byte) {
	fake.getMetadataMutex.Lock()
	defer fake.getMetadataMutex.Unlock()
	fake.GetMetadataStub = stub
}

func (fake *TransferAction) GetMetadataReturns(result1 map[string][]byte) {
	fake.getMetadataMutex.Lock()
	defer fake.getMetadataMutex.Unlock()
	fake.GetMetadataStub = nil
	fake.getMetadataReturns = struct {
		result1 map[string][]byte
	}{result1}
}

func (fake *TransferAction) GetMetadataReturnsOnCall(i int, result1 map[string][]byte) {
	fake.getMetadataMutex.Lock()
	defer fake.getMetadataMutex.Unlock()
	fake.GetMetadataStub = nil
	if fake.getMetadataReturnsOnCall == nil {
		fake.getMetadataReturnsOnCall = make(map[int]struct {
			result1 map[string][]byte
		})
	}
	fake.getMetadataReturnsOnCall[i] = struct {
		result1 map[string][]byte
	}{result1}
}

func (fake *TransferAction) GetSerialNumbers() []string {
	fake.getSerialNumbersMutex.Lock()
	ret, specificReturn := fake.getSerialNumbersReturnsOnCall[len(fake.getSerialNumbersArgsForCall)]
	fake.getSerialNumbersArgsForCall = append(fake.getSerialNumbersArgsForCall, struct {
	}{})
	stub := fake.GetSerialNumbersStub
	fakeReturns := fake.getSerialNumbersReturns
	fake.recordInvocation("GetSerialNumbers", []interface{}{})
	fake.getSerialNumbersMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TransferAction) GetSerialNumbersCallCount() int {
	fake.getSerialNumbersMutex.RLock()
	defer fake.getSerialNumbersMutex.RUnlock()
	return len(fake.getSerialNumbersArgsForCall)
}

func (fake *TransferAction) GetSerialNumbersCalls(stub func() []string) {
	fake.getSerialNumbersMutex.Lock()
	defer fake.getSerialNumbersMutex.Unlock()
	fake.GetSerialNumbersStub = stub
}

func (fake *TransferAction) GetSerialNumbersReturns(result1 []string) {
	fake.getSerialNumbersMutex.Lock()
	defer fake.getSerialNumbersMutex.Unlock()
	fake.GetSerialNumbersStub = nil
	fake.getSerialNumbersReturns = struct {
		result1 []string
	}{result1}
}

func (fake *TransferAction) GetSerialNumbersReturnsOnCall(i int, result1 []string) {
	fake.getSerialNumbersMutex.Lock()
	defer fake.getSerialNumbersMutex.Unlock()
	fake.GetSerialNumbersStub = nil
	if fake.getSerialNumbersReturnsOnCall == nil {
		fake.getSerialNumbersReturnsOnCall = make(map[int]struct {
			result1 []string
		})
	}
	fake.getSerialNumbersReturnsOnCall[i] = struct {
		result1 []string
	}{result1}
}

func (fake *TransferAction) GetSerializedInputs() ([][]byte, error) {
	fake.getSerializedInputsMutex.Lock()
	ret, specificReturn := fake.getSerializedInputsReturnsOnCall[len(fake.getSerializedInputsArgsForCall)]
	fake.getSerializedInputsArgsForCall = append(fake.getSerializedInputsArgsForCall, struct {
	}{})
	stub := fake.GetSerializedInputsStub
	fakeReturns := fake.getSerializedInputsReturns
	fake.recordInvocation("GetSerializedInputs", []interface{}{})
	fake.getSerializedInputsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TransferAction) GetSerializedInputsCallCount() int {
	fake.getSerializedInputsMutex.RLock()
	defer fake.getSerializedInputsMutex.RUnlock()
	return len(fake.getSerializedInputsArgsForCall)
}

func (fake *TransferAction) GetSerializedInputsCalls(stub func() ([][]byte, error)) {
	fake.getSerializedInputsMutex.Lock()
	defer fake.getSerializedInputsMutex.Unlock()
	fake.GetSerializedInputsStub = stub
}

func (fake *TransferAction) GetSerializedInputsReturns(result1 [][]byte, result2 error) {
	fake.getSerializedInputsMutex.Lock()
	defer fake.getSerializedInputsMutex.Unlock()
	fake.GetSerializedInputsStub = nil
	fake.getSerializedInputsReturns = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *TransferAction) GetSerializedInputsReturnsOnCall(i int, result1 [][]byte, result2 error) {
	fake.getSerializedInputsMutex.Lock()
	defer fake.getSerializedInputsMutex.Unlock()
	fake.GetSerializedInputsStub = nil
	if fake.getSerializedInputsReturnsOnCall == nil {
		fake.getSerializedInputsReturnsOnCall = make(map[int]struct {
			result1 [][]byte
			result2 error
		})
	}
	fake.getSerializedInputsReturnsOnCall[i] = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *TransferAction) GetSerializedOutputs() ([][]byte, error) {
	fake.getSerializedOutputsMutex.Lock()
	ret, specificReturn := fake.getSerializedOutputsReturnsOnCall[len(fake.getSerializedOutputsArgsForCall)]
	fake.getSerializedOutputsArgsForCall = append(fake.getSerializedOutputsArgsForCall, struct {
	}{})
	stub := fake.GetSerializedOutputsStub
	fakeReturns := fake.getSerializedOutputsReturns
	fake.recordInvocation("GetSerializedOutputs", []interface{}{})
	fake.getSerializedOutputsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TransferAction) GetSerializedOutputsCallCount() int {
	fake.getSerializedOutputsMutex.RLock()
	defer fake.getSerializedOutputsMutex.RUnlock()
	return len(fake.getSerializedOutputsArgsForCall)
}

func (fake *TransferAction) GetSerializedOutputsCalls(stub func() ([][]byte, error)) {
	fake.getSerializedOutputsMutex.Lock()
	defer fake.getSerializedOutputsMutex.Unlock()
	fake.GetSerializedOutputsStub = stub
}

func (fake *TransferAction) GetSerializedOutputsReturns(result1 [][]byte, result2 error) {
	fake.getSerializedOutputsMutex.Lock()
	defer fake.getSerializedOutputsMutex.Unlock()
	fake.GetSerializedOutputsStub = nil
	fake.getSerializedOutputsReturns = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *TransferAction) GetSerializedOutputsReturnsOnCall(i int, result1 [][]byte, result2 error) {
	fake.getSerializedOutputsMutex.Lock()
	defer fake.getSerializedOutputsMutex.Unlock()
	fake.GetSerializedOutputsStub = nil
	if fake.getSerializedOutputsReturnsOnCall == nil {
		fake.getSerializedOutputsReturnsOnCall = make(map[int]struct {
			result1 [][]byte
			result2 error
		})
	}
	fake.getSerializedOutputsReturnsOnCall[i] = struct {
		result1 [][]byte
		result2 error
	}{result1, result2}
}

func (fake *TransferAction) IsGraphHiding() bool {
	fake.isGraphHidingMutex.Lock()
	ret, specificReturn := fake.isGraphHidingReturnsOnCall[len(fake.isGraphHidingArgsForCall)]
	fake.isGraphHidingArgsForCall = append(fake.isGraphHidingArgsForCall, struct {
	}{})
	stub := fake.IsGraphHidingStub
	fakeReturns := fake.isGraphHidingReturns
	fake.recordInvocation("IsGraphHiding", []interface{}{})
	fake.isGraphHidingMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TransferAction) IsGraphHidingCallCount() int {
	fake.isGraphHidingMutex.RLock()
	defer fake.isGraphHidingMutex.RUnlock()
	return len(fake.isGraphHidingArgsForCall)
}

func (fake *TransferAction) IsGraphHidingCalls(stub func() bool) {
	fake.isGraphHidingMutex.Lock()
	defer fake.isGraphHidingMutex.Unlock()
	fake.IsGraphHidingStub = stub
}

func (fake *TransferAction) IsGraphHidingReturns(result1 bool) {
	fake.isGraphHidingMutex.Lock()
	defer fake.isGraphHidingMutex.Unlock()
	fake.IsGraphHidingStub = nil
	fake.isGraphHidingReturns = struct {
		result1 bool
	}{result1}
}

func (fake *TransferAction) IsGraphHidingReturnsOnCall(i int, result1 bool) {
	fake.isGraphHidingMutex.Lock()
	defer fake.isGraphHidingMutex.Unlock()
	fake.IsGraphHidingStub = nil
	if fake.isGraphHidingReturnsOnCall == nil {
		fake.isGraphHidingReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isGraphHidingReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *TransferAction) IsRedeemAt(arg1 int) bool {
	fake.isRedeemAtMutex.Lock()
	ret, specificReturn := fake.isRedeemAtReturnsOnCall[len(fake.isRedeemAtArgsForCall)]
	fake.isRedeemAtArgsForCall = append(fake.isRedeemAtArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.IsRedeemAtStub
	fakeReturns := fake.isRedeemAtReturns
	fake.recordInvocation("IsRedeemAt", []interface{}{arg1})
	fake.isRedeemAtMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TransferAction) IsRedeemAtCallCount() int {
	fake.isRedeemAtMutex.RLock()
	defer fake.isRedeemAtMutex.RUnlock()
	return len(fake.isRedeemAtArgsForCall)
}

func (fake *TransferAction) IsRedeemAtCalls(stub func(int) bool) {
	fake.isRedeemAtMutex.Lock()
	defer fake.isRedeemAtMutex.Unlock()
	fake.IsRedeemAtStub = stub
}

func (fake *TransferAction) IsRedeemAtArgsForCall(i int) int {
	fake.isRedeemAtMutex.RLock()
	defer fake.isRedeemAtMutex.RUnlock()
	argsForCall := fake.isRedeemAtArgsForCall[i]
	return argsForCall.arg1
}

func (fake *TransferAction) IsRedeemAtReturns(result1 bool) {
	fake.isRedeemAtMutex.Lock()
	defer fake.isRedeemAtMutex.Unlock()
	fake.IsRedeemAtStub = nil
	fake.isRedeemAtReturns = struct {
		result1 bool
	}{result1}
}

func (fake *TransferAction) IsRedeemAtReturnsOnCall(i int, result1 bool) {
	fake.isRedeemAtMutex.Lock()
	defer fake.isRedeemAtMutex.Unlock()
	fake.IsRedeemAtStub = nil
	if fake.isRedeemAtReturnsOnCall == nil {
		fake.isRedeemAtReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isRedeemAtReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *TransferAction) NumOutputs() int {
	fake.numOutputsMutex.Lock()
	ret, specificReturn := fake.numOutputsReturnsOnCall[len(fake.numOutputsArgsForCall)]
	fake.numOutputsArgsForCall = append(fake.numOutputsArgsForCall, struct {
	}{})
	stub := fake.NumOutputsStub
	fakeReturns := fake.numOutputsReturns
	fake.recordInvocation("NumOutputs", []interface{}{})
	fake.numOutputsMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TransferAction) NumOutputsCallCount() int {
	fake.numOutputsMutex.RLock()
	defer fake.numOutputsMutex.RUnlock()
	return len(fake.numOutputsArgsForCall)
}

func (fake *TransferAction) NumOutputsCalls(stub func() int) {
	fake.numOutputsMutex.Lock()
	defer fake.numOutputsMutex.Unlock()
	fake.NumOutputsStub = stub
}

func (fake *TransferAction) NumOutputsReturns(result1 int) {
	fake.numOutputsMutex.Lock()
	defer fake.numOutputsMutex.Unlock()
	fake.NumOutputsStub = nil
	fake.numOutputsReturns = struct {
		result1 int
	}{result1}
}

func (fake *TransferAction) NumOutputsReturnsOnCall(i int, result1 int) {
	fake.numOutputsMutex.Lock()
	defer fake.numOutputsMutex.Unlock()
	fake.NumOutputsStub = nil
	if fake.numOutputsReturnsOnCall == nil {
		fake.numOutputsReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.numOutputsReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *TransferAction) Serialize() ([]byte, error) {
	fake.serializeMutex.Lock()
	ret, specificReturn := fake.serializeReturnsOnCall[len(fake.serializeArgsForCall)]
	fake.serializeArgsForCall = append(fake.serializeArgsForCall, struct {
	}{})
	stub := fake.SerializeStub
	fakeReturns := fake.serializeReturns
	fake.recordInvocation("Serialize", []interface{}{})
	fake.serializeMutex.Unlock()
	if stub != nil {
		return stub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TransferAction) SerializeCallCount() int {
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	return len(fake.serializeArgsForCall)
}

func (fake *TransferAction) SerializeCalls(stub func() ([]byte, error)) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = stub
}

func (fake *TransferAction) SerializeReturns(result1 []byte, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	fake.serializeReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *TransferAction) SerializeReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.serializeMutex.Lock()
	defer fake.serializeMutex.Unlock()
	fake.SerializeStub = nil
	if fake.serializeReturnsOnCall == nil {
		fake.serializeReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.serializeReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *TransferAction) SerializeOutputAt(arg1 int) ([]byte, error) {
	fake.serializeOutputAtMutex.Lock()
	ret, specificReturn := fake.serializeOutputAtReturnsOnCall[len(fake.serializeOutputAtArgsForCall)]
	fake.serializeOutputAtArgsForCall = append(fake.serializeOutputAtArgsForCall, struct {
		arg1 int
	}{arg1})
	stub := fake.SerializeOutputAtStub
	fakeReturns := fake.serializeOutputAtReturns
	fake.recordInvocation("SerializeOutputAt", []interface{}{arg1})
	fake.serializeOutputAtMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TransferAction) SerializeOutputAtCallCount() int {
	fake.serializeOutputAtMutex.RLock()
	defer fake.serializeOutputAtMutex.RUnlock()
	return len(fake.serializeOutputAtArgsForCall)
}

func (fake *TransferAction) SerializeOutputAtCalls(stub func(int) ([]byte, error)) {
	fake.serializeOutputAtMutex.Lock()
	defer fake.serializeOutputAtMutex.Unlock()
	fake.SerializeOutputAtStub = stub
}

func (fake *TransferAction) SerializeOutputAtArgsForCall(i int) int {
	fake.serializeOutputAtMutex.RLock()
	defer fake.serializeOutputAtMutex.RUnlock()
	argsForCall := fake.serializeOutputAtArgsForCall[i]
	return argsForCall.arg1
}

func (fake *TransferAction) SerializeOutputAtReturns(result1 []byte, result2 error) {
	fake.serializeOutputAtMutex.Lock()
	defer fake.serializeOutputAtMutex.Unlock()
	fake.SerializeOutputAtStub = nil
	fake.serializeOutputAtReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *TransferAction) SerializeOutputAtReturnsOnCall(i int, result1 []byte, result2 error) {
	fake.serializeOutputAtMutex.Lock()
	defer fake.serializeOutputAtMutex.Unlock()
	fake.SerializeOutputAtStub = nil
	if fake.serializeOutputAtReturnsOnCall == nil {
		fake.serializeOutputAtReturnsOnCall = make(map[int]struct {
			result1 []byte
			result2 error
		})
	}
	fake.serializeOutputAtReturnsOnCall[i] = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

func (fake *TransferAction) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.getInputsMutex.RLock()
	defer fake.getInputsMutex.RUnlock()
	fake.getMetadataMutex.RLock()
	defer fake.getMetadataMutex.RUnlock()
	fake.getSerialNumbersMutex.RLock()
	defer fake.getSerialNumbersMutex.RUnlock()
	fake.getSerializedInputsMutex.RLock()
	defer fake.getSerializedInputsMutex.RUnlock()
	fake.getSerializedOutputsMutex.RLock()
	defer fake.getSerializedOutputsMutex.RUnlock()
	fake.isGraphHidingMutex.RLock()
	defer fake.isGraphHidingMutex.RUnlock()
	fake.isRedeemAtMutex.RLock()
	defer fake.isRedeemAtMutex.RUnlock()
	fake.numOutputsMutex.RLock()
	defer fake.numOutputsMutex.RUnlock()
	fake.serializeMutex.RLock()
	defer fake.serializeMutex.RUnlock()
	fake.serializeOutputAtMutex.RLock()
	defer fake.serializeOutputAtMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TransferAction) recordInvocation(key string, args []interface{}) {
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

var _ translator.TransferAction = new(TransferAction)
