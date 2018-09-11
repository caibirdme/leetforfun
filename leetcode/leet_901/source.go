package leet_901

type StockSpanner struct {
	f     []int
	price []int
}

func Constructor() StockSpanner {
	return StockSpanner{}
}

func (this *StockSpanner) Next(price int) int {
	this.price = append(this.price, price)
	if len(this.f) == 0 {
		this.f = append(this.f, 1)
		return 1
	}
	ans := 1
	i := len(this.price) - 2
	for i >= 0 && this.price[i] <= price {
		ans += this.f[i]
		i -= this.f[i]
	}
	this.f = append(this.f, ans)
	return ans
}

/**
 * Your StockSpanner object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Next(price);
 */
