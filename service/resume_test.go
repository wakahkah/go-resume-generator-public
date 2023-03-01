package service_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wakahkah/go-resume-generator/service"
)

func TestGetResumePDFSuccess(t *testing.T) {
	templatePath := "../templates/resume.html"
	outputPath := "../output-test/resume-test.pdf"
	sourcePath := "../source/test-data.json"

	ok, err := service.GetResumePDF(sourcePath, outputPath, templatePath)
	require.NoError(t, err)
	require.Equal(t, ok, true)

	// check file exist
	_, error := os.Stat(outputPath)
	require.NoError(t, error)
}

func TestGetResumePDFFail(t *testing.T) {
	templatePath := "../templates/resume.html"
	outputPath := "../output-test/resume-test.pdf"
	sourcePath := ""

	ok, err := service.GetResumePDF(sourcePath, outputPath, templatePath)
	require.ErrorContains(t, err, "no such file or directory")
	require.Equal(t, ok, false)
}

func TestGetResumePDFParseFail(t *testing.T) {
	templatePath := ""
	outputPath := "../output-test/resume-test.pdf"
	sourcePath := "../source/test-data.json"

	ok, err := service.GetResumePDF(sourcePath, outputPath, templatePath)
	require.NotEmpty(t, err)
	require.Equal(t, ok, false)
}
