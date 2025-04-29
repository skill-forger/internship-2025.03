package post

import (
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
	var result model.Post

	query := r.db.Model(&model.Post{}).
		Preload("User").
		Where("id = ? AND is_published = ?", id, true).
		First(&result)

	if err := query.Error; err != nil {
		return nil, err
	}

	// Get tags for the post
	tags, err := r.GetTags(id)
	if err != nil {
		return nil, err
	}
	result.Tags = tags

	return &result, nil
}

// Insert creates a new post in the database
func (r *repository) Insert(post *model.Post) (*model.Post, error) {
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

// AddPostTags adds multiple tag associations to a post
func (r *repository) AddPostTags(postID int, tagIDs []int) error {
	// Remove duplicate tag IDs using empty struct for memory efficiency
	uniq := make(map[int]struct{}, len(tagIDs))
	pivots := make([]*model.PostTag, 0, len(tagIDs))

	for _, id := range tagIDs {
		if _, dup := uniq[id]; dup {
			continue
		}
		uniq[id] = struct{}{}
		pivots = append(pivots, &model.PostTag{
			PostID: postID,
			TagID:  id,
		})
	}

	// If no associations to create, return early
	if len(pivots) == 0 {
		return nil
	}

	// Create post tag associations
	return r.db.Create(&pivots).Error
}

// FindSlugsLike retrieves all slugs that start with the given baseSlug
func (r *repository) FindSlugsLike(baseSlug string) ([]string, error) {
	var slugs []string
	pattern := baseSlug + "%"
	err := r.db.Model(&model.Post{}).Where("slug LIKE ?", pattern).Pluck("slug", &slugs).Error
	if err != nil {
		return nil, err
	}
	return slugs, nil
}

// GetTags retrieves all tags associated with a post
func (r *repository) GetTags(postID int) ([]*model.Tag, error) {
	var tags []*model.Tag

	err := r.db.Table("tags").
		Joins("JOIN post_tag ON tags.id = post_tag.tag_id").
		Where("post_tag.post_id = ? AND tags.deleted_at IS NULL", postID).
		Find(&tags).Error

	if err != nil {
		return nil, err
	}

	return tags, nil
}

// ReadByCondition finds a post based on provided conditions with optional preloads
func (r *repository) ReadByCondition(condition map[string]interface{}, preloads ...string) (*model.Post, error) {
	var post model.Post

	// Build query with preloads
	query := r.db.Model(&model.Post{})

	// Apply preloads
	for _, preload := range preloads {
		query = query.Preload(preload)
	}

	// Apply conditions
	for field, value := range condition {
		query = query.Where(field+" = ?", value)
	}

	// Execute query
	if err := query.First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}
