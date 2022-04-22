package server

var (
	pagesMap map[string]*Page // generated at the init() function
	pages    = []*Page{
		{
			URL:         "quere.ndo.cafe",
			Title:       "QUERO CAFÉ!1!!",
			Subtitle:    "ÉÉÉÉÉÉÉÉ",
			Description: "QUERO CAFÉÉÉÉÉÉÉÉÉÉ",
			Image:       "quero-cafe.gif",
		},
		{
			URL:         "faze.ndo.cafe",
			Title:       "Fazendo Café",
			Subtitle:    "Já volto",
			Description: "Fui fazer café... Já volto",
			Image:       "cafe.gif",
		},
		{
			URL:         "toma.ndo.cafe",
			Title:       "Tomando Café",
			Subtitle:    "Aguarde um minuto, estou tomando café",
			Description: "Estou tomando café... Já volto",
			Image:       "blink.gif",
		},
		{
			URL:         "bebe.ndo.cafe",
			Title:       "Bebendo Café...",
			Subtitle:    "Recarregando as energias e matando a sede",
			Description: "Recarregando as energias",
			Image:       "glub.gif",
		},
		{
			URL:         "come.ndo.cafe",
			Title:       "Comendo Café",
			Subtitle:    "Yummy",
			Description: "Comendo Café",
			Image:       "yummy.gif",
		},
		{
			URL:         "derrama.ndo.cafe",
			Title:       "Derramando Café",
			Subtitle:    "Infelizmente, acabei derramando café",
			Description: "Acabei derramando café, estarei de volta assim que superar",
			Image:       "bad.gif",
		},
		{
			URL:         "queima.ndo.cafe",
			Title:       "Queimei o café",
			Subtitle:    "=(",
			Description: "Queimei o café",
			Image:       "oof.png",
		},
	}
)

func init() {
	pagesMap = make(map[string]*Page)
	for _, page := range pages {
		pagesMap[page.URL] = page
	}
}
