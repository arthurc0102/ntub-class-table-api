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

// PersonalClassTableByDay service
func PersonalClassTableByDay(studentID string, today int) ([]map[string]string, error) {
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

	var classTableOfDay []map[string]string

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

		classTableOfDay = append(classTableOfDay, class)
	})

	return classTableOfDay, nil
}

// PersonalClassTable service
func PersonalClassTable(studentID string) (classTable [7][]map[string]string, errorList []error) {
	var wg sync.WaitGroup

	wg.Add(7)

	for i := 1; i < 8; i++ {
		go func(today int) {
			classList, err := PersonalClassTableByDay(studentID, today)

			if err != nil {
				errorList = append(errorList, err)
			}

			classTable[today-1] = classList
			wg.Done()
		}(i)
	}

	wg.Wait()
	return
}
