package network

import "errors"

// ErrPeerAndPeerGroups is returned when a network router sets both a peer and peer groups.
var ErrPeerAndPeerGroups = errors.New("a network router supports either a peer or peer groups, not both")
