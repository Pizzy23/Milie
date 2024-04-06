package service

import (
	"encoding/json"
	"fmt"
	"millieMind/db"
	inter "millieMind/internal/questions/interface"
	"millieMind/internal/questions/storage"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func InputQuestionsDB(c *gin.Context, data inter.Resp) {

	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}

	var response inter.Resp
	err = json.Unmarshal(jsonData, &response)
	if err != nil {
		panic(err)
	}

	for category, questions := range response {

		newCategory, minAge, maxAge, err := extractAgeInfo(category)
		if err != nil {
			panic("Error not is possible extract.")

		}
		newCategory = strings.Replace(newCategory, "–", "", 1)
		newCategory = strings.TrimSpace(newCategory)

		ageRange := strconv.Itoa(minAge) + " - " + strconv.Itoa(maxAge)
		for _, q := range questions {
			newQuestion := db.BaseQuestions{
				FormsName:  "Portage",
				Categories: newCategory,
				Question:   q.Desc,
				Age:        ageRange,
			}
			storage.AddQuestion(db.Repo, newQuestion)
		}
	}
	c.Set("Response", "Questions add in database :)")
	c.Status(http.StatusOK)
}

func extractAgeInfo(category string) (string, int, int, error) {
	re := regexp.MustCompile(`\s*–\s*(\d+)\s*(?:a\s*(\d+)\s+anos?|ano)$`)

	matches := re.FindStringSubmatch(category)
	if len(matches) < 2 || len(matches) > 3 {
		return category, 0, 0, fmt.Errorf("formato de idade inválido: %s", category)
	}

	minAge, err := strconv.Atoi(matches[1])
	if err != nil {
		return "", 0, 0, err
	}

	var maxAge int
	if len(matches) == 3 {
		maxAge, err = strconv.Atoi(matches[2])
		if err != nil {
			return "", 0, 0, err
		}
	} else {
		maxAge = minAge
	}

	newCategory := strings.TrimSpace(re.ReplaceAllString(category, ""))
	return newCategory, minAge, maxAge, nil
}
