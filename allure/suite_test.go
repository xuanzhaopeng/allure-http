package allure

import (
	"testing"
	"time"
	"encoding/xml"
	"fmt"
	"github.com/stretchr/testify/assert"
)

func TestNewSuite(t *testing.T) {
	startTime := time.Now()
	name := "suite1"
	suite := NewSuite(name, startTime)
	suiteData,err := xml.Marshal(suite)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<ns2:test-suite xmlns:ns2="urn:model.allure.qatools.yandex.ru" start="%d" stop="0"><title>%s</title><test-cases></test-cases><labels></labels></ns2:test-suite>`, startTime.Unix(), name), string(suiteData))
}

func TestSuite_AddLabel(t *testing.T) {
	startTime := time.Now()
	name := "suite1"
	suite := NewSuite(name, startTime)

	labelKey := "framework"
	labelValue := "JUnit"
	label := &Label{Name: labelKey, Value: labelValue}

	suite.AddLabel(label)

	suiteData,err := xml.Marshal(suite)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<ns2:test-suite xmlns:ns2="urn:model.allure.qatools.yandex.ru" start="%d" stop="0"><title>%s</title><test-cases></test-cases><labels><label name="%s" value="%s"></label></labels></ns2:test-suite>`, startTime.Unix(), name, labelKey, labelValue), string(suiteData))
}

func TestSuite_AddTest(t *testing.T) {
	testName := "t1"
	testStartTime := time.Now()
	testcase := NewTestCase(testName, testStartTime)

	startTime := time.Now()
	name := "suite1"
	suite := NewSuite(name, startTime)

	suite.AddTest(testcase)
	_,err := xml.Marshal(suite)
	assert.Nil(t, err)
	//TODO: check data
}