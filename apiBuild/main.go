package main

import "github.com/reyesGeorge/taubyte/apiBuild/getContent"

func main() {
	path := "QmP8jTG1m9GSDJLCbeWhVSVgEzCPPwXRdCRuJtQ5Tz9Kc9"
	getContent.HashGrabber(path)
	getContent.StoreNReadString()
}
