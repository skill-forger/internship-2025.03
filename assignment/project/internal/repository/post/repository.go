package post

import (
	"fmt"

	"gorm.io/gorm"

	"golang-project/internal/contract"
	"golang-project/internal/model"
	repo "golang-project/internal/repository"
	"golang-project/util/pagination"
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
		Preload("Tags", "deleted_at IS NULL").
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
	pattern := fmt.Sprintf("%s%%", baseSlug)
	err := r.db.Model(&model.Post{}).Where("slug LIKE ?", pattern).Pluck("slug", &slugs).Error
	if err != nil {
		return nil, err
	}
	return slugs, nil
}

// GetTags retrieves all tags associated with a post
func (r *repository) GetTags(postID int) ([]*model.Tag, error) {
	var tags []*model.Tag
	err := r.db.
		Model(&model.Tag{}).
		Joins("JOIN post_tag ON post_tag.tag_id = tags.id").
		Where("post_tag.post_id = ? AND tags.deleted_at IS NULL", postID).
		Find(&tags).Error
	if err != nil {
		return nil, err
	}
	return tags, nil
}

// Select retrieves all posts from the database with optional filters
func (r *repository) Select(filter *contract.ListPostRequest) ([]*model.Post, error) {
	var posts []*model.Post

	query := r.db.Model(&model.Post{}).
		Preload("User").
		Preload("Tags", "deleted_at IS NULL").
		Where("is_published = ?", true)

	// Apply filters if provided
	if filter != nil {
		if filter.Title != "" {
			title := fmt.Sprintf("%%%s%%", filter.Title)
			query = query.Where("title LIKE ?", title)
		}
		if filter.Pseudonym != "" {
			pseudonym := fmt.Sprintf("%%%s%%", filter.Pseudonym)
			query = query.Joins("JOIN users ON posts.user_id = users.id").
				Where("users.pseudonym LIKE ?", pseudonym)
		}
		if filter.Tag != "" {
			query = query.Joins("JOIN post_tag ON posts.id = post_tag.post_id").
				Joins("JOIN tags ON post_tag.tag_id = tags.id").
				Where("tags.name = ? AND tags.deleted_at IS NULL", filter.Tag)
		}
	}

	// Apply pagination
	if filter != nil && filter.Page > 0 && filter.PageSize > 0 {
		offset := pagination.CalculateOffset(filter.Page, filter.PageSize)
		query = query.Offset(offset).Limit(filter.PageSize)
	}

	err := query.Find(&posts).Error
	if err != nil {
		return nil, err
	}

	return posts, nil
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
	if len(condition) > 0 {
		query = query.Where(condition)
	}

	// Execute query
	if err := query.First(&post).Error; err != nil {
		return nil, err
	}

	return &post, nil
}

// UpdatePost updates the post model in the database
func (r *repository) UpdatePost(post *model.Post, updatePost map[string]interface{}) error {
	return r.db.Model(&post).Omit("user_id").Updates(updatePost).Error
}

// UpdatePostTag updates the post_tag table by association with post
func (r *repository) UpdatePostTag(post *model.Post, tags []*model.Tag) error {
	return r.db.Model(&post).Association("Tags").Replace(tags)
}

// Delete delete the post model, related comments, and related post_tag relations in the database
// If err on any step => rollback on all delete
func (r *repository) Delete(postID int) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete(&model.Post{}, postID).Error; err != nil {
			return err
		}

		if err := tx.Delete(&model.PostTag{}, "post_id = ?", postID).Error; err != nil {
			return err
		}

		if err := tx.Delete(&model.Comment{}, "post_id = ?", postID).Error; err != nil {
			return err
		}

		return nil
	})
}
