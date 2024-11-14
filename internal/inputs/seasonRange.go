package inputs

type SeasonRangeParams struct {
	SeasonFrom int `query:"seasonFrom" default:"1800"`
	SeasonTo   int `query:"seasonTo" default:"2100"`
}

type SeasonRanges string

const (
	SeasonFromKey SeasonRanges = "seasonFrom"
	SeasonToKey   SeasonRanges = "seasonTo"
)
