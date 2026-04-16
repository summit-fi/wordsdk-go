package unifiedTime

import "iter"

// UnifiedTimeIterator iterates over UnifiedTime values.
type UnifiedTimeIterator struct {
	current UnifiedTime
	till    UnifiedTime
	unit    TimeUnit
	amount  int
	isFirst bool
}

// NewUnifiedTimeIterator creates a new UnifiedTimeIterator starting from 'start' until 'till', incrementing by 'amount' of 'unit'.
func NewUnifiedTimeIterator(start, till UnifiedTime, unit TimeUnit, amount int) *UnifiedTimeIterator {
	return &UnifiedTimeIterator{
		current: start,
		till:    till,
		unit:    unit,
		amount:  amount,
		isFirst: true,
	}
}

func (it *UnifiedTimeIterator) Current() UnifiedTime {
	return it.current
}

func (it *UnifiedTimeIterator) MoveNext() bool {
	if it.isFirst {
		it.isFirst = false
		return !it.current.Is(RelationAfter, &it.till)
	}
	it.current = it.current.Add(it.unit, it.amount)
	return !it.current.Is(RelationAfter, &it.till)
}

func (it *UnifiedTimeIterator) IteratorSeq() iter.Seq[UnifiedTime] {
	return func(yield func(UnifiedTime) bool) {

		for it.MoveNext() {
			current := it.Current()
			if !yield(current) {
				return
			}
		}
	}
}

func (it *UnifiedTimeIterator) ToList() []UnifiedTime {
	seq := it.IteratorSeq()
	var result []UnifiedTime
	for cur := range seq {
		result = append(result, cur)
	}
	return result
}
