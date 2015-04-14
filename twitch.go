package twitch

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	CHANNELS_BASE_URL = "https://api.twitch.tv/kraken/channels/"
	STREAMS_BASE_URL  = "https://api.twitch.tv/kraken/streams/"
)

type StreamsData struct {
	Links struct {
		Channel string `json:"channel"`
		Self    string `json:"self"`
	} `json:"_links"`
	Stream struct {
		ID    float64 `json:"_id"`
		Links struct {
			Self string `json:"self"`
		} `json:"_links"`
		AverageFps float64 `json:"average_fps"`
		Channel    struct {
			ID    float64 `json:"_id"`
			Links struct {
				Chat          string `json:"chat"`
				Commercial    string `json:"commercial"`
				Editors       string `json:"editors"`
				Features      string `json:"features"`
				Follows       string `json:"follows"`
				Self          string `json:"self"`
				StreamKey     string `json:"stream_key"`
				Subscriptions string `json:"subscriptions"`
				Teams         string `json:"teams"`
				Videos        string `json:"videos"`
			} `json:"_links"`
			Background                   interface{} `json:"background"`
			Banner                       string      `json:"banner"`
			BroadcasterLanguage          string      `json:"broadcaster_language"`
			CreatedAt                    string      `json:"created_at"`
			Delay                        float64     `json:"delay"`
			DisplayName                  string      `json:"display_name"`
			Followers                    float64     `json:"followers"`
			Game                         string      `json:"game"`
			Language                     string      `json:"language"`
			Logo                         string      `json:"logo"`
			Mature                       bool        `json:"mature"`
			Name                         string      `json:"name"`
			Partner                      bool        `json:"partner"`
			ProfileBanner                string      `json:"profile_banner"`
			ProfileBannerBackgroundColor string      `json:"profile_banner_background_color"`
			Status                       string      `json:"status"`
			UpdatedAt                    string      `json:"updated_at"`
			URL                          string      `json:"url"`
			VideoBanner                  string      `json:"video_banner"`
			Views                        float64     `json:"views"`
		} `json:"channel"`
		CreatedAt string `json:"created_at"`
		Game      string `json:"game"`
		Preview   struct {
			Large    string `json:"large"`
			Medium   string `json:"medium"`
			Small    string `json:"small"`
			Template string `json:"template"`
		} `json:"preview"`
		VideoHeight float64 `json:"video_height"`
		Viewers     float64 `json:"viewers"`
	} `json:"stream"`
}

type TwitchApi struct {
	Channel      string
	ChannelOauth string
	ChannelsUrl  string
	StreamsURL   string
}

func New(ChannelName, Oauth string) *TwitchApi {
	return &TwitchApi{
		Channel:      ChannelName,
		ChannelOauth: Oauth,
		ChannelsUrl:  CHANNELS_BASE_URL + ChannelName,
		StreamsURL:   STREAMS_BASE_URL + ChannelName,
	}
}

func (tw *TwitchApi) UpdateStatus(status string) error {
	client := &http.Client{}
	urlp, err := url.Parse(tw.ChannelsUrl)
	if err != nil {
		return err
	}

	param := url.Values{}
	param.Add("channel[status]", status)
	param.Add("_method", "put")
	param.Add("oauth_token", tw.ChannelOauth)
	urlp.RawQuery = param.Encode()

	_, err = client.Get(urlp.String())
	if err != nil {
		return err
	}

	return nil
}

func (tw *TwitchApi) UpdateGame(game string) error {
	client := &http.Client{}
	urlp, err := url.Parse(tw.ChannelsUrl)
	if err != nil {
		return err
	}

	param := url.Values{}
	param.Add("channel[game]", game)
	param.Add("_method", "put")
	param.Add("oauth_token", tw.ChannelOauth)
	urlp.RawQuery = param.Encode()

	_, err = client.Get(urlp.String())
	if err != nil {
		return err
	}

	return nil
}

func (tw *TwitchApi) Uptime() (string, error) {
	var stream *StreamsData
	res, err := http.Get(tw.StreamsURL)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	if err := json.Unmarshal(body, &stream); err != nil {
		return "", err
	}

	layout := "2006-01-02T15:04:05Z"
	startTime := stream.Stream.CreatedAt
	if startTime != "" {
		parsedTime, err := time.Parse(layout, startTime)
		if err != nil {
			return "", err
		}
		duration := time.Since(parsedTime)
		return duration.String(), nil
	} else {
		return "", nil
	}
}
