package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"
)

type StubResponseFetcher struct{}

func (mrf StubResponseFetcher) Fetch(url string) (int, error) {
	fmt.Printf("I am in tests")
	return 0, nil
}

type MockResponseFetcher struct {
	mock.Mock
}

func (mrf *MockResponseFetcher) Fetch(url string) (int, error) {
	args := mrf.Called(url)
	return args.Int(0), nil
}

func TestMockTest(t *testing.T) {
	mockFetcher := &MockResponseFetcher{} //Arrange
	mockFetcher.On("Fetch", "http://example.com/info").Return(100)

	i, err := getResponse(mockFetcher) //Act
	if err != nil {
		t.Errorf("Did not expect error, Bro!")
	}

	if i != 100 {
		t.Errorf("Expected 100 bro, but got %d", i)
	}

	mockFetcher.AssertExpectations(t) //Assert
}
