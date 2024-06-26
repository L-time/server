package res

import (
	"time"

	"github.com/bangumi/server/internal/model"
)

type IndexCommentResp struct {
	ID        model.CommentID `json:"id"`
	FieldID   model.IndexID   `json:"field_id"`
	UserID    model.UserID    `json:"user_id"`
	Related   model.CommentID `json:"related"`
	CreatedAt time.Time       `json:"created_at"`
	Content   string          `json:"content"`
}

func ConventIndexCommit2Resp(model model.IndexComment) IndexCommentResp {
	return IndexCommentResp{
		ID:        model.ID,
		FieldID:   model.Field,
		UserID:    model.User,
		Content:   model.Content,
		CreatedAt: model.CreatedAt,
		Related:   model.Related,
	}
}

type EpisodeCommentResp struct {
	ID        model.CommentID `json:"id"`
	FieldID   model.EpisodeID `json:"field_id"`
	UserID    model.UserID    `json:"user_id"`
	Related   model.CommentID `json:"related"`
	CreatedAt time.Time       `json:"created_at"`
	Content   string          `json:"content"`
}

func ConventEpisodeComment2Resp(model model.EpisodeComment) EpisodeCommentResp {
	return EpisodeCommentResp{
		ID:        model.ID,
		FieldID:   model.Field,
		UserID:    model.User,
		Content:   model.Content,
		CreatedAt: model.CreatedAt,
		Related:   model.Related,
	}
}
