package registartion

import (
	"go_restart/5_part_user_storage/internal/storage"
	"go_restart/5_part_user_storage/internal/user"
)

func UserRegistration(storage storage.UserStorage, user user.User) {
	storage.Save(user)
}