package main

import (
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/p2p"
	// "fmt"
	// "github.com/jackpal/gateway"
	// natpmp "github.com/jackpal/go-nat-pmp"
)

func main() {
	// gatewayIP, err := gateway.DiscoverGateway()
	// if err != nil {
	// 	fmt.Println("error1")
	// }

	// client := natpmp.NewClient(gatewayIP)
	// response, err := client.GetExternalAddress()
	// if err != nil {
	// 	fmt.Println("error2")
	// }
	// fmt.Println("External IP address:", response.ExternalIPAddress)
	nodekey, _ := crypto.GenerateKey()
	config := p2p.Config{
		MaxPeers:   10,
		PrivateKey: nodekey,
		Name:       "my node name",
		ListenAddr: ":30300",
		Protocols:  []p2p.Protocol{},
	}
	srv := &p2p.Server{
		Config: config,
	}
	srv.Start()
	select {}
}
