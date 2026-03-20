package controller

import (
	"fmt"
	"log"
	"strconv"

	dripdatasource "github.com/tomasvalettini/latte/drips/data/data-source"
	dripdatamodel "github.com/tomasvalettini/latte/drips/data/model"
	drippath "github.com/tomasvalettini/latte/drips/path"
)

type DripController struct {
	dataSource *dripdatasource.DripBacklog
}

func NewDripController(dripPath drippath.DripPath) *DripController {
	dataSource := dripdatasource.NewDripBacklog(dripPath.GetDripPath())

	return &DripController{
		dataSource: dataSource,
	}
}

func (tc *DripController) ListDrips() {
	drips, notified := tc.loadDrips()
	if !notified {
		return
	}

	fmt.Println("===========")
	fmt.Println(" DRIP LIST ")
	fmt.Println("===========")

	w := dripdatamodel.MaxIdWidth(drips)
	for _, t := range drips {
		fmt.Printf("  [%*d]  %s\n", w, t.Id, t.Text)
	}
}

// adding a new drip based on the dripText
func (tc *DripController) AddDrip(dripText string) {
	drips := tc.dataSource.Load()
	nextId := dripdatamodel.GetNextId(drips)

	drip := dripdatamodel.Drip{
		Id:   nextId,
		Text: dripText,
	}
	drips = append(drips, drip)

	tc.dataSource.Save(drips)
	fmt.Println("Drip added successfully!!!")
}

func (tc *DripController) DeleteDrip(idStr string) {
	id, ok := tc.getDripId(idStr)
	if !ok {
		return
	}

	drips, notified := tc.loadDrips()
	if !notified {
		return
	}

	idx := dripdatamodel.FindIndexFromId(drips, id)
	drips = append(drips[:idx], drips[idx+1:]...)
	tc.dataSource.Save(drips)

	fmt.Printf("Drip with id: %d was successfully removed!!\n", id)
}

func (tc *DripController) UpdateDrip(idStr string, newText string) {
	id, ok := tc.getDripId(idStr)
	if !ok {
		return
	}

	drips := tc.dataSource.Load()
	index := dripdatamodel.FindIndexFromId(drips, id)

	drips[index].Text = newText
	tc.dataSource.Save(drips)

	fmt.Printf("Drip with id: %d was successfully modified with new text:%s\n", id, newText)
}

func (tc *DripController) loadDrips() ([]dripdatamodel.Drip, bool) {
	drips := tc.dataSource.Load()
	dripsCount := len(drips)

	if dripsCount <= 0 {
		fmt.Println("No drips yet.")
		return nil, false
	}

	return drips, true
}

func (tc *DripController) getDripId(idStr string) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Fatalln("Id entered is not a number.")
	}

	return id, true
}
