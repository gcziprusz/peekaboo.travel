package main

import (
    "html/template"
    "net/http"
)
type Email struct {
	Sections []EmailSection
}

type EmailSection struct {
		URL  string
		IMG  string // size 580x435
		Title string
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
					URL: "https://peekaboo.travel/family-weekend-on-the-danube-bend",
					IMG: "https://peekaboo.travel/img/eml/visegrad.jpg" ,
					Title: "NEW Peekabo Travel post: Family Weekend on the Danube Bend",
					Description: "Family time is the best of time. This particular weekend trip was just that.  Family on both sides traveled from 3 different countries to meet up for a delightful weekend on the Danube Bend.",
					},
					{
					URL: "https://peekaboo.travel/seaside-escape-to-cascais-portugal",
					IMG: "https://peekaboo.travel/img/eml/1.jpg" ,
					Title: "Seaside Escape to Cascais Portugal",
					Description: "We had such a lovely day exploring the beaches and coastline of Cascais. If you're looking for warmer weather this winter, do keep Portugal in mind.",
					},{
					IMG:"https://peekaboo.travel/img/eml/2.jpg",
					URL: "https://peekaboo.travel/a-magical-day-in-sintra-with-kids/",
					Title: "A Magical Day in Sintra with Kids",
					Description:"Sintra's medieval castles, lavish palaces, are Portuguese charm at it's best. Sintra is the most popular day trip from Lisbon.",
					},{
					IMG: "https://peekaboo.travel/img/eml/3.jpg", // size 220x220
					URL: "https://peekaboo.travel/little-explorers-in-lisbon/",
					Title: "Little Explorers in Lisbon",
					Description: "Lisbon stole our hearts and we are already dreaming of a return trip to this incredible place.",
					},
				},
			}
        tmpl.Execute(w, data)
    })
    http.ListenAndServe(":80", nil)
}