package internal

import "net/http"



func NewClient(client *Client)(*Client , error){
	
	c := Client{
		url:    client.url ,
		timeout: client.timeout,
		httpClient: &http.Client{},
	}

	return &c, nil

}