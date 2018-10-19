package gopager

import (
	"fmt"
	"strconv"
)

type TLengthAwarePaginator struct {
	TPaginator
	onEachSide int
}

func NewLengthAwarePaginator(items interface{}, total int, perPage int, currentPage int, options map[string]string) *TLengthAwarePaginator {
	paginator := NewPaginator(items, total, perPage, currentPage, options)

	lengthAwarePaginator := &TLengthAwarePaginator{TPaginator: *paginator}

	if options != nil {
		for optionKey, optionValue := range options {
			if optionKey == "onEachSide" {
				i, _ := strconv.Atoi(optionValue)
				lengthAwarePaginator.onEachSide = i
			}
		}
	}

	if lengthAwarePaginator.onEachSide == 0 {
		lengthAwarePaginator.onEachSide = 3
	}

	return lengthAwarePaginator
}

func (this *TLengthAwarePaginator) getSmallSlider() []map[int]string {
	return []map[int]string{this.getUrlRange(1, this.LastPage), {}, {},}
}

func (this *TLengthAwarePaginator) getUrlSlider(onEachSide int) []map[int]string {
	window := onEachSide * 2

	if ! this.HasPage() {
		return []map[int]string{{}, {}, {},}
	}

	if this.CurrentPage <= window {
		return this.getSliderTooCloseToBeginning(window)
	} else if this.CurrentPage > (this.LastPage - window) {
		return this.getSliderTooCloseToEnding(window)
	}

	return this.getFullSlider(onEachSide)
}

func (this *TLengthAwarePaginator) getSliderTooCloseToBeginning(window int) []map[int]string {
	return []map[int]string{
		this.getUrlRange(1, window+2),
		{},
		this.getFinish(),
	}
}

func (this *TLengthAwarePaginator) getSliderTooCloseToEnding(window int) []map[int]string {
	last := this.getUrlRange(this.LastPage-(window+2), this.LastPage)

	return []map[int]string{this.getStart(), {}, last,}
}

func (this *TLengthAwarePaginator) getFullSlider(onEachSide int) []map[int]string {
	return []map[int]string{this.getStart(), this.getAdjacentUrlRange(onEachSide), this.getFinish(),}
}

func (this *TLengthAwarePaginator) getAdjacentUrlRange(onEachSide int) map[int]string {
	return this.getUrlRange(this.CurrentPage-onEachSide, this.CurrentPage+onEachSide)
}

func (this *TLengthAwarePaginator) getStart() map[int]string {
	return this.getUrlRange(1, 2)
}

func (this *TLengthAwarePaginator) getFinish() map[int]string {
	return this.getUrlRange(this.LastPage-1, this.LastPage)
}

func (this *TLengthAwarePaginator) Elements() []map[int]string {
	if this.LastPage < (this.onEachSide*2)+6 {
		fmt.Println(this.getSmallSlider())
		return this.getSmallSlider()
	}

	return this.getUrlSlider(this.onEachSide)
}
