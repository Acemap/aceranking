package model

type ID uint32

func (id ID) GetIDType() string {
	if id == 0 {
		return "None"
	}
	if id <= 1000000000 {
		return "paper"
	}
	if id <= 2000000000 {
		return "author"
	}
	if id <= 2100000000 {
		return "field"
	}
	if id <= 2110000000 {
		return "affiliation"
	}
	if id <= 2120000000 {
		return "journal"
	}
	if id <= 2130000000 {
		return "conference"
	}
	if id <= 2140000000 {
		return "conference_instance"
	}
	if id <= 2141000000 {
		return "country"
	}
	return "undefined"
}

type IDList []ID
type IDSet map[ID]bool

func (l IDList) ToSet() IDSet {
	s := make(map[ID]bool)
	for _, v := range l {
		s[v] = true
	}
	return s
}

//===========================================================

type Response struct {
	Status  bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Filter struct {
	Area          string `form:"area"`
	StartYear     int    `form:"start_year"`
	EndYear       int    `form:"end_year"`
	FirstAuthor   int    `form:"first_author"`
	SpecificAffID ID     `form:"affiliation_id"`
	//MinAcademicAge int    `form:"min_academic_age"`
	//MaxAcademicAge int    `form:"max_academic_age"`
}

//===========================================================
// /v1/ace_ranking/filter

type FilterResp struct {
	AreaList       []string `json:"area_list"`
	MinYear        int      `json:"min_year"`
	MaxYear        int      `json:"max_year"`
	MinAcademicAge int      `json:"min_academic_age"`
	MaxAcademicAge int      `json:"max_academic_age"`
}

//===========================================================
// /v1/ace_ranking/venue

type VenueReq struct {
	Filter
	Type string `form:"type"`
}

type VenueResp []VenueContainer

type VenueContainer struct {
	ID       string      `json:"venue_id"`
	Name     string      `json:"name"`
	Count    int         `json:"-"`
	CountStr string      `json:"count"`
	List     interface{} `json:"list"`
}

type VenueInfo struct {
	VenueID  string `json:"venue_id"`
	Type     string `json:"type"`
	Name     string `json:"name"`
	Count    int    `json:"-"`
	CountStr string `json:"count"`
}

//===========================================================
// /v1/ace_ranking/author_list

type AuthorListReq struct {
	Filter
	Type     string `form:"type"`
	VenueIDs string `form:"venue_ids"`
	OrderBy  int    `form:"order_by"`
}

type AuthorListResp struct {
	AuthorList    []AuthorStatistics  `json:"author_list"`
	CountAnalysis []CountAnalysisType `json:"count_analysis"`
	FieldAnalysis []FieldAnalysisType `json:"field_analysis"`
	HasFigure     bool
}

type AuthorStatistics struct {
	AuthorID      ID                  `json:"author_id"`
	AuthorName    string              `json:"author_name"`
	CountAnalysis []CountAnalysisType `json:"count_analysis"`
	AffCount      map[ID]int          `json:"-"`
	AffList       []AuthorAffInfo     `json:"affiliation_list"`
	Statistics
}

type AuthorAffInfo struct {
	AffiliationID   ID     `json:"affiliation_id"`
	Abbreviation    string `json:"abbr"`
	AffiliationName string `json:"name"`
	Count           int    `json:"count"`
}

type CountAnalysisType struct {
	Year  int `json:"year"`
	Count int `json:"count"`
}

func NewCountAnalysis(startYear, endYear int) []CountAnalysisType {
	c := make([]CountAnalysisType, endYear-startYear+1)
	for i := range c {
		c[i].Year = startYear + i
	}
	return c
}

type FieldAnalysisType struct {
	FieldName string `json:"name"`
	Count     int    `json:"count"`
}

//===========================================================
// /v1/ace_ranking/author

type AuthorReq struct {
	Filter
	Type     string `form:"type"`
	VenueIDs string `form:"venue_ids"`
	AuthorID ID     `form:"author_id"`
}

type AuthorResp struct {
	PaperList     []AuthorPaperInfo   `json:"paper_list"`
	FieldAnalysis []FieldAnalysisType `json:"field_analysis"`
}

type AuthorPaperInfo struct {
	PaperID  ID     `json:"paper_id"`
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Venue    string `json:"venue"`
	Citation int    `json:"citation"`
}

//===========================================================
// /v1/ace_ranking/affiliation_list

type AffiliationListReq struct {
	Filter
	Type     string `form:"type"`
	VenueIDs string `form:"venue_ids"`
	OrderBy  int    `form:"order_by"`
}

type AffiliationListResp []AffiliationStatistics

type AffiliationStatistics struct {
	AffiliationID   ID     `json:"affiliation_id"`
	AffiliationName string `json:"name"`
	Abbreviation    string `json:"abbr"`
	Statistics
}
