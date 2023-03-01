package main

import (
	"fmt"

	"github.com/wakahkah/go-resume-generator/service"
)

func main() {
	templatePath := "templates/resume.html"

	outputPath := "output/resume.pdf"

	sourcePath := "./source/resume-data.json"

	ok, err := service.GetResumePDF(sourcePath, outputPath, templatePath)

	if err == nil {
		fmt.Println(ok, "pdf generated successfully")
	} else {
		fmt.Println(err)
	}
}
