package main

import (
	"io"
	"time"
)

func DialSock(file, sock, queryString string) (got bool, body []byte) {
	// Create a new FastCGI client
	fcgi, err := DialTimeout("unix", sock, 5*time.Second)
	if err != nil {
		// log error
		return false, nil
	}

	// Create a new FastCGI request
	env := make(map[string]string)
	env["SCRIPT_FILENAME"] = file // This should be the absolute path to your PHP file
	env["REQUEST_METHOD"] = "GET"
	env["QUERY_STRING"] = queryString

	resp, err := fcgi.Request(env, nil)
	if err != nil {
		// log error
		return false, nil
	}
	defer resp.Body.Close()

	body, err = io.ReadAll(resp.Body)
	if err != nil {
		// log error
		return false, nil
	}

	return true, body
}
