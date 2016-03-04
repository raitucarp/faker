package faker

import (
	"math/rand"
	"reflect"
	"strconv"
)

var imageCategories = []string{
	"abstract", "animals", "business", "cats", "city",
	"food", "nightlife", "fashion",
	"people", "nature", "sports",
	"technics", "transport",
}

type Image struct {
	Category string
	Width    int
	Height   int
	Url      string
}

func (image *Image) Imageʹ(params ...interface{}) string {
	category := sample(imageCategories)

	switch category {
	case "abstract":
		image.Abstractʹ(params...)
	case "animals":
		image.Animalsʹ(params...)
	case "business":
		image.Businessʹ(params...)
	case "cats":
		image.Catsʹ(params...)
	case "city":
		image.Cityʹ(params...)
	case "food":
		image.Foodʹ(params...)
	case "nightlife":
		image.Nightlifeʹ(params...)
	case "fashion":
		image.Fashionʹ(params...)
	case "people":
		image.Peopleʹ(params...)
	case "nature":
		image.Natureʹ(params...)
	case "sports":
		image.Sportsʹ(params...)
	case "technics":
		image.Technicsʹ(params...)
	case "transport":
		image.Transportʹ(params...)
	}

	return image.Url
}

func (image *Image) imageURL(category string, params ...interface{}) string {
	randomize := false
	image.Width = 640
	image.Height = 480
	image.Category = ""

	types := kindOf(params...)

	if len(types) >= 1 && types[0] == reflect.Int {
		image.Width = params[0].(int)
	}

	if len(types) >= 2 && types[1] == reflect.Int {
		image.Height = params[1].(int)
	}

	if len(types) >= 3 && types[2] == reflect.Bool {
		randomize = params[2].(bool)
	}

	image.Url = "http://lorempixel.com/" +
		strconv.Itoa(image.Width) + "/" +
		strconv.Itoa(image.Height)

	if image.Category != "" {
		image.Url += "/" + image.Category
	}

	if randomize {
		image.Url += "?" + strconv.Itoa(rand.Int())
	}

	return image.Url
}

func (image *Image) Avatarʹ(params ...interface{}) string {

	internet := Internet{}
	return internet.Avatarʹ()
}

func (image *Image) Abstractʹ(params ...interface{}) string {
	image.imageURL("abstract", params...)
	return image.Url
}

func (image *Image) Animalsʹ(params ...interface{}) string {
	image.imageURL("animals", params...)
	return image.Url
}

func (image *Image) Businessʹ(params ...interface{}) string {
	image.imageURL("business", params...)
	return image.Url
}

func (image *Image) Catsʹ(params ...interface{}) string {
	image.imageURL("cats", params...)
	return image.Url
}

func (image *Image) Cityʹ(params ...interface{}) string {
	image.imageURL("city", params...)
	return image.Url
}

func (image *Image) Foodʹ(params ...interface{}) string {
	image.imageURL("food", params...)
	return image.Url
}
func (image *Image) Nightlifeʹ(params ...interface{}) string {
	image.imageURL("nightlife", params...)
	return image.Url
}
func (image *Image) Fashionʹ(params ...interface{}) string {
	image.imageURL("fashion", params...)
	return image.Url
}
func (image *Image) Peopleʹ(params ...interface{}) string {
	image.imageURL("people", params...)
	return image.Url
}
func (image *Image) Natureʹ(params ...interface{}) string {
	image.imageURL("nature", params...)
	return image.Url
}
func (image *Image) Sportsʹ(params ...interface{}) string {
	image.imageURL("sports", params...)
	return image.Url
}
func (image *Image) Technicsʹ(params ...interface{}) string {
	image.imageURL("technics", params...)
	return image.Url
}
func (image *Image) Transportʹ(params ...interface{}) string {
	image.imageURL("transport", params...)
	return image.Url
}

// Fake Generate random data for all field
func (image *Image) Fake() {
	image.Imageʹ()
}

// ToJSON Encode name and its fields to JSON.
func (image *Image) ToJSON() (string, error) {
	return ToJSON(image)
}

// ToJSON Encode name and its fields to JSON.
func (image *Image) ToXML() (string, error) {
	return ToJSON(image)
}
