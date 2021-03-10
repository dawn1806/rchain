package client

import (
	"context"
	casper_v1 "dawn1806/rchain/client/pb"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
)

func Run() {

	const target = "observer-asia.services.mainnet.rchain.coop:40401"

	conn, err := grpc.Dial(target, grpc.WithInsecure())
	if err != nil {
		log.Fatal("grpc Dial error ", err)
	}
	defer conn.Close()

	client := casper_v1.NewDeployServiceClient(conn)
	blocksClient, err := client.GetBlocks(context.Background(), &casper_v1.BlocksQuery{Depth: 20})
	if err != nil {
		log.Fatal("client getBlocks error ", err)
	}

	for {
		blocksInfo, err := blocksClient.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal("blocksClient.Recv error ", err)
			return
		}

		lightBlockInfo := blocksInfo.GetBlockInfo()
		fmt.Printf("blockHash=%s \t blockNumber=%d \n", lightBlockInfo.BlockHash, lightBlockInfo.BlockNumber)
	}
}






























