package service

// 灰度策略常量。
const (
	StrategyCanary    = "canary"
	StrategyBlueGreen = "blue_green"
)

// DefaultSchedule 返回指定策略的默认放量百分比序列。
func DefaultSchedule(strategy string) []int {
	switch strategy {
	case StrategyBlueGreen:
		return []int{50, 100}
	default:
		return []int{5, 20, 50, 100}
	}
}

// StrategyInfo 灰度策略描述。
type StrategyInfo struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Schedule    []int  `json:"schedule"`
}

// ListStrategies 返回支持的灰度策略。
func (s *Service) ListStrategies() []StrategyInfo {
	return []StrategyInfo{
		{Name: StrategyCanary, Description: "金丝雀：小流量逐步放大，风险可控", Schedule: DefaultSchedule(StrategyCanary)},
		{Name: StrategyBlueGreen, Description: "蓝绿：先切一半流量，验证后全量", Schedule: DefaultSchedule(StrategyBlueGreen)},
	}
}

// PlanByStrategy 按策略为发布单自动生成放量步骤。
func (s *Service) PlanByStrategy(releaseID, strategy string) (int, error) {
	schedule := DefaultSchedule(strategy)
	return s.PlanRollout(releaseID, schedule, "")
}
