package services

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

type SpotifyService struct {
	ClientID     string
	ClientSecret string
	AccessToken  string
	ExpiresAt    time.Time
}

type SpotifyToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type SpotifySearchResponse struct {
	Tracks struct {
		Items []SpotifyTrack `json:"items"`
	} `json:"tracks"`
}

type SpotifyTrack struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	URI        string `json:"uri"`
	PreviewURL string `json:"preview_url"`
	DurationMS int    `json:"duration_ms"`
	Artists    []struct {
		Name string `json:"name"`
	} `json:"artists"`
	Album struct {
		Name   string `json:"name"`
		Images []struct {
			URL string `json:"url"`
		} `json:"images"`
	} `json:"album"`
}

func NewSpotifyService(clientID, clientSecret string) *SpotifyService {
	return &SpotifyService{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}
}

func (s *SpotifyService) GetToken() (string, error) {
	if s.AccessToken != "" && time.Now().Before(s.ExpiresAt) {
		return s.AccessToken, nil
	}

	client := &http.Client{}
	data := url.Values{}
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", err
	}

	auth := base64.StdEncoding.EncodeToString([]byte(s.ClientID + ":" + s.ClientSecret))
	req.Header.Add("Authorization", "Basic "+auth)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var token SpotifyToken
	err = json.Unmarshal(body, &token)
	if err != nil {
		return "", err
	}

	s.AccessToken = token.AccessToken
	s.ExpiresAt = time.Now().Add(time.Duration(token.ExpiresIn-60) * time.Second)

	return s.AccessToken, nil
}

func (s *SpotifyService) SearchTracks(query string, limit int) ([]SpotifyTrack, error) {
	token, err := s.GetToken()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	searchURL := fmt.Sprintf("https://api.spotify.com/v1/search?q=%s&type=track&limit=%d",
		url.QueryEscape(query), limit)

	req, err := http.NewRequest("GET", searchURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var searchResponse SpotifySearchResponse
	err = json.Unmarshal(body, &searchResponse)
	if err != nil {
		return nil, err
	}

	return searchResponse.Tracks.Items, nil
}

func (s *SpotifyService) GetTrack(trackID string) (*SpotifyTrack, error) {
	token, err := s.GetToken()
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	trackURL := fmt.Sprintf("https://api.spotify.com/v1/tracks/%s", trackID)

	req, err := http.NewRequest("GET", trackURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", "Bearer "+token)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var track SpotifyTrack
	err = json.Unmarshal(body, &track)
	if err != nil {
		return nil, err
	}

	return &track, nil
}
