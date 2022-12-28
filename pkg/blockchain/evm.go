package blockchain

import (
	"github.com/doge-verse/zk-doge-backend/pkg/conf"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

const (
	MainnetEth     uint = 1
	MainnetPolygon uint = 137
	TestnetGoerli  uint = 5
	TestnetWallaby uint = 31415
	TestnetMumbai  uint = 80001
)

var (
	ClientList map[uint]*ethclient.Client
)

func Init() {
	// create client from the beginning
	clients := make(map[uint]*ethclient.Client)
	clientEth, err := ethclient.Dial(conf.Cfg.RPC.MainnetEth)
	if err != nil {
		log.Fatalln("clientEth Dial err:", err)
	}
	clients[MainnetEth] = clientEth
	clientPolygon, err := ethclient.Dial(conf.Cfg.RPC.MainnetPolygon)
	if err != nil {
		log.Fatalln("clientPolygon Dial err:", err)
	}
	clients[MainnetPolygon] = clientPolygon
	clientGoerli, err := ethclient.Dial(conf.Cfg.RPC.TestnetGoerli)
	if err != nil {
		log.Fatalln("clientGoerli Dial err:", err)
	}
	clients[TestnetGoerli] = clientGoerli
	clientWallaby, err := ethclient.Dial(conf.Cfg.RPC.TestnetWallaby)
	if err != nil {
		log.Fatalln("clientWallaby Dial err:", err)
	}
	clients[TestnetWallaby] = clientWallaby
	clientMumbai, err := ethclient.Dial(conf.Cfg.RPC.TestnetWallaby)
	if err != nil {
		log.Fatalln("clientWallaby Dial err:", err)
	}
	clients[TestnetMumbai] = clientMumbai
	ClientList = clients
}
