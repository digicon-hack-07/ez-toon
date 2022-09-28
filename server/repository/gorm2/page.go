package gorm2

import (
	"context"

	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/oklog/ulid/v2"
)

func (repo *Repository) SelectProjectPageNum(ctx context.Context, projectID ulid.ULID) (int, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	err = tx.Model(&repository.Page{}).Where("project_id = ?", projectID).Count(&count).Error
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

func (repo *Repository) SelectProjectPages(ctx context.Context, projectID ulid.ULID) ([]*repository.Page, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}

	var pages []*repository.Page
	err = tx.Where("project_id = ?", projectID).Order("index asc").Find(&pages).Error
	if err != nil {
		return nil, err
	}

	return pages, nil
}
