package enhanced

import (
	"fmt"
	"strconv"

	dmm "github.com/usk81/go-dmm"
)

// Actress represents a actress data
type Actress struct {
	ID         string       `json:"id"`
	Name       string       `json:"name"`
	Ruby       string       `json:"ruby"`
	Bust       int          `json:"bust"`
	Cup        string       `json:"cup,omitempty"`
	Waist      int          `json:"waist"`
	Hip        int          `json:"hip"`
	Height     int          `json:"height"`
	Birthday   string       `json:"birthday"`
	BloodType  string       `json:"blood_type"`
	Hobby      string       `json:"hobby"`
	Prefecture string       `json:"prefecture"`
	ImageURL   dmm.ImageURL `json:"imageURL,omitempty"`
	ListURL    dmm.ListURL  `json:"listURL"`
}

// ConvertActress converts from dmm.Actress to Actress
func ConvertActress(r dmm.Actress) (result Actress, err error) {
	var bust, waist, hip, height int
	if r.Bust != "" {
		if bust, err = strconv.Atoi(r.Bust); err != nil {
			err = fmt.Errorf("bust is not numeric; %s; %v", r.Bust, err)
			return
		}
	}
	if r.Waist != "" {
		if waist, err = strconv.Atoi(r.Waist); err != nil {
			err = fmt.Errorf("waist is not numeric; %s; %v", r.Waist, err)
			return
		}
	}
	if r.Hip != "" {
		if hip, err = strconv.Atoi(r.Hip); err != nil {
			err = fmt.Errorf("hip is not numeric; %s; %v", r.Hip, err)
			return
		}
	}
	if r.Height != "" {
		if height, err = strconv.Atoi(r.Height); err != nil {
			err = fmt.Errorf("height is not numeric; %s; %v", r.Height, err)
			return
		}
	}

	result = Actress{
		ID:         r.ID,
		Name:       r.Name,
		Ruby:       r.Ruby,
		Bust:       bust,
		Cup:        r.Cup,
		Waist:      waist,
		Hip:        hip,
		Height:     height,
		Birthday:   r.Birthday,
		BloodType:  r.BloodType,
		Hobby:      r.Hobby,
		Prefecture: r.Prefectures,
		ImageURL:   r.ImageURL,
		ListURL:    r.ListURL,
	}
	return
}
