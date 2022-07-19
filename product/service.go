package product

type Service interface {
	FindAll() ([]Product, error)
	FindByID(ID int) (Product, error)
	Create(productRequest ProductRequest) (Product, error)
	Update(ID int, productRequest ProductRequest) (Product, error)
	Delete(ID int) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Product, error) {
	products, err := s.repository.FindAll()
	return products, err
}

func (s *service) FindByID(ID int) (Product, error) {
	product, err := s.repository.FindByID(ID)
	return product, err
}

func (s *service) Create(productRequest ProductRequest) (Product, error) {
	price, _ := productRequest.Price.Int64()

	product := Product{
		Title:       productRequest.Title,
		Price:       int(price),
		Description: productRequest.Description,
	}

	newProduct, err := s.repository.Create(product)

	return newProduct, err
}

func (s *service) Update(ID int, productRequest ProductRequest) (Product, error) {
	product, err := s.repository.FindByID(ID)

	price, _ := productRequest.Price.Int64()

	product.Title = productRequest.Title
	product.Price = int(price)
	product.Description = productRequest.Description

	newProduct, err := s.repository.Update(product)

	return newProduct, err
}

func (s *service) Delete(ID int) (Product, error) {
	product, err := s.repository.FindByID(ID)
	newProduct, err := s.repository.Delete(product)
	return newProduct, err
}
