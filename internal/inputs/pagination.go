package inputs

type PaginationParams struct {
	Cursor string `query:"pageCursor"`
	Limit  int    `query:"pageSize" default:"10" minimum:"10" maximum:"100"`
}
