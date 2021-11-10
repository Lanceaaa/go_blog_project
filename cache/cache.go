package cache

// 设置/添加一个缓存，如果 key 存在，用新值覆盖旧值；
// 通过 key 获取一个缓存值；
// 通过 key 删除一个缓存值；
// 删除最“无用”的一个缓存值；
// 获取缓存已存在的记录数；
type Cache interface {
	Set(key string, value interface{})
	Get(key string) interface{}
	Del(key string)
	DelOldest()
	Len() int
}
