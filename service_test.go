package news

import (
	"os"
	"reflect"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/jmoiron/sqlx"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

var (
	s1 = NewsItem{ID: 1, Title: "name1", Contents: "description1"}
	s2 = NewsItem{ID: 2, Title: "name2", Contents: "description2"}
	s3 = NewsItem{ID: 3, Title: "name3", Contents: "description3"}
	s4 = NewsItem{ID: 4, Title: "name4", Contents: "description4"}
	s5 = NewsItem{ID: 5, Title: "name5", Contents: "description5"}
)

var logger log.Logger

func TestNewsServiceList(t *testing.T) {
	logger = log.NewLogfmtLogger(os.Stderr)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening stub database connection", err)
	}
	defer db.Close()
	sqlxDB := sqlx.NewDb(db, "sqlmock")

	var cols = []string{"id", "title", "contents"}

	// Test Case 1
	mock.ExpectQuery("SELECT *").WillReturnRows(sqlmock.NewRows(cols).
		AddRow(s1.ID, s1.Title, s1.Contents).
		AddRow(s2.ID, s2.Title, s2.Contents).
		AddRow(s3.ID, s3.Title, s3.Contents).
		AddRow(s4.ID, s4.Title, s4.Contents).
		AddRow(s5.ID, s5.Title, s5.Contents))

	s := NewNewsService(sqlxDB, logger)
	for _, testcase := range []struct {
		want []NewsItem
	}{
		{
			want: []NewsItem{s1, s2, s3, s4, s5},
		},
	} {
		have, err := s.List()
		if err != nil {
			t.Errorf(
				"List(): returned error %s",
				err.Error(),
			)
		}
		if want := testcase.want; !reflect.DeepEqual(want, have) {
			t.Errorf("List(): want %v, have %v", want, have)
		}
	}
}
