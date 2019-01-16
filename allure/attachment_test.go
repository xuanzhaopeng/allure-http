package allure

import (
	"testing"
	"github.com/stretchr/testify/assert"
	"encoding/xml"
	"fmt"
)

var (
	title  = "my attachment"
	source = "abc.png"
	mime   = "image/png"
	size   = 1024
)

func TestNewAttachment(t *testing.T) {
	attachment := NewAttachment(title, mime, source, size)
	data, err := xml.Marshal(attachment)
	assert.Nil(t, err)
	assert.Equal(t, fmt.Sprintf(`<attachment title="%s" type="%s" size="%d" source="%s"></attachment>`, title, mime, size, source), string(data))
}
