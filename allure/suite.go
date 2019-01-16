package allure

// Reference from https://github.com/GabbyyLS/allure-go-common/blob/master/beans/suite.go

import (
	"encoding/xml"
	"time"
)

const NsModel = `urn:model.allure.qatools.yandex.ru`

type Suite struct {
	XMLName   xml.Name    `xml:"ns2:test-suite"`
	NsAttr    string      `xml:"xmlns:ns2,attr"`
	Start     int64       `xml:"start,attr"`
	Stop      int64       `xml:"stop,attr"`
	Title     string      `xml:"title"`
	TestCases []*TestCase `xml:"test-cases>test-case"`
	Labels    []*Label    `xml:"labels>label"`
}

func NewSuite(name string, start time.Time) *Suite {
	s := new(Suite)

	s.NsAttr = NsModel
	s.Title = name

	if !start.IsZero() {
		s.Start = start.UTC().Unix()
	} else {
		s.Start = time.Now().UTC().Unix()
	}

	return s
}

func (s *Suite) End(endTime time.Time) {
	if !endTime.IsZero() {
		//strict UTC
		s.Stop = endTime.UTC().Unix()
	} else {
		s.Stop = time.Now().UTC().Unix()
	}
}

func (s Suite) HasTests() bool {
	return len(s.TestCases) > 0
}

func (s *Suite) AddTest(test *TestCase) {
	s.TestCases = append(s.TestCases, test)
}

func (s *Suite) AddLabel(label *Label) {
	s.Labels = append(s.Labels, label)
}
