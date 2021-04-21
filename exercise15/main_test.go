package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TestDevMv() is testcase for api:"http://localhost:3001"
func TestDevMv(t *testing.T) {
	go main()

	 // A Client is an HTTP client. Its zero value (DefaultClient) is a usable client that uses DefaultTransport.
	 
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	res, _ := http.NewRequest("GET", "http://localhost:3001", nil)

	// Do sends an HTTP request and returns an HTTP response, following
	// policy (such as redirects, cookies, auth) as configured on the client.
	// An error is returned if caused by client policy (such as CheckRedirect), 
	//or failure to speak HTTP (such as a network connectivity problem).
	//Generally Get, Post, or PostForm will be used instead of Do.
	// Any returned error will be of type *url.Error. 
	//The url.Error value's Timeout method will report true if request timed out or was canceled.
	response, err := client.Do(res)
	if err != nil {
		panic(err)
	}

	// ReadAll reads from r until an error or EOF and returns the data it read.
	// A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF,
	// it does not treat an EOF from Read as an error to be reported.
	resbody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	 
	// Sprintf formats according to a format specifier and returns the resulting string.
	// Contains reports whether substr is within string.
	str := fmt.Sprintf("%s", resbody)
	flag := strings.Contains(str, "Hello Gopher!")
	assert.Equal(t,true, flag, "pass")

}

// TestPanicAfterDemo is also cover testcase when error encounterd during server connection and also cover testcases for makeLinks()
// TestPanicAfterDemo to test the panic condition
// TestPanicAfterDemo() is testcase for api:"http://localhost:3001/panic-after"
func TestPanicAfterDemo(t *testing.T) {

	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	res, _ := http.NewRequest("GET", "http://localhost:3001/panic-after", nil)
	response, err := client.Do(res)
	if err != nil {
		panic(err)
	}
	resbody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("%s", resbody)
	flag := strings.Contains(str, "Oh no!")
	assert.Equal(t, true, flag, "pass")
}

// TestSourceCodeHandler to test source code handler (positive testcases)
// TestSourceCodeHandlerNegative() is testcase for api:"http://localhost:3001/debug"
func TestSourceCodeHandler(t *testing.T) {

	client := &http.Client{
		Timeout: 1 * time.Second,
	}
    // url with demo file path to make link of it
	d := "http://localhost:3001/debug?path=/home/gs-4117/go/src/github.ibm.com/gophercises/exercise15/main.go&line=200"
	res, _ := http.NewRequest("GET", d, nil)
	response, err := client.Do(res)
	if err != nil {
		panic(err)
	}
	resbody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("%s", resbody)
	flag := len(str) > 10000
	assert.Equal(t, true, flag, "pass")
}

// TestSourceCodeHandlerNegative to test sourcecodehandler() when error encounter
//  TestSourceCodeHandlerNegative() is testcase for api:"http://localhost:3001/debug"
func TestSourceCodeHandlerNegative(t *testing.T) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	res, _ := http.NewRequest("GET", "http://localhost:3001/debug", nil)
	
	response, err := client.Do(res)
	if err != nil {
		panic(err)
	}
	resbody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	str := fmt.Sprintf("%s", resbody)
	flag := len(str) > 10000
	assert.Equal(t, false, flag, "pass")
}

// mockFileCopy is return error after copying a file

func mockFileCopy(buf *bytes.Buffer, file *os.File) (int64, error) {
	return 0, errors.New("Encounter error while copying file")

}

//TestSourceCodeHandlerFileCopyError() is testcase for api:"http://localhost:3001/debug"
// used when error encounter on copying a file.
func TestSourceCodeHandlerFileCopyError(t *testing.T) {
	client := &http.Client{
		Timeout: 1 * time.Second,
	}

	temp := fileCopyHandler

	defer func() {
		fileCopyHandler = temp
	}()

	fileCopyHandler = mockFileCopy        // mockFileCopy will throw a error

	// demo file path with url to makelink of it
	d := "http://localhost:3001/debug?path=/home/gs-4117/go/src/github.ibm.com/gophercises/exercise15/main.go&line=200"
	
	res, _ := http.NewRequest("GET", d, nil)
	response, _ := client.Do(res)

	flag := response.StatusCode == 500
	assert.Equal(t, true, flag, "pass")

}
