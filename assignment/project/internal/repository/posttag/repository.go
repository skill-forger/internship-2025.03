package posttag

import (
	"strings"

	"gorm.io/gorm"

	"golang-project/internal/model"
	repo "golang-project/internal/repository"
)

// repository represents the implementation of repository.PostTag
type repository struct {
	db *gorm.DB
}

// NewRepository returns a new implementation of repository.PostTag
func NewRepository(db *gorm.DB) repo.PostTag {
	return &repository{db: db}
}

// Select finds and returns all post_tag models by post ids
func (r *repository) Select(id []int) ([]*model.PostTag, error) {
	var result []*model.PostTag

	query := r.db.Table("post_tag").
		Where("post_id IN (?)", id).
		Find(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// Delete delete the post_tag models from the database
func (r *repository) Delete(post_tags []*model.PostTag) error {
	if len(post_tags) == 0 {
		return nil
	}
	conditions := make([]string, 0, len(post_tags))
	args := make([]interface{}, 0, len(post_tags)*2)

	for _, pt := range post_tags {
		conditions = append(conditions, "(post_id = ? AND tag_id = ?)")
		args = append(args, pt.PostID, pt.TagID)
	}

	whereClause := strings.Join(conditions, " OR ")

	return r.db.Where(whereClause, args...).Delete(&model.PostTag{}).Error
}

// Insert insert the post_tag models into the database
func (r *repository) Insert(post_tags []*model.PostTag) error {
	return r.db.Create(&post_tags).Error
}

// Update updates the post_tag models by post ids
func (r *repository) Update(postID int, tagIDs []int) error {
	exist, existErr := r.Select([]int{postID})
	if existErr != nil {
		return existErr
	}

	checkMap := make(map[int]bool)
	deletePT := make([]*model.PostTag, 0, len(exist))
	addPT := make([]*model.PostTag, 0, len(tagIDs))

	for _, tagID := range tagIDs {
		checkMap[tagID] = false
	}

	for _, postTag := range exist {
		if _, ok := checkMap[postTag.TagID]; ok {
			checkMap[postTag.TagID] = true
		} else {
			deletePT = append(deletePT, postTag)
		}
	}

	for tagID, isExist := range checkMap {
		newPT := &model.PostTag{
			PostID: postID,
			TagID:  tagID,
		}
		if !isExist {
			addPT = append(addPT, newPT)
		}
	}

	if deleteErr := r.Delete(deletePT); deleteErr != nil {
		return deleteErr
	}

	if addErr := r.Insert(addPT); addErr != nil {
		return addErr
	}
	return nil
}
