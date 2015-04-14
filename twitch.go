package twitch

import (
	"net/http"
	"net/url"
)

const (
	BASE_URL = "https://api.twitch.tv/kraken/channels/"
)

type TwitchApi struct {
	Channel      string
	ChannelOauth string
	Url          string
}

func New(ChannelName, Oauth string) *TwitchApi {
	return &TwitchApi{
		Channel:      ChannelName,
		ChannelOauth: Oauth,
		Url:          BASE_URL + ChannelName,
	}
}

func (tw *TwitchApi) UpdateStatus(status string) error {
	client := &http.Client{}
	urlp, err := url.Parse(tw.Url)
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
	urlp, err := url.Parse(tw.Url)
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
