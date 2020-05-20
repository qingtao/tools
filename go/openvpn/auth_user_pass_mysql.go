//Package main used for openvpn auth-user-pass-verify verify the client's username and password
//set server.ovpn
//  auth-user-pass-verfiy auth_user_pass_mysql.exe via-file
//set client.ovpn
//  auth-user-pass
package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

//before run: go build -v -ldflags="-w -s", set these parameter below
const (
	dbNet   = "tcp"
	dbAddr  = "127.0.0.1:3306"
	dbName  = "openvpn"
	dbTable = "user"
	dbUser  = "openvpn"
	dbPass  = "openvpn"
)

//use mysql verify username and password
func main() {
	//only need two parameter, one is program self
	args := os.Args
	if len(args) != 2 {
		os.Exit(1)
	}
	//read content of file, then return []byte; exit status is 1, if failed
	b, err := ioutil.ReadFile(args[1])
	if err != nil {
		os.Exit(1)
	}
	//split string(b), tested with openvpn, used static username and password
	user := strings.Fields(string(b))
	//if len(user) is not 2, exit status is 1
	if len(user) != 2 {
		os.Exit(1)
	}

	if err := verifyWithMysql(user[0], user[1]); err != nil {
		os.Exit(1)
	}
}

//select pass from mysql, and check the user is active or not
func verifyWithMysql(username, password string) error {
	//not allow empty username or password
	if username == "" || password == "" {
		return errors.New("username or password is empty")
	}
	connectString := fmt.Sprintf("%s:%s@%s(%s)/%s", dbUser, dbPass, dbNet, dbAddr, dbName)
	db, err := sql.Open("mysql", connectString)
	if err != nil {
		return errors.Wrap(err, "mysql Open")
	}
	defer db.Close()
	if err := db.Ping(); err != nil {
		return errors.Wrap(err, "mysql Ping")
	}
	var (
		pass   string
		active int
	)
	queryString := fmt.Sprintf("SELECT password, active FROM %s WHERE username = ?", dbTable)
	if err := db.QueryRow(queryString, username).Scan(&pass, &active); err != nil {
		return errors.Wrap(err, "mysql QueryRow")
	}
	//check username and password, if use mysql, ldap etc. can check other parameters is valid or not
	if password != pass || active != 1 {
		return errors.New("password invalid or user is not active")
	}
	return nil
}
