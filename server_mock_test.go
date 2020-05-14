// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package freya

import (
	"context"
	"sync"
)

var (
	lockServerMockListenAndServe sync.RWMutex
	lockServerMockShutdown       sync.RWMutex
)

// Ensure, that ServerMock does implement Server.
// If this is not the case, regenerate this file with moq.
var _ Server = &ServerMock{}

// ServerMock is a mock implementation of Server.
//
//     func TestSomethingThatUsesServer(t *testing.T) {
//
//         // make and configure a mocked Server
//         mockedServer := &ServerMock{
//             ListenAndServeFunc: func() error {
// 	               panic("mock out the ListenAndServe method")
//             },
//             ShutdownFunc: func(in1 context.Context) error {
// 	               panic("mock out the Shutdown method")
//             },
//         }
//
//         // use mockedServer in code that requires Server
//         // and then make assertions.
//
//     }
type ServerMock struct {
	// ListenAndServeFunc mocks the ListenAndServe method.
	ListenAndServeFunc func() error

	// ShutdownFunc mocks the Shutdown method.
	ShutdownFunc func(in1 context.Context) error

	// calls tracks calls to the methods.
	calls struct {
		// ListenAndServe holds details about calls to the ListenAndServe method.
		ListenAndServe []struct {
		}
		// Shutdown holds details about calls to the Shutdown method.
		Shutdown []struct {
			// In1 is the in1 argument value.
			In1 context.Context
		}
	}
}

// ListenAndServe calls ListenAndServeFunc.
func (mock *ServerMock) ListenAndServe() error {
	if mock.ListenAndServeFunc == nil {
		panic("ServerMock.ListenAndServeFunc: method is nil but Server.ListenAndServe was just called")
	}
	callInfo := struct {
	}{}
	lockServerMockListenAndServe.Lock()
	mock.calls.ListenAndServe = append(mock.calls.ListenAndServe, callInfo)
	lockServerMockListenAndServe.Unlock()
	return mock.ListenAndServeFunc()
}

// ListenAndServeCalls gets all the calls that were made to ListenAndServe.
// Check the length with:
//     len(mockedServer.ListenAndServeCalls())
func (mock *ServerMock) ListenAndServeCalls() []struct {
} {
	var calls []struct {
	}
	lockServerMockListenAndServe.RLock()
	calls = mock.calls.ListenAndServe
	lockServerMockListenAndServe.RUnlock()
	return calls
}

// Shutdown calls ShutdownFunc.
func (mock *ServerMock) Shutdown(in1 context.Context) error {
	if mock.ShutdownFunc == nil {
		panic("ServerMock.ShutdownFunc: method is nil but Server.Shutdown was just called")
	}
	callInfo := struct {
		In1 context.Context
	}{
		In1: in1,
	}
	lockServerMockShutdown.Lock()
	mock.calls.Shutdown = append(mock.calls.Shutdown, callInfo)
	lockServerMockShutdown.Unlock()
	return mock.ShutdownFunc(in1)
}

// ShutdownCalls gets all the calls that were made to Shutdown.
// Check the length with:
//     len(mockedServer.ShutdownCalls())
func (mock *ServerMock) ShutdownCalls() []struct {
	In1 context.Context
} {
	var calls []struct {
		In1 context.Context
	}
	lockServerMockShutdown.RLock()
	calls = mock.calls.Shutdown
	lockServerMockShutdown.RUnlock()
	return calls
}
