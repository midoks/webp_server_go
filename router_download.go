package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"unsafe"

	"github.com/gofiber/fiber/v2"
)

func Md5Byte(buf []byte) string {
	hash := md5.New()
	hash.Write(buf)
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func Md5(s string) string {
	return Md5Byte([]byte(s))
}

// IsExist returns true if a file or directory exists.
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func Download(url, path string) error {
	if IsExist(path) {
		return nil
	}

	os.MkdirAll(filepath.Dir(path), os.ModePerm)

	current, err := http.Get(url)
	if err != nil {
		return err
	}
	imgbytes, err := ioutil.ReadAll(current.Body)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(path,
		imgbytes, 0644)
	if err != nil {
		return err
	}
	return nil
}

// Byte to string, only read-only
func BytesToString(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}

func ping(c *fiber.Ctx) error {
	c.SendString("ok")
	return nil
}

func img_download(c *fiber.Ctx) error {
	var err error
	reqURI := c.Params("URL")
	decoded, err := base64.StdEncoding.DecodeString(reqURI)
	if err != nil {
		return err
	}

	reqURI = BytesToString(decoded)
	fileSuffix := path.Ext(reqURI)
	a := Md5(reqURI)

	uri_path := "/" + a[0:1] + "/" + a[1:10] + "/" + a[11:20] + "/" + a[21:32]
	uri := uri_path + "/" + a + fileSuffix
	file_path := config.ImgPath + uri

	err = Download(reqURI, file_path)
	if err != nil {
		return err
	}

	c.Redirect(uri, 301)
	return nil
}
