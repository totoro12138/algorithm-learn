package main

import (
	"algorithm-learn/demo/upload/config"
	"crypto/md5"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	files, err := ioutil.ReadDir("file")
	if err != nil {
		panic(err)
	}

	for i, file := range files {
		if file.IsDir() || i == 0 {
			continue
		}
		client, err := oss.New(
			config.AliyunOssEndpoint,
			config.AliyunAccessKey,
			config.AliyunSecret)
		if err != nil {
			panic(err)
		}
		f, err := os.Open("file/" + file.Name())
		if err != nil {
			log.Printf("打开文件:%s失败  error:%s\n", file.Name(), err.Error())
			continue
		}

		md5obj := md5.New()
		if _, err := io.Copy(md5obj, f); err != nil {
			log.Printf("文件:%s生成hash失败  error:%s\n", file.Name(), err.Error())
			f.Close()
			continue
		}

		bucket, _ := client.Bucket(config.AliyunOssBucket)
		_, _ = f.Seek(0, 0)
		//md5str := string(md5obj.Sum(nil))

		//signedURL, err := bucket.SignURL(md5str, oss.HTTPPut, 60,
		//	oss.MaxUploads(10*1024*1024),
		//		oss.ContentMD5(fileHash),
		//oss.ContentType("image/png"))

		filename := time.Now().Format("20060102150405") + string(Krand(5, 0)) + ".png"
		err = bucket.PutObject(filename, f)
		if err != nil {
			log.Printf("file : %s  error : %s\n", file.Name(), err)
			f.Close()
			return
		}

		fmt.Printf("filename : %s  url : http://easyliveimg.oss-cn-shenzhen.aliyuncs.com/%s\n\n", f.Name(), filename)
		//fmt.Printf("filename : %s  url : %s\n\n", f.Name(), signedURL)
		f.Close()
	}

}

func Krand(size int, kind int) []byte {
	ikind, kinds, result := kind, [][]int{[]int{10, 48}, []int{26, 97}, []int{26, 65}}, make([]byte, size)

	is_all := kind > 2 || kind < 0
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		if is_all { // random ikind
			ikind = rand.Intn(3)
		}
		scope, base := kinds[ikind][0], kinds[ikind][1]
		result[i] = uint8(base + rand.Intn(scope))
	}
	return result
}
