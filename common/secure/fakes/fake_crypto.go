// This file was generated by counterfeiter
package fakes

import (
	"sync"

	"github.com/cloudfoundry/gorouter/common/secure"
)

type FakeCrypto struct {
	EncryptStub        func(plainText []byte) (cipherText []byte, nonce []byte, err error)
	encryptMutex       sync.RWMutex
	encryptArgsForCall []struct {
		plainText []byte
	}
	encryptReturns struct {
		result1 []byte
		result2 []byte
		result3 error
	}
	DecryptStub        func(cipherText, nonce []byte) ([]byte, error)
	decryptMutex       sync.RWMutex
	decryptArgsForCall []struct {
		cipherText []byte
		nonce      []byte
	}
	decryptReturns struct {
		result1 []byte
		result2 error
	}
}

func (fake *FakeCrypto) Encrypt(plainText []byte) (cipherText []byte, nonce []byte, err error) {
	fake.encryptMutex.Lock()
	fake.encryptArgsForCall = append(fake.encryptArgsForCall, struct {
		plainText []byte
	}{plainText})
	fake.encryptMutex.Unlock()
	if fake.EncryptStub != nil {
		return fake.EncryptStub(plainText)
	} else {
		return fake.encryptReturns.result1, fake.encryptReturns.result2, fake.encryptReturns.result3
	}
}

func (fake *FakeCrypto) EncryptCallCount() int {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	return len(fake.encryptArgsForCall)
}

func (fake *FakeCrypto) EncryptArgsForCall(i int) []byte {
	fake.encryptMutex.RLock()
	defer fake.encryptMutex.RUnlock()
	return fake.encryptArgsForCall[i].plainText
}

func (fake *FakeCrypto) EncryptReturns(result1 []byte, result2 []byte, result3 error) {
	fake.EncryptStub = nil
	fake.encryptReturns = struct {
		result1 []byte
		result2 []byte
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeCrypto) Decrypt(cipherText []byte, nonce []byte) ([]byte, error) {
	fake.decryptMutex.Lock()
	fake.decryptArgsForCall = append(fake.decryptArgsForCall, struct {
		cipherText []byte
		nonce      []byte
	}{cipherText, nonce})
	fake.decryptMutex.Unlock()
	if fake.DecryptStub != nil {
		return fake.DecryptStub(cipherText, nonce)
	} else {
		return fake.decryptReturns.result1, fake.decryptReturns.result2
	}
}

func (fake *FakeCrypto) DecryptCallCount() int {
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	return len(fake.decryptArgsForCall)
}

func (fake *FakeCrypto) DecryptArgsForCall(i int) ([]byte, []byte) {
	fake.decryptMutex.RLock()
	defer fake.decryptMutex.RUnlock()
	return fake.decryptArgsForCall[i].cipherText, fake.decryptArgsForCall[i].nonce
}

func (fake *FakeCrypto) DecryptReturns(result1 []byte, result2 error) {
	fake.DecryptStub = nil
	fake.decryptReturns = struct {
		result1 []byte
		result2 error
	}{result1, result2}
}

var _ secure.Crypto = new(FakeCrypto)