package handlers

import (
	"fmt"
	"regexp"

	"github.com/gocolly/colly/v2"
)

func ExtractURLs(data interface{}) []string {
    urls := []string{}

    switch v := data.(type) {
    case map[string]interface{}:
        for _, value := range v {
            urls = append(urls, ExtractURLs(value)...)
        }
    case []interface{}:
        for _, value := range v {
            urls = append(urls, ExtractURLs(value)...)
        }
    case string:
        urls = append(urls, v)
    }

    return urls
}

//working, kinda, deberia ahora esperar los parametros w y r
func ScrapHandler(){
	fmt.Println("Comenzando a scrappear")

	c := colly.NewCollector()

	c.OnHTML("body",func(h *colly.HTMLElement) {
		jsonString := h.Text

		re := regexp.MustCompile(`"url"\s*:\s*"([^"]+)"`)

		// Encontrar todas las coincidencias en el texto
		matches := re.FindAllStringSubmatch(jsonString, -1)
	
		// Imprimir los resultados
		for _, match := range matches {
			fmt.Println("Clave:", match[0])
			fmt.Println("Valor:", match[1])
		}
	})

	c.Visit("https://www3.animeflv.net/ver/dosanko-gal-wa-namara-menkoi-6")

}