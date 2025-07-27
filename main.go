package main

import (
	"encoding/base64"
	"encoding/json"
	"encoding/pem"
	"flag"
	"log"
	"net/http"
	"os"
	"path"
)

const (
	CTV1GetRootsEndpoint = "/ct/v1/get-roots"
)

func main() {
	logURL := flag.String("l", "tuscolo2026h2.sunlight.geomys.org", "CT Log URL")
	outName := flag.String("o", "roots.pem", "Output PEM file")
	flag.Parse()

	rsp, err := http.Get("https://" + path.Join(*logURL, CTV1GetRootsEndpoint))
	if err != nil {
		log.Fatalf("failed to download CT log roots: %v", err)
	}

	if rsp.StatusCode != http.StatusOK {
		log.Fatalf("failed to download CT log roots: Not 200 OK")
	}

	jsonDec := json.NewDecoder(rsp.Body)
	var ctRoots CTLogV1GetRootsResponse
	err = jsonDec.Decode(&ctRoots)
	if err != nil {
		log.Fatalf("failed to decode CT log response: %v", err)
	}

	outFile, err := os.Create(*outName)
	if err != nil {
		log.Fatalf("failed to open %q for writing output PEM: %v", *outName, err)
	}

	for i, root := range ctRoots.Certificates {
		der, err := base64.StdEncoding.DecodeString(root)
		if err != nil {
			log.Fatalf("failed to decode certificate in position index %d: %v", i, err)
		}

		err = pem.Encode(outFile, &pem.Block{Type: "CERTIFICATE", Bytes: der})
		if err != nil {
			log.Fatalf("failed to encode PEM certificate in position index %d: %v", i, err)
		}
	}
}

type CTLogV1GetRootsResponse struct {
	Certificates []string `json:"certificates"`
}
