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
	"github.com/ainsleyclark/go-mail"
)

var (
	smtpCfg = mail.Config{
		URL:         "CHANGE ME",
		FromAddress: "CHANGE ME",
		FromName:    "CHANGE ME",
		Password:    "CHANGE ME",
		Port:        587,
	}
)

func (t *MailTestSuite) Test_SMTP() {
	tx := t.GetTransmission()

	driver, err := mail.NewClient(mail.SMTP, smtpCfg)
	if err != nil {
		t.Fail("error creating client", err)
		return
	}

	result, err := driver.Send(tx)
	if err != nil {
		t.Fail("error sending smtp email", err)
		return
	}

	t.Equal(200, result.StatusCode)
	t.NotEmpty(result.Message)
}
