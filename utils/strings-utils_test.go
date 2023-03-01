package utils_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wakahkah/go-resume-generator/utils"
)

func TestArrayToStr(t *testing.T) {
	ary := []string{"a", "b", "c"}

	result := utils.ArrayToStr(ary)

	require.Equal(t, result, "a, b and c")
}

func TestReplaceLastSuccess(t *testing.T) {
	comma := ", "

	testString := "a, b, c"

	result := utils.ReplaceLast(testString, comma, " and ")

	require.Equal(t, result, "a, b and c")
}

func TestReplaceLastFail(t *testing.T) {
	comma := ", "

	testString := "abc"

	result := utils.ReplaceLast(testString, comma, " and ")

	require.Equal(t, result, testString)
}
