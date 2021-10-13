package metatext

import (
	"strings"

	"google.golang.org/grpc/metadata"
)

type MetadataTextMap struct {
	metadata.MD
}

// 基于 TextMap 模式，对照实现了 metadata 的设置和读取方法
func (m MetadataTextMap) ForeachKey(handler func(key, val string) error) error {
	for k, vs := range m.MD {
		for _, v := range vs {
			if err := handler(k, v); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m MetadataTextMap) Set(key, val string) {
	key = strings.ToLower(key)
	m.MD[key] = append(m.MD[key], val)
}
