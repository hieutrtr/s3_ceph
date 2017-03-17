package main

import (
	"fmt"
	"log"
	"os"

	"github.com/minio/minio-go"
)

func main() {
	endpoint := "10.60.6.99:7480"
	accessKeyID := "TEKA586YD30JBUV1B5GD"
	secretAccessKey := "JXkcNb5iF4f6dt7k4MADMF5AooKkRDcnwrZ33a2P"
	useSSL := false

	// Initialize minio client object.
	minioClient, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
	if err != nil {
		log.Fatalln(err)
	}

	// Make a new bucked called mymusic.
	bucketName := "hello1"
	location := ""

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			log.Printf("We already own %s\n", bucketName)
		} else {
			log.Fatalln(err)
		}
	}
	log.Printf("Successfully created %s\n", bucketName)

	file, err := os.Open("/Users/hieutrtr/Downloads/sidaday.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	n, err := minioClient.PutObject(bucketName, "sida6", file, "application/octet-stream")
	fmt.Println(n)
	if err != nil {
		fmt.Println(err)
		return
	}

	object, err := minioClient.GetObject(bucketName, "sida6")
	if err != nil {
		fmt.Println(err)
		return
	}
	stat, _ := object.Stat()
	var bb = make([]byte, stat.Size)
	length, _ := object.Read(bb)
	fmt.Println(length)
}
