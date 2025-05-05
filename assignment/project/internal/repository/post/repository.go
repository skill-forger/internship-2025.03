package tag

import (
	"strings"

	"gorm.io/gorm"

	"golang-project/internal/model"
	repo "golang-project/internal/repository"
)

// repository represents the implementation of repository.Post
type repository struct {
	db *gorm.DB
}

// NewRepository returns a new implementation of repository.Post
func NewRepository(db *gorm.DB) repo.Post {
	return &repository{db: db}
}

// Read finds and returns the post model by id
func (r *repository) Read(id int) (*model.Post, error) {
	var post model.Post
	err := r.db.First(&post, id).Error
	if err != nil {
		return nil, err
	}
	return &post, nil
}

// Update updates the post model in the database
func (r *repository) Update(post *model.Post) error {
	return r.db.Save(post).Error
}

// SelectPostTag finds and returns all post_tag models by post ids
func (r *repository) SelectPostTag(id []int) ([]*model.PostTag, error) {
	var result []*model.PostTag

	query := r.db.Table("post_tag").
		Where("post_id IN (?)", id).
		Find(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	return result, nil
}

// DeletePostTag delete the post_tag models in the database
func (r *repository) DeletePostTag(post_tags []*model.PostTag) error {
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

// InsertPostTag insert the post_tag models in the database
func (r *repository) InsertPostTag(post_tags []*model.PostTag) error {
	return r.db.Create(&post_tags).Error
}

// UpdatePostTag updates the post_tag models in the database
func (r *repository) UpdatePostTag(postID int, tagIDs []int) ([]*model.PostTag, error) {
	exist, existErr := r.SelectPostTag([]int{postID})
	if existErr != nil {
		return nil, existErr
	}

	checkMap := make(map[int]bool)
	deletePT := make([]*model.PostTag, 0, len(exist))
	addPT := make([]*model.PostTag, 0, len(tagIDs))
	notChangePT := make([]*model.PostTag, 0, len(exist))

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
		if isExist {
			notChangePT = append(notChangePT, newPT)
		} else {
			addPT = append(addPT, newPT)
		}
	}

	if deleteErr := r.DeletePostTag(deletePT); deleteErr != nil {
		return nil, deleteErr
	}

	if addErr := r.InsertPostTag(addPT); addErr != nil {
		return nil, addErr
	}
	return append(notChangePT, addPT...), nil
}
