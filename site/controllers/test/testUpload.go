package test

import (
	"./aws"
	"bufio"
	"fmt"
	"os"
)

func testupload() {
	f := aws.FileUpload{}
	// err := f.FileUploadFromPath("tumbudu.png", "testCampaign", "testbadge")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	file, err := os.Open("tumbudu.png")

	if err != nil {
		fmt.Println(err)
		// return err
	}

	defer file.Close()

	//get file size
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	bytes := make([]byte, size)

	//Prepare buffer to post
	buffer := bufio.NewReader(file)
	_, err = buffer.Read(bytes)

	// func (this *FileUpload) UploadToS3(data []byte, campaign string, badge string) (err error) {
	err1 := f.UploadToS3(bytes, "testCampaign", "badge1")
	if err1 != nil {
		fmt.Println(err)
	}
}
