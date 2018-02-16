package leaf

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func openRemoteFile(name string) (io.ReadCloser, error) {
	resp, err := http.Get(name)
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(resp.Body); err != nil {
		return nil, err
	}
	if err := resp.Body.Close(); err != nil {
		return nil, err
	}
	return ioutil.NopCloser(buf), nil
}

func openLocalFile(name string) (io.ReadCloser, error) {
	f, err := os.Open(name)
	if err != nil {
		log.Printf("local file open : err [%s]", err)
		return nil, err
	}
	buf := new(bytes.Buffer)
	if _, err := buf.ReadFrom(f); err != nil {
		log.Printf("local file read : err [%s]", err)
		return nil, err
	}
	if err := f.Close(); err != nil {
		log.Printf("local file close : err [%s]", err)
		return nil, err
	}
	return ioutil.NopCloser(buf), nil
}
