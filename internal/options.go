package internal

import "net/http"



func NewClient(options ...option)(*Client , error){
	
	c := Client{
		url: productionURL,
		httpClient: &http.Client{},
	}
}