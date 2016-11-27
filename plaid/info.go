package plaid

import (
	"bytes"
	"encoding/json"
)

// InfoAddUser (POST /info) submits a set of user credentials to add an Info user.
//
// See https://plaid.com/docs/api/#add-info-user.
func (c *Client) InfoAddUser(username, password, institutionType string,
	options *InfoOptions) (postRes *postResponse, mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(infoJson{
		c.clientID,
		c.secret,
		institutionType,
		username,
		password,
		options,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/info", bytes.NewReader(jsonText))
}

// InfoStepSendMethod (POST /info/step) specifies a particular send method for MFA,
// e.g. `{"mask":"xxx-xxx-5309"}`.
//
// See https://plaid.com/docs/api/#info-mfa
func (c *Client) infoStepSendMethod(accessToken, key, value string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	sendMethod := map[string]string{key: value}
	jsonText, err := json.Marshal(infoStepSendMethodJson{
		c.clientID,
		c.secret,
		accessToken,
		infoStepOptions{sendMethod},
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/info/step", bytes.NewReader(jsonText))
}

// InfoStep (POST /info/step) submits an MFA answer for a given access token.
//
// See https://plaid.com/docs/api/#mfa-authentication.
func (c *Client) InfoStep(accessToken, answer string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(infoStepJson{
		c.clientID,
		c.secret,
		accessToken,
		answer,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/info/step", bytes.NewReader(jsonText))
}

// InfoGet (POST /info/get) retrieves account holder information data for a given access token.
//
// See https://plaid.com/docs/api/#get-info-data.
func (c *Client) InfoGet(accessToken string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(infoGetJson{
		c.clientID,
		c.secret,
		accessToken,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.postAndUnmarshal("/info/get", bytes.NewReader(jsonText))
}

// InfoUpdate (PATCH /info) updates user credentials for a given access token.
//
// See https://plaid.com/docs/api/#update-user.
func (c *Client) InfoUpdate(username, password, pin, accessToken string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(infoUpdateJson{
		c.clientID,
		c.secret,
		username,
		password,
		pin,
		accessToken,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.patchAndUnmarshal("/info", bytes.NewReader(jsonText))
}

// InfoUpdateStep (PATCH /info/step) updates user credentials and MFA for a given access token.
//
// TODO: Documentation link once it is posted.
func (c *Client) InfoUpdateStep(username, password, pin, mfa, accessToken string) (postRes *postResponse,
	mfaRes *mfaResponse, err error) {

	jsonText, err := json.Marshal(infoUpdateStepJson{
		c.clientID,
		c.secret,
		username,
		password,
		pin,
		mfa,
		accessToken,
	})
	if err != nil {
		return nil, nil, err
	}
	return c.patchAndUnmarshal("/info/step", bytes.NewReader(jsonText))
}

// InfoDelete (DELETE /info) deletes data for a given access token.
//
// See https://plaid.com/docs/api/#delete-info-user.
func (c *Client) InfoDelete(accessToken string) (deleteRes *deleteResponse, err error) {
	jsonText, err := json.Marshal(infoDeleteJson{
		c.clientID,
		c.secret,
		accessToken,
	})
	if err != nil {
		return nil, err
	}
	return c.deleteAndUnmarshal("/info", bytes.NewReader(jsonText))
}

// InfoOptions represents options associated with adding an Info user.
//
// See https://plaid.com/docs/api/#add-info-user.
type InfoOptions struct {
	List bool `json:"list,omitempty"`
}
type infoJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`
	Type     string `json:"type"`

	Username string       `json:"username"`
	Password string       `json:"password"`
	Options  *InfoOptions `json:"options,omitempty"`
}

type infoStepOptions struct {
	SendMethod map[string]string `json:"send_method"`
}
type infoStepSendMethodJson struct {
	ClientID    string          `json:"client_id"`
	Secret      string          `json:"secret"`
	AccessToken string          `json:"access_token"`
	Options     infoStepOptions `json:"options"`
}

type infoStepJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`

	MFA string `json:"mfa"`
}

type infoGetJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}

type infoUpdateJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`

	Username    string `json:"username"`
	Password    string `json:"password"`
	PIN         string `json:"pin,omitempty"`
	AccessToken string `json:"access_token"`
}

type infoUpdateStepJson struct {
	ClientID string `json:"client_id"`
	Secret   string `json:"secret"`

	Username    string `json:"username"`
	Password    string `json:"password"`
	PIN         string `json:"pin,omitempty"`
	MFA         string `json:"mfa"`
	AccessToken string `json:"access_token"`
}

type infoDeleteJson struct {
	ClientID    string `json:"client_id"`
	Secret      string `json:"secret"`
	AccessToken string `json:"access_token"`
}
