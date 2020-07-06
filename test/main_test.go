/*
 * @Version: 0.0.1
 * @Author: ider
 * @Date: 2020-07-05 21:27:15
 * @LastEditors: ider
 * @LastEditTime: 2020-07-06 12:20:27
 * @Description:
 */
package test

import (
	"io/ioutil"
	"testing"
	"time"

	"github.com/iwuzhen/filehub/seaweeds3"
)

func Test_Seaweed(t *testing.T) {
	// f, _ := awss3.NewAwsS3("awss3://dafqeQG2141RFq:gegio4gnberianowanewF@127.0.0.1:8333/bucket1/top")
	swf, _ := seaweeds3.NewSeaweedS3("seaweeds3://dafqeQG2141RFq:gegio4gnberianowanewF@127.0.0.1:8333/bucket1/top")
	dataImage, _ := ioutil.ReadFile("/home/ider/Pictures/girl_1.jpg")
	// swf.Put("cc.jpg", dataImage, "image/jpeg")
	swf.PutExpire("image/c2.jpg", dataImage, "image/jpeg", time.Second*5)

	fs, err := swf.List("/")
	t.Log(fs, err)

	// 	exs, _ := swf.Exists("image/c1.jpg")
	// 	t.Log(exs)
	// 	exs, _ = swf.Exists("image/c11.jpg")
	// 	t.Log(exs)
	// 	data, cont, _ := swf.Get("image/c1.jpg")
	// 	t.Log(cont)
	// 	ioutil.WriteFile("/tmp/cc.jpg", data, 0777)
	// 	err = swf.Del("c1.jpg")
}
