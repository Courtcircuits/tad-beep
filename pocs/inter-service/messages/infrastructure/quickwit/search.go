package quickwit

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

type SearchResult struct {
	Hits              []map[string]any `json:"hits"`
	NumHits           int   `json:"num_hits"`
	ElapsedTimeMicros int   `json:"elapsed_time_micros"`
}

type SearchQuery struct {
	Query          string
	StartTimestamp string
	EndTimestamp   string
	StartOffset    int
	MaxHits        int
	SearchField    []string
	SnippetFields  []string
	SortBy         []string
	Format         string
	Aggs           map[string]any
}

func NewSearchQuery(query string) *SearchQuery {
	return &SearchQuery{
		Query: query,
	}
}

func (q *SearchQuery) WithStartTimestamp(startTimestamp string) *SearchQuery {
	q.StartTimestamp = startTimestamp
	return q
}

func (q *SearchQuery) WithEndTimestamp(endTimestamp string) *SearchQuery {
	q.EndTimestamp = endTimestamp
	return q
}

func (q *SearchQuery) WithStartOffset(startOffset int) *SearchQuery {
	q.StartOffset = startOffset
	return q
}

func (q *SearchQuery) WithMaxHits(maxHits int) *SearchQuery {
	q.MaxHits = maxHits
	return q
}

func (q *SearchQuery) WithSearchField(searchField []string) *SearchQuery {
	q.SearchField = searchField
	return q
}

func (q *SearchQuery) WithSnippetFields(snippetFields []string) *SearchQuery {
	q.SnippetFields = snippetFields
	return q
}

func (q *SearchQuery) WithSortBy(sortBy []string) *SearchQuery {
	q.SortBy = sortBy
	return q
}

func (q *SearchQuery) WithFormat(format string) *SearchQuery {
	q.Format = format
	return q
}

func (q *SearchQuery) WithAggs(aggs map[string]any) *SearchQuery {
	q.Aggs = aggs
	return q
}

func (q *SearchQuery) Build() (string, error) {
	query := fmt.Sprintf("?query=%s", url.QueryEscape(q.Query))
	if q.StartTimestamp != "" {
		query += fmt.Sprintf("&start_timestamp=%s", url.QueryEscape(q.StartTimestamp))
	}
	if q.EndTimestamp != "" {
		query += fmt.Sprintf("&end_timestamp=%s", url.QueryEscape(q.EndTimestamp))
	}
	if q.StartOffset != 0 {
		query += fmt.Sprintf("&start_offset=%d", q.StartOffset)
	}
	if q.MaxHits != 0 {
		query += fmt.Sprintf("&max_hits=%d", q.MaxHits)
	}
	if len(q.SearchField) > 0 {
		query += fmt.Sprintf("&search_field=%s", url.QueryEscape(FormatStringArray(q.SearchField)))
	}
	if len(q.SnippetFields) > 0 {
		query += fmt.Sprintf("&snippet_fields=%s", url.QueryEscape(FormatStringArray(q.SnippetFields)))
	}
	if len(q.SortBy) > 0 {
		query += fmt.Sprintf("&sort_by=%s", url.QueryEscape(FormatStringArray(q.SortBy)))
	}
	if q.Format != "" {
		query += fmt.Sprintf("&format=%s", url.QueryEscape(q.Format))
	}
	if len(q.Aggs) > 0 {
		aggs, err := FormatMap(q.Aggs)
		if err != nil {
			return "", err
		}
		query += fmt.Sprintf("&aggs=%s", url.QueryEscape(aggs))
	}
	return query, nil
}

func FormatStringArray(arr []string) string {
	var formattedArr []string
	for _, item := range arr {
		formattedArr = append(formattedArr, fmt.Sprintf("\"%s\"", item))
	}
	return strings.Join(formattedArr, ",")
}

func FormatMap(m map[string]any) (string, error) {
	res, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(res), nil
}

func Search(indexID string, endpoint string, query *SearchQuery) (*SearchResult, error) {
	search_query, err := query.Build()
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("%s/api/v1/%s/search%s", endpoint, indexID, search_query)
	fmt.Println(url)
	r, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	readbody, err := io.ReadAll(r.Body)
	fmt.Println(r.Status)
	if r.StatusCode != 200 {
		return nil, fmt.Errorf("%s", readbody)
	}
	defer r.Body.Close()
	var result_request *SearchResult = new(SearchResult)
	err = json.Unmarshal(readbody, result_request)
	if err != nil {
		return nil, err
	}
	return result_request, nil
}
