package n3508_implement_router

type Packet struct {
	source      int
	destination int
	timestamp   int
}

type Router struct {
	queue       []Packet
	packets     map[Packet]bool
	timestamps  map[int][]int
	memoryLimit int
}

func Constructor(memoryLimit int) Router {
	return Router{
		queue:       make([]Packet, 0, memoryLimit),
		packets:     make(map[Packet]bool),
		timestamps:  make(map[int][]int),
		memoryLimit: memoryLimit,
	}
}

func (s *Router) AddPacket(source int, destination int, timestamp int) bool {
	np := Packet{
		source,
		destination,
		timestamp,
	}
	if s.packets[np] {
		return false
	}
	s.packets[np] = true
	s.queue = append(s.queue, np)
	s.timestamps[destination] = append(s.timestamps[destination], timestamp)

	if len(s.queue) > s.memoryLimit {
		s.Forward()
	}
	return true
}

func (s *Router) Forward() Packet {
	p := s.queue[0]
	s.packets[p] = false
	s.timestamps[p.destination] = s.timestamps[p.destination][1:]
	s.queue = s.queue[1:]
	return p
}

func (s *Router) ForwardPacket() []int {
	if len(s.queue) == 0 {
		return []int{}
	}
	p := s.Forward()
	return []int{p.source, p.destination, p.timestamp}
}

func (s *Router) GetCount(destination int, startTime int, endTime int) int {
	ts := s.timestamps[destination]
	if len(ts) == 0 {
		return 0
	}
	l := lowerBound(ts, 0, len(ts), startTime)
	if l == len(ts) {
		return 0
	}
	r := upperBound(ts, l, len(ts), endTime)
	return r - l
}

func lowerBound(a []int, left, right, value int) int {
	l, r := left, right-1
	for l < r {
		m := l + (r-l)/2
		if a[m] < value {
			l = m + 1
		} else {
			r = m
		}
	}
	if a[l] < value {
		return right
	}
	return l
}

func upperBound(a []int, left, right, value int) int {
	l, r := left, right-1
	for l < r {
		m := l + (r-l)/2
		if a[m] <= value {
			l = m + 1
		} else {
			r = m
		}
	}
	if a[l] <= value {
		return right
	}
	return l
}
