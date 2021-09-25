package entity

import (
	"math"
	"net/url"
	"os"
	"strconv"

	"bitbucket.org/matchmove/platform-utilities-sgprefund-service/modules/constant"
)

type Pagination struct {
	TotalRecords   int `json:"total_records"`
	RecordsPerPage int `json:"records_per_page"`
	TotalPages     int `json:"total_pages"`
	CurrentPage    int `json:"current_page"`
	Links          `json:"links"`
	Records        interface{} `json:"records"`
}

type Links struct {
	Self  string `json:"self"`
	First string `json:"first"`
	Last  string `json:"last"`
}

const (
	QueryOffset  = "page"
	QueryLimit   = "records_per_page"
	DefaultLimit = 10
)

// FormatPagination function - creates and formats pagination object
func FormatPagination(url *url.URL, records interface{}, totalRecords int) (p Pagination) {
	var (
		currentPage    int
		recordsPerPage int
		err            error
	)

	currentPage, _ = strconv.Atoi(url.Query().Get(QueryOffset))
	if currentPage == 0 {
		currentPage = 1
	}

	if recordsPerPage, err = strconv.Atoi(url.Query().Get(QueryLimit)); err != nil || recordsPerPage == 0 {
		recordsPerPage = DefaultLimit
	}

	p.Records = records
	p.TotalRecords = totalRecords
	p.RecordsPerPage = recordsPerPage
	p.TotalPages = int(math.Ceil(float64(totalRecords) / float64(recordsPerPage)))
	p.CurrentPage = currentPage

	uri, _ := url.Parse(os.Getenv(constant.EnvAppDomain))
	path := "https://" + uri.Hostname() + url.Path + "?"
	query := url.Query()

	query.Set(QueryOffset, strconv.Itoa(currentPage))
	query.Set(QueryLimit, strconv.Itoa(recordsPerPage))
	self := query.Encode()

	query.Set(QueryOffset, "1")
	first := query.Encode()

	query.Set(QueryOffset, strconv.Itoa(p.TotalPages))
	last := query.Encode()

	p.Links = Links{
		Self:  path + self,
		First: path + first,
		Last:  path + last,
	}

	return
}
