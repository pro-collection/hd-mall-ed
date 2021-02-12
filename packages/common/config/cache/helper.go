package cache

import "fmt"

// @title           设置缓存key的 title
// @description
// @auth            晴小篆  331393627@qq.com
// @param           id
// @return          key
func AminCacheKey(id int) string {
	return fmt.Sprintf("admin - %d", id)
}
