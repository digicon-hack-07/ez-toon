package gorm2

import (
	"context"
	"errors"
	"time"

	"github.com/digicon-hack-07/ez-toon/server/repository"
	"github.com/digicon-hack-07/ez-toon/server/utils/ulid"
	"gorm.io/gorm"
)

func (repo *Repository) SelectProjects(ctx context.Context) ([]*repository.Project, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var projects []*repository.Project
	err = tx.Order("updated_at desc").Find(&projects).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return projects, nil
}

func (repo *Repository) InsertProject(ctx context.Context, id ulid.ULID, name string, thumbnail string) (*repository.Project, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	project := repository.Project{
		ID:        id,
		Name:      name,
		Thumbnail: thumbnail,
	}

	err = tx.Create(&project).Error
	if err != nil {
		return nil, err
	}

	// project = repository.Project{}
	// err = tx.Where("id = ?", id).First(&project).Error
	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return nil, repository.ErrNotFound
	// 	}
	// 	return nil, err
	// }

	project.CreatedAt = time.Now()
	project.UpdatedAt = time.Now()

	tx.Commit()
	return &project, nil
}

func (repo *Repository) SelectProject(ctx context.Context, id ulid.ULID) (*repository.Project, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	var project repository.Project
	err = tx.Where("id = ?", id).First(&project).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	tx.Commit()
	return &project, nil
}

func (repo *Repository) UpdateProject(ctx context.Context, id ulid.ULID, name string) (*repository.Project, error) {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = tx.Model(&repository.Project{}).Where("id = ?", id).Updates(repository.Project{
		Name: name,
	}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	var project repository.Project
	err = tx.Where("id = ?", id).First(&project).Error
	if err != nil {
		return nil, err
	}

	tx.Commit()
	return &project, nil
}

func (repo *Repository) DeleteProject(ctx context.Context, id ulid.ULID) error {
	tx, err := repo.getTX(ctx)
	if err != nil {
		return err
	}

	err = tx.Where("id = ?", id).Delete(&repository.Project{}).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return repository.ErrNotFound
		}
		return err
	}

	tx.Commit()
	return nil
}
