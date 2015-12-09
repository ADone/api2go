package api2go

import "net/http"

// The CRUD interface MUST be implemented in order to use the api2go api.
// Use Responder for success status codes and content/meta data. In case of an error,
// use the error return value preferrably with an instance of our HTTPError struct.
type CRUD interface {
	// FindOne returns an object by its ID
	// Possible Responder success status code 200
	FindOne(ID string, req Request) (Responder, error)

	// Create a new object. Newly created object/struct must be in Responder.
	// Possible Responder status codes are:
	// - 201 Created: Resource was created and needs to be returned
	// - 202 Accepted: Processing is delayed, return nothing
	// - 204 No Content: Resource created with a client generated ID, and no fields were modified by
	//   the server
	Create(obj interface{}, req Request) (Responder, error)

	// Delete an object
	// Possible Responder status codes are:
	// - 200 OK: Deletion was a success, returns meta information, currently not implemented! Do not use this
	// - 202 Accepted: Processing is delayed, return nothing
	// - 204 No Content: Deletion was successful, return nothing
	Delete(id string, req Request) (Responder, error)

	// Update an object
	// Possible Responder status codes are:
	// - 200 OK: Update successful, however some field(s) were changed, returns updates source
	// - 202 Accepted: Processing is delayed, return nothing
	// - 204 No Content: Update was successful, no fields were changed by the server, return nothing
	Update(obj interface{}, req Request) (Responder, error)
}

// The PaginatedFindAll interface can be optionally implemented to fetch a subset of all records.
// Pagination query parameters must be used to limit the result. Pagination URLs will automatically
// be generated by the api. You can use a combination of the following 2 query parameters:
// page[number] AND page[size]
// OR page[offset] AND page[limit]
type PaginatedFindAll interface {
	PaginatedFindAll(req Request) (totalCount uint, response Responder, err error)
}

// The FindAll interface can be optionally implemented to fetch all records at once.
type FindAll interface {
	// FindAll returns all objects
	FindAll(req Request) (Responder, error)
}

//URLResolver allows you to implement a static
//way to return a baseURL for all incoming
//requests for one api2go instance.
type URLResolver interface {
	GetBaseURL() string
}

// RequestAwareURLResolver allows you to dynamically change
// generated urls.
//
// This is particulary useful if you have the same
// API answering to multiple domains, or subdomains
// e.g customer[1,2,3,4].yourapi.example.com
//
// SetRequest will always be called prior to
// the GetBaseURL() from `URLResolver` so you
// have to change the result value based on the last
// request.
type RequestAwareURLResolver interface {
	URLResolver
	SetRequest(http.Request)
}

// The Responder interface is used by all Resource Methods as a container for the Response.
// Metadata is additional Metadata. You can put anything you like into it, see jsonapi spec.
// Result returns the actual payload. For FindOne, put only one entry in it.
// StatusCode sets the http status code.
type Responder interface {
	Metadata() map[string]interface{}
	Result() interface{}
	StatusCode() int
}
