package textract

import (
	"errors"
	"io"
)

func RetrieveTextFromFile(path string) (string, error) {
	ft, err := GetTrueFileType(path)
	if err != nil {
		return "", err
	}

	ext := GetFileExtension(path)

	docx := DocxParser{}

	parsers := []DocumentParser{&docx}

	for _, p := range parsers {
		if ft == p.trueType() && ext == p.extension() {
			err := p.readFile(path)
			if err != nil {
				return "", nil
			}

			return p.retrieveTextFromFile()
		}
	}

	return "", errors.New("unsupported file format")
}

func RetrieveText(r io.Reader, size int64) (string, error) {
	docx := DocxParser{}

	parsers := []DocumentParser{&docx}
	ft := DetectFileType(r)

	for _, p := range parsers {
		if ft == p.trueType() || ft == "application/octet-stream" {
			err := p.readFromReader(r, size)
			if err != nil {
				return "", nil
			}

			return p.retrieveTextFromFile()
		}
	}

	return "", errors.New("unsupported file format")
}
