package gorm2

import (
	"context"
	"errors"

	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
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

func (repo *Repository) InsertDialogue(
	ctx context.Context,
	id ulid.ULID,
	pageID ulid.ULID,
	dialogue string,
	top float64,
	bottom float64,
	left float64,
	right float64,
) (*repository.Dialogue, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	dial := repository.Dialogue{
		ID:       id,
		PageID:   pageID,
		Dialogue: dialogue,
		Top:      top,
		Bottom:   bottom,
		Left:     left,
		Right:    right,
	}

	err = tx.Create(&dial).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &dial, nil
}

func (repo *Repository) UpdateDialogue(
	ctx context.Context,
	id ulid.ULID,
	dialogue string,
	top float64,
	bottom float64,
	left float64,
	right float64,
) (*repository.Dialogue, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}

	err = tx.Model(&repository.Dialogue{}).Where("id = ?", id).Updates(repository.Dialogue{
		Dialogue: dialogue,
		Top: top,
		Bottom: bottom,
		Left: left,
		Right: right,
	}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	dial := repository.Dialogue{}
	err = tx.Where("id = ?", id).First(&dial).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &dial, nil
}

func (repo *Repository) DeleteDialogue(ctx context.Context, id ulid.ULID) error {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = tx.Where("id = ?", id).Delete(&repository.Dialogue{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repository.ErrNotFound
		}
		return err
	}

	tx.Commit()
	return nil
}
