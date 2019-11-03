package parser

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRaiseAnErrorOnEmptyLine(t *testing.T) {
	str := ""

	_, err := readAttributeLine(str)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), " 2 "))
}

func TestRaiseAnErrorOnLineWithoutName(t *testing.T) {
	str := ":2.0"

	_, err := readAttributeLine(str)
	assert.Error(t, err)
	assert.True(t, strings.Contains(err.Error(), "name"))
}

func TestReadSimpleLine(t *testing.T) {
	str := "VERSION:2.0"

	attribute, err := readAttributeLine(str)
	assert.NoError(t, err)
	assert.Equal(t, "VERSION", attribute.Name)
	assert.Equal(t, "2.0", attribute.Value)
	assert.Equal(t, 0, len(attribute.Info))
}

func TestReadLineWithoutValue(t *testing.T) {
	str := "VERSION:"

	attribute, err := readAttributeLine(str)
	assert.NoError(t, err)
	assert.Equal(t, "VERSION", attribute.Name)
	assert.Equal(t, 0, len(attribute.Value))
	assert.Equal(t, 0, len(attribute.Info))
}

func TestReadLineWithSingleParam(t *testing.T) {
	str := "VERSION;OPT=1:2.0"

	attribute, err := readAttributeLine(str)
	assert.NoError(t, err)
	assert.Equal(t, "VERSION", attribute.Name)
	assert.Equal(t, "2.0", attribute.Value)
	assert.Equal(t, 1, len(attribute.Info))

	assert.Equal(t, "1", attribute.Info["OPT"])
}

func TestReadLineWithTwoParams(t *testing.T) {
	str := "VERSION;OPT1=1;OPT2=2:2.0"

	attribute, err := readAttributeLine(str)
	assert.NoError(t, err)
	assert.Equal(t, "VERSION", attribute.Name)
	assert.Equal(t, "2.0", attribute.Value)
	assert.Equal(t, 2, len(attribute.Info))

	assert.Equal(t, "1", attribute.Info["OPT1"])
	assert.Equal(t, "2", attribute.Info["OPT2"])
}

func TestReadLineWithQuotedParams(t *testing.T) {
	str := "VERSION;OPT1=\";:,\";OPT2=2:2.0"

	attribute, err := readAttributeLine(str)
	assert.NoError(t, err)
	assert.Equal(t, "VERSION", attribute.Name)
	assert.Equal(t, "2.0", attribute.Value)
	assert.Equal(t, 2, len(attribute.Info))

	assert.Equal(t, ";:,", attribute.Info["OPT1"])
	assert.Equal(t, "2", attribute.Info["OPT2"])
}

func TestReadLineWithTwoNonValuatedParams(t *testing.T) {
	str := "VERSION;OPT1;OPT2:2.0"

	attribute, err := readAttributeLine(str)
	assert.NoError(t, err)
	assert.Equal(t, "VERSION", attribute.Name)
	assert.Equal(t, "2.0", attribute.Value)
	assert.Equal(t, 2, len(attribute.Info))

	assert.Equal(t, "", attribute.Info["OPT1"])
	assert.Equal(t, "", attribute.Info["OPT2"])
}
