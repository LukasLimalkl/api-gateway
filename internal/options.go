package internal



func NewClient(url string, port  int)  *Client {
	
	return &Client{url, port,timeout, httpClient}

}