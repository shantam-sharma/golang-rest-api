package users

/*
	|| Module Docs ||
	- Business Logic here
	- what should happen - in what order - what rules are valid
	- Brain of the app , generate user id , validate , check etc.
*/
type Service struct {
	store  *Store
	nextID int
}

// equv. constructor . parameter . return type is pointere to service
func NewService(store *Store) *Service {
	return &Service{
		store:  store,
		nextID: 1,
	}
}

//belongs to service str . func name and parametre . return type user
func (s *Service) CreateUser(name string, email string) User {
	// user is a variable that stores User .
	user := User{
		ID:    s.nextID,
		Name:  name,
		Email: email,
	}
	s.nextID++

	s.store.users[user.ID] = user
	return user
}

//check in storage if the user id exists
func (s *Service) GetUser(userID int) (User, bool) {
	return s.store.GetUser(userID)
}

//delete user
func (s *Service) DeleteUser(userID int) {
	s.store.DeleteUser(userID)
}

func (s *Service) UpdateUser(userID int, name string, email string) User {
	updatedUser := User{
		ID:    userID,
		Name:  name,
		Email: email,
	}
	s.store.UpdateUser(updatedUser)
	return updatedUser
}
