package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"time"

	"bittorrent/torrent"

	"github.com/gosuri/uilive"
)

var torrentFile = flag.String("torrentfile", "", "read the contents of the torrent `file`")
var magnet = flag.String("magnet", "", "read the contents of the magnet")
