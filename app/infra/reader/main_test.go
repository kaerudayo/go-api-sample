package reader_test

import (
	"fmt"
	"testing"

	"github.com/api-sample/app/infra"
	"github.com/api-sample/app/pkg/consts"
)

func TestMain(t *testing.M) {

	fmt.Println("run reader test 🏃")
	infra.TestDatabaseInit("../../scripts/db")
	if err := infra.Migrate.Up(); err != nil {
		fmt.Println(err)
	}
	t.Run()
	fmt.Println("done reader test ✨")
}

func loadFixtures(t *testing.T) {
	if !consts.IsTest() {
		panic("Only the test environment can be run")
	}
	if err := infra.Fixture.Load(); err != nil {
		panic(err)
	}
}
