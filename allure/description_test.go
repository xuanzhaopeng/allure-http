package allure

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
	"fmt"
)

var (
	htmlData = `<html><body><h1>test</h1></body></html>`
	markDownData = `#Test
		* test1

		> Hello

		## Test2
	`
)

func TestNewDescription(t *testing.T) {
	htmlDescription := NewDescription(htmlData, HTML)
	html, err := xml.Marshal(htmlDescription)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<description type="html">%s</description>`, htmlData), string(html))

	mdDescription := NewDescription(markDownData, MARKDOWN)
	md, err := xml.Marshal(mdDescription)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<description type="markdown">%s</description>`, markDownData), string(md))
}
