package model

type IDName struct {
	ID   ID     `bson:"id"`
	Name string `bson:"name"`
}

type IDNameAbbr struct {
	IDName `bson:",inline"`
	Abbr   string `bson:"abbreviation"`
}

type IDsNames struct {
	IDs   IDList   `bson:"id"`
	Names []string `bson:"name"`
}

type IDsNamesAbbrs struct {
	IDsNames `bson:",inline"`
	Abbrs    []string `bson:"abbreviation"`
}

type PaperAuthor struct {
	IDsNames `bson:",inline"`
	AffIndex [][]int `bson:"aff_index"`
}

type PaperAnalysis struct {
	ReferenceCount  int `bson:"reference_count"`
	CitationCount   int `bson:"citation_count"`
	T2CitationCount int `bson:"t2_citation_count"`
}

// The lines commented out are the field in mongoDB but do not use in this project.
// Comments are for speeding up.
type Paper struct {
	PaperID          ID            `bson:"_id"`
	Title            string        `bson:"title"`
	Year             int           `bson:"year"`
	Author           PaperAuthor   `bson:"author"`
	FirstAuthor      IDName        `bson:"first_author"`
	Affiliation      IDsNamesAbbrs `bson:"affiliation"`
	FirstAffiliation IDNameAbbr    `bson:"first_affiliation"`
	ReferenceList    IDList        `bson:"reference_list"`
	Analysis         PaperAnalysis `bson:"analysis"`
	ConferenceSeries IDNameAbbr    `bson:"conference_series"`
	Journal          IDName        `bson:"journal"`
	//ConferenceInstance IDName        `bson:"conference_instance"`
	//Doi                string            `bson:"doi"`
	//DocType            int               `bson:"doc_type"`
	//Volume             int               `bson:"volume"`
	//Issue              int               `bson:"issue"`
	//FirstPage          int               `bson:"first_page"`
	//LastPage           int               `bson:"last_page"`
	//Abstract           string            `bson:"abstract"`
	//Field              IDsNames          `bson:"field"`
	//MachineReading     map[string]string `bson:"machine_reading"`
}

type VenuePaperCount struct {
	VenueID ID  `bson:"_id"`
	Count   int `bson:"count"`
}

type RefCount struct {
	ReferenceID ID  `bson:"_id"`
	Count       int `bson:"count"`
}
