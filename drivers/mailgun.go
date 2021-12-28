// Copyright 2020 The Go Mail Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package drivers

import (
	"context"
	"errors"
	"fmt"
	"github.com/ainsleyclark/go-mail/internal/httputil"
	"github.com/ainsleyclark/go-mail/mail"
	"net/http"
	"strings"
)

// mailGun represents the data for sending mail via the
// MailGun API. Configuration, the client and the
// main send function are parsed for sending
// data.
type mailGun struct {
	cfg    mail.Config
	client httputil.Requester
}

const (
	// mailgunEndpoint defines the endpoint to POST to.
	mailgunEndpoint = "/v3/%s/messages"
)

// NewMailGun creates a new MailGun client. Configuration
// is validated before initialisation.
func NewMailGun(cfg mail.Config) (mail.Mailer, error) {
	err := cfg.Validate()
	if err != nil {
		return nil, err
	}
	if cfg.Domain == "" {
		return nil, errors.New("driver requires a domain")
	}
	return &mailGun{
		cfg:    cfg,
		client: httputil.NewClient(),
	}, nil
}

func (m *mailGun) Send(t *mail.Transmission) (mail.Response, error) {
	err := t.Validate()
	if err != nil {
		return mail.Response{}, err
	}

	f := httputil.NewFormData()
	f.AddValue("from", fmt.Sprintf("%s <%s>", m.cfg.FromName, m.cfg.FromAddress))
	f.AddValue("subject", t.Subject)
	f.AddValue("html", t.HTML)
	f.AddValue("text", t.PlainText)

	for _, to := range t.Recipients {
		f.AddValue("to", to)
	}

	if t.HasCC() {
		for _, c := range t.CC {
			f.AddValue("cc", c)
		}
	}

	if t.HasBCC() {
		for _, b := range t.BCC {
			f.AddValue("bcc", b)
		}
	}

	if t.Attachments.Exists() {
		for _, v := range t.Attachments {
			f.AddBuffer("attachment", v.Filename, v.Bytes)
		}
	}

	url := fmt.Sprintf("%s/%s", m.cfg.URL, strings.TrimPrefix(fmt.Sprintf(mailgunEndpoint, m.cfg.Domain), "/"))
	req := httputil.NewHTTPRequest(http.MethodPost, url)
	req.SetBasicAuth("api", m.cfg.APIKey)

	buf, resp, err := m.client.Do(context.Background(), req, f)
	if err != nil {
		return mail.Response{}, err
	}

	fmt.Println("here", string(buf))
	fmt.Printf("%+v\n", resp)

	return mail.Response{}, nil
}
