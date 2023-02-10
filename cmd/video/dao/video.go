package dao

import (
	"context"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	Id       int64  `json:"id"`
	AuthorId int64  `json:"author_id"`
	PlayUrl  string `json:"play_url"`
	CoverUrl string `json:"cover_url"`
	Title    string `json:"title"`
}

func (v *Video) TableName() string {
	return "video_module"
}
func GetFeed(ctx context.Context, latestTime int64, token string) ([]*Video, error) {
	var videos []*Video
	condition := make(map[string]interface{})
	if latestTime != 0 {
		condition["created_at"] = latestTime
	}
	err := DB.WithContext(ctx).Where(condition).Find(&videos).Error
	if err != nil {
		return nil, err
	}
	return videos, nil
}
