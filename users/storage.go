package users

/*
	|| Module Docs ||
	- Queries , all the CURD happens here
*/

type Store struct {
	users map[int]User
}

func NewStore() *Store {
	return &Store{
		users: make(map[int]User),
	}
}

func (s *Store) GetUser(userID int) (User, bool) {
	user, exists := s.users[userID]
	return user, exists
}

func (s *Store) DeleteUser(userID int) {
	delete(s.users, userID)
}
func (s *Store) UpdateUser(user User) {
	s.users[user.ID] = user
}
