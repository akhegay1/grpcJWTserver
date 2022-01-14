package main

import (
	"bufio"
	"log"
	"os"
	"strings"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"grpcJWTserver/pkg/jwtserver"
)

func main() {

	////////////////FILE/////
	conf, err := os.Open("../confJWT")
	if err != nil {
		log.Println("failed opening file conf: %s", err)
	}
	defer conf.Close()

	sc := bufio.NewScanner(conf)

	var params []string

	for sc.Scan() {
		str := sc.Text() // GET the line string
		val := str[strings.Index(str, "=")+1:]
		params = append(params, val)
	}

	if err := sc.Err(); err != nil {
		log.Println("scan file error: %v", err)
	}
	log.Println("params:", params)
	jwthost := params[1]

	//////////////////
	var conn *grpc.ClientConn
	conn, err = grpc.Dial(jwthost+":9011", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := jwtserver.NewJwtServerServiceClient(conn)

	//response, err := c.GetToken(context.Background(), &jwtserver.Reqtoken{User: "alek"})
	response, err := c.CheckToken(context.Background(), &jwtserver.CheckAuth{TokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NzU5OTM4NDYsImlhdCI6MTYzOTk5Mzg0NiwidXNlciI6ImFkbWluX3VzZXIifQ.jSG5cKyIVd3sqS2Qri64qCpemTNbuxlGPRMXkw8hVJM"})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}

	//log.Printf("Response from server: %s", response.TokenString)
	log.Printf("Response from server: %t %s", response.Tokenvalid, response.User)
}
