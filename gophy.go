package gophy

const GIPHY_API = "http://api.giphy.com/v1/gifs/search"

// Taken from github.com/poptip/buster
type GiphyImageData struct {
	URL    string
	Width  string
	Height string
	Size   string
	Frames string
}

type GiphyGif struct {
	Type               string
	Id                 string
	URL                string
	Tags               string
	BitlyGifURL        string `json:"bitly_gif_url"`
	BitlyFullscreenURL string `json:"bitly_fullscreen_url"`
	BitlyTiledURL      string `json:"bitly_tiled_url"`
	Images             struct {
		Original               GiphyImageData
		FixedHeight            GiphyImageData `json:"fixed_height"`
		FixedHeightStill       GiphyImageData `json:"fixed_height_still"`
		FixedHeightDownsampled GiphyImageData `json:"fixed_height_downsampled"`
		FixedWidth             GiphyImageData `json:"fixed_width"`
		FixedwidthStill        GiphyImageData `json:"fixed_width_still"`
		FixedwidthDownsampled  GiphyImageData `json:"fixed_width_downsampled"`
	}
}
