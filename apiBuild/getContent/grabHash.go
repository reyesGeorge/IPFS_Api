package getContent

import (
	"fmt"

	httpapi "github.com/ipfs/go-ipfs-http-client"
)

func HashGrabber(x string) {

	fmt.Println(httpapi.NewLocalApi())

	// path := "QmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9"

	// httpapi.NewURLApiWithClient(path)

	// fmt.Print(httpapi.NewRequest(context.Background(), path, "get"))
}
