package server

var (
	pages = map[string]*Page{
		"faze.ndo.cafe": {
			Title:       "Fazendo Café",
			Subtitle:    "Já volto",
			Description: "Fui fazer café... Já volto",
			Gif:         "cafe",
		},
		"toma.ndo.cafe": {
			Title:       "Tomando Café",
			Subtitle:    "Aguarde um minuto, estou tomando café",
			Description: "Estou tomando café... Já volto",
			Gif:         "blink",
		},
		"derrama.ndo.cafe": {
			Title:       "Derramando Café",
			Subtitle:    "Infelizmente, acabei derramando café",
			Description: "Acabei derramando café, estarei de volta assim que superar",
			Gif:         "bad",
		},
	}
)
