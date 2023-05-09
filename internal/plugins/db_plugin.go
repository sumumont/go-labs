package plugins

import (
	"fmt"
	"github.com/go-labs/internal/logging"
	"gorm.io/gorm"
	"gorm.io/gorm/callbacks"
)

func New() gorm.Plugin {
	return &sumumont{}
}

type sumumont struct {
}

func (s *sumumont) Initialize(db *gorm.DB) error {
	return db.Callback().Query().Replace("gorm:query", s.QueryDB)
}

func (s *sumumont) Name() string {
	return "gorm:sumumont"
}

func (p *sumumont) QueryDB(tx *gorm.DB) {

	if tx.Error != nil || tx.DryRun {
		return
	}
	//for k, v := range tx.Statement.Clauses {
	//	v.Name = strings.ToUpper(v.Name)
	//	tx.Statement.Clauses[k] = v
	//}
	callbacks.BuildQuerySQL(tx)

	sql := fmt.Sprintf(tx.Statement.SQL.String())
	logging.Debug().Str("sql", sql).Send()

	rows, err := tx.Statement.ConnPool.QueryContext(tx.Statement.Context, sql, tx.Statement.Vars...)
	if err != nil {
		_ = tx.AddError(err)
		return
	}

	defer func() {
		_ = tx.AddError(rows.Close())
	}()

	gorm.Scan(rows, tx, 0)
}
