package models

import (
	"gorm.io/gorm"
)

type Problem struct {
	gorm.Model
	Identity          string            `json:"identity"`
	Title             string            `json:"title"`
	Content           string            `json:"content"`
	MaxRuntime        int               `json:"max_runtime"`
	MaxMem            int               `json:"max_mem"`
	PassNum           int               `json:"pass_num"`
	SubmitNum         int               `json:"submit_num"`
	ProblemCategories []ProblemCategory `gorm:"foreignKey:problem_id;references:id"`
}

func (Problem) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword string, categoryIdentity string) *gorm.DB {
	var problem Problem
	tx := DB.Model(&problem).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title LIKE ? OR content LIKE ?", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx.Joins("RIGHT JOIN problem_category pc on pc.problem_id = problem_basic.id").
			Where("pc.category_id = (SELECT cb.id from category_basic cb WHERE cb.identity = ?)", categoryIdentity)
	}
	return tx
}
