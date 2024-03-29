package repository

import (
	"errors"
	"fmt"

	"github.com/qin-team-recipe/02-recipe-api/internal/domain"
	"gorm.io/gorm"
)

type RecipeRepository struct{}

func (rr *RecipeRepository) FirstByWatchID(db *gorm.DB, watchID string) (*domain.Recipes, error) {
	recipe := &domain.Recipes{}
	if err := db.Where("watch_id = ?", watchID).First(&recipe).Error; err != nil {
		return &domain.Recipes{}, fmt.Errorf("Not found: %w", errors.New("recipe is not found"))
	}
	return recipe, nil
}

func (rr *RecipeRepository) Find(db *gorm.DB) ([]*domain.Recipes, error) {
	recipes := []*domain.Recipes{}
	db.Find(&recipes)
	if len(recipes) <= 0 {
		return []*domain.Recipes{}, fmt.Errorf("Not found: %w", errors.New("recipes is not found"))
	}
	return recipes, nil
}

func (rr *RecipeRepository) FindByQuery(db *gorm.DB, userID, cursor, limit int, q string) ([]*domain.Recipes, error) {
	recipes := []*domain.Recipes{}

	query := db.
		Joins("left outer join chef_recipes as cr on recipes.id = cr.recipe_id").
		Limit(limit)

	if q != "" {
		q = "%" + q + "%"
		query = query.Where("recipes.title like ? or recipes.description like ?", q, q)
	}
	if 0 < userID {
		query = query.
			Joins("left outer join user_recipes as ur on recipes.id = ur.recipe_id").
			Or("ur.user_id = ?", userID)
	}

	if 0 < cursor {
		query = query.Where("recipes.published_status = ? and 0 < cr.chef_id and recipes.id < ?", "public", cursor)
		// query = query.Where("recipes.id < ?", cursor)
	} else {
		query = query.Where("recipes.published_status = ? and 0 < cr.chef_id", "public")
	}

	query.Order("recipes.created_at desc").Find(&recipes)
	if len(recipes) <= 0 {
		return []*domain.Recipes{}, fmt.Errorf("Not found: %w", errors.New("recipes is not found"))
	}
	return recipes, nil
}

// func (rr *RecipeRepository) FindByChefID(db *gorm.DB, chefID int) ([]*domain.Recipes, error) {
// 	recipes := []*domain.Recipes{}
// 	db.Where("chef_id = ?", chefID).Find(&recipes)
// 	if len(recipes) <= 0 {
// 		return []*domain.Recipes{}, fmt.Errorf("Not found: %w", errors.New("recipes is not found"))
// 	}
// 	return recipes, nil
// }

func (rr *RecipeRepository) FindByUserID(db *gorm.DB, userID int) ([]*domain.Recipes, error) {
	recipes := []*domain.Recipes{}
	db.Where("user_id = ?", userID).Find(&recipes)
	if len(recipes) <= 0 {
		return []*domain.Recipes{}, fmt.Errorf("Not found: %w", errors.New("recipes is not found"))
	}
	return recipes, nil
}

func (rr *RecipeRepository) FindInRecipeIDs(db *gorm.DB, ids []int) ([]*domain.Recipes, error) {
	recipes := []*domain.Recipes{}
	db.Where("id in ?", ids).Find(&recipes)
	if len(recipes) <= 0 {
		return []*domain.Recipes{}, fmt.Errorf("Not found: %w", errors.New("recipes is not found"))
	}
	return recipes, nil
}

func (rr *RecipeRepository) FirstByID(db *gorm.DB, id int) (*domain.Recipes, error) {
	recipe := &domain.Recipes{}
	if err := db.First(recipe, id).Error; err != nil {
		return &domain.Recipes{}, fmt.Errorf("not found: %w", err)
	}
	return recipe, nil
}

func (rr *RecipeRepository) Create(db *gorm.DB, recipe *domain.Recipes) (*domain.Recipes, error) {
	recipe.PublishedStatus = "public"
	if err := db.Create(recipe).Error; err != nil {
		return &domain.Recipes{}, fmt.Errorf("failed recipe create: %w", err)
	}
	return recipe, nil
}

func (rr *RecipeRepository) Save(db *gorm.DB, recipe *domain.Recipes) (*domain.Recipes, error) {
	if err := db.Save(recipe).Error; err != nil {
		return &domain.Recipes{}, fmt.Errorf("failed recipe save: %w", err)
	}
	return recipe, nil
}
