package compress

import (
	"bytes"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/disintegration/imaging"
	"image/jpeg"
	"io"
	"os"
)

func main() {
	f, err := os.Open("photo.jpg")
	//f, h, err := c.GetFile("file")
	if err != nil {
		beego.Error(err)
		//c.ReturnJson(10999, "SYSTEM_ERROR")
	}
	realBuffer := bytes.NewBuffer(nil)
	_, _ = io.Copy(realBuffer, f)
	result := bytes.NewBuffer(nil)
	for {
		if realBuffer.Len() < 1024*1024 {
			_, _ = io.Copy(result, realBuffer)
			break
		}
		dst := bytes.NewBuffer(nil)
		img, err := jpeg.Decode(realBuffer)
		if err != nil {
			//c.ReturnJson(10001, "系统繁忙")
			return
		}
		kuan := img.Bounds().Dx()
		gao := img.Bounds().Dy()
		m := imaging.Resize(img, int(float64(kuan)*0.8), int(float64(gao)*0.8), imaging.Lanczos)
		err = jpeg.Encode(dst, m, nil)
		if err != nil {
			//c.ReturnJson(10002, "系统错误")
			return
		}
		realBuffer = dst
	}

	fmt.Println(float64(result.Len()) / 1024)
	f2, err := os.Create("photo2.jpg")
	if err != nil {
		panic(err)
	}

	_, _ = io.Copy(f2, result)
	f2.Close()
}
