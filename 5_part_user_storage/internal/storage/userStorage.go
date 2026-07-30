package storage

import "go_restart/5_part_user_storage/internal/user"

type UserStorage interface {
	Save(user.User)
}