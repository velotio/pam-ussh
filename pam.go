// +build darwin linux

/*
Copyright (c) 2017 Uber Technologies, Inc.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package main

// code in here can't be tested because it relies on cgo. :(

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"unsafe"

	"github.com/google/logger"
)

/*
#cgo LDFLAGS: -lpam -fPIC
#include <security/pam_appl.h>
#include <stdlib.h>
#include <string.h>

char *get_user_name(pam_handle_t *pamh);
*/
import "C"

const logPath = "/var/tmp/authz.log"

var verbose = flag.Bool("verbose", false, "print info level logs to stdout")

func authorize(username string) bool {
	jsonFile, err := os.Open("/etc/authz/users.json")
	if err != nil {
		panic(err)
	}

	jsonStr, _ := ioutil.ReadAll(jsonFile)
	jsonMap := make(map[string]interface{})

	err = json.Unmarshal([]byte(jsonStr), &jsonMap)
	if err != nil {
		panic(err)
	}

	if jsonMap[username] == "admin" {
		fmt.Printf("User %s is in `admin` group\n", username)
		return true
	}

	fmt.Printf("User %s is not in `admin` group\n", username)
	return false
}

//export pam_sm_acct_mgmt
func pam_sm_acct_mgmt(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	lf, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		fmt.Printf("Failed to open log file: %v", err)
	}
	defer lf.Close()

	defer logger.Init("LoggerExample", *verbose, true, lf).Close()

	cUsername := C.get_user_name(pamh)
	username := C.GoString(cUsername)
	defer logger.Info("Username is " + username)
	defer C.free(unsafe.Pointer(cUsername))

	return C.PAM_SUCCESS
}

//export pam_sm_setcred
func pam_sm_setcred(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
	return C.PAM_IGNORE
}

func main() {
	username := "admin"
	authorize(username)
}
