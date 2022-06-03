package m3u8

import "log"

type Parser struct {
	Config *Config
}

func NewParser(c *Config) *Parser {
	if c.Num == 0 {
		c.Num = DefaultConfig.Num
	}
	if c.Host == "" {
		c.Host = DefaultConfig.Host
	}
	if c.Out == "" {
		c.Out = DefaultConfig.Out
	}
	return &Parser{
		Config: c,
	}
}

func (p *Parser) Parse() {
	defer func() {
		if r := recover(); r != nil {
			log.Println(r)
		}
	}()
	Run(p.Config)
}
