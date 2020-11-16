package get

import (
	"context"
	"fmt"
	"testing"

	"github.com/rs/xid"
	"github.com/stretchr/testify/mock"

	operations "api-first-01/api/swagger/restapi/operations/hello"
	"api-first-01/storage"
)

func TestHandlerBuilder_handle(t *testing.T) {
	type resultTestFunc func(result interface{}) bool
	type mockStorageFunc func(m *MockHelloStorage) HelloStorage

	testRequest := operations.V1ReadHelloParams{
		HelloID: "id-of-the-hello-to-read",
	}

	type fields struct {
		storage mockStorageFunc
	}
	type args struct {
		params operations.V1ReadHelloParams
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		test   resultTestFunc
	}{
		{
			name: "Storage failed for some reason",
			fields: fields{
				storage: func(m *MockHelloStorage) HelloStorage {
					m.On("ReadHello", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("error"))
					return m
				},
			},
			args: args{
				params: testRequest,
			},
			test: func(result interface{}) bool {
				_, ok := result.(*operations.V1ReadHelloInternalServerError)
				return ok
			},
		},
		{
			name: "Storage cannot find the hello",
			fields: fields{
				storage: func(m *MockHelloStorage) HelloStorage {
					m.On("ReadHello", mock.Anything, mock.Anything).Return(nil, storage.ErrRecordNotFound)
					return m
				},
			},
			args: args{
				params: testRequest,
			},
			test: func(result interface{}) bool {
				_, ok := result.(*operations.V1ReadHelloNotFound)
				return ok
			},
		},
		{
			name: "Working normal",
			fields: fields{
				storage: func(m *MockHelloStorage) HelloStorage {
					m.On("ReadHello", mock.Anything, mock.Anything).
						Return(&storage.ReadHelloResponse{
							Hello: &storage.Hello{
								ID:      xid.New().String(),
								Message: "create hello response message from test",
							},
						}, nil)
					return m
				},
			},
			args: args{
				params: testRequest,
			},
			test: func(result interface{}) bool {
				_, ok := result.(*operations.V1ReadHelloOK)
				return ok
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockStorage := &MockHelloStorage{}

			got := Dependency{
				Storage: tt.fields.storage(mockStorage),
			}.handle(context.Background(), tt.args.params)

			if !tt.test(got) {
				t.Errorf("handle() failed: got = %#v", got)
			}
		})
	}
}
