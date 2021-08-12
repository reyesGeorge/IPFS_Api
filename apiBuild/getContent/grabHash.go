package getContent

import (
	"fmt"
	"os"
	"strings"

	shell "github.com/ipfs/go-ipfs-api"
)

func HashGrabber(path string) {

	sh := shell.NewShell("localhost:5001")
	cid, err := sh.BlockGet(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}

	fmt.Printf("added %s", cid)
}

func StoreNReadString() {

	sh := shell.NewShell("localhost:5001")
	cid, err := sh.Add(strings.NewReader("Hello it Works!"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %s", err)
		os.Exit(1)
	}
	HashGrabber(cid)
}
