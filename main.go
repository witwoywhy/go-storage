package main

import (
	"context"
	"fmt"
	"go-storage/storage"
	"os"
	"time"

	"github.com/witwoywhy/go-cores/vipers"
)

func init() {
	vipers.Init()
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func upload(key string) {
	store := storage.Init(key)

	f, err := os.Open("./object/cat.jpg")
	must(err)

	must(store.Upload(context.Background(), &storage.UploadRequest{
		FilePath:    "cats/cat.jpg",
		File:        f,
		ContentType: "image/jpg",
	}))
}

func getUrl(key string) {
	store := storage.Init(key)

	fmt.Println(store.GetUrl("cats/cat.jpg"))
}

func getSignedUrl(key string) {
	store := storage.Init(key)

	r, e := store.GetSingedUrl(context.Background(), "cats/cat.jpg", 1*time.Hour)
	must(e)

	fmt.Println(r)
}

func delete(key string) {
	store := storage.Init(key)

	must(store.Delete(context.Background(), "cats/cat.jpg"))
}

func main() {

}
