package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

var rdb *redis.ClusterClient

// 初始化连接 10.15.1.48:7003  pass.123 ok;10.15.1.30:6379 kfzisgreatman100 is down
func initClient()(err error){
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"10.15.1.30:6379"},
		Password: "kfzisgreatman100", // no password set
	})

	fmt.Println(rdb.ClusterNodes())

	_, err = rdb.Ping().Result()
	fmt.Println(rdb.Ping().Result())
	if err != nil {
		return err
	}


	err = rdb.Set("score", 100, 0).Err()
	if err != nil {
		fmt.Printf("set score failed, err:%v\n", err)
		return
	}

	val, err := rdb.Get("score").Result()
	if err != nil {
		fmt.Printf("get score failed, err:%v\n", err)
		return
	}
	fmt.Println("score", val)

	return nil
}

func main() {

	err:=initClient()
	if err != nil {
		fmt.Println(err)
	}



}