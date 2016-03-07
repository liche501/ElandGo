package models

import (
	"fmt"
	"time"
)

type Template struct {
	TemplateId   int
	TemplateName string
	Desc         string
	UseChk       bool
	InUserId     string
	InDatetime   time.Time
	ModiUserId   string
	ModiDatetime time.Time
}

func init() {
	fmt.Println("======================models.Template.init 已执行")
}
