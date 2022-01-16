package items

//CalendarItem calendar是一个列表，每一个列表项是由weekday信息和items信息组成的。一个calendar里由7个calender item
type CalendarItem struct {
	WeekDay WeekdayBgm       `json:"weekday"`
	Items   []*SubjectMedium `json:"items,omitempty"`
}
