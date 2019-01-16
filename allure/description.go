package allure

import "encoding/xml"

type DescriptionType string

const (
	TEXT     DescriptionType = "text"
	HTML     DescriptionType = "html"
	MARKDOWN DescriptionType = "markdown"
)

type Description struct {
	XMLName   xml.Name `xml:"description"`
	Value string `xml:",innerxml"`
	Type  DescriptionType `xml:"type,attr"`
}

func NewDescription(value string, descriptionType DescriptionType) *Description {
	return &Description{
		Value: value,
		Type:  descriptionType,
	}
}
