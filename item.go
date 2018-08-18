package enhanced

import (
	"fmt"
	"strconv"
	"strings"

	dmm "github.com/usk81/go-dmm"
)

// Item represents a DMM product
type Item struct {
	AffiliateURL       string                     `json:"affiliateURL"`
	AffiliateURLMobile string                     `json:"affiliateURLsp"`
	BandaiInformation  dmm.BandaiInformation      `json:"bandaiinfo"`
	CategoryName       string                     `json:"category_name"`
	CdInformation      dmm.CdInformation          `json:"cdinfo"`
	Comment            string                     `json:"comment"`
	ContentID          string                     `json:"content_id"`
	Date               string                     `json:"date"`
	FloorCode          string                     `json:"floor_code"`
	FloorName          string                     `json:"floor_name"`
	ImageURL           dmm.ImageURL               `json:"imageURL"`
	ISBN               string                     `json:"isbn,omitempty"`
	ItemInfo           map[string][]ItemComponent `json:"iteminfo"`
	JANCode            string                     `json:"jancode,omitempty"`
	MakerProduct       string                     `json:"maker_product"`
	Prices             Prices                     `json:"prices"`
	ProductID          string                     `json:"product_id"`
	Review             Review                     `json:"review"`
	SampleImageURL     dmm.SampleImage            `json:"sampleImageURL,omitempty"`
	SampleMovieURL     dmm.SampleMovie            `mapstructure:"sampleMovieURL"`
	ServiceCode        string                     `json:"service_code"`
	ServiceName        string                     `json:"service_name"`
	Stock              string                     `json:"stock"`
	Title              string                     `json:"title"`
	URL                string                     `json:"URL"`
	URLMobile          string                     `json:"URLsp"`
	Volume             string                     `json:"volume"`
}

// ItemComponent is a product detail
type ItemComponent struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Ruby     string `json:"ruby,omitempty"`
	Category string `json:"category,omitempty"`
}

// Review is a review for a product
type Review struct {
	Count   int     `json:"count"`
	Average float64 `json:"average"`
}

// Prices is price information for a product
type Prices struct {
	Retail   int            `json:"retail"`
	Display  string         `json:"display"`
	Contents map[string]int `json:"contents"`
}

// ConvertItem converts from dmm.Review to Review
func ConvertItem(r dmm.Item) (result Item, err error) {
	var prices Prices
	var review Review
	var itemInfo map[string][]ItemComponent

	for k, v := range r.ItemInfo {
		var ics []ItemComponent
		if ics, err = convertItemComponents(v); err != nil {
			return
		}
		itemInfo[k] = ics
	}
	if review, err = convertReview(r.Review); err != nil {
		return
	}

	result = Item{
		AffiliateURL:       r.AffiliateURL,
		AffiliateURLMobile: r.AffiliateURLMobile,
		BandaiInformation:  r.BandaiInformation,
		CategoryName:       r.CategoryName,
		CdInformation:      r.CdInformation,
		Comment:            r.Comment,
		ContentID:          r.ContentID,
		Date:               r.Date,
		FloorCode:          r.FloorCode,
		FloorName:          r.FloorName,
		ImageURL:           r.ImageURL,
		ISBN:               r.ISBN,
		ItemInfo:           itemInfo,
		JANCode:            r.JANCode,
		MakerProduct:       r.MakerProduct,
		Prices:             prices,
		ProductID:          r.ProductID,
		Review:             review,
		SampleImageURL:     r.SampleImageURL,
		SampleMovieURL:     r.SampleMovieURL,
		ServiceCode:        r.ServiceCode,
		ServiceName:        r.ServiceName,
		Stock:              r.Stock,
		Title:              r.Title,
		URL:                r.URL,
		URLMobile:          r.URLMobile,
		Volume:             r.Volume,
	}
	return
}

func convertReview(r dmm.Review) (result Review, err error) {
	var a float64
	if r.Average != "" {
		if a, err = strconv.ParseFloat(r.Average, 64); err != nil {
			err = fmt.Errorf("Review.Average is not float; %s; %v", r.Average, err)
			return
		}
	}
	result = Review{
		Count:   r.Count,
		Average: a,
	}
	return
}

func convertItemComponents(r []dmm.ItemComponent) (result []ItemComponent, err error) {
	m := map[string]ItemComponent{}

	for _, v := range r {
		if v.ID.String() == "" {
			continue
		}
		ss := strings.Split(v.ID.String(), "_")
		id := ss[0]
		i := m[id]
		if len(ss) == 1 {
			i.ID = id
			i.Name = v.Name
		} else if len(ss) == 2 {
			switch ss[1] {
			case "ruby":
				i.ID = id
				i.Ruby = v.Name
			case "classify":
				i.ID = id
				i.Category = v.Name
			default:
				err = fmt.Errorf("item component stores unxepected data; %v", v)
				return
			}
		} else {
			err = fmt.Errorf("item component's id is unxepected format; %s", v.ID.String())
			return
		}
		m[id] = i
	}

	return
}

func convertPrices(r dmm.Prices) (result Prices, err error) {
	var retail int
	if r.ListPrice != "" {
		if retail, err = strconv.Atoi(r.ListPrice); err != nil {
			err = fmt.Errorf("listPrice is not numeric; %s; %v", r.ListPrice, err)
			return
		}
	}

	ss := map[string]int{}
	for i, d := range r.Deliveries.Delivery {
		var p int
		if d.Price != "" {
			if p, err = strconv.Atoi(r.ListPrice); err != nil {
				err = fmt.Errorf("Deliveries.Delivery (%d) is not numeric; %v; %v", i, d, err)
				return
			}
		}
		if d.Type == "" {
			err = fmt.Errorf("Deliveries.Delivery (%d)'s Nmae is empty; %v; %v", i, d, err)
			return
		}
		ss[d.Type] = p
	}
	result = Prices{
		Retail:   retail,
		Display:  r.Price,
		Contents: ss,
	}
	return
}
