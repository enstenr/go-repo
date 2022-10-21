package main

import (
	"context"
	"fmt"

	"cloud.google.com/go/storage"
)

func main() {
	fmt.Println("hallo")
	checkGCS("mercari-custom-item-us-item-metadata-tree-detail-dev",
		"802ea447267a46f9ba5ff86e7d39ef23_pokemon_trading_card_games.json")
}
func checkGCS(bucketName, fileName string) error {
	// Create GCS connection
	ctx := context.Background()
	client, _ := storage.NewClient(ctx)
	bucket := client.Bucket(bucketName) // Get the url response
	fmt.Println(bucket)

	return nil
}
