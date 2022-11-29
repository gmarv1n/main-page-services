package mps

type Quote struct {
	Id                int    `json:"id" db:"id"`
	Quote             string `json:"quote" db:"quote" binding: "required"`
	Author            string `json:"author" db:"author"`
	AuthorDescription string `json:"author_description" db:"author_description"`
	RecordId          int    `json:"record_id" db:"record_id"`
	RegionId          int    `json:"region_id" db:"region_id"`
}

type PictureOfTheDayTabsList struct {
	Id       int `json:"id" db:"id"`
	RegionId int `json:"region_id" db:"region_id"`
}

type PictureOfTheDayTabsListsItem struct {
	Id         int `json:"tab_id" db:"id"`
	Type       int `json:"type" db:"type" binding:"required"`
	EntityId   int `json:"id" db:"entity_id"`
	Title      int `json:"item_id" db:"item_id"`
	OrderPlace int `json:"order_place" db:"order_place"`
	RegionId   int `json:"region_id" db:"region_id"`
}

type Theme struct {
	Id       int `json:"id" db:"id"`
	ThemeId  int `json:"theme_id" db:"theme_id" binding:"required"`
	RegionId int `json:"region_id" db:"region_id"`
}

type UpdateThemeInput struct {
	ThemeId  *int `json:"theme_id"`
	RegionId *int `json:"region_id"`
}

type UpdateQuoteInput struct {
	Quote             *string `json:"qoute"`
	Author            *string `json:"author"`
	AuthorDescription *string `json:"author_descritpion"`
	RecordId          *int    `json:"recrd_id"`
	RegionId          *int    `json:"region_id"`
}

type UpdateTabsListInput struct {
	RegionId *int `json:"region_id"`
}

type UpdateTabItemInput struct {
	Title      *string `json:"title"`
	Type       *string `json:"type"`
	EntityId   *int    `json:"entity_id"`
	OrderPlace *int    `json:"order_place" db:"order_place"`
	RegionId   *int    `json:"region_id" db:"region_id"`
}

func (i UpdateThemeInput) Validate() error {

	return nil
}

func (i UpdateQuoteInput) Validate() error {

	return nil
}

func (i UpdateTabItemInput) Validate() error {

	return nil
}
