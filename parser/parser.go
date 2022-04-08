package parser

import (
	"basis-parser/db"
	"errors"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type ClothingParameter int

const (
	PageUrlAtCategoryPage ClothingParameter = iota
	ImageUrl
	Price
	Description
	Brand
	Color
	FitTypeShirts
	FitTypePants
	Composition
	Pattern
	Season
	LegOpeningCm
	SleeveLengthCm
	LengthCm
	CollarOrCutout
	InsulationComposition
	HoodType
)

var regexpByParameter = initRegexp()

func initRegexp() map[ClothingParameter]*regexp.Regexp {
	result := map[ClothingParameter]*regexp.Regexp{}
	result[PageUrlAtCategoryPage], _ = regexp.Compile(`class="x-product-card__card"><a href="([^"]*)"`)
	result[ImageUrl], _ = regexp.Compile(`property="og:image" content="([^"]*)"`)
	result[Price], _ = regexp.Compile(`"price_amount":([^,]*),`)
	result[Color], _ = regexp.Compile(`"color_family":"([\p{L}]+)"`)
	result[FitTypeShirts], _ = regexp.Compile(`"type_of_knitwear","text":"([^"]*)"`)
	result[FitTypePants], _ = regexp.Compile(`"type_of_knitwear","text":"([^"]*)"`)
	result[Description], _ = regexp.Compile(`data-name="([^"]*)"`)
	result[Brand], _ = regexp.Compile(`data-brand="([^"]*)"`)
	result[Season], _ = regexp.Compile(`"season_wear","text":"([^"]*)"`)
	result[Pattern], _ = regexp.Compile(`"print","text":"([^"]*)"`)
	result[Composition], _ = regexp.Compile(`"material_filling","text":"([^"]*)"`)
	result[InsulationComposition], _ = regexp.Compile(`"material_filler","text":"([^"]*)"`)
	result[HoodType], _ = regexp.Compile(`"hood_features","text":"([^"]*)"`)
	result[LegOpeningCm], _ = regexp.Compile(`"bottom_width","text":"([^"]*)"`)
	result[SleeveLengthCm], _ = regexp.Compile(`"sleeve_length","text":"([^"]*)"`)
	result[LengthCm], _ = regexp.Compile(`"length","text":"([^"]*)"`)
	return result
}

func ParseOuterwearSubcategory(url string, baselineItem db.Outerwear, output func(db.Outerwear)) {
	urls := extractSubcategoryUrls(url)
	for i := range urls {
		pageText, err := loadPageText(urls[i])
		if err != nil {
			continue
		}
		item, err := extractOuterwear(pageText)
		if err != nil {
			continue
		}
		item.PageUrl = urls[i]
		item.Subcategory = baselineItem.Subcategory
		output(item)
	}
}

func ParseShirtSubcategory(url string, baselineItem db.Shirt, output func(db.Shirt)) {
	urls := extractSubcategoryUrls(url)
	for i := range urls {
		pageText, err := loadPageText(urls[i])
		if err != nil {
			continue
		}
		item, err := extractShirt(pageText)
		if err != nil {
			continue
		}
		item.PageUrl = urls[i]
		item.Subcategory = baselineItem.Subcategory
		item.CollarOrCutout = baselineItem.CollarOrCutout
		output(item)
	}
}

func ParsePantsSubcategory(url string, baselineItem db.Pants, output func(db.Pants)) {
	urls := extractSubcategoryUrls(url)
	for i := range urls {
		pageText, err := loadPageText(urls[i])
		if err != nil {
			continue
		}
		item, err := extractPants(pageText)
		if err != nil {
			continue
		}
		item.PageUrl = urls[i]
		item.Subcategory = baselineItem.Subcategory
		output(item)
	}
}

func extractCommonParameters(pageText string) (db.ClothingItem, error) {
	result := db.ClothingItem{}
	var err error
	result.Color, err = extractParameter(pageText, Color)
	if err != nil {
		return db.ClothingItem{}, err
	}
	priceStr, err := extractParameter(pageText, Price)
	if err != nil {
		return db.ClothingItem{}, err
	}
	result.Price, err = strconv.Atoi(priceStr)
	if err != nil {
		return db.ClothingItem{}, err
	}
	result.ImageUrl, err = extractParameter(pageText, ImageUrl)
	if err != nil {
		return db.ClothingItem{}, err
	}
	result.Brand, err = extractParameter(pageText, Brand)
	if err != nil {
		return db.ClothingItem{}, err
	}
	result.Description, err = extractParameter(pageText, Description)
	if err != nil {
		return db.ClothingItem{}, err
	}
	result.Pattern, err = extractParameter(pageText, Pattern)
	if err != nil {
		return db.ClothingItem{}, err
	}
	result.Season, err = extractParameter(pageText, Season)
	if err != nil {
		return db.ClothingItem{}, err
	}
	return result, nil
}

func extractOuterwear(pageText string) (db.Outerwear, error) {
	commonParams, err := extractCommonParameters(pageText)
	if err != nil {
		return db.Outerwear{}, err
	}
	result := db.Outerwear{ClothingItem: commonParams}
	result.HoodType, err = extractParameter(pageText, HoodType)
	length, _ := extractParameter(pageText, LengthCm)
	result.LengthCm, _ = strconv.Atoi(length)
	sleeveLength, _ := extractParameter(pageText, SleeveLengthCm)
	result.SleeveLengthCm, _ = strconv.Atoi(sleeveLength)
	result.InsulationComposition, _ = extractParameter(pageText, InsulationComposition)
	return result, nil
}

func extractShirt(pageText string) (db.Shirt, error) {
	commonParams, err := extractCommonParameters(pageText)
	if err != nil {
		return db.Shirt{}, err
	}
	result := db.Shirt{ClothingItem: commonParams}
	result.FitType, _ = extractParameter(pageText, FitTypeShirts)
	length, _ := extractParameter(pageText, LengthCm)
	result.LengthCm, _ = strconv.Atoi(length)
	sleeveLength, _ := extractParameter(pageText, SleeveLengthCm)
	result.SleeveLengthCm, _ = strconv.Atoi(sleeveLength)
	return result, nil
}

func extractPants(pageText string) (db.Pants, error) {
	commonParams, err := extractCommonParameters(pageText)
	if err != nil {
		return db.Pants{}, err
	}
	result := db.Pants{ClothingItem: commonParams}
	result.FitType, _ = extractParameter(pageText, FitTypePants)
	legOpening, _ := extractParameter(pageText, LegOpeningCm)
	result.LegOpeningCm, _ = strconv.Atoi(legOpening)
	return result, nil
}

func extractSubcategoryUrls(categoryUrl string) []string {
	urls := map[string]struct{}{}
	for i := 1; i <= 2; i++ {
		currentPageUrls := extractCategoryPageUrls(categoryUrl + "?&page=" + strconv.Itoa(i))
		urlQtyBefore := len(urls)
		for j := range currentPageUrls {
			urls[currentPageUrls[j]] = struct{}{}
		}
		if len(urls) == urlQtyBefore {
			break
		}
	}
	var urlList []string
	for i := range urls {
		urlList = append(urlList, i)
	}
	return urlList
}

func extractCategoryPageUrls(url string) []string {
	var urls []string
	pageText, _ := loadPageText(url)
	matches := regexpByParameter[PageUrlAtCategoryPage].FindAllStringSubmatch(pageText, 4)
	for i := range matches {
		urls = append(urls, "https://www.lamoda.by"+matches[i][1])
	}
	return urls
}

func extractParameter(pageText string, param ClothingParameter) (string, error) {
	result := regexpByParameter[param].FindStringSubmatch(pageText)
	if len(result) == 2 {
		return result[1], nil
	}
	return "", errors.New("parameter not found")
}

func loadPageText(url string) (string, error) {
	client := http.Client{}
	response, err := client.Get(url)
	if err != nil {
		return "", err
	}
	log.Printf("Loaded with response: %s at: %s\n", response.Status, url)
	if response.StatusCode != 200 {
		return "", err
	}
	b, _ := io.ReadAll(response.Body)
	response.Body.Close()
	return string(b), nil
}
