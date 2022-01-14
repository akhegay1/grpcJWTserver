package main

import (
	"bufio"
	"grpcJWTserver/pkg/jwtserver"
	"log"
	"net"
	"os"
	"strings"

	"google.golang.org/grpc"
)

func main() {

	////////////////FILE/////
	conf, err := os.Open("confJWT")
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
	jwtserver.Appkey = params[0]
	host := params[1]

	//////////////////

	lis, err := net.Listen("tcp", host+":9011")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := jwtserver.JwtServer{}
	grpcServer := grpc.NewServer()
	jwtserver.RegisterJwtServerServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
