package textract

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadMissingFile(t *testing.T) {
	text, err := RetrieveTextFromFile("testdats/thisfiledoesntexist")

	assert.Equal(t, "", text, "Reading a non-existent file should return empty string")
	assert.NotNil(t, err, "Error expected reading a non-existent file")
}

// func TestReadEmptyMSDocx(t *testing.T) {
// 	text, err := RetrieveTextFromFile("testdats/empty.docx")
//
// 	text = strings.TrimSpace(text)
//
// 	assert.Equal(t, "", text, "Reading an empty file should return empty string")
// 	assert.Nil(t, err, "Nil error expected")
// }

func TestReadMSDocxSample1(t *testing.T) {
	const correct = `P1PLAIN-1234567890 EOL
P1PLAIN-abcdefghijklmnopqrstuvwxyz MIDDLE EOL
P1BOLD-1234567890 EOL

P2BOLD-abcdefghijklmnopqrstuvwxyz
P2ITALIC-1234567890
P2ITALIC-abcdefghijklmnopqrstuvwxyz

P3UNDERSCORE-1234567890
P3UNDERSCORE-abcdefghijklmnopqrstuvwxyz
P3BOLD-ITALIC-1234567890

P4BOLD-ITALIC-abcdefghijklmnopqrstuvwxyz`

	text, err := RetrieveTextFromFile("testdats/sample1.docx")

	text = strings.TrimSpace(text)

	assert.Equal(t, correct, text, "Incorrect content extracted from sample1.docx")
	assert.Nil(t, err, "Nil error expected")
}

func TestReadMSDocxFromReader(t *testing.T) {
	const correct = `P1PLAIN-1234567890 EOL
P1PLAIN-abcdefghijklmnopqrstuvwxyz MIDDLE EOL
P1BOLD-1234567890 EOL

P2BOLD-abcdefghijklmnopqrstuvwxyz
P2ITALIC-1234567890
P2ITALIC-abcdefghijklmnopqrstuvwxyz

P3UNDERSCORE-1234567890
P3UNDERSCORE-abcdefghijklmnopqrstuvwxyz
P3BOLD-ITALIC-1234567890

P4BOLD-ITALIC-abcdefghijklmnopqrstuvwxyz`

	filepath := "testdats/sample1.docx"
	file, err := os.Open(filepath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()
	buf, err := io.ReadAll(file)
	if err != nil {
		t.Fatal(err)
	}

	r := bytes.NewReader(buf)
	text, err := RetrieveText(r, int64(len(buf)))

	text = strings.TrimSpace(text)

	assert.Equal(t, correct, text, "Incorrect content extracted from sample1.docx")
	assert.Nil(t, err, "Nil error expected")
}
