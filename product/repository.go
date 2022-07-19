package product

import "gorm.io/gorm"

type Repository interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	Create(product Product) (Product, error)
	Update(product Product) (Product, error)
	Delete(product Product) (Product, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]Product, error) {
	var products []Product
	err := r.db.Find(&products).Error
	return products, err
}

func (r *repository) FindByID(ID int) (Product, error) {
	var product Product
	err := r.db.Find(&product, ID).Error
	return product, err
}

func (r *repository) Create(product Product) (Product, error) {
	err := r.db.Create(&product).Error
	return product, err
}

func (r *repository) Update(product Product) (Product, error) {
	err := r.db.Save(&product).Error
	return product, err
}

func (r *repository) Delete(product Product) (Product, error) {
	err := r.db.Delete(&product).Error
	return product, err
}
