package usecase

import "strconv"

type PageInfo struct {
	Length          int    `json:"length"`
	HasNextPage     bool   `json:"has_next_page"`
	HasPreviousPage bool   `json:"has_previous_page"`
	StartCursor     string `json:"start_cursor"`
	EndCursor       string `json:"end_cursor"`
}

func NewPageInfo(limit, lenght, cursor, endCursor, startCursor int) PageInfo {
	return PageInfo{
		HasNextPage:     limit < lenght,
		HasPreviousPage: 0 < cursor,
		StartCursor:     strconv.Itoa(startCursor),
		EndCursor:       strconv.Itoa(endCursor),
	}
}

func (p PageInfo) setStartCursor(startCursor int) string {
	if p.Length > 0 {
		return strconv.Itoa(startCursor)
	}
	return ""
}

func (p PageInfo) setEndCursor(endCursor int) string {
	if p.HasNextPage {
		return strconv.Itoa(endCursor)
	}
	return ""
}
