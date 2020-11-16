package list

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

	testRequest := operations.V1ListHellosParams{
		Count: 1,
	}

	type fields struct {
		storage mockStorageFunc
	}
	type args struct {
		params operations.V1ListHellosParams
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
					m.On("ListHellos", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("error"))
					return m
				},
			},
			args: args{
				params: testRequest,
			},
			test: func(result interface{}) bool {
				_, ok := result.(*operations.V1ListHellosInternalServerError)
				return ok
			},
		},
		{
			name: "Working normal",
			fields: fields{
				storage: func(m *MockHelloStorage) HelloStorage {
					hellosToReturn := []*storage.Hello{
						{
							ID:      xid.New().String(),
							Message: "hello!",
						},
					}
					m.On("ListHellos", mock.Anything, mock.Anything).
						Return(&storage.ListHellosResponse{Hellos: hellosToReturn}, nil)
					return m
				},
			},
			args: args{
				params: testRequest,
			},
			test: func(result interface{}) bool {
				_, ok := result.(*operations.V1ListHellosOK)
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
