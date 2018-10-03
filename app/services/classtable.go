package services

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/arthurc0102/ntub-class-table-api/config"
)

var classMapKey = []string{"name", "teacher", "room"}

// GetPersonalClassTable send request to get class table and return doc
func GetPersonalClassTable(studentID string, today int) (*goquery.Document, error) {
	client := &http.Client{}
	data := url.Values{"StdNo": {studentID}, "today": {fmt.Sprint(today)}}

	req, err := http.NewRequest(http.MethodPost, config.ClassTableURL, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("X-Requested-With", "com.hanglong.NTUBStdApp")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return doc, nil
}

// PersonalClassTableByDay service
func PersonalClassTableByDay(doc *goquery.Document) (classTableOfDay []map[string]string) {
	doc.Find("td.Stdtd001").Each(func(_ int, s *goquery.Selection) {
		name := s.Find("a").First().Text()

		if name == "" {
			classTableOfDay = append(classTableOfDay, nil)
			return
		}

		classInfo := []string{}
		classInfoHTML, _ := s.Html()
		classInfo = strings.Split(classInfoHTML, "<br/>")
		classInfo[0] = name

		class := map[string]string{}

		for i := 0; i < 3; i++ {
			class[classMapKey[i]] = classInfo[i]
		}

		class["room"] = strings.Split(class["room"], "<")[0] // 修復一堂課多教師時教室顯示異常
		classTableOfDay = append(classTableOfDay, class)
	})

	return
}

// PersonalClassTableTime return time off class
func PersonalClassTableTime(doc *goquery.Document) []map[string]string {
	var timeList []map[string]string

	doc.Find("th.Stdth003").Each(func(_ int, s *goquery.Selection) {
		timeInfoHTML, _ := s.Html()
		timeInfo := strings.Split(timeInfoHTML, "<br/>")

		if l := len(timeInfo); l < 3 {
			for i := 0; i < 3-l; i++ {
				timeInfo = append(timeInfo, "")
			}
		}

		timeList = append(timeList, map[string]string{
			"class_no": timeInfo[0],
			"start_at": timeInfo[1],
			"end_at":   timeInfo[2],
		})
	})

	return timeList
}

// PersonalClassTable service
func PersonalClassTable(studentID string) ([7][]map[string]string, []map[string]string, []error) {
	var wg sync.WaitGroup
	var classTable [7][]map[string]string
	var classTime []map[string]string
	var errorList []error

	wg.Add(7)

	for i := 1; i < 8; i++ {
		go func(today int) {
			doc, err := GetPersonalClassTable(studentID, today)
			if err != nil {
				errorList = append(errorList, err)
				return
			}

			classTable[today-1] = PersonalClassTableByDay(doc)

			if today == 1 {
				classTime = PersonalClassTableTime(doc)
			}

			wg.Done()
		}(i)
	}

	wg.Wait()
	return classTable, classTime, errorList
}
