package models

import (
	"golang-web/src/config"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name  string `gorm:"type:varchar(255)"`
	Price int
	Stock int
}

func SelectALL() *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("SELECT * FROM articles").Scan(&items)
	// return items

	items := []Product{}
	return config.DB.Find(&items)

}

func Select(id string) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("SELECT * FROM articles WHERE id = ?", id).Scan(&items)
	// return items
	var item Product
	return config.DB.First(&item, "id = ?", id)
}

func Create(newProduct *Product) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("INSERT INTO `articles` (`id`, `created_at`, `updated_at`, `deleted_at`, `title`, `slug`, `description`) VALUES(Null, NULL, NULL, NULL, ?, ?, ?)", Title, Slug, Desc).Scan(&items)
	// return items
	return config.DB.Create(&newProduct)
}

func Updates(id string, updateProduct *Product) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("UPDATE articles SET title = ?, id = ? , description = ? WHERE id = ?", Title, id, Desc, Id).Scan(&items)
	// return items
	var item Product
	return config.DB.Model(&item).Where("id = ?", id).Updates(&updateProduct)
}

func Deletes(id string) *gorm.DB {
	// items := []Article{}
	// config.DB.Raw("DELETE FROM articles WHERE id = ?", Id).Scan(&items)
	// return items
	var item Product
	return config.DB.Delete(&item, "id = ?", id)
}
