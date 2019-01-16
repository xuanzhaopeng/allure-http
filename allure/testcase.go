package allure

// Reference from https://github.com/GabbyyLS/allure-go-common/blob/master/beans/test.go

import (
	"time"
	"encoding/xml"
)

func NewTestCase(name string, start time.Time) *TestCase {
	test := new(TestCase)
	test.Name = name

	if !start.IsZero() {
		test.Start = start.Unix()
	} else {
		test.Start = time.Now().Unix()
	}

	return test
}

type Failure struct {
	XMLName xml.Name `xml:"failure"`
	Msg   string `xml:"message"`
	Trace string `xml:"stack-trace"`
}

type TestCase struct {
	XMLName     xml.Name      `xml:"test-case"`
	Status      Status        `xml:"status,attr"`
	Start       int64         `xml:"start,attr"`
	Stop        int64         `xml:"stop,attr"`
	Name        string        `xml:"name"`
	Steps       []*Step       `xml:"steps>step"`
	Labels      []*Label      `xml:"labels>label"`
	Attachments []*Attachment `xml:"attachments>attachment"`
	Desc        *Description  `xml:"description"`
	Failure     *Failure      `xml:"failure,omitempty"`
}

func (t *TestCase) SetDescription(desc string, descType DescriptionType) {
	t.Desc = NewDescription(desc, descType)
}

func (t *TestCase) AddLabel(label *Label) {
	t.Labels = append(t.Labels, label)
}

func (t *TestCase) AddStep(step *Step) {
	t.Steps = append(t.Steps, step)
}

func (t *TestCase) AddAttachment(attach *Attachment) {
	t.Attachments = append(t.Attachments, attach)
}

func (t *TestCase) End(status Status, failure *Failure, end time.Time) {
	if !end.IsZero() {
		t.Stop = end.Unix()
	} else {
		t.Stop = time.Now().Unix()
	}
	t.Status = status
	if failure != nil {
		t.Failure = failure
	}
}