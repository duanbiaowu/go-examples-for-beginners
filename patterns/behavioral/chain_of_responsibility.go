package behavioral

type SensitiveWordFilter interface {
	Filter(string) bool
}

type SensitiveWordFilterChain struct {
	filters []*SensitiveWordFilter
}

type AdWordFilter struct{}

type PornographicWordFilter struct{}

func (c *SensitiveWordFilterChain) AddFilter(f SensitiveWordFilter) {
	c.filters = append(c.filters, &f)
}

func (c *SensitiveWordFilterChain) Filter(content string) bool {
	for i := range c.filters {
		if (*c.filters[i]).Filter(content) {
			return true
		}
	}
	return false
}

func (ad *AdWordFilter) Filter(content string) bool {
	return false
}

func (p *PornographicWordFilter) Filter(content string) bool {
	return true
}
