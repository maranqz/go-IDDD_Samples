package domain

type Count int64

type ProductID int64

type Product struct {
	ProductID ProductID
	Count     Count
}

const MaxProducts = 100

type Basket struct {
	UserID   UserID
	Products []Product
}

func NewBasket(userID UserID, products []Product) (Basket, error) {
	return Basket{UserID: userID, Products: products}, nil
}

func (b *Basket) IsValid() bool {
	return b.Validate() == nil
}

func (b *Basket) Validate() error {
	if err := b.UserID.Validate(); err != nil {
		return b.UserID.Validate()
	}

	return nil
}

func checkMaxProducts(pp []Product) bool {
	count := Count(0)

	for _, p := range pp {
		count += p.Count
	}

	return count < MaxProducts
}
