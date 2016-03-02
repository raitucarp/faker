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

func (image *Image) Image_(params ...interface{}) string {
	category := sample(imageCategories)

	switch category {
	case "abstract":
		image.Abstract_(params...)
	case "animals":
		image.Animals_(params...)
	case "business":
		image.Business_(params...)
	case "cats":
		image.Cats_(params...)
	case "city":
		image.City_(params...)
	case "food":
		image.Food_(params...)
	case "nightlife":
		image.Nightlife_(params...)
	case "fashion":
		image.Fashion_(params...)
	case "people":
		image.People_(params...)
	case "nature":
		image.Nature_(params...)
	case "sports":
		image.Sports_(params...)
	case "technics":
		image.Technics_(params...)
	case "transport":
		image.Transport_(params...)
	}

	return image.Url
}

func (image *Image) ImageURL(category string, params ...interface{}) string {
	randomize := false
	image.Width = 640
	image.Height = 480
	image.Category = category

	types := typeof(params...)

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

func (image *Image) Avatar_(params ...interface{}) string {

	internet := Internet{}
	return internet.Avatar_()
}

func (image *Image) Abstract_(params ...interface{}) string {
	image.ImageURL("abstract", params...)
	return image.Url
}

func (image *Image) Animals_(params ...interface{}) string {
	image.ImageURL("animals", params...)
	return image.Url
}

func (image *Image) Business_(params ...interface{}) string {
	image.ImageURL("business", params...)
	return image.Url
}

func (image *Image) Cats_(params ...interface{}) string {
	image.ImageURL("cats", params...)
	return image.Url
}

func (image *Image) City_(params ...interface{}) string {
	image.ImageURL("city", params...)
	return image.Url
}

func (image *Image) Food_(params ...interface{}) string {
	image.ImageURL("food", params...)
	return image.Url
}
func (image *Image) Nightlife_(params ...interface{}) string {
	image.ImageURL("nightlife", params...)
	return image.Url
}
func (image *Image) Fashion_(params ...interface{}) string {
	image.ImageURL("fashion", params...)
	return image.Url
}
func (image *Image) People_(params ...interface{}) string {
	image.ImageURL("people", params...)
	return image.Url
}
func (image *Image) Nature_(params ...interface{}) string {
	image.ImageURL("nature", params...)
	return image.Url
}
func (image *Image) Sports_(params ...interface{}) string {
	image.ImageURL("sports", params...)
	return image.Url
}
func (image *Image) Technics_(params ...interface{}) string {
	image.ImageURL("technics", params...)
	return image.Url
}
func (image *Image) Transport_(params ...interface{}) string {
	image.ImageURL("transport", params...)
	return image.Url
}
