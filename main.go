package main

import (
	"log"

	pb "github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	asset "github.com/mtyrrell/buf-example/pkg/asset/v1"
	v1 "github.com/mtyrrell/buf-example/pkg/semver/v1"
)

func main() {
	asset := &asset.Asset{
		Version: &v1.Version{
			Label: "Major",
			Major: 0,
			Minor: 1,
			Micro: 0,
		},
		Id:         uuid.New().String(),
		CreateTime: pb.TimestampNow(),
	}
	log.Printf("Asset: %v", asset)

}
