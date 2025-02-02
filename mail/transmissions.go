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

package mail

import (
	"errors"
)

// Transmission represents the JSON structure accepted by
// and returned from the driver's API. Recipients,
// HTML and a subject is required to send the
// email.
type Transmission struct {
	Recipients  []string
	CC          []string
	BCC         []string
	Subject     string
	HTML        string
	PlainText   string
	Attachments []Attachment
}

// Validate runs sanity checks of a Transmission struct.
// This is run before any email sending to ensure
// there are no invalid API calls.
func (t *Transmission) Validate() error {
	if t == nil {
		return errors.New("can't validate a nil transmission")
	}

	if len(t.Recipients) == 0 {
		return errors.New("transmission requires recipients")
	}

	if t.Subject == "" {
		return errors.New("transmission requires a subject")
	}

	if t.HTML == "" {
		return errors.New("transmission requires html content")
	}

	return nil
}

// HasCC determines if there are any CC recipients
// attached to the transmission.
func (t *Transmission) HasCC() bool {
	return len(t.CC) != 0
}

// HasBCC determines if there are any BCC recipients
// attached to the transmission.
func (t *Transmission) HasBCC() bool {
	return len(t.BCC) != 0
}

// HasAttachments determines if there are any attachments in
// the transmission.
func (t Transmission) HasAttachments() bool {
	return len(t.Attachments) != 0
}
