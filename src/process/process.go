package process

import (
	"fmt"

	"github.com/huelet/encode/src/utils"
	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Encode(videoPath string, fileName string) (location string) {
	err := ffmpeg.Input(videoPath).
		Output(fmt.Sprintf(`./content/processed/%s`, fileName), ffmpeg.KwArgs{"c:v": "libx264"}).
		OverWriteOutput().ErrorToStdOut().Run()
	utils.HandleError(err)
	return fmt.Sprintf("./content/processed/%s", fileName)
}
func EncodeGIF(videoPath string, fileName string) (location string) {
	err := ffmpeg.Input(videoPath, ffmpeg.KwArgs{"ss": "1"}).
		Output(fmt.Sprintf("./content/processed/.gif", fileName), ffmpeg.KwArgs{"s": "320x240", "pix_fmt": "rgb24", "t": "3", "r": "3"}).
		OverWriteOutput().ErrorToStdOut().Run()
	utils.HandleError(err)
	return fmt.Sprintf("./content/processed/.gif", fileName)
}
