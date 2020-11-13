package shoppingCart

import "fmt"

type MakeChart interface {
	NewChart() map[string]int
}

type Cart struct {
	Items map[string]int
}

func (cart Cart) NewCart() Cart {
	cart.Items = make(map[string]int)
	return cart
}

func (cart Cart) TambahProduk(produk string, ammount int) Cart {

	switch {
	case cart.Items[produk] == 0:
		cart.Items[produk] = ammount
	default:
		cart.Items[produk] = cart.Items[produk] + ammount
	}
	return cart
}

func (cart Cart) HapusProduk(produk string) {

	delete(cart.Items, produk)
	return
}

func (cart Cart) TampilkanCart() {
	fmt.Println(cart.Items)
	return
}

func shoppingCart() {
	var keranjang Cart
	keranjang = Cart.NewCart(keranjang)

	keranjang.TambahProduk("Pisang Hijau", 2)
	keranjang.TambahProduk("Semangka Kuning", 3)
	keranjang.TambahProduk("Apel Merah", 1)
	keranjang.TambahProduk("Apel Merah", 4)
	keranjang.TambahProduk("Apel Merah", 2)

	keranjang.HapusProduk("Semangka Kuning")
	keranjang.HapusProduk("Semangka Merah")

	keranjang.TampilkanCart()
}
