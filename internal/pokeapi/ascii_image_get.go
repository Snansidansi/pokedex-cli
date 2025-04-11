package pokeapi

import (
	"image"
	"net/http"

	"github.com/qeesung/image2ascii/convert"
)

func (c *Client) GetAsciiImage(inputURL string) (string, error) {
	if data, ok := c.cache.Get(inputURL); ok {
		asciiImage := string(data)
		return asciiImage, nil
	}

	img, err := c.getImage(inputURL)
	if err != nil {
		return "", err
	}

	asciiImage := generateAsciiImage(img)
	c.cache.Add(inputURL, []byte(asciiImage))
	return asciiImage, nil
}

func (c *Client) getImage(inputURL string) (image.Image, error) {
	req, err := http.NewRequest("GET", inputURL, nil)
	if err != nil {
		return nil, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	img, _, err := image.Decode(resp.Body)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func generateAsciiImage(img image.Image) string {

	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = 80
	convertOptions.FixedHeight = 40

	converter := convert.NewImageConverter()
	asciiImage := converter.Image2ASCIIString(img, &convertOptions)
	return asciiImage
}
