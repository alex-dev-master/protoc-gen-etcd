package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	/*	c, err := proto.NewUserEtcdClient("127.0.0.1:2379")
		fmt.Println(err)
		err = c.Put(context.Background(), &proto.User{
			Id:   1,
			Name: "",
		})
		fmt.Println(err)
		fmt.Println("Stop")*/
}
