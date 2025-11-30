- Создание класса


```
type UsersRepository struct {
	Database *db.Db
}

<!-- Конструктор -->
func NewUsersRepository(database *db.Db) *UsersRepository {
	return &UsersRepository{
		Database: database,
	}
}
```