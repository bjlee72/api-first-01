package create

import (
	"context"
	"fmt"
	"testing"

	"github.com/go-openapi/swag"
	"github.com/rs/xid"
	"github.com/stretchr/testify/mock"

	"api-first-01/api/swagger/models"
	operations "api-first-01/api/swagger/restapi/operations/hello"
	"api-first-01/storage"
)

func TestHandlerBuilder_handle(t *testing.T) {
	type resultTestFunc func(result interface{}) bool
	type mockStorageFunc func(m *MockHelloStorage) HelloStorage

	testRequest := operations.V1CreateHelloParams{
		Body: &models.CreateHelloRequest{
			Hello: &models.Hello{
				Message: swag.String("test message"),
			},
		},
	}

	testBadRequest := operations.V1CreateHelloParams{
		Body: &models.CreateHelloRequest{
			Hello: &models.Hello{
				ID:      "test-id",
				Message: swag.String("test message"),
			},
		},
	}

	type fields struct {
		storage mockStorageFunc
	}
	type args struct {
		params operations.V1CreateHelloParams
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		test   resultTestFunc
	}{
		{
			name: "Bad request",
			fields: fields{
				storage: func(m *MockHelloStorage) HelloStorage { return m },
			},
			args: args{
				params: testBadRequest,
			},
			test: func(result interface{}) bool {
				_, ok := result.(*operations.V1CreateHelloBadRequest)
				return ok
			},
		},
		{
			name: "Storage fails for some unknown reason",
			fields: fields{
				storage: func(m *MockHelloStorage) HelloStorage {
					m.On("CreateHello", mock.Anything, mock.Anything).Return(nil, fmt.Errorf("error"))
					return m
				},
			},
			args: args{
				params: testRequest,
			},
			test: func(result interface{}) bool {
				_, ok := result.(*operations.V1CreateHelloInternalServerError)
				return ok
			},
		},
		{
			name: "Working normal",
			fields: fields{
				storage: func(m *MockHelloStorage) HelloStorage {
					m.On("CreateHello", mock.Anything, mock.Anything).
						Return(&storage.CreateHelloResponse{
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
				_, ok := result.(*operations.V1CreateHelloOK)
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
