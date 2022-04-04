package parser

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
)

type ClothingParameter int

const (
	PageUrlAtCategoryPage ClothingParameter = iota
	ImageUrl
	Price
	Color
)

type ClothingType int

const (
	OuterwearType ClothingType = iota
	ShirtType
	PantsType
)

var regexpByParameter = initRegexp()

func initRegexp() map[ClothingParameter]*regexp.Regexp {
	result := map[ClothingParameter]*regexp.Regexp{}
	result[PageUrlAtCategoryPage], _ = regexp.Compile(`class="x-product-card__card"><a href="([^"]*)"`)
	result[ImageUrl], _ = regexp.Compile(`property="og:image" content="([^"]*)"`)
	result[Price], _ = regexp.Compile(`"price_amount":([^,]*),`)
	result[Color], _ = regexp.Compile(`"color_family":"([\p{L}]+)"`)
	return result
}

func saveAsJson(data ClothingItem, path string) {
	file, _ := json.MarshalIndent(data, "", " ")
	_ = ioutil.WriteFile(path, file, 0644)
}

func extractCategoryLinks(categoryUrl string) []string {
	links := map[string]struct{}{}
	for i := 1; i <= 2; i++ {
		currentPageLinks := extractCategoryPageLinks(categoryUrl + "?&page=" + strconv.Itoa(i))
		linksQtyBefore := len(links)
		for j := range currentPageLinks {
			links[currentPageLinks[j]] = struct{}{}
		}
		if len(links) == linksQtyBefore {
			break
		}
		fmt.Printf("page: %d\t\t links: %d\n", i, len(links)-linksQtyBefore)
	}
	var linksList []string
	for i := range links {
		linksList = append(linksList, i)
	}
	return linksList
}

func extractCategoryPageLinks(url string) []string {
	var links []string
	pageText, _ := loadPageText(url)
	l := regexpByParameter[PageUrlAtCategoryPage].FindAllStringSubmatch(pageText, 4)
	for i := range l {
		links = append(links, "https://www.lamoda.by"+l[i][1])
	}
	return links
}

func extractParameter(pageText string, param ClothingParameter) (string, error) {
	result := regexpByParameter[param].FindStringSubmatch(pageText)
	if len(result) == 2 {
		return result[1], nil
	}
	return "", errors.New("color not found")
}

func loadPageText(url string) (string, error) {
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}
	if response.StatusCode != 200 {
		fmt.Println(response.Status)
	}
	defer response.Body.Close()
	b, _ := io.ReadAll(response.Body)
	return string(b), nil
}

func ParseCategory(url string, categoryType ClothingType) {
	links := extractCategoryLinks(url)
	for i := range links {
		pageText, err := loadPageText(links[i])
		if err != nil {
			continue
		}
		item := ClothingItem{}
		item.Color, _ = extractParameter(pageText, Color)
		if err != nil {
			continue
		}
		item.Price, err = extractParameter(pageText, Price)
		if err != nil {
			continue
		}
		item.ImageUrl, err = extractParameter(pageText, ImageUrl)
		if err != nil {
			continue
		}
		item.PageUrl = links[i]
		fmt.Println(item)
	}
}
