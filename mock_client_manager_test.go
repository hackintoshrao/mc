package main

import "github.com/stretchr/testify/mock"

import "io"

type mockClientManager struct {
	mock.Mock
}

func (m *mockClientManager) getSourceReader(sourceURLParser *parsedURL) (io.ReadCloser, int64, string, error) {
	ret := m.Called(sourceURLParser)

	r0 := ret.Get(0).(io.ReadCloser)
	r1 := ret.Get(1).(int64)
	r2 := ret.Get(2).(string)
	r3 := ret.Error(3)

	return r0, r1, r2, r3
}
func (m *mockClientManager) getTargetWriter(targetURLParser *parsedURL, md5Hex string, length int64) (io.WriteCloser, error) {
	ret := m.Called(targetURLParser, md5Hex, length)

	r0 := ret.Get(0).(io.WriteCloser)
	r1 := ret.Error(1)

	return r0, r1
}
