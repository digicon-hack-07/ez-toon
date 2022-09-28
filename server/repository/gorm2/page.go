package gorm2

import (
	"context"
	"errors"

	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
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

	tx.Commit()
	return int(count), nil
}

func (repo *Repository) SelectProjectPages(ctx context.Context, projectID ulid.ULID) ([]*repository.Page, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var pages []*repository.Page
	err = tx.Where("project_id = ?", projectID).Order("`index` ASC").Find(&pages).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return pages, nil
}

func (repo *Repository) InsertPage(ctx context.Context, id ulid.ULID, projectID ulid.ULID, height int, width int) (*repository.Page, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	lastIndex := 0
	err = tx.Table("pages").Select("`index`").Where("project_id = ?", projectID).Order("`index` DESC").Limit(1).Scan(&lastIndex).Error
	if err != nil {
		return nil, err
	}

	page := repository.Page{
		ID:        id,
		ProjectID: projectID,
		Index:     lastIndex + 1,
		Height:    height,
		Width:     width,
	}

	err = tx.Create(&page).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &page, nil
}

func (repo *Repository) SelectPage(ctx context.Context, id ulid.ULID) (*repository.Page, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var page repository.Page
	err = tx.Where("id = ?", id).First(&page).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	tx.Commit()
	return &page, nil
}

func (repo *Repository) UpdateIndex(ctx context.Context, id ulid.ULID, operation string) (*repository.Page, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var page repository.Page
	err = tx.Where("id = ?", id).First(&page).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	switch operation {
	case "inc":
		err = tx.Model(&page).Update("index", page.Index+1).Error
		if err != nil {
			return nil, err
		}
		err := tx.Model(&repository.Page{}).Where("project_id = ? AND `index` = ?", page.ProjectID, page.Index+1).Update("index", page.Index).Error
		if err != nil {
			return nil, err
		}

	case "dec":
		err = tx.Model(&page).Update("index", page.Index-1).Error
		if err != nil {
			return nil, err
		}
		err := tx.Model(&repository.Page{}).Where("project_id = ? AND `index` = ?", page.ProjectID, page.Index-1).Update("index", page.Index).Error
		if err != nil {
			return nil, err
		}

	default:
		return nil, errors.New("invalid operation")
	}

	tx.Commit()
	return &page, nil
}
