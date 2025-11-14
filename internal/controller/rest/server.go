package rest

import (
	"context"
)

// Server is interface to control rest server
//
//	@title		Booking Service API
//	@version	dev
type Server interface {
	// Run starts server with context
	Run(ctx context.Context) error
}
