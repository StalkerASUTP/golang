package user

import (
	"go/adv-api/pkg/db"
)

type UserRepository struct {
	Database *db.Db
}

func NewUserRepository(database *db.Db) *UserRepository {
	return &UserRepository{
		Database: database,
	}
}

func (repo *UserRepository) Create(user *User) (*User, error) {
	result := repo.Database.DB.Create(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
func (repo *UserRepository) FindByEmail(email string) (*User, error) {
	var user User
	result := repo.Database.DB.First(&user, "email = ?", email)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

// func (repo *LinkRepository) GetById(id uint) (*Link, error) {
// 	var link Link
// 	result := repo.Database.DB.First(&link, id)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return &link, nil
// }
// func (repo *LinkRepository) Update(link *Link) (*Link, error) {
// 	result := repo.Database.DB.Clauses(clause.Returning{}).Updates(link)
// 	if result.Error != nil {
// 		return nil, result.Error
// 	}
// 	return link, nil
// }
// func (repo *LinkRepository) Delete(id uint) error {
// 	result := repo.Database.DB.Delete(&Link{}, id)
// 	if result.Error != nil {
// 		return result.Error
// 	}
// 	return nil
// }
