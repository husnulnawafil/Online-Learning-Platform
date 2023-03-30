package services

import (
	"context"

	statisticRepositories "github.com/husnulnawafil/online-learning-platform/repositories/statistics"
)

type StatisticServiceInstance struct {
	statisticRepo statisticRepositories.StatisticRepository
}

func NewStatisticService() StatisticService {
	repoStatistic := statisticRepositories.NewStatisticRepository()
	return &StatisticServiceInstance{
		statisticRepo: repoStatistic,
	}
}

func (s *StatisticServiceInstance) CountCourse(ctx context.Context, isFree bool) (int64, error) {
	return s.statisticRepo.CountCourse(ctx, isFree)
}
func (s *StatisticServiceInstance) CountUser(ctx context.Context) (int64, error) {
	return s.statisticRepo.CountUser(ctx)
}
