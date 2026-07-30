package order

type Payment interface {
	Pay(*Order) bool
}