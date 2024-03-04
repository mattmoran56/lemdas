package repositories

import (
	"github.com/mattmoran/fyp/api/pkg/database/models"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type SearchRepository struct {
	db *gorm.DB
}

type SearchQuery struct {
	Operand string `json:"operand"`
	Object  string `json:"object"`
	Value   string `json:"value"`
	Field   string `json:"field"`
}

func NewSearchRepository(db *gorm.DB) *SearchRepository {
	return &SearchRepository{db: db}
}

func (s *SearchRepository) Search(querys []SearchQuery, userID string) ([]models.Dataset, []models.File, error) {
	var files []models.File
	var datasets []models.Dataset

	fileSearch := false
	datasetSearch := false

	filesQuery := s.db.Table("files")
	datasetsQuery := s.db.Table("datasets")

	for _, query := range querys {
		if query.Object == "file" {
			fileSearch = true
			switch query.Operand {
			case "AND":
				zap.S().Debug("AND")
				filesQuery = filesQuery.Where("files.id IN (?)", s.db.Table("file_attributes").
					Select("file_id").Where("file_attributes.attribute_name = ? AND file_attributes.attribute_value = ?", query.Field, query.Value))
			case "OR":
				filesQuery = filesQuery.Or("files.id IN (?)", s.db.Table("file_attributes").
					Select("file_id").Where("file_attributes.attribute_name = ? AND file_attributes.attribute_value = ?", query.Field, query.Value))
			case "NOT":
				filesQuery = filesQuery.Not("files.id IN (?)", s.db.Table("file_attributes").
					Select("file_id").Where("file_attributes.attribute_name = ? AND file_attributes.attribute_value = ?", query.Field, query.Value))
			}
		} else if query.Object == "file_value" {
			fileSearch = true
			switch query.Operand {
			case "AND":
				filesQuery = filesQuery.Where("? = ?", query.Field, query.Value)
			case "OR":
				filesQuery = filesQuery.Or("? = ?", query.Field, query.Value)
			case "NOT":
				filesQuery = filesQuery.Not("? = ?", query.Field, query.Value)
			}
		} else if query.Object == "dataset" {
			datasetSearch = true
			switch query.Operand {
			case "AND":
				datasetsQuery = datasetsQuery.Where("dataset.id IN (?)", s.db.Table("dataset_attributes").
					Select("dataset_id").Where("dataset_attributes.attribute_name = ? AND dataset_attributes.attribute_value = ?", query.Field, query.Value))
			case "OR":
				datasetsQuery = datasetsQuery.Or("dataset.id IN (?)", s.db.Table("dataset_attributes").
					Select("dataset_id").Where("dataset_attributes.attribute_name = ? AND dataset_attributes.attribute_value = ?", query.Field, query.Value))
			case "NOT":
				datasetsQuery = datasetsQuery.Not("dataset.id IN (?)", s.db.Table("dataset_attributes").
					Select("dataset_id").Where("dataset_attributes.attribute_name = ? AND dataset_attributes.attribute_value = ?", query.Field, query.Value))
			}
		} else if query.Object == "dataset_value" {
			datasetSearch = true
			switch query.Operand {
			case "AND":
				datasetsQuery = datasetsQuery.Where("? = ?", query.Field, query.Value)
			case "OR":
				datasetsQuery = datasetsQuery.Or("? = ?", query.Field, query.Value)
			case "NOT":
				datasetsQuery = datasetsQuery.Not("? = ?", query.Field, query.Value)
			}
		}
	}

	userSharedDatasets := s.db.Table("user_share_datasets").Select("dataset_id").Where("user_id = ?", userID)
	groupShareDatasets := s.db.Table("group_share_datasets").Select("dataset_id").Where("group_id IN (?)", s.db.Table("group_members").Select("group_id").Where("user_id = ?", userID))

	if fileSearch {
		filesQuery.
			Where("files.status = 'processed'").
			Where("files.owner_id = ? OR files.is_public = 1 OR files.dataset_id IN (?) OR files.dataset_id IN (?)", userID, userSharedDatasets, groupShareDatasets).
			Find(&files)
	}
	if datasetSearch {
		datasetsQuery.Where("datasets.owner_id = ? OR datasets.is_public = 1 OR datasets.id IN (?) OR datasets.id IN (?)", userID, userSharedDatasets, groupShareDatasets).Find(&datasets)
	}

	return datasets, files, nil
}
