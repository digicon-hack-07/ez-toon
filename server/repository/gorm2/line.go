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

func (repo *Repository) InsertLine(ctx context.Context, id ulid.ULID, pageID ulid.ULID, penSize int, points []repository.Point) (*repository.Line, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	line := repository.Line{
		ID:      id,
		PageID:  pageID,
		PenSize: penSize,
		Points:  points,
	}

	err = tx.Create(&line).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &line, nil
}

func (repo *Repository) DeleteLine(ctx context.Context, id ulid.ULID) error {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.Where("id = ?", id).Delete(&repository.Line{}).Error
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}
