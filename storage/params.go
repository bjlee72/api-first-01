package storage

// Hello represents a hello.
type Hello struct {
	ID      string
	Message string
}

// HelloStatus represents the status of a hello.
type HelloStatus struct {
	ID      string
	Enabled bool
}

// CreateHelloRequest represents an input to create a hello.
type CreateHelloRequest struct {
	Hello *Hello
}

// CreateHelloResponse represents a response to a CreateHelloRequest.
type CreateHelloResponse struct {
	Hello *Hello
}

// ReadHelloRequest represents a request to read a hello.
type ReadHelloRequest struct {
	ID string
}

// ReadHelloResponse is a response to a ReadHelloRequest.
type ReadHelloResponse struct {
	Hello *Hello
}

// ListHellosRequest is a request to list hellos.
type ListHellosRequest struct {
	Count int32
}

// ListHellosResponse is a response to a ListHellosRequest.
type ListHellosResponse struct {
	Hellos []*Hello
}

// ReadHelloStatusRequest is a request to read the status of a hello.
type ReadHelloStatusRequest struct {
	ID string
}

// ReadHelloStatusResponse is a response to ReadHelloStatusRequest.
type ReadHelloStatusResponse struct {
	Status *HelloStatus
}
