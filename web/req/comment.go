package req

import "github.com/bangumi/server/internal/model"

type EpisodeComment struct {
	ID      model.CommentID `json:"id,omitempty"`
	FieldID model.CommentID `json:"field_id,omitempty"`
	Comment string          `json:"comment"`
}
