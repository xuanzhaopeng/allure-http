package allure

import (
	"testing"
	"time"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestNewTestCase(t *testing.T) {
	testName := "t1"
	startTime := time.Now()
	testcase := NewTestCase(testName, startTime)
	testData, err := xml.Marshal(testcase)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<test-case status="" start="%d" stop="0"><name>%s</name><steps></steps><labels></labels><attachments></attachments></test-case>`, startTime.Unix(), testName), string(testData))
}

func TestTestCase_SetDescription(t *testing.T) {
	testName := "t1"
	startTime := time.Now()
	desc := "hello"
	testcase := NewTestCase(testName, startTime)
	testcase.SetDescription(desc, TEXT)

	testData, err := xml.Marshal(testcase)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<test-case status="" start="%d" stop="0"><name>%s</name><steps></steps><labels></labels><attachments></attachments><description type="text">%s</description></test-case>`, startTime.Unix(), testName, desc), string(testData))
}

func TestTestCase_AddAttachment(t *testing.T) {
	testName := "t1"
	startTime := time.Now()
	testcase := NewTestCase(testName, startTime)

	attachmentTitle := "log"
	source := "abc.png"
	mime   := "image/png"
	size   := 1024
	attachment := NewAttachment(attachmentTitle, mime,source, size)

	testcase.AddAttachment(attachment)

	testData, err := xml.Marshal(testcase)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<test-case status="" start="%d" stop="0"><name>%s</name><steps></steps><labels></labels><attachments><attachment title="%s" type="%s" size="%d" source="%s"></attachment></attachments></test-case>`, startTime.Unix(), testName, attachmentTitle, mime, size, source), string(testData))
}

func TestTestCase_AddLabel(t *testing.T) {
	testName := "t1"
	startTime := time.Now()
	testcase := NewTestCase(testName, startTime)

	labelKey := "feature"
	labelValue := "My feature"
	label := &Label{Name: labelKey, Value: labelValue}

	testcase.AddLabel(label)

	testData, err := xml.Marshal(testcase)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<test-case status="" start="%d" stop="0"><name>%s</name><steps></steps><labels><label name="%s" value="%s"></label></labels><attachments></attachments></test-case>`, startTime.Unix(), testName, labelKey, labelValue), string(testData))
}

func TestTestCase_AddStep(t *testing.T) {
	stepTitle := "s1"
	stepStartTime := time.Now()
	stepEndTime := stepStartTime.Add(10 * time.Second)
	step := NewStep(stepTitle, stepStartTime)
	step.End(PASSED, stepEndTime)

	testName := "t1"
	startTime := time.Now()
	testcase := NewTestCase(testName, startTime)
	testcase.AddStep(step)

	testData, err := xml.Marshal(testcase)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<test-case status="" start="%d" stop="0"><name>%s</name><steps><step status="%s" start="%d" stop="%d"><title>%s</title><steps></steps><attachments></attachments></step></steps><labels></labels><attachments></attachments></test-case>`, startTime.Unix(), testName, PASSED, stepStartTime.Unix(),stepEndTime.Unix(), stepTitle), string(testData))
}

func TestTestCase_EndFailed(t *testing.T) {
	testName := "t1"
	startTime := time.Now()
	endTime := startTime.Add(10 * time.Second)
	testcase := NewTestCase(testName, startTime)
	failure := &Failure{Msg: "error message", Trace: "a\nb\nc\n"}
	testcase.End(FAILED, failure , endTime)

	_, err := xml.Marshal(testcase)
	assert.Nil(t, err)
	//TODO: Check data
}

