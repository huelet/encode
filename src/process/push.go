package process

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/huelet/encode/src/utils"
)

type Video struct {
	VideoUrl string `json:"vurl"`
	Success  bool   `json:"success"`
}

func UploadToAzBlob(fileLocation string) (url any) {
	client := &http.Client{}
	var video []Video

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("file", fileLocation)
	utils.HandleError(err)
	file, err := os.Open(fileLocation)
	utils.HandleError(err)
	_, err = io.Copy(fw, file)
	utils.HandleError(err)
	writer.Close()
	req, err := http.NewRequest("POST", "https://api.huelet.net/videos/upload/item", bytes.NewReader(body.Bytes()))
	utils.HandleError(err)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	rsp, err := client.Do(req)
	utils.HandleError(err)
	bodyData, err := ioutil.ReadAll(rsp.Body)
	utils.HandleError(err)
	// this is the most inefficient stupid code ive ever written
	err = json.Unmarshal([]byte(fmt.Sprintf(`[%s]`, string(bodyData))), &video)
	utils.HandleError(err)
	fmt.Println(rsp)

	return video
}
