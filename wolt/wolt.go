package wolt

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/polisko/rosnagen/readexcel"
)

type WoltImport struct {
	Items   []MenuItem    `json:"items"`
	Menu    interface{}   `json:"menu"`
	Options []interface{} `json:"options"`
}
type MenuItem struct {
	ID struct {
		Oid string `json:"$oid"`
	} `json:"_id"`
	Additives                 []interface{} `json:"additives"`
	AlcoholPercentage         int           `json:"alcohol_percentage"`
	Allergens                 []interface{} `json:"allergens"`
	BasePrice                 int           `json:"baseprice"`
	ConditionsOfUseAndStorage []interface{} `json:"conditions_of_use_and_storage"`
	CountryOfOrigin           []interface{} `json:"country_of_origin"`
	CourierRestrictions       []interface{} `json:"courier_restrictions"`
	DeliveryMethods           []string      `json:"delivery_methods"`
	Description               []interface{} `json:"description"`
	DietaryPreferences        []interface{} `json:"dietary_preferences"`
	DisplayUnit               string        `json:"display_unit,omitempty"`
	DistributorInformation    []interface{} `json:"distributor_information"`
	Enabled                   struct {
		After   interface{} `json:"after,omitempty"`
		Enabled bool        `json:"enabled"`
	} `json:"enabled"`
	Image             string        `json:"image,omitempty"`
	ImageBlurhash     string        `json:"image_blurhash,omitempty"`
	Images            []interface{} `json:"images"`
	Ingredients       []interface{} `json:"ingredients"`
	IsBundleOffer     bool          `json:"is_bundle_offer"`
	IsOverTheCounter  bool          `json:"is_over_the_counter"`
	MandatoryWarnings []interface{} `json:"mandatory_warnings"`
	MenuID            struct {
		Oid string `json:"$oid"`
	} `json:"menu_id"`
	MoreInformation []interface{} `json:"more_information"`
	Name            []struct {
		Lang     string `json:"lang"`
		Value    string `json:"value"`
		Verified bool   `json:"verified"`
	} `json:"name"`
	NumberOfUnits            int           `json:"number_of_units,omitempty"`
	NutritionFacts           []interface{} `json:"nutrition_facts"`
	NutritionValues          []interface{} `json:"nutrition_values"`
	OfferingPlatformMetadata struct {
		ID struct {
			Oid string `json:"$oid"`
		} `json:"id"`
	} `json:"offering_platform_metadata"`
	Options                    []interface{} `json:"options"`
	ProducerInformation        []interface{} `json:"producer_information"`
	ProductHierarchicalLevelID string        `json:"product_hierarchy_level_id,omitempty"`
	UseInventory               *bool         `json:"use_inventory,omitempty"`
	UserRequirements           []interface{} `json:"user_requirements"`
	V                          struct {
		Author struct {
			ID   string `json:"id"`
			Kind string `json:"kind"`
		} `json:"author"`
		CreatedAt struct {
			Date int64 `json:"$date"`
		} `json:"created_at"`
		IsRemoved bool `json:"is_removed"`
		Num       int  `json:"num"`
		OrigID    struct {
			Oid string `json:"$oid"`
		} `json:"orig_id"`
	} `json:"v"`
	VatPercentage        int           `json:"vat_percentage"`
	VatCategoryCode      string        `json:"vat_category_code,omitempty"`
	VatPercentageDecimal int           `json:"vat_percentage_decimal"`
	VolumeInMl           int           `json:"volume_in_ml,omitempty"`
	WeeklyAvailability   []interface{} `json:"weekly_availability,omitempty"`
	WeightInGrams        int           `json:"weight_in_grams,omitempty"`
	WoltPlusOnly         bool          `json:"wolt_plus_only"`
}

func NewWoltMenu(f string) (*WoltImport, error) {
	file, err := os.Open(filepath.Clean(f))
	if err != nil {
		return nil, err
	}
	defer file.Close()
	res := WoltImport{}
	byteValue, _ := io.ReadAll(file)
	err = json.Unmarshal(byteValue, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (w *WoltImport) SaveToFile(s string) error {
	//fmt.Println(w)
	file, _ := json.Marshal(w)

	err := os.WriteFile(s, file, 0644)
	return err
}

func (w *WoltImport) UpdateByExcel(ex []readexcel.Item, packagePrice float64) error {
	for _, v := range ex {
		for kk, vv := range w.Items {
			if vv.ID.Oid == v.OID {
				//fmt.Printf("\n%s (%q)\n", v.Description, v.OID)
				if w.Items[kk].Name[0].Value != v.Description {
					//	fmt.Printf("Menim polozku z %q na %q\n", w.Items[kk].Name[0].Value, v.Description)
					w.Items[kk].Name[0].Value = v.Description
				}
				if w.Items[kk].BasePrice != int(v.Price+packagePrice)*100 {
					//	fmt.Printf("Menim cenu z %d na %d\n", w.Items[kk].BasePrice/100, int(v.Price))
					w.Items[kk].BasePrice = int(v.Price+packagePrice) * 100
				}
				fmt.Printf("%s %d\n", w.Items[kk].Name[0].Value, w.Items[kk].BasePrice/100)
			}
		}
	}
	return nil
}
