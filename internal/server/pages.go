package server

var (
	pages = map[string]*Page{
		"quere.ndo.cafe": {
			Title:       "QUERO CAFÉ!1!!",
			Subtitle:    "ÉÉÉÉÉÉÉÉ",
			Description: "QUERO CAFÉÉÉÉÉÉÉÉÉÉ",
			Image:       "quero-cafe.gif",
		},
		"faze.ndo.cafe": {
			Title:       "Fazendo Café",
			Subtitle:    "Já volto",
			Description: "Fui fazer café... Já volto",
			Image:       "cafe.gif",
		},
		"toma.ndo.cafe": {
			Title:       "Tomando Café",
			Subtitle:    "Aguarde um minuto, estou tomando café",
			Description: "Estou tomando café... Já volto",
			Image:       "blink.gif",
		},
		"derrama.ndo.cafe": {
			Title:       "Derramando Café",
			Subtitle:    "Infelizmente, acabei derramando café",
			Description: "Acabei derramando café, estarei de volta assim que superar",
			Image:       "bad.gif",
		},
		"queima.ndo.cafe": {
			Title:       "Queimei o café",
			Subtitle:    "=(",
			Description: "Queimei o café",
			Image:       "oof.png",
		},
	}
)
