package models

import (
	"fmt"
	"time"
)

type Contract struct {
	ContractId    int
	ContractNo    string
	TemplateId    int
	PrepareChk    bool
	InternalChk   bool
	UseChk        bool
	Title         string
	Desc          string
	Sum           float64
	StartDatetime time.Time
	EndDatetime   time.Time
	InUserId      string
	InDatetime    time.Time
	ModiUserId    string
	ModiDatetime  time.Time
}

func init() {
	fmt.Println("======================models.Contract.init 已执行")
}
