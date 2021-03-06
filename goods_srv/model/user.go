package model

/*
	密码不可反解

*/

type Category struct {
	BaseModel
	Name             string `gorm:"type:varchar(20); not null"`
	ParentCategoryID int32
	ParentCategory   *Category
	Level            int32 `gorm:"type:int;not null;default:1"`
	IsTab            bool  `gorm:"default:false;not null"`
}
type Brands struct {
	BaseModel
	Name string `gorm:"type:varchar(20);not null"`
	Logo string `gorm:"type:varchar(200);default:'';not null"`
}

type GoodsCategoryBrand struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands
}

func (GoodsCategoryBrand) TableName() string {
	return "goodscategorybrand"
}

type Banner struct {
	BaseModel
	Image string `gorm:"type:varchar(200);not null"`
	Url   string `gorm:"type:varchar(200);not null"`
	Index int32  `gorm:"type:int;default:1;not null"`
}

type Goods struct {
	BaseModel
	CategoryID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Category   Category

	BrandsID int32 `gorm:"type:int;index:idx_category_brand,unique"`
	Brands   Brands

	OnSale   bool `gorm:"type:bool;default:false;not null"`
	ShipFree bool `gorm:"type:bool;default:false;not null"`
	IsNew    bool `"type:bool;default:false;not null"`
	IsHot    bool `"type:bool;default:false;not null"`

	Name     string `"type:string;type:varchar(50);not null"`
	GoodsSn  string `"type:string;type:varchar(50);not null"`
	ClickNum int32  `gorm:"type:int;default:0;not null"`
	SoldNum  int32  `gorm:"type:int;default:0;not null"`
	FavNum   int32  `gorm:"type:int;default:0;not null"`

	MarketPrice float32 `gorm:"not null"`
	ShopPrice   float32 `gorm:"not null"`

	GoodsBrief      string   `"type:string;type:varchar(100);not null"`
	Images          GormList `"type:string;type:varchar(1000);not null"`
	DescImages      GormList `"type:string;type:varchar(1000);not null"`
	GoodsFrontImage string   `"type:string;type:varchar(200);not null"`
}
