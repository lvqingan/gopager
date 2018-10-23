# GoPager

A Go (golang) port of the Laravel pagination

## Installation

As a library

```bash
go get github.com/lvqingan/gopager
```

Use govendor

```bash
govendor fetch github.com/lvqingan/gopager
```

## Usage

### Paginator

Create paginator instance in handler and pass it to template
```golang
result := make([]int, 100)
paginator := gopager.NewPaginator(result, 100, 5, 1, nil)
// ...
tpl.ExecuteTemplate(w, "list", map[string]interface{}{
   "paginator": paginator,
})
```

And the pagination bar could be a list of links with page number
```html
<nav class="pagination" role="navigation">
    {{if not .paginator.OnFirstPage}}
        <a class="pagination-previous" href="{{.paginator.PreviousPageUrl}}">Previous</a>
    {{end}}
    {{if .paginator.HasMorePages}}
        <a class="pagination-next" href="{{.paginator.NextPageUrl}}">Next</a>
    {{end}}

    {{if gt .paginator.LastPage 1}}
    <ul class="pagination-list">
        {{range $page, $url := .paginator.Elements }}
        <li>
            {{ if eq $page $.paginator.CurrentPage}}
                <a class="pagination-link is-current" href="{{$url}}">{{$page}}</a>
            {{else}}
                <a class="pagination-link" href="{{$url}}">{{$page}}</a>
            {{end}}
        </li>
        {{end}}
    </ul>
    {{end}}
</nav>
```
Or you can put them inside a select tag

### LengthAwarePaginator

LengthAwarePaginator can give you a fix length pagination bar no matter how many pages there will be

```golang
result := make([]int, 100)
paginator := gopager.NewLengthAwarePaginator(result, 100, 5, 1, nil)
// ...
tpl.ExecuteTemplate(w, "list", map[string]interface{}{
   "paginator": paginator,
})
```

The template should be
```html
<nav class="pagination" role="navigation">
{{if not .paginator.OnFirstPage}}
    <a class="pagination-previous" href="{{.paginator.PreviousPageUrl}}">Previous</a>
{{end}}
{{if .paginator.HasMorePages}}
    <a class="pagination-next" href="{{.paginator.NextPageUrl}}">Next</a>
{{end}}

{{if gt .paginator.LastPage 1}}
    <ul class="pagination-list">
       {{range $element := paginator.Elements}}
           {{if $element.Show}}
               {{if $element.IsDots }}
                   <li class="pagination-link is-disabled"><span>...</span></li>
               {{else}}
                   {{range $page, $url := $element.Items}}
                       <li>
                       {{ if eq $page $.paginator.CurrentPage}}
                           <a class="pagination-link is-current" href="{{$url}}">{{$page}}</a>
                       {{else}}
                           <a class="pagination-link" href="{{$url}}">{{$page}}</a>
                       {{end}}
                       </li>
                   {{end}}
               {{end}}
           {{end}}
       {{end}}
    </ul>
{{end}}
</nav>
```

### Change url path

The default path is `/` and you can change it by yourself

```golang
paginator := gopager.NewLengthAwarePaginator(result, 100, 5, 1, map[string]string{
    "path": "/foo/bar",
})
```

### Change LengthAwarePaginator length

```golang
paginator := gopager.NewLengthAwarePaginator(result, 100, 5, 1, map[string]string{
    "onEachSide": "4",
})
```

Then, the pagination bar will be longer than before

### Change page name

The url parameter name of page is `page` by default. For example `http://my.site/foo/bar?page=1`

If you wanna change the parameter name, you could set a `pageName`

```golang
paginator := gopager.NewLengthAwarePaginator(result, 100, 5, 1, map[string]string{
    "pageName": "p",
})
```
### Appends query string

```golang
paginator := gopager.NewPaginator(make([]int, 20), 20, 10, 1, nil)

paginator.Appends(map[string][]string{
    "keyword": {"andy"},
    "names":   {"tom", "jack"},
})
```

The path will be `/?keyword=andy&names[]=tom&names[]=jack&page=1`

### Get map data for json

```golang
mapData := paginator.GetStringMap()
```

It can be convert to json:

```json
{
   "total": 50,
   "per_page": 15,
   "current_page": 1,
   "last_page": 4,
   "first_page_url": "http://localhost?page=1",
   "last_page_url": "http://localhost?page=4",
   "next_page_url": "http://localhost?page=2",
   "prev_page_url": "",
   "path": "http://localhost",
   "from": 1,
   "to": 15,
   "data":[
        {
            // Result Object
        },
        {
            // Result Object
        }
   ]
}
```