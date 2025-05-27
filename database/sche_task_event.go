package database

import (
	"context"

	"github.com/jianbo-zh/jydata/database/ent"
	"github.com/jianbo-zh/jydata/database/ent/schetaskevent"

	"entgo.io/ent/dialect/sql"
)

func (m *Database) AddScheTaskEvent(ctx context.Context, event *ent.ScheTaskEvent) (*ent.ScheTaskEvent, error) {
	return m.MainDB().ScheTaskEvent.Create().
		SetScheTaskID(event.ScheTaskID).
		SetState(event.State).
		SetAbnormalState(event.AbnormalState).
		SetRemark(event.Remark).
		Save(ctx)
}

func (m *Database) GetScheTaskEvents(ctx context.Context, taskID int) ([]*ent.ScheTaskEvent, error) {
	return m.MainDB().ScheTaskEvent.Query().
		Where(schetaskevent.ScheTaskIDEQ(taskID)).
		Order(schetaskevent.ByID(sql.OrderDesc())).
		All(ctx)
}
