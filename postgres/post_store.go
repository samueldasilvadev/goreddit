package postgres

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"samueldasilvadev.com/goreddit"
)

type PostStore struct {
	*sqlx.DB
}

func (s *PostStore) Post(id uuid.UUID) (goreddit.Post, error) {
	var p goreddit.Post
	if err := s.Get(&p, `SELECT * FROM posts WHERE id = $1`, id); err != nil {
		return goreddit.Post{}, fmt.Errorf("error getting posts: %w", err)
	}
	return p, nil
}

func (s *PostStore) PostsByThread(threadId uuid.UUID) ([]goreddit.Post, error) {
	var pp []goreddit.Post
	if err := s.Select(&pp,
		`SELECT * FROM posts WHERE thread_id = $1`, threadId); err != nil {
		return []goreddit.Post{}, fmt.Errorf("error getting posts by thread: %w", err)
	}
	return pp, nil
}

func (s *PostStore) CreatePost(p *goreddit.Post) error {
	if err := s.Get(
		p,
		`INSERT INTO posts VALUES (thread_id = $1, title = $2, content = $3, votes = $4) RETURNING *`,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes,
	); err != nil {
		return fmt.Errorf("error inserting post: %w", err)
	}
	return nil
}

func (s *PostStore) UpdatePost(p *goreddit.Post) error {
	if err := s.Get(
		p,
		`UPDATE posts SET thread_id = $1, title = $2, content = $3, votes = $4 WHERE id = $5 RETURNING *`,
		p.ThreadID,
		p.Title,
		p.Content,
		p.Votes,
		p.ID,
	); err != nil {
		return fmt.Errorf("error updating posts: %w", err)
	}
	return nil
}

func (s *PostStore) DeletePost(id uuid.UUID) error {
	if _, err := s.Exec(
		`DELETE FROM posts WHERE id = $1`,
		id,
	); err != nil {
		return fmt.Errorf("error deleting post: %w", err)
	}
	return nil
}
