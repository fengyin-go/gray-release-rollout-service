package service

import (
	"grayrelease/internal/model"
	"grayrelease/pkg/idgen"
)

// PresetTemplate 内置放量模板定义。
type PresetTemplate struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Schedule    []int  `json:"schedule"`
}

// PresetTemplates 返回内置放量模板清单。
func PresetTemplates() []PresetTemplate {
	return []PresetTemplate{
		{Name: "金丝雀-激进", Description: "5%→20%→50%→100%，快速验证", Schedule: []int{5, 20, 50, 100}},
		{Name: "金丝雀-保守", Description: "1%→5%→10%→25%→50%→100%，稳妥推进", Schedule: []int{1, 5, 10, 25, 50, 100}},
		{Name: "蓝绿发布", Description: "50%→100%，两阶段切换", Schedule: []int{50, 100}},
		{Name: "三步放量", Description: "10%→50%→100%", Schedule: []int{10, 50, 100}},
	}
}

// SeedPresetTemplates 将内置模板写入存储（幂等，已存在则跳过）。
func (s *Service) SeedPresetTemplates() int {
	created := 0
	for _, p := range PresetTemplates() {
		t := &model.ReleaseTemplate{
			ID:          idgen.Hex(),
			Name:        p.Name,
			Description: p.Description,
			Schedule:    p.Schedule,
		}
		if err := s.store.CreateReleaseTemplate(t); err == nil {
			created++
		}
	}
	return created
}
