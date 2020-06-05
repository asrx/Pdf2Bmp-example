package pdf

import (
	"errors"
	"fmt"
	"github.com/gen2brain/go-fitz"
	"golang.org/x/image/bmp"
	_ "image/jpeg"
	"os"
	"path/filepath"
	//"github.com/golang/image"
)

/**
pathName: pdf文件
sPathName: 保存后的文件（不含后缀）
startPage: 起始页
endPage: 结束页
*/
func Handle(pathName,sPathName string,startPage, endPage int) error {
	doc,err := fitz.New(pathName)
	if err != nil {
		return errors.New(pathName + ": 打开错误")
	}
	defer doc.Close()


	for n := 0; n <= startPage;n++  {
		if n == (endPage + 1) {
			break
		}

		img, err:= doc.Image(n)

		if err != nil {
			return errors.New(pathName + ": 初始化图片文件失败")
		}
		var f *os.File;
		if endPage == 0{
			f, err = os.Create(filepath.Join("",sPathName + ".bmp"))
		}else {
			f, err = os.Create(filepath.Join("",fmt.Sprintf(sPathName + "%03d.bmp",n)))
		}

		if err != nil {
			return errors.New(pathName + ": 创建-BMP文件失败")
		}
		defer f.Close()


		err = bmp.Encode(f,img)
		//err = jpeg.Encode(f, img, &jpeg.Options{jpeg.DefaultQuality})
		if err != nil {
			return errors.New(pathName + ": 保存-转BMP文件失败")
		}
	}

	return nil
}
