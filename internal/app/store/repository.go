package store

import "github.com/katelinlis/Wallbackend/internal/app/model"

//WallRepository ...
type WallRepository interface {
	Create(*model.Wall) error                                                //Создание пользователя
	GetByAuthor(offset int, limit int, userid int) ([]model.Wall, error)     // Получить новости за определенного пользователя
	GetByFriends(offset int, limit int, userids []int) ([]model.Wall, error) // Получить новости друзей и людей на которых подписан пользователь
	GetPost(PostID string) (model.Wall, []model.Wall, error)                 // Получение определенного поста
	ScanAndCreateUUID() error                                                // Сканирование и создание UUID если пусто
	GetAnswersCount(PostID string) (int, error)
	GetAnswers(PostID string) ([]model.Wall, error)
}

//UserRepository ...
type UserRepository interface {
	GetUser(AuthorID int) model.UserObj //Получение данных о пользователе
	GetFriends(int) []int               //Получение списка друзей пользователя
}
