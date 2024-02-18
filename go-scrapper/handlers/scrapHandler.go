package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
	"strings"

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
func ScrapHandler(w http.ResponseWriter, r *http.Request){
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al parsear el cuerpo de la solicitud", http.StatusBadRequest)
		return
	}

	// Obtener el valor del parámetro "q" (url)
	url := r.Form.Get("q")
	if url == "" {
		http.Error(w, "Falta el parámetro 'q' en la solicitud", http.StatusBadRequest)
		return
	}
	c := colly.NewCollector()

	c.OnHTML("body",func(h *colly.HTMLElement) {
		jsonString := h.Text

		
		re := regexp.MustCompile(`"url"\s*:\s*"([^"]+)"`)
		// Encontrar todas las coincidencias en el texto
		matches := re.FindAllStringSubmatch(jsonString, -1)
		if len(matches) == 0{
			re := regexp.MustCompile(`"src"\s*:\s*"([^"]+)"`)
			// re = regexp.MustCompile(`<video[^>]*src="([^"]+)"[^>]*>`)
			// Encontrar todas las coincidencias en el texto HTML
			matches = re.FindAllStringSubmatch(jsonString, -1)
			
		}
		
		fmt.Println("matches:", len(matches))
		// Imprimir los resultados
		var validURLs []string
		for _, match := range matches {
			url := match[1]
			if strings.HasSuffix(url, ".jpg") || strings.HasSuffix(url, ".png") || strings.HasSuffix(url, ".jpeg") {
				continue
			}
			if strings.Contains(url,"youtube.com"){
				continue
			}
			validURLs = append(validURLs, url)
		}
	
		urlsArr := make(map[string]interface{})
		// Imprimir los resultados
		for i, url := range validURLs {
			key := fmt.Sprintf("url-%d", i+1)
			urlsArr[key] = url
		}
		// Convertir urlsArr a JSON
		jsonData, err := json.Marshal(urlsArr)
		if err != nil {
			fmt.Println("Error al convertir a JSON:", err)
			return
		}
		sendResponse(w, jsonData)
		
	})

	c.OnScraped(func(r *colly.Response) {
		// sendResponse(w, string(r.Body))
	})

	c.OnError(func(r *colly.Response, err error) {
		// sendResponse(w, "Error")
		fmt.Println("error", err)
	})

	c.Visit(url)
}

func sendResponse(w http.ResponseWriter, text []byte){
	w.Header().Set("Content-Type", "application/json") // Establecer el tipo de contenido como JSON
	w.WriteHeader(http.StatusOK)                       // Establecer el código de estado HTTP 200
	w.Write([]byte(text)) 
}