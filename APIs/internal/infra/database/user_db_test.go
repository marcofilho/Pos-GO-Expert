package database

import (
	"testing"

	"github.com/marcofilho/Pos-GO-Expert/APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupUserTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect database: %v", err)
	}
	db.AutoMigrate(&entity.User{})
	return db
}

func TestUserDB(t *testing.T) {

	t.Run("TestNewUser", func(t *testing.T) {
		db := setupUserTestDB(t)

		user, err := entity.NewUser("Maya", "maya_lima@gmail.com", "123456")
		assert.Nil(t, err)

		userDB := NewUser(db)
		err = userDB.Create(user)
		assert.Nil(t, err)

		var userFound entity.User
		err = db.First(&userFound, "id = ?", user.ID).Error
		assert.Nil(t, err)
		assert.Equal(t, user, &userFound)
		assert.Equal(t, user.Name, userFound.Name)
		assert.Equal(t, user.Email, userFound.Email)
		assert.NotNil(t, user.Password)
	})

	t.Run("TestFindByEmail", func(t *testing.T) {
		db := setupUserTestDB(t)

		user, err := entity.NewUser("Maya", "maya_lima@gmail.com", "123456")
		assert.Nil(t, err)

		userDB := NewUser(db)
		err = userDB.Create(user)
		assert.Nil(t, err)

		userFound, err := userDB.FindByEmail(user.Email)
		assert.Nil(t, err)
		assert.Equal(t, user.ID, userFound.ID)
		assert.Equal(t, user.Name, userFound.Name)
		assert.Equal(t, user.Email, userFound.Email)
		assert.NotNil(t, userFound.Password)
	})
}
