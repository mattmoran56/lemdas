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
	Operator        string `json:"operator"`
	Object          string `json:"object"`
	OperatorOperand string `json:"operand"`
	Value           string `json:"value"`
	Field           string `json:"field"`
}

func NewSearchRepository(db *gorm.DB) *SearchRepository {
	return &SearchRepository{db: db}
}

func (s *SearchRepository) Search(querys []SearchQuery) ([]models.Dataset, []models.File, error) {
	var files []models.File
	var datasets []models.Dataset

	filesQuery := s.db.Table("files")
	datasetsQuery := s.db.Table("datasets")

	for _, query := range querys {
		if query.Object == "file" {
			switch query.Operator {
			case "AND":
				zap.S().Debug("AND")
				filesQuery = filesQuery.Where("id IN (?)", s.db.Table("file_attributes").
					Select("file_id").Where("file_attributes.attribute_name = ? AND file_attributes.attribute_value = ?", query.Field, query.Value))
			case "OR":
				filesQuery = filesQuery.Or("id IN (?)", s.db.Table("file_attributes").
					Select("file_id").Where("file_attributes.attribute_name = ? AND file_attributes.attribute_value = ?", query.Field, query.Value))
			case "NOT":
				filesQuery = filesQuery.Not("id IN (?)", s.db.Table("file_attributes").
					Select("file_id").Where("file_attributes.attribute_name = ? AND file_attributes.attribute_value = ?", query.Field, query.Value))
			}
		} else if query.Object == "file_value" {
			switch query.Operator {
			case "AND":
				filesQuery = filesQuery.Where("? = ?", query.Field, query.Value)
			case "OR":
				filesQuery = filesQuery.Or("? = ?", query.Field, query.Value)
			case "NOT":
				filesQuery = filesQuery.Not("? = ?", query.Field, query.Value)
			}
		} else if query.Object == "dataset" {
			switch query.Operator {
			case "AND":
				datasetsQuery = datasetsQuery.Where("id IN (?)", s.db.Table("dataset_attributes").
					Select("dataset_id").Where("dataset_attributes.attribute_name = ? AND dataset_attributes.attribute_value = ?", query.Field, query.Value))
			case "OR":
				datasetsQuery = datasetsQuery.Or("id IN (?)", s.db.Table("dataset_attributes").
					Select("dataset_id").Where("dataset_attributes.attribute_name = ? AND dataset_attributes.attribute_value = ?", query.Field, query.Value))
			case "NOT":
				datasetsQuery = datasetsQuery.Not("id IN (?)", s.db.Table("dataset_attributes").
					Select("dataset_id").Where("dataset_attributes.attribute_name = ? AND dataset_attributes.attribute_value = ?", query.Field, query.Value))
			}
		} else if query.Object == "dataset_value" {
			switch query.Operator {
			case "AND":
				datasetsQuery = datasetsQuery.Where("? = ?", query.Field, query.Value)
			case "OR":
				datasetsQuery = datasetsQuery.Or("? = ?", query.Field, query.Value)
			case "NOT":
				datasetsQuery = datasetsQuery.Not("? = ?", query.Field, query.Value)
			}
		}
	}
	filesQuery.Find(&files)

	return datasets, files, nil
}
