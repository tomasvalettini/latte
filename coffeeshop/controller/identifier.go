package controller

type Identifier struct {
	Id    int
	Title string
}

func (bi Identifier) IsValid() bool {
	return bi.IsIdValid() || bi.IsTitleValid()
}

func (bi Identifier) IsIdValid() bool {
	return bi.Id >= 0
}

func (bi Identifier) IsTitleValid() bool {
	return bi.Title != ""
}

// add method to that validates the blend identifier and returns itself if valid, nil otherwise
func (bi Identifier) Validate() *Identifier {
	if bi.IsValid() {
		return &bi
	}

	return nil
}
