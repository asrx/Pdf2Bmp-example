package main

import (
	"Pdf2Bmp/conf"
	"Pdf2Bmp/pdf"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	targetSuffix := ".pdf"
	conf := conf.ConfigXML()
	pdfPath := conf.PdfPath
	pdfPathOk := pdfPath + "ok/"
	bmpPath := conf.BmpPath


	// 获取所有pdf文件
	files,err := ioutil.ReadDir(pdfPath)
	if err != nil {
		fmt.Println("路径错误!")
		return
	}

	for _, file := range files{
		if !file.IsDir() && strings.HasSuffix(file.Name(),targetSuffix){
			name := strings.TrimSuffix(file.Name(),targetSuffix)
			fName := pdfPath + file.Name()

			err = pdf.Handle(fName,bmpPath + name,0,0)
			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			// move to ok
			err = os.Rename(fName,pdfPathOk + file.Name())
			if err != nil {
				fmt.Println(fName + "移除失败！")
			}else{
				fmt.Println("Conversion to complete: " + fName)
			}

		}
	}

	if !conf.Enable {
		fmt.Println("\n\nFilish！ \t回车退出程序...")
		var input string
		fmt.Scanln(&input)
	}
}

