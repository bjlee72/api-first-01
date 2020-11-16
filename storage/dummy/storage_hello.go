package dummy

import (
	"context"
	"fmt"

	"github.com/rs/xid"

	"api-first-01/storage"
)

// CreateHello creates a hello in the storage.
func (s *Storage) CreateHello(_ context.Context, req *storage.CreateHelloRequest) (*storage.CreateHelloResponse, error) {
	return &storage.CreateHelloResponse{
		Hello: &storage.Hello{
			ID:      req.Hello.ID,
			Message: req.Hello.Message,
		},
	}, nil
}

// ReadHello reads a hello from the storage.
func (s *Storage) ReadHello(_ context.Context, req *storage.ReadHelloRequest) (*storage.ReadHelloResponse, error) {
	return &storage.ReadHelloResponse{
		Hello: &storage.Hello{
			ID:      req.ID,
			Message: "hello message read from the storage",
		},
	}, nil
}

// ListHellos lists hellos in the storage.
func (s *Storage) ListHellos(_ context.Context, req *storage.ListHellosRequest) (*storage.ListHellosResponse, error) {
	hellos := make([]*storage.Hello, 0, req.Count)
	for i := int32(0); i < req.Count; i++ {
		hellos = append(hellos, &storage.Hello{
			ID:      xid.New().String(),
			Message: fmt.Sprintf("hello message #%d", i),
		})
	}

	return &storage.ListHellosResponse{Hellos: hellos}, nil
}

// ReadHelloStatus reads the status of a hello.
func (s *Storage) ReadHelloStatus(_ context.Context, req *storage.ReadHelloStatusRequest) (*storage.ReadHelloStatusResponse, error) {
	return &storage.ReadHelloStatusResponse{
		Status: &storage.HelloStatus{
			ID:      req.ID,
			Enabled: true,
		},
	}, nil
}
