package main

import (
	"html/template"
	"net/http"
)

type Email struct {
	Sections []EmailSection
}

type EmailSection struct {
	URL         string
	IMG         string // size 580x435
	Title       string
	Description string

	// SecondIMG string // size 220x220
	// SecondURL string
	// SecondTitle string
	// SecondDescription string

	// ThirdIMG string // size 220x220
	// ThirdURL string
	// ThirdTitle string
	// ThirdDescription string

	// FourthIMG string // size 220x220
	// FourthURL string
	// FourthTitle string
	// FourthDescription string
}

func main() {
	tmpl := template.Must(template.ParseFiles("layout.html"))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := Email{
			Sections: []EmailSection{
				{
					IMG:         "https://gcziprusz.github.io/peekaboo.travel/img/eml/top.jpg",
					URL:         "https://gcziprusz.github.io/peekaboo.travel/greece-family-road-trip/",
					Title:       "Greece Family Road Trip",
					Description: "We booked an off-season trip to Greece in search of some winter sun and were dazzled by all that Greece has to offer.  From the amazing food to the diverse landscape and the juxtaposition of ancient and modern, Greece is a fantastic country for exploring as a family.",
				},
				{
					URL:         "https://gcziprusz.github.io/peekaboo.travel/a-memorable-weekend-in-krakow-poland",
					IMG:         "https://gcziprusz.github.io/peekaboo.travel/img/eml/IMG_20191228_100722939_HDR.jpg",
					Title:       "A Memorable Weekend in Krakow, Poland",
					Description: "We spent a chilly weekend getting to know Krakow and its intricate history.  Once a home to royalty, Krakow is a well-preserved medieval city located on the bank of the Vistula River in southern Poland."},
				{
					URL:         "https://gcziprusz.github.io/peekaboo.travel/family-weekend-on-the-danube-bend",
					IMG:         "https://gcziprusz.github.io/peekaboo.travel/img/eml/5.jpg",
					Title:       "Family Weekend on the Danube Bend",
					Description: "Family time is the best of time. This particular weekend trip was just that.  Family on both sides traveled from 3 different countries to meet up for a delightful weekend on the Danube Bend.",
				},
				{
					URL:         "https://gcziprusz.github.io/peekaboo.travel/seaside-escape-to-cascais-portugal",
					IMG:         "https://gcziprusz.github.io/peekaboo.travel/img/eml/1.jpg",
					Title:       "Seaside Escape to Cascais Portugal",
					Description: "We had such a lovely day exploring the beaches and coastline of Cascais. If you're looking for warmer weather this winter ...",
				},
			},
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":80", nil)
}
