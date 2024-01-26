package utils

import (
	"bufio"
	"encoding/base64"
	"os"
	"strconv"
	"strings"
	"time"
)

func WriteImages(data string) (string, error) {
	DataArr := strings.Split(data, ",")
	imgBase64 := DataArr[1]
	imgs, err := base64.StdEncoding.DecodeString(imgBase64)
	if err != nil {
		return "", err
	}
	timenow := time.Now().Unix()
	imgname := strconv.FormatInt(timenow, 10) + ".jpg"
	file, err := os.OpenFile("./images/"+imgname, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return "", err
	}
	w := bufio.NewWriter(file)
	_, err = w.WriteString(string(imgs))
	if err != nil {
		return "", err
	}
	w.Flush()
	defer file.Close()
	return imgname, nil
}
