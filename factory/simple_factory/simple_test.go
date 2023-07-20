package simplefactory

import (
	"testing"
)

func TestJSON(t *testing.T) {
	f := NewRuleConfigParserFactory()
	parser := f.CreateParser("json")
	rc := parser.parse("")
	if rc.format != "json" {
		t.Fatal("JSON test fail")
	}
}

func TestXML(t *testing.T) {
	f := NewRuleConfigParserFactory()
	parser := f.CreateParser("xml")
	rc := parser.parse("")
	if rc.format != "xml" {
		t.Fatal("XML test fail")
	}
}
