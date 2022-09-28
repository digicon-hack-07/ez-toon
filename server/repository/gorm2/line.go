package gorm2

import (
	"context"

	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/oklog/ulid/v2"
)

func (repo *Repository) SelectLines(ctx context.Context, pageID ulid.ULID) ([]*repository.Line, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var lines []*repository.Line
	err = tx.Where("page_id = ?", pageID).Find(&lines).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return lines, nil
}
