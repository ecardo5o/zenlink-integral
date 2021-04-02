package test

import (
	"encoding/json"
	"fmt"
	"testing"
	"zenlink-integral/config"
)

func TestConfig(t *testing.T) {
	c := make([]config.ChainCfg, 1)
	nodes := make([]string, 1)
	n := config.ChainCfg{
		ChainId: 0,
		Nodes:   nodes,
	}
	c = append(c, n)
	cfg := config.DbConfig{
		Chains:      c,
	}
	b,_:= json.Marshal(cfg)
	fmt.Println(string(b))
}