// Package route contains the REST API route handler helpers.
// These helpers use generics to provide a common interface for handling routes.
// The underlying implementation uses the Gin framework.
//
// The generic route handlers include `Handle`, `HandleIn`, `HandleOut`, and `HandleVoid`. These handlers are designed to
// handle different types of requests:
//
// - `Handle` is a generic handler that can handle both input and output types.
// - `HandleIn` is a generic handler that only handles input types.
// - `HandleOut` is a generic handler that only handles output types.
// - `HandleVoid` is a generic handler that does not handle any input or output types but may return an error.
//
// These handlers do not set an error status code or send any response in case of an error. This allows for middleware to
// catch these errors and respond accordingly.
package route
