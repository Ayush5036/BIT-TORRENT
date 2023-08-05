package main

import (
	// "fmt"
	"sync"
)

// constants
var piecelength = int64(0)
var PEER_ID = make([]byte, 20)
var pieceDone = make(map[int]bool)
var listOfPeers = make(map[string]bool)
var mutex = &sync.Mutex{}
var pieces []*Piece
var path string
var info bencodeTorrent

