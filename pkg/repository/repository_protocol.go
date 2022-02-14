package repository

import "fmt"

// Protocol is an enum that represents all
// allowed protocols in the Remote Import Paths rule
type Protocol string

const (
	// ProtocolBzr is the bzr protocol
	ProtocolBzr Protocol = "bzr"
	// ProtocolFossil is the fossil protocol
	ProtocolFossil Protocol = "fossil"
	// ProtocolGit is the git protocol
	ProtocolGit Protocol = "git"
	// ProtocolHg is the hg protocol
	ProtocolHg Protocol = "hg"
	// ProtocolSvn is the svn protocol
	ProtocolSvn Protocol = "svn"
)

// Transforms the protocol to a string
func (rp Protocol) String() string {
	if rp.Validate() != nil {
		return ""
	}

	return string(rp)
}

// Validate the protocol given is valid or not
func (rp Protocol) Validate() error {
	switch rp {
	case ProtocolBzr, ProtocolFossil, ProtocolGit, ProtocolHg, ProtocolSvn:
		return nil
	default:
		return fmt.Errorf("invalid protocol")
	}
}
