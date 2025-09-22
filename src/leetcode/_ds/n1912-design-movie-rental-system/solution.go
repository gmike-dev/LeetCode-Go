package n1912_design_movie_rental_system

import (
	"slices"
)

const MaxMovieId = 10_000

type MovieRentingSystem struct {
	unrented [][]unrentedItem
	price    map[movieKey]int
	rented   []rentedItem
}

type movieKey struct {
	movie int
	shop  int
}

type unrentedItem struct {
	price int
	shop  int
}

type rentedItem struct {
	price int
	shop  int
	movie int
}

func compareUnrentedItem(a, b unrentedItem) int {
	c := a.price - b.price
	if c != 0 {
		return c
	}
	return a.shop - b.shop
}

func compareRentedItem(a, b rentedItem) int {
	c := a.price - b.price
	if c != 0 {
		return c
	}
	c = a.shop - b.shop
	if c != 0 {
		return c
	}
	return a.movie - b.movie
}

func Constructor(n int, entries [][]int) MovieRentingSystem {
	s := MovieRentingSystem{
		unrented: make([][]unrentedItem, MaxMovieId+1),
		price:    make(map[movieKey]int),
		rented:   make([]rentedItem, 0),
	}
	for _, e := range entries {
		shop, movie, price := e[0], e[1], e[2]
		s.unrented[movie] = append(s.unrented[movie], unrentedItem{
			price: price,
			shop:  shop,
		})
		s.price[movieKey{
			movie: movie,
			shop:  shop,
		}] = price
	}
	for _, u := range s.unrented {
		slices.SortFunc(u, compareUnrentedItem)
	}
	return s
}

func (s *MovieRentingSystem) Search(movie int) []int {
	var res []int
	for _, u := range s.unrented[movie][:min(len(s.unrented[movie]), 5)] {
		res = append(res, u.shop)
	}
	return res
}

func (s *MovieRentingSystem) Rent(shop int, movie int) {
	price := s.price[movieKey{
		movie: movie,
		shop:  shop,
	}]
	u := unrentedItem{
		price: price,
		shop:  shop,
	}
	i, _ := slices.BinarySearchFunc(s.unrented[movie], u, compareUnrentedItem)
	s.unrented[movie] = append(s.unrented[movie][:i], s.unrented[movie][i+1:]...)
	r := rentedItem{
		price: price,
		shop:  shop,
		movie: movie,
	}
	i, _ = slices.BinarySearchFunc(s.rented, r, compareRentedItem)
	s.rented = append(s.rented, rentedItem{})
	copy(s.rented[i+1:], s.rented[i:])
	s.rented[i] = r
}

func (s *MovieRentingSystem) Drop(shop int, movie int) {
	price := s.price[movieKey{
		movie: movie,
		shop:  shop,
	}]
	r := rentedItem{
		price: price,
		shop:  shop,
		movie: movie,
	}
	i, _ := slices.BinarySearchFunc(s.rented, r, compareRentedItem)
	s.rented = append(s.rented[:i], s.rented[i+1:]...)

	u := unrentedItem{
		price: price,
		shop:  shop,
	}
	i, _ = slices.BinarySearchFunc(s.unrented[movie], u, compareUnrentedItem)
	s.unrented[movie] = append(s.unrented[movie], unrentedItem{})
	copy(s.unrented[movie][i+1:], s.unrented[movie][i:])
	s.unrented[movie][i] = u
}

func (s *MovieRentingSystem) Report() [][]int {
	var res [][]int
	for _, r := range s.rented[:min(len(s.rented), 5)] {
		res = append(res, []int{r.shop, r.movie})
	}
	return res
}
