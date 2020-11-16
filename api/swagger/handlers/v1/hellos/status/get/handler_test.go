package get

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/mock"

	operations "api-first-01/api/swagger/restapi/operations/hello"
	"api-first-01/storage"
)

func TestHandlerBuilder_handle(t *testing.T) {
	type resultTestFunc func(result interface{}) bool
	type mockStorageFunc func(m *MockHelloStorage) HelloStorage

	testRequest := operations.V1ReadHelloStatusParams{
		HelloID: "id-of-the-hello-to-read",
	}

	type fields struct {
		storage mockStorageFunc
	}
	type args struct {
		params operations.V1ReadHelloStatusParams
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
					m.On("ReadHelloStatus", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("error"))
					return m
				},
			},
			args: args{
				params: testRequest,
			},
			test: func(result interface{}) bool {
				_, ok := result.(*operations.V1ReadHelloStatusInternalServerError)
				return ok
			},
		},
		{
			name: "Storage cannot find the hello",
			fields: fields{
				storage: func(m *MockHelloStorage) HelloStorage {
					m.On("ReadHelloStatus", mock.Anything, mock.Anything).Return(nil, storage.ErrRecordNotFound)
					return m
				},
			},
			args: args{
				params: testRequest,
			},
			test: func(result interface{}) bool {
				_, ok := result.(*operations.V1ReadHelloStatusNotFound)
				return ok
			},
		},
		{
			name: "Working normal",
			fields: fields{
				storage: func(m *MockHelloStorage) HelloStorage {
					m.On("ReadHelloStatus", mock.Anything, mock.Anything).
						Return(&storage.ReadHelloStatusResponse{
							Status: &storage.HelloStatus{
								ID:      "id-of-the-hello-to-read",
								Enabled: true,
							},
						}, nil)
					return m
				},
			},
			args: args{
				params: testRequest,
			},
			test: func(result interface{}) bool {
				_, ok := result.(*operations.V1ReadHelloStatusOK)
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
