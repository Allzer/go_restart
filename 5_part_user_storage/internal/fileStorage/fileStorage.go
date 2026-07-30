package filestorage

import (
	"fmt"
	"go_restart/5_part_user_storage/internal/user"
)

type FileStorage struct {}

func (FileStorage) Save(user user.User) {
	fmt.Printf("Пользователь %s сохранён", user.Name)
}