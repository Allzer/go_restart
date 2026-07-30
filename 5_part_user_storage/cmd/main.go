package main

import (
	"go_restart/5_part_user_storage/internal/fileStorage"
	"go_restart/5_part_user_storage/internal/memoryStorage"
	"go_restart/5_part_user_storage/internal/registration"
	"go_restart/5_part_user_storage/internal/storage"
	"go_restart/5_part_user_storage/internal/user"
)

func main() {
	var memorySt storage.UserStorage = &memorystorage.MemoryStorage{}
	var fileSt storage.UserStorage = filestorage.FileStorage{}

	user := user.User{
		Name: "Doni",
		Email: "shapkro@gmail.com",
	}
	registartion.UserRegistration(memorySt, user)
	registartion.UserRegistration(fileSt, user)
}
