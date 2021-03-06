package soundcloud

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/pquerna/ffjson/ffjson"
)

const (
	baseAPIURL = "https://api.soundcloud.com/%s"
)

//Auth token SoundCloud
type auth struct {
	AccessToken  string `json:"access_token"`
	Scope        string `json:"scope"`
	ExpiresIn    string `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

//Client client struct
type Client struct {
	// Keep client credentials internal
	clientID, clientSecret, username, password string
	authToken                                  auth
}

//InitClient init client with authtoken
func InitClient(username, authtoken string) *Client {
	c := &Client{}
	c.authToken.AccessToken = authtoken
	c.username = username
	fmt.Printf("%s: %s", c.username, c.authToken.AccessToken)
	return c
}

//NewClient initialized a new client
func NewClient(clientID, clientSecret, username, password string) *Client {
	baseURL := fmt.Sprintf(baseAPIURL, "oauth2/token?")
	params := url.Values{}
	params.Add("client_id", clientID)
	params.Add("client_secret", clientSecret)
	params.Add("grant_type", "password")
	params.Add("username", username)
	params.Add("password", password)
	finalURL := baseURL + params.Encode()
	response, err := http.Post(finalURL, "application/json", nil)
	defer response.Body.Close()
	//AC SoundCloud access token
	var ac auth
	if err != nil {
		log.Fatal("Error in client authorization", err)
		return nil
	}
	decoder := ffjson.NewDecoder()
	decoder.Decode(streamToByte(response.Body), &ac)
	return &Client{clientID: clientID, clientSecret: clientSecret,
		username: username, password: password, authToken: ac}
}

//GetUser gets the user info
func (c Client) GetUser(user *User) {
	t := time.Now()
	baseURL := fmt.Sprintf(baseAPIURL, "me?")
	params := url.Values{}
	params.Add("oauth_token", c.authToken.AccessToken)
	finalURL := baseURL + params.Encode()
	response, err := http.Get(finalURL)
	if err != nil {
		log.Fatal("Error in grabbing user info", err)
		return
	}
	defer response.Body.Close()
	decoder := ffjson.NewDecoder()
	decoder.Decode(streamToByte(response.Body), &user)
	log.Printf("Retrieved User in:%s\n", time.Since(t))
}

//StreamToByte convert io.reader to byte aray for json encoding
func streamToByte(stream io.Reader) []byte {
	buf := new(bytes.Buffer)
	buf.ReadFrom(stream)
	return buf.Bytes()
}

//GetUserGroups gets a list of the user's groups
func (c Client) GetUserGroups(groups *[]Group) []byte {
	t := time.Now()
	baseURL := fmt.Sprintf(baseAPIURL, "me/groups?")
	params := url.Values{}
	params.Add("limit", "75")
	params.Add("oauth_token", c.authToken.AccessToken)
	finalURL := baseURL + params.Encode()
	response, err := http.Get(finalURL)
	if err != nil {
		log.Fatal("Error in grabbing user groups", err)
		return nil
	}
	defer response.Body.Close()
	decoder := ffjson.NewDecoder()
	var respInBytes = streamToByte(response.Body)
	decoder.Decode(respInBytes, &groups)
	log.Printf("Retrieved Groups in:%s\n", time.Since(t))
	return respInBytes
}

//GetTracks gets the user's tracks
func (c Client) GetTracks(tracks *[]Track) []byte {
	t := time.Now()
	baseURL := fmt.Sprintf(baseAPIURL, "me/tracks?")
	params := url.Values{}
	params.Add("oauth_token", c.authToken.AccessToken)
	finalURL := baseURL + params.Encode()
	response, err := http.Get(finalURL)
	if err != nil {
		log.Fatal("Error in grabbing user info", err)
		return nil
	}
	defer response.Body.Close()
	decoder := ffjson.NewDecoder()
	var respInBytes = streamToByte(response.Body)
	decoder.Decode(respInBytes, &tracks)
	log.Printf("Retrieved Tracks in:%s\n", time.Since(t))
	return respInBytes
}

//GetNewGroups gets a new set of groups to upload to
func (c Client) GetNewGroups(query string, groups *[]Group) []byte {
	t := time.Now()
	baseURL := fmt.Sprintf(baseAPIURL, "groups?")
	params := url.Values{}
	params.Add("q", query)
	params.Add("limit", "75")
	params.Add("client_id", c.clientID)
	finalURL := baseURL + params.Encode()
	log.Println(finalURL)
	response, err := http.Get(finalURL)
	if err != nil {
		log.Fatal("Error in grabbing user info", err)
		return nil
	}
	defer response.Body.Close()
	decoder := ffjson.NewDecoder()
	var respInBytes = streamToByte(response.Body)
	decoder.Decode(respInBytes, &groups)
	log.Println(response.StatusCode)
	log.Printf("Retrieved Tracks in:%s\n", time.Since(t))
	return respInBytes
}

//AddToGroup adds track to gorup
func (c Client) AddToGroup(groupID, trackID int) (int, error) {
	responseCode, err := groupReq(groupID, trackID, http.MethodPut, c.authToken.AccessToken)
	return responseCode, err
}

//RemoveFromGroup Remove track from a Group
func (c Client) RemoveFromGroup(groupID, trackID int) (int, error) {
	responseCode, err := groupReq(groupID, trackID, http.MethodDelete, c.authToken.AccessToken)
	return responseCode, err
}

//GroupReq takes in the http method and token
func groupReq(groupID, trackID int, httpMethod, accessToken string) (int, error) {
	URLToBuild := fmt.Sprintf(baseAPIURL, "groups/%d/contributions/%d?")
	baseURL := fmt.Sprintf(URLToBuild, groupID, trackID)
	params := url.Values{}
	params.Add("oauth_token", accessToken)
	finalURL := baseURL + params.Encode()
	req, err := http.NewRequest(httpMethod, finalURL, nil)
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalf("failure in %s request\n", httpMethod)
	}
	response, err := http.DefaultClient.Do(req)
	return response.StatusCode, err
}

//GetClientName returns the given client user name
func (c Client) GetClientName() string {
	return c.username
}
