package main

import (
	"crypto/rand"
	"fmt"
	"os"
	"time"

	gotorrentparser "github.com/j-muller/go-torrent-parser"
	bencode "github.com/jackpal/bencode-go"
)

func main() {
	// Generating a random peer id 
	rand.Read(PEER_ID)


	arg := os.Args[1:]
	file, err := os.Open(arg[0])
	if err != nil {
		panic(err)
	}
	defer file.Close()

	info := bencodeTorrent{}
	err = bencode.Unmarshal(file, &info)
	if err != nil {
		panic(err)
	}

	fmt.Println(arg)
	fmt.Println(file)
	fmt.Println(info.Announce)

	for i := range info.Info.Files {
		info.Info.Length += info.Info.Files[i].Length
	}

	lastLen := info.Info.Length % info.Info.PieceLength
	if lastLen == 0 {
		lastLen = info.Info.PieceLength
	}

	piecesString := info.Info.Pieces
	pieces = make([]*Piece, len(piecesString)/20)

	for i := 0; i < len(piecesString); i += 20 {
		var temp Piece
		temp.index = i / 20
		for j := 0; j < 20; j++ {
			temp.hash[j] = piecesString[i+j]
		}
		if i+20 == len(piecesString) {
			temp.length = lastLen
		} else {
			temp.length = info.Info.PieceLength
		}
		temp.data = nil
		pieces[i/20] = &temp
	}

	torrent, err := gotorrentparser.ParseFromFile(arg[0])
	if err != nil {
		panic(err)
	}

	fmt.Println(torrent)
	fmt.Println(arg[0])


	workQueue := make(chan *Piece, len(pieces))
	for i := range pieces {
		workQueue <- pieces[i]
	}

	

	// fmt.Println("Length :", info.Info.PieceLength, "\nFiles:", info.Info.Files)
}
