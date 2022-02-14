package repository

import (
	"fmt"
	"strings"
)

// Repository is a struct that holds the configuration for the proxy
// You can see more about the Remote iomport paths here:
//
// @see https://pkg.go.dev/cmd/go#hdr-Remote_import_paths
type Repository struct {
	// Entrypoint is the entrypoint of the repository.
	// Represents the desired url of the repository.
	EntryPoint string
	// Protocol is the protocol of the repository.
	// Allowed values are: "bzr", "fossil", "git", "hg", "svn".
	Protocol Protocol
	// Destination of the repository.
	Destination string
}

// New creates a new repository and auto fill the protocol
// based on the destination. The Destination must be have the
// protocol included at the end of URL.
func New(entrypoint, destination string) (*Repository, error) {
	if destination == "" {
		return nil, fmt.Errorf("destination is required")
	}

	if entrypoint == "" {
		return nil, fmt.Errorf("entrypoint is required")
	}

	repository := &Repository{
		EntryPoint:  entrypoint,
		Destination: destination,
	}

	repository.Protocol = Protocol(destination[strings.LastIndex(destination, ".")+1:])
	if err := repository.Protocol.Validate(); err != nil {
		return nil, err
	}

	return repository, nil
}
