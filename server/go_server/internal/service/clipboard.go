/*
业务逻辑层
*/
package service

import "go_server/internal/model"

type ClipboardService struct{}

func NewClipboardService() *ClipboardService {
	return &ClipboardService{}
}

func (s *ClipboardService) GetLatestClip() (*model.Clip, error) {
	// 模拟获取最新剪贴板内容
	return &model.Clip{Content: "Hello, world!"}, nil
}
