// Package simplefactory
// 示例代码：根据配置文件的扩展名，如 json、xml、yaml、toml、ini、properties，选择不同的解析器，如 JSONRuleConfigParser、XMLRuleConfigParser等，
// 将存储在文件中的配置解析成内存对象：RuleConfig。、
package simplefactory

import (
	"fmt"
)

type RuleConfig struct {
	format  string
	content any
}

type IRuleConfigParser interface {
	// 将存储在文件中的配置解析成内存对象
	parse(configText string) *RuleConfig
}

// NewRuleConfigParser return parser instance by format,
// 但是每次调用都会创建一个新的 Parser，应当复用并缓存，见 RuleConfigParserFactory.
func NewRuleConfigParser(configFormat string) IRuleConfigParser {
	if configFormat == "" {
		return nil
	}

	return createParser(configFormat)
}

type RuleConfigSource struct{}

// load
//
//	 @Date 20230720
//		@Description: load and parse config file
//		@receiver rc
//		@param ruleConfigFilePath
//		@return *RuleConfig
//		@return error
func (rc *RuleConfigSource) load(ruleConfigFilePath string) (*RuleConfig, error) {
	ruleConfigFileExtension := rc.getFileExtension(ruleConfigFilePath)
	parser := createParser(ruleConfigFileExtension)
	if parser == nil {
		return nil, fmt.Errorf("rule config file format is not supported: %s", ruleConfigFileExtension)
	}

	configText := ""
	return parser.parse(configText), nil
}

func (rc *RuleConfigSource) getFileExtension(ruleConfigFilePath string) string {
	// 解析文件名，以获取扩展名，如 rule.json，返回 json
	return "json"
}

func createParser(configFormat string) IRuleConfigParser {
	var parser IRuleConfigParser
	// 实际上，如果 switch 分支并不是很多，代码中存在 switch 分支也是可以接受的。
	switch configFormat {
	case "json":
		parser = &JSONRuleConfigParser{}
	default:
		parser = nil
	}

	return parser
}

type RuleConfigParserFactory struct {
	// 复用 parser：事先创建并缓存 parser
	cachedParsers map[string]IRuleConfigParser
}

var cachedParsers = map[string]IRuleConfigParser{
	"json": new(JSONRuleConfigParser),
	"xml":  new(XMLRuleConfigParser),
}

func NewRuleConfigParserFactory() *RuleConfigParserFactory {
	// 如果添加新的 parser，势必改动此处的代码，是不是违反了开闭原则？
	// 实际上，如果并非频繁地添加新的 parser，只是偶尔修改，即便实现不符合开闭原则，也是可以接受的。
	pf := &RuleConfigParserFactory{}
	if len(cachedParsers) > 0 {
		pf.cachedParsers = cachedParsers
	}
	return pf
}

func (f *RuleConfigParserFactory) CreateParser(configFormat string) IRuleConfigParser {
	if configFormat == "" {
		return nil
	}

	if len(f.cachedParsers) == 0 {
		return nil
	}

	parser, ok := f.cachedParsers[configFormat]
	if !ok {
		return nil
	}

	return parser
}

type JSONRuleConfigParser struct{}

func (P *JSONRuleConfigParser) parse(configText string) *RuleConfig {
	return &RuleConfig{
		format: "json",
	}
}

type XMLRuleConfigParser struct{}

func (p *XMLRuleConfigParser) parse(configText string) *RuleConfig {
	return &RuleConfig{format: "xml"}
}
