package memorystorage

import "go_restart/5_part_user_storage/internal/user"

type MemoryStorage struct {
	Users []user.User
}

func (m MemoryStorage) Save(user user.User) {
	m.Users = append(m.Users, user)
}