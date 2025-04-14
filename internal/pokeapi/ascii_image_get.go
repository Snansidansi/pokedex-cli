package pokeapi

import (
	"fmt"
	"image"
	"net/http"

	"github.com/qeesung/image2ascii/convert"
)

func (c *Client) GetAsciiImage(inputURL string, size int) (string, error) {
	identifier := fmt.Sprintf("%s#%v", inputURL, size)
	if data, ok := c.cache.Get(identifier); ok {
		asciiImage := string(data)
		return asciiImage, nil
	}

	img, err := c.getImage(inputURL)
	if err != nil {
		return "", err
	}

	asciiImage := generateAsciiImage(img, size)
	c.cache.Add(identifier, []byte(asciiImage))
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

func generateAsciiImage(img image.Image, size int) string {
	if size == 0 {
		size = 40
	}

	convertOptions := convert.DefaultOptions
	convertOptions.FixedWidth = size * 2
	convertOptions.FixedHeight = size

	converter := convert.NewImageConverter()
	asciiImage := converter.Image2ASCIIString(img, &convertOptions)
	return asciiImage
}
