package controller

import (
	"os"
	"testing"

	dripdatasource "github.com/tomasvalettini/latte/drips/data/data-source"
	drippath "github.com/tomasvalettini/latte/drips/path"
)

func GetTestDripController() *DripController {
	tp := drippath.GetTestingDripPath()
	bl := dripdatasource.NewDripBacklog(tp.GetDripPath())
	return &DripController{
		dripPath:   tp,
		dataSource: bl,
	}
}

func TestDripController(t *testing.T) {
	tc := GetTestDripController()

	tc.ListDrips()

	tc.AddDrip("test")
	tc.AddDrip("test")
	tc.AddDrip("test")
	tc.AddDrip("test")
	tc.AddDrip("test")

	tc.ListDrips()

	tc.DeleteDrip("1")
	tc.DeleteDrip("0")

	tc.UpdateDrip("2", "new test")

	t.Cleanup(func() {
		os.RemoveAll(drippath.TMP)
	})
}
