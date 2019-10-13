package data

import (
	"github.com/tinylib/msgp/msgp"
	"io/ioutil"
	"log"
	"sync"
)

var (
	VenueList         IDList
	AffiliationList   IDList
	CountryList       IDList
	AuthorList        IDList
	PaperList         IDList
	PaperInfoList     []PaperInfo
	NumOfVenues       int
	NumOfAffiliations int
	NumOfAuthors      int
	NumOfPapers       int
)

var (
	VenueIDToIndex       map[ID]int
	AffiliationIDToIndex map[ID]int
	AuthorIDToIndex      map[ID]int
	VenuePaperMap        [][]int
	AffiliationPaperList [][]int
	//AuthorPaperList [][]int
)

func LoadDataFromFiles() {
	log.Println("Data pre-processing Start!")
	var raw RawData
	LoadObjectFromFile(&raw, "data/static/data.bin")
	log.Println("Raw data was loaded successfully!")

	VenueList = raw.VenueList
	AffiliationList = raw.AffiliationList
	CountryList = raw.CountryList
	AuthorList = raw.AuthorList
	PaperList = raw.PaperList
	PaperInfoList = raw.PaperInfoList

	NumOfVenues = len(VenueList)
	NumOfAffiliations = len(AffiliationList)
	NumOfAuthors = len(AuthorList)
	NumOfPapers = len(PaperList)

	wg := sync.WaitGroup{}
	wg.Add(4)

	go func() {
		VenueIDToIndex = make(map[ID]int)
		for venueIndex, venueID := range VenueList {
			VenueIDToIndex[venueID] = venueIndex
		}
		AffiliationIDToIndex = make(map[ID]int)
		for affIndex, affID := range AffiliationList {
			AffiliationIDToIndex[affID] = affIndex
		}
		wg.Done()
	}()

	go func() {
		AuthorIDToIndex = make(map[ID]int)
		for authorIndex, authorID := range AuthorList {
			AuthorIDToIndex[authorID] = authorIndex
		}
		wg.Done()
	}()

	go func() {
		VenuePaperMap = make([][]int, NumOfVenues)
		for i := range VenuePaperMap {
			VenuePaperMap[i] = make([]int, 0)
		}
		for i, info := range PaperInfoList {
			venueIndex := info.VenueIndex
			if venueIndex != 0 {
				VenuePaperMap[venueIndex] = append(VenuePaperMap[venueIndex], i)
			}
		}
		wg.Done()
	}()

	//log.Println("All map was generated successfully!")

	go func() {
		AffiliationPaperList = make([][]int, NumOfAffiliations)
		for i := range AffiliationPaperList {
			AffiliationPaperList[i] = make([]int, 0)
		}
		for i, info := range PaperInfoList {
			//PaperInfoList[i].AffList = make([]int, 0)
			affMark := make(map[int]bool)
			authorMark := make(map[int]bool)
			for _, authorAffPair := range info.AuthorAffList {
				affIndex := authorAffPair.AffiliationIndex
				authorIndex := authorAffPair.AuthorIndex
				if !affMark[affIndex] {
					affMark[affIndex] = true
					//PaperInfoList[i].AffList = append(PaperInfoList[i].AffList, affIndex)
					AffiliationPaperList[affIndex] = append(AffiliationPaperList[affIndex], i)
				}
				authorMark[authorIndex] = true
			}
			PaperInfoList[i].NumOfAuthors = len(authorMark)
		}
		wg.Done()
	}()
	wg.Wait()

	for _, info := range PaperInfoList {
		if info.Citation == 0 {
			continue
		}
		authorAffList := info.AuthorAffList
		sharePerAuthor := float64(info.Citation) / float64(info.NumOfAuthors)
		nowAuthorIndex := authorAffList[0].AuthorIndex
		begin := 0
		cnt := 0
		for i := range authorAffList {
			authorIndex := authorAffList[i].AuthorIndex
			if authorIndex != nowAuthorIndex {
				share := sharePerAuthor / float64(cnt)
				for j := begin; j < i; j++ {
					authorAffList[j].CitationShare = share
				}
				begin = i
				cnt = 0
				nowAuthorIndex = authorIndex
			}
			cnt++
		}
		share := sharePerAuthor / float64(cnt)
		for j := begin; j < len(authorAffList); j++ {
			authorAffList[j].CitationShare = share
		}
	}
	log.Println("Data pre-processing Done!!")
}

func LoadObjectFromFile(o msgp.Unmarshaler, filename string) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = o.UnmarshalMsg(b)
	if err != nil {
		log.Fatal(err)
	}
}
