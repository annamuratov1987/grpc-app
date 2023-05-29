package main

import (
	"context"
	"google.golang.org/grpc"
	"grpc-app/proto/notification"
	"log"
)

func main() {
	//var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := notification.NewNotificationServiceClient(conn)

	response, err := c.Notify(context.Background(), &notification.NotificationRequest{Message: "Я на пути к становлению ниндзей!"})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("STATUS:", response.Status)
}