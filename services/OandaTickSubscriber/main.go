package main

import (
	"log"
	"math/rand"
	"net"
	"sync"
	"time"

	"github.com/micro/go-micro/client"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/metadata"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
	proto "github.com/nii236/nii-finance/services/TickRecorder/proto"
	"golang.org/x/net/context"
)

type pairslice []string

var pairs pairslice

func main() {
	cmd.Init()
	log.Println("Starting up Oanda Tick Subscriber...")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
	}
	log.Println("Interfaces:")
	for _, add := range addrs {
		log.Println(add.Network()+":", add.String())
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	wg := sync.WaitGroup{}
	wg.Add(1)
	t := time.NewTicker(10 * time.Second)
	for range t.C {
		log.Println("Publishing mock tick data...")
		ctx := metadata.NewContext(context.Background(), map[string]string{
			"X-User-Id": "john",
			"X-From-Id": "script",
		})
		tmpbid := 100.0 + rand.Float64()
		now := time.Now().UnixNano()
		msg := client.NewPublication("go.micro.srv.TickRecorder.Tick", &proto.Tick{
			Time:   now,
			Bid:    tmpbid,
			Ask:    tmpbid + r.Float64(),
			Last:   100.0 + r.Float64(),
			Pair:   "AUDUSD",
			Broker: "oanda",
		})
		if err := client.Publish(ctx, msg); err != nil {
			log.Println("publish err: ", err)
		}
	}
	wg.Wait()
}
