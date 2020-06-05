package conf

import (
	"encoding/xml"
	"os"
)

type configuration struct {
	PdfPath string	`xml:"pdf"`
	BmpPath string	`xml:"bmp"`
	Enable bool `xml:"enable"`
}

func ConfigXML() configuration {
	xmlFile, err := os.Open("conf.xml")
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()

	var conf configuration
	if err = xml.NewDecoder(xmlFile).Decode(&conf); err != nil {
		panic(err)
	}
	return conf
}

func GetPdfPath() string{
	return ConfigXML().PdfPath
}

func GetSavePath() string{
	return ConfigXML().BmpPath
}