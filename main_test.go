package main

import (
	"testing"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
	"gorm.io/playground/models"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	userId := uuid.Must(uuid.NewV4())
	projectId := uuid.Must(uuid.NewV4())
	layouts := []models.Layout{}

	sql1 := DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		byUser := func(dbScope *gorm.DB) *gorm.DB {
			if userId.Valid {
				return dbScope.Where(dbScope.Where("private_for = ?", userId).Or("private_for IS NULL"))
			}
			return dbScope
		}
		return tx.Table("layouts").Scopes(byUser).Find(&layouts, map[string]any{"project_id": projectId})
	})
	t.Logf(sql1)

	sql2 := DB.ToSQL(func(tx *gorm.DB) *gorm.DB {
		byUser := func(dbScope *gorm.DB) *gorm.DB {
			if userId.Valid {
				return dbScope.Where("(private_for = ? OR private_for IS NULL)", userId)
			}
			return dbScope
		}
		return tx.Table("layouts").Scopes(byUser).Find(&layouts, map[string]any{"project_id": projectId})
	})
	t.Logf(sql2)
}

// func TestGORMGen(t *testing.T) {
// 	user := models.User{Name: "jinzhu2"}
// 	ctx := context.Background()

// 	gorm.G[models.User](DB).Create(ctx, &user)

// 	if u, err := gorm.G[models.User](DB).Where(g.User.ID.Eq(user.ID)).First(ctx); err != nil {
// 		t.Errorf("Failed, got error: %v", err)
// 	} else if u.Name != user.Name {
// 		t.Errorf("Failed, got user name: %v", u.Name)
// 	}
// }
