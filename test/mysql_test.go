// Copyright 2017 AMIS Technologies
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package test

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/assert"
)

func TestMySQLSetupAndTeardown(t *testing.T) {
	mysql, err := SetupMySQL()
	assert.NoError(t, err, "mysql connection handle should be created.")

	err = mysql.Teardown()
	assert.NoError(t, err, "mysql connection handle should be torn down.")
}

func TestMySQLContainer(t *testing.T) {
	container, err := NewMySQLContainer(DefaultMySQLOptions)
	assert.NoError(t, err, "mysql container should be created.")
	assert.NotNil(t, container)
	assert.NoError(t, container.Start(), "mysql container should be started")

	db, err := gorm.Open("mysql", container.URL)
	assert.NoError(t, err, "mysql connection should work")
	db.Close()

	// stop MySQL
	assert.NoError(t, container.Suspend())
	time.Sleep(100 * time.Millisecond)
	_, err = gorm.Open("mysql", container.URL)
	assert.Error(t, err, "should got error")

	// restart MySQL
	assert.NoError(t, container.Start())
	db, err = gorm.Open("mysql", container.URL)
	assert.NoError(t, err, "should be no error")
	db.Close()

	// close MySQL
	assert.NoError(t, container.Stop())
	time.Sleep(100 * time.Millisecond)

	_, err = gorm.Open("mysql", container.URL)
	assert.Error(t, err, "should got error")
}
