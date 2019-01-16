package allure

// Reference from https://github.com/GabbyyLS/allure-go-common/blob/master/beans/step.go

import (
	"time"
	"encoding/xml"
)

type Status string

const (
	PASSED  Status = "passed"
	FAILED  Status = "failed"
	PENDING Status = "pending"
	BROKEN  Status = "broken"
)

func NewStep(title string, start time.Time) *Step {
	step := new(Step)
	step.Title = title
	step.Attachments = make([]*Attachment, 0)
	step.Steps = make([]*Step, 0)

	if !start.IsZero() {
		step.Start = start.Unix()
	} else {
		step.Start = time.Now().Unix()
	}

	return step
}

type Step struct {
	XMLName xml.Name `xml:"step"`
	Parent  *Step    `xml:"-"`

	Status      Status        `xml:"status,attr"`
	Start       int64         `xml:"start,attr"`
	Stop        int64         `xml:"stop,attr"`
	Title       string        `xml:"title"`
	Steps       []*Step       `xml:"steps>step"`
	Attachments []*Attachment `xml:"attachments>attachments"`
}

func (s *Step) End(status Status, end time.Time) {
	if !end.IsZero() {
		s.Stop = end.Unix()
	} else {
		s.Stop = time.Now().Unix()
	}
	s.Status = status
}

func (s *Step) AddStep(step *Step) {
	if step != nil {
		s.Steps = append(s.Steps, step)
	}
}

func (s *Step) AddAttachment(attachment *Attachment) {
	if attachment != nil {
		s.Attachments = append(s.Attachments, attachment)
	}
}
