package main

import (
	"crypto/tls"
	"log"
	"net"
	"os"

	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session

func getMongoSession() *mgo.Session {
	if mgoSession == nil {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		connString := os.Getenv("MONGODB_CONN_STRING")

		dialInfo, err := mgo.ParseURL(connString)

		if err != nil {
			log.Println(err)
		}
		dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
			tlsConfig := &tls.Config{}
			conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
			if err != nil {
				log.Println(err)
			}
			return conn, err
		}
		session, err := mgo.DialWithInfo(dialInfo)
		if err != nil {
			panic(err)
		}

		mgoSession = session
	}

	return mgoSession.Clone()
}
