package models

import (
	"fmt"
)

type ContractDtl struct {
	ContractNo string
	BookmarkId int
	SeqNo      int
	Value      string
}

func init() {
	fmt.Println("======================models.Bookmark.init 已执行")
}
