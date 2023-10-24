package post

type PostService interface {
	CreatePost(title string, body string, userId int64) (PostWithUsername, error)
	GetPost(postId int64) (PostWithUsername, error)
	ListPosts() ([]PostWithUsername, error)
}

type postService struct {
	repo PostRepository
}

func NewPostService(repo PostRepository) PostService {
	return &postService{repo: repo}
}

func (s *postService) CreatePost(title string, body string, userId int64) (PostWithUsername, error) {
	post := Post{Title: title, Body: body, UserID: userId}

	newPost, err := s.repo.CreatePost(post)
	if err != nil {
		return PostWithUsername{}, err
	}

	return newPost, nil

}

func (s *postService) GetPost(postId int64) (PostWithUsername, error) {
	post, err := s.repo.GetPost(postId)
	if err != nil {
		return PostWithUsername{}, err
	}

	return post, nil

}

func (s *postService) ListPosts() ([]PostWithUsername, error) {
	posts, err := s.repo.ListPosts()
	if err != nil {
		return []PostWithUsername{}, err
	}

	return posts, nil
}
