// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/bosh-agent/platform/cert"
)

type FakeManager struct {
	UpdateCertificatesStub        func(certs string) error
	updateCertificatesMutex       sync.RWMutex
	updateCertificatesArgsForCall []struct {
		certs string
	}
	updateCertificatesReturns struct {
		result1 error
	}
}

func (fake *FakeManager) UpdateCertificates(certs string) error {
	fake.updateCertificatesMutex.Lock()
	fake.updateCertificatesArgsForCall = append(fake.updateCertificatesArgsForCall, struct {
		certs string
	}{certs})
	fake.updateCertificatesMutex.Unlock()
	if fake.UpdateCertificatesStub != nil {
		return fake.UpdateCertificatesStub(certs)
	} else {
		return fake.updateCertificatesReturns.result1
	}
}

func (fake *FakeManager) UpdateCertificatesCallCount() int {
	fake.updateCertificatesMutex.RLock()
	defer fake.updateCertificatesMutex.RUnlock()
	return len(fake.updateCertificatesArgsForCall)
}

func (fake *FakeManager) UpdateCertificatesArgsForCall(i int) string {
	fake.updateCertificatesMutex.RLock()
	defer fake.updateCertificatesMutex.RUnlock()
	return fake.updateCertificatesArgsForCall[i].certs
}

func (fake *FakeManager) UpdateCertificatesReturns(result1 error) {
	fake.UpdateCertificatesStub = nil
	fake.updateCertificatesReturns = struct {
		result1 error
	}{result1}
}

var _ cert.Manager = new(FakeManager)
