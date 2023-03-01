package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wakahkah/go-resume-generator/utils"
)

func TestGetDataFromFileSuccess(t *testing.T) {
	testFilePath := "../source/test-data.json"

	resumeData, err := utils.GetDataFromFile(testFilePath)

	require.NoError(t, err)
	require.NotEmpty(t, resumeData)

	require.NotEmpty(t, resumeData.PersonalInfo)
	require.NotEmpty(t, resumeData.PersonalInfo.ContactInfo)

	require.NotEmpty(t, resumeData.Skills)
	require.NotEmpty(t, resumeData.WorkExperience)
	require.NotEmpty(t, resumeData.Certificates)
	require.NotEmpty(t, resumeData.Awards)
	require.NotEmpty(t, resumeData.Education)
	require.NotEmpty(t, resumeData.PersonalProject)
	require.NotEmpty(t, resumeData.Interest)
}

func TestGetDataFromFileFail(t *testing.T) {
	testFilePath := ""

	resumeData, err := utils.GetDataFromFile(testFilePath)

	require.ErrorContains(t, err, "no such file or directory")
	require.Empty(t, resumeData)
}
