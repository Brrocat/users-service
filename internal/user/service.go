package user

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateUser(user *User) error {
	return s.repo.CreateUser(user)
}

func (s *Service) GetUserByID(id uint) (*User, error) {
	return s.repo.GetUserByID(id)
}

func (s *Service) GetAllUsers() ([]User, error) {
	return s.repo.GetAllUsers()
}

func (s *Service) UpdateUserByID(id uint, user *User) (*User, error) {
	return s.repo.UpdateUserByID(id, user)
}

func (s *Service) DeleteUserByID(id uint) error {
	return s.repo.DeleteUserByID(id)
}
