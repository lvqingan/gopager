package gopager

import (
	"fmt"
	"math"
	"net/http"
)

type TPaginator struct {
	Items       interface{}
	Total       int
	PerPage     int
	CurrentPage int
	Path        string
	LastPage    int
	Query       map[string][]string
	PageName    string
}

func NewPaginator(items interface{}, total int, perPage int, currentPage int, options map[string]string) *TPaginator {
	paginator := &TPaginator{
		Items:       items,
		Total:       total,
		PerPage:     perPage,
		CurrentPage: currentPage,
	}

	paginator.LastPage = int(math.Max(math.Ceil(float64(total/perPage)), 1))

	if options != nil {
		for optionKey, optionValue := range options {
			if optionKey == "path" {
				paginator.Path = optionValue
			}

			if optionKey == "pageName" {
				paginator.PageName = optionValue
			}
		}
	}

	if paginator.Path == "" {
		paginator.Path = "/"
	}

	if paginator.PageName == "" {
		paginator.PageName = "page"
	}

	return paginator
}

func (this *TPaginator) Appends(query map[string][]string) {
	this.Query = query
}

func (this *TPaginator) url(page int) string {
	if page <= 0 {
		page = 1
	}

	sPage := fmt.Sprintf("%v", page)

	parameters := map[string][]string{
		this.PageName: []string{sPage},
	}

	if len(this.Query) > 0 {
		for key, value := range this.Query {
			if key == "page" {
				continue
			}

			parameters[key] = value
		}
	}

	return this.buildQuery(parameters)
}

func (this *TPaginator) NextPageUrl() string {
	if this.LastPage > this.CurrentPage {
		return this.url(this.CurrentPage + 1)
	} else {
		return ""
	}
}

func (this *TPaginator) PreviousPageUrl() string {
	if this.CurrentPage > 1 {
		return this.url(this.CurrentPage - 1)
	} else {
		return ""
	}
}

func (this *TPaginator) buildQuery(parameters map[string][]string) string {
	req, _ := http.NewRequest("GET", this.Path, nil)

	q := req.URL.Query()
	for key, value := range parameters {
		if len(value) > 1 {
			for _, val := range value {
				q.Add(key + "[]", val)
			}
		} else {
			q.Add(key, value[0])
		}
	}

	req.URL.RawQuery = q.Encode()

	return req.URL.String()
}

func (this *TPaginator) OnFirstPage() bool {
	return this.CurrentPage <= 1
}

func (this *TPaginator) HasMorePages() bool {
	return this.CurrentPage < this.LastPage
}

func (this *TPaginator) Elements() map[int]string {
	elements := make(map[int]string)

	for i := 1; i <= this.LastPage; i++ {
		elements[i] = this.url(i)
	}

	return elements
}

func (this *TPaginator) getUrlRange(start int, end int) map[int]string {
	urlRange := make(map[int]string)

	for i := start; i <= end; i++ {
		urlRange[i] = this.url(i)
	}

	return urlRange
}

func (this *TPaginator) HasPage() bool {
	return this.LastPage > 1
}
