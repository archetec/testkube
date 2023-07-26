package parser

import (
	"encoding/xml"
)

// TestResults is a root element of junit xml report
type TestResults struct {
	XMLName     xml.Name     `xml:"testResults"`
	HTTPSamples []HTTPSample `xml:"httpSample,omitempty"`
}

// HTTPSample is http sample details
type HTTPSample struct {
	XMLName         xml.Name         `xml:"httpSample"`
	Time            int              `xml:"t,attr"`
	Success         bool             `xml:"s,attr"`
	Label           string           `xml:"lb,attr"`
	ResponseCode    string           `xml:"rc,attr"`
	AssertionResult *AssertionResult `xml:"assertionResult"`
}

// AssertionResult contains assertion
type AssertionResult struct {
	XMLName        xml.Name `xml:"assertionResult"`
	Name           string   `xml:"name"`
	Failure        bool     `xml:"failure"`
	Error          bool     `xml:"error"`
	FailureMessage string   `xml:"failureMessage"`
}

func ParseXML(data []byte) (results TestResults, err error) {
	err = xml.Unmarshal(data, &results)

	return results, err
}
