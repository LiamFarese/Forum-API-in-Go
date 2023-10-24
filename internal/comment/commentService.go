package comment

import "database/sql"

type CommentService interface {
	CreateComment(body string, userId int64, postId int64) (CommentWithUsername, error)
	Reply(body string, userId int64, postId int64, parentComment sql.NullInt64) (CommentWithUsername, error)
	GetCommentsFromPost(postId int64) ([]CommentWithUsername, error)
}

type commentService struct {
	repo CommentRepository
}

func NewCommentService(repo CommentRepository) CommentService {
	return &commentService{repo: repo}
}

func (s *commentService) CreateComment(body string, userId int64, postId int64) (CommentWithUsername, error) {
	comment := Comment{Body: body, UserID: userId, PostID: postId}

	newCommet, err := s.repo.CreateComment(comment)
	if err != nil {
		return CommentWithUsername{}, err
	}

	return newCommet, nil

}

func (s *commentService) Reply(body string, userId int64, postId int64, parentComment sql.NullInt64) (CommentWithUsername, error) {
	comment := Comment{Body: body, UserID: userId, PostID: postId, ParentComment: parentComment}

	newComment, err := s.repo.Reply(comment)
	if err != nil {
		return CommentWithUsername{}, err
	}

	return newComment, nil
}

func (s *commentService) GetCommentsFromPost(postId int64) ([]CommentWithUsername, error) {
	comments, err := s.repo.GetCommentsFromPost(postId)
	if err != nil {
		return []CommentWithUsername{}, err
	}

	return comments, nil
}
