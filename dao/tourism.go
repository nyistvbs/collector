package dao

import (
	"collector/model"
	"context"
)

func (d *Dao) TourismInsert(ctx context.Context, row *model.TourismDB) {
	d.db.Create(row)
}
