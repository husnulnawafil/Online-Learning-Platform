package repositories

import "context"

type StatisticRepository interface {
	CountCourse(ctx context.Context, isFree bool) (int64, error)
	CountUser(ctx context.Context) (int64, error)
}
