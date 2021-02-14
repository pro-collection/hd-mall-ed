package categoryController

import (
	"github.com/thoas/go-funk"
	"github.com/ulule/deepcopier"
	"hd-mall-ed/packages/common/database/tableModel"
)

func handleListHelper(list []*tableModel.CategoryBase) []*listItemResultStruct {
	var resultList []*listItemResultStruct
	var subordinate []*tableModel.CategoryBase

	for _, base := range list {
		if base.ParentId == 0 {
			item := &listItemResultStruct{}
			_ = deepcopier.Copy(base).To(item)
			resultList = append(resultList, item)
		} else {
			subordinate = append(subordinate, base)
		}
	}

	// 这是一个可行的办法1
	//for _, base := range subordinate {
	//	for _, result := range resultList {
	//		if result.ID == base.ParentId {
	//			result.Children = append(result.Children, base)
	//		}
	//	}
	//
	//}

	// 实现方法2
	for _, base := range subordinate {
		result := funk.Find(resultList, func(result *listItemResultStruct) bool {
			if result.ID == base.ParentId {
				return true
			}
			return false
		})

		if !funk.IsEmpty(result) {
			result.(*listItemResultStruct).Children = append(result.(*listItemResultStruct).Children, base)
		}
	}

	return resultList
}
