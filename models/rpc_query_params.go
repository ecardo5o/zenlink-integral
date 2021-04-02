package models

import (
	"github.com/gin-gonic/gin"
	"log"
)

type RpcQueryParams struct {
	Jsonrpc string   `json:"jsonrpc"`
	ID      int      `json:"id"`
	Method  string   `json:"method"`
	Params  []string `json:"params"`
}

func (p *RpcQueryParams) Parse(c *gin.Context) *RpcQueryParams {
	if c.ShouldBind(&p) == nil {
		log.Println(p.Method)
	}
	return p
}
