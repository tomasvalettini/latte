package controller

type BlendIdentifier struct {
	Id    int
	Title string
}

func (bi BlendIdentifier) IsValid() bool {
	return bi.IsIdValid() || bi.IsTitleValid()
}

func (bi BlendIdentifier) IsIdValid() bool {
	return bi.Id >= 0
}

func (bi BlendIdentifier) IsTitleValid() bool {
	return bi.Title != ""
}

// add method to that validates the blend identifier and returns itself if valid, nil otherwise
func (bi BlendIdentifier) Validate() *BlendIdentifier {
	if bi.IsValid() {
		return &bi
	}

	return nil
}
