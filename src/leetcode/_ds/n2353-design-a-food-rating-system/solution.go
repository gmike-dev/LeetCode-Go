package n2353_design_a_food_rating_system

import "container/heap"

type HeapItem struct {
	rating int
	food   string
	index  int
}

type MaxHeap []HeapItem

func (h *MaxHeap) Len() int {
	return len(*h)
}
func (h *MaxHeap) Less(i, j int) bool {
	x := (*h)[i]
	y := (*h)[j]
	return x.rating > y.rating || x.rating == y.rating && x.food < y.food
}
func (h *MaxHeap) Swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
	(*h)[i].index = i
	(*h)[j].index = j
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(HeapItem))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type FoodRatings struct {
	cuisine      map[string]string
	rating       map[string]int
	cuisineFoods map[string]*MaxHeap
}

func Constructor(foods []string, cuisines []string, ratings []int) FoodRatings {
	rating := make(map[string]int)
	cuisine := make(map[string]string)
	cuisineFoods := make(map[string]*MaxHeap)
	for i, food := range foods {
		rating[food] = ratings[i]
		cuisine[food] = cuisines[i]
		item := HeapItem{
			rating: ratings[i],
			food:   food,
		}
		if h, ok := cuisineFoods[cuisines[i]]; ok {
			*h = append(*h, item)
		} else {
			cuisineFoods[cuisines[i]] = &MaxHeap{item}
		}
	}
	for _, h := range cuisineFoods {
		heap.Init(h)
	}
	return FoodRatings{
		cuisine:      cuisine,
		rating:       rating,
		cuisineFoods: cuisineFoods,
	}
}

func (fr *FoodRatings) ChangeRating(food string, newRating int) {
	fr.rating[food] = newRating
	heap.Push(fr.cuisineFoods[fr.cuisine[food]], HeapItem{
		rating: newRating,
		food:   food,
	})
}

func (fr *FoodRatings) HighestRated(cuisine string) string {
	cuisineFoods := fr.cuisineFoods[cuisine]
	for (*cuisineFoods)[0].rating != fr.rating[(*cuisineFoods)[0].food] {
		heap.Pop(cuisineFoods)
	}
	return (*cuisineFoods)[0].food
}
