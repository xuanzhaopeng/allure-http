package allure

import "encoding/xml"

func NewAttachment(title, mime, source string, size int) *Attachment {
	return &Attachment{
		Title:  title,
		Type:   mime,
		Source: source,
		Size:   size,
	}
}

type Attachment struct {
	XMLName   xml.Name `xml:"attachment"`
	Title  string `xml:"title,attr"`
	Type   string `xml:"type,attr"`
	Size   int    `xml:"size,attr"`
	Source string `xml:"source,attr"`
}

