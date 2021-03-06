package test

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	casper_v1 "rchain/api/pb"
	"rchain/internal/service"
	"testing"
)

func TestPullBlocks(t *testing.T) {
	service.PullBlocks()
}

func TestGetBlockInfo(t *testing.T) {

	const target1 = "observer-asia.services.mainnet.rchain.coop:40401"
	const target2 = "observer-eu.services.mainnet.rchain.coop:40401"
	const target3 = "observer-us.services.mainnet.rchain.coop:40401"

	conn, err := grpc.Dial(target1, grpc.WithInsecure())
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
