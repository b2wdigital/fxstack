package util

import (
	"encoding/json"
	"fmt"

	logger "github.com/sirupsen/logrus"
	"github.com/yudai/gojsondiff/formatter"

	diff "github.com/yudai/gojsondiff"
)

// JSONComparator holds a source and a expected value, as well as de comparator
type JSONComparator struct {
	source   []byte
	expected []byte
	diff     diff.Diff
}

// Equals check  if its equal
func (c *JSONComparator) Equals() bool {
	return !c.diff.Modified()
}

// Print shows the diff
func (c *JSONComparator) Print() {

	var diffString string
	var aJSON map[string]interface{}

	err := json.Unmarshal(c.source, &aJSON)
	if err != nil {
		logger.Error(err)
		return
	}

	config := formatter.AsciiFormatterConfig{
		ShowArrayIndex: false,
		Coloring:       true,
	}

	f := formatter.NewAsciiFormatter(aJSON, config)
	diffString, _ = f.Format(c.diff)
	fmt.Println(diffString)
}

// NewJSONComparator instatiates a json comparator
func NewJSONComparator(source []byte, expected []byte) (*JSONComparator, error) {

	differ := diff.New()

	d, err := differ.Compare(source, expected)
	if err != nil {
		return nil, err
	}

	return &JSONComparator{
		diff:     d,
		source:   source,
		expected: expected,
	}, nil
}
