package data

//go:generate msgp -tests=false

type ID uint32

type IDList []ID

//msgp:tuple AuthorAffiliationPair
type AuthorAffiliationPair struct {
	AuthorIndex      int
	AffiliationIndex int
	CitationShare    float64 `msg:"-"`
}

//msgp:tuple PaperInfo
type PaperInfo struct {
	Year          int
	VenueIndex    int
	Citation      int
	NumOfAuthors  int `msg:"-"`
	RefList       []int
	AuthorAffList []AuthorAffiliationPair
	//AffList []int `msg:"-"`
}

//type PaperInfoListType []PaperInfo
//msgp:tuple RawData
type RawData struct {
	VenueList       []ID
	AffiliationList []ID
	CountryList     []ID
	AuthorList      []ID
	PaperList       []ID
	PaperInfoList   []PaperInfo
}

//type PaperAuthorListType [][]AuthorAffiliationPair
