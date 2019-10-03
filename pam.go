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

// /*
// #cgo LDFLAGS: -lpam -fPIC
// #include <security/pam_appl.h>
// #include <stdlib.h>
// #include <string.h>

// char *get_user_name(pam_handle_t *pamh);
// char *get_auth_tok(pam_handle_t *pamh);
// */
// import "C"

// var logger syslog.Writer

// func init() {
// 	logger, err := syslog.New(syslog.LOG_AUTH|syslog.LOG_WARNING, "pam-ussh")
// 	xt := reflect.TypeOf(logger).Kind()
// 	fmt.Println(xt)
// 	if err != nil {
// 		fmt.Printf("error opening file: %v", err)
// 		os.Exit(1)
// 	}
// }

// func authorize(username string) bool {
// 	logger.Debug(fmt.Sprintf("Authorizing user *%s*", username))

// 	jsonFile, err := os.Open("/etc/authz/users.json")
// 	if err != nil {
// 		panic(err)
// 	}

// 	jsonStr, _ := ioutil.ReadAll(jsonFile)
// 	jsonMap := make(map[string]interface{})

// 	err = json.Unmarshal([]byte(jsonStr), &jsonMap)
// 	if err != nil {
// 		panic(err)
// 	}

// 	if jsonMap[username] == "admin" || jsonMap[username] == "dev" {
// 		logger.Alert(fmt.Sprintf("User %s is present in groups %s\n", username, jsonMap[username]))
// 		return true
// 	}

// 	logger.Alert(fmt.Sprintf("User %s is not in groups %s\n", username, "admin"+", dev"))
// 	return false
// }

// //export pam_sm_authenticate
// func pam_sm_authenticate(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
// 	logger.Debug("Inside AuthZ PAM : Auth Management")

// 	// cUsername := C.get_user_name(pamh)
// 	// username := C.GoString(cUsername)

// 	cAuthTok := C.get_auth_tok(pamh)
// 	authTok := C.GoString(cAuthTok)

// 	logger.Debug(authTok)
// 	return C.PAM_SUCCESS

// 	// logger.Debug(fmt.Sprintf("User *%s* is trying to login", username))

// 	// if authorize(username) {
// 	// 	logger.Info(fmt.Sprintf("User *%s* is authorized", username))
// 	// 	return C.PAM_SUCCESS
// 	// }
// 	// logger.Info(fmt.Sprintf("User *%s* is un-authorized", username))
// 	// return C.PAM_AUTH_ERR
// }

// //export pam_sm_acct_mgmt
// func pam_sm_acct_mgmt(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
// 	logger.Debug("Inside AuthZ PAM : Account Management")

// 	// cUsername := C.get_user_name(pamh)
// 	// username := C.GoString(cUsername)

// 	cAuthTok := C.get_auth_tok(pamh)
// 	authTok := C.GoString(cAuthTok)

// 	logger.Debug(authTok)
// 	return C.PAM_SUCCESS

// 	// logger.Debug(fmt.Sprintf("User *%s* is trying to login", username))

// 	// if authorize(username) {
// 	// 	logger.Info(fmt.Sprintf("User *%s* is authorized", username))
// 	// 	return C.PAM_SUCCESS
// 	// }
// 	// logger.Info(fmt.Sprintf("User *%s* is un-authorized", username))
// 	// return C.PAM_AUTH_ERR
// }

// //export pam_sm_open_session
// func pam_sm_open_session(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
// 	logger.Debug("Inside AuthZ PAM : Open Session")
// 	return C.PAM_IGNORE
// }

// //export pam_sm_close_session
// func pam_sm_close_session(pamh *C.pam_handle_t, flags, argc C.int, argv **C.char) C.int {
// 	logger.Debug("Inside AuthZ PAM : Close Session")
// 	return C.PAM_IGNORE
// }

func main() {
	//username := "admin"
	logger.Err("Hello from PAM")
}
