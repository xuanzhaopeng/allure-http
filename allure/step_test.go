package allure

import (
	"testing"
	"time"
	"encoding/xml"
	"github.com/stretchr/testify/assert"
	"fmt"
)

func TestNewStep(t *testing.T) {
	stepTitle := "t1"
	startTime := time.Now()
	step := NewStep(stepTitle, startTime)
	stepData, err := xml.Marshal(step)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<step status="" start="%d" stop="0"><title>%s</title><steps></steps><attachments></attachments></step>`, startTime.Unix(), stepTitle), string(stepData))
}

func TestStep_AddStep(t *testing.T) {
	stepTitle := "t1"
	startTime := time.Now()
	step := NewStep(stepTitle, startTime)

	subStepTitle := "s1"
	subStartTime := time.Now()
	subStep := NewStep(subStepTitle, subStartTime)

	step.AddStep(subStep)

	stepData, err := xml.Marshal(step)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<step status="" start="%d" stop="0"><title>%s</title><steps><step status="" start="%d" stop="0"><title>%s</title><steps></steps><attachments></attachments></step></steps><attachments></attachments></step>`, startTime.Unix(),stepTitle, subStartTime.Unix(), subStepTitle), string(stepData))
}

func TestStep_AddAttachment(t *testing.T) {
	stepTitle := "t1"
	startTime := time.Now()
	step := NewStep(stepTitle, startTime)

	attachmentTitle := "log"
	source := "abc.png"
	mime   := "image/png"
	size   := 1024
	attachment := NewAttachment(attachmentTitle, mime,source, size)
	step.AddAttachment(attachment)

	stepData, err := xml.Marshal(step)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<step status="" start="%d" stop="0"><title>%s</title><steps></steps><attachments><attachment title="%s" type="%s" size="%d" source="%s"></attachment></attachments></step>`, startTime.Unix(), stepTitle, attachmentTitle, mime, size, source), string(stepData))
}

func TestStep_End(t *testing.T) {
	stepTitle := "t1"
	startTime := time.Now()
	endTime := startTime.Add(10 * time.Second)
	step := NewStep(stepTitle, startTime)
	step.End(PASSED, endTime)

	stepData, err := xml.Marshal(step)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<step status="passed" start="%d" stop="%d"><title>%s</title><steps></steps><attachments></attachments></step>`, startTime.Unix(), endTime.Unix(), stepTitle), string(stepData))
}
