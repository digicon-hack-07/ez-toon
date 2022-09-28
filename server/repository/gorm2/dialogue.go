package gorm2

import (
	"context"

	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/oklog/ulid/v2"
)

func (repo *Repository) SelectDialogues(ctx context.Context, pageID ulid.ULID) ([]*repository.Dialogue, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var dialogues []*repository.Dialogue
	err = tx.Where("page_id = ?", pageID).Find(&dialogues).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return dialogues, nil
}
