package winestro

import (
	"encoding/xml"
	"fmt"
	"html"
	"net/url"
)

type BundleItem struct {
	SKU      string  `xml:"artikel_sort_item_weinnr" json:"sku"`
	Quantity float32 `xml:"artikel_sort_item_anzahl" json:"quantity"`
}

type ProductNuance struct {
	Name  string `xml:"artikel_nuancen_item_name" json:"name"`
	Image string `xml:"artikel_nuancen_item_image" json:"image,omitempty"`
}

type ProductAward struct {
	Name  string `xml:"artikel_auszeichnungen_name" json:"name"`
	Image string `xml:"artikel_auszeichnungen_image" json:"image,omitempty"`
}

type ProductFoodPairing struct {
	Name string `xml:"artikel_speisen_name" json:"name"`
}

type Product struct {
	SKU                string               `xml:"artikel_nr" json:"sku"`
	Name               string               `xml:"artikel_name" json:"name"`
	Description        string               `xml:"artikel_beschreibung" json:"description,omitempty"`
	Vintage            string               `xml:"artikel_jahrgang" json:"vintage,omitempty"`
	Variety            string               `xml:"artikel_sorte" json:"variety,omitempty"`
	Quality            string               `xml:"artikel_qualitaet" json:"quality,omitempty"`
	Taste              string               `xml:"artikel_geschmack" json:"taste,omitempty"`
	Sugar              string               `xml:"artikel_zucker" json:"sugar,omitempty"`
	Alcohol            string               `xml:"artikel_alkohol" json:"alcohol,omitempty"`
	Acid               string               `xml:"artikel_saeure" json:"acid,omitempty"`
	VolumeLiter        float32              `xml:"artikel_liter" json:"volume_liter,omitempty"`
	WeightKg           float32              `xml:"artikel_gewicht" json:"weight,omitempty"`
	Sulfites           bool                 `xml:"artikel_sulfite" json:"sulfites,omitempty"`
	Image              string               `xml:"artikel_bild" json:"image,omitempty"`
	ImageLarge         string               `xml:"artikel_bild_big" json:"image_large,omitempty"`
	Image2             string               `xml:"artikel_bild_2" json:"image_2,omitempty"`
	Image2Large        string               `xml:"artikel_bild_big_2" json:"image_2_large,omitempty"`
	Image3             string               `xml:"artikel_bild_3" json:"image_3,omitempty"`
	Image3Large        string               `xml:"artikel_bild_big_3" json:"image_3_large,omitempty"`
	Image4             string               `xml:"artikel_bild_4" json:"image_4,omitempty"`
	Image4Large        string               `xml:"artikel_bild_big_4" json:"image_4_large,omitempty"`
	ShippingQuantity   int                  `xml:"artikel_versandmenge" json:"shipping_quantity,omitempty"`
	BundleItems        []BundleItem         `xml:"artikel_sort_items>artikel_sort_item" json:"bundle_items,omitempty"`
	Price              float32              `xml:"artikel_preis" json:"price,omitempty"`
	PriceLiter         float32              `xml:"artikel_literpreis" json:"price_liter,omitempty"`
	VAT                float32              `xml:"artikel_mwst" json:"vat,omitempty"`
	Calories           string               `xml:"artikel_brennwert" json:"calories,omitempty"`
	Protein            string               `xml:"artikel_eiweiss" json:"protein,omitempty"`
	FreeShipping       bool                 `xml:"artikel_versandfrei" json:"free_shipping,omitempty"`
	HidePriceLiter     bool                 `xml:"artikel_keinliterpreis" json:"hide_price_liter,omitempty"`
	FillWeightGram     int                  `xml:"artikel_fuellgewicht" json:"fill_weight_gram,omitempty"`
	PriceKg            float32              `xml:"artikel_kilopreis" json:"price_kg,omitempty"`
	OutOfStock         bool                 `xml:"artikel_ausgetrunken" json:"out_of_stock,omitempty"`
	APNR               string               `xml:"artikel_apnr" json:"apnr,omitempty"`
	Vineyard           string               `xml:"artikel_lage" json:"vineyard,omitempty"`
	Expertise          string               `xml:"artikel_expertise" json:"expertise,omitempty"`
	TypeName           string               `xml:"artikel_typ" json:"type_name,omitempty"`
	TypeID             int                  `xml:"artikel_typ_id" json:"type_id,omitempty"`
	Color              string               `xml:"artikel_farbe" json:"color,omitempty"`
	DrinkTemperature   string               `xml:"artikel_trinktemperatur" json:"drink_temperature,omitempty"`
	StorageTemperature string               `xml:"artikel_lagertemperatur" json:"storage_temperature,omitempty"`
	Elaboration        string               `xml:"artikel_ausbau" json:"elaboration,omitempty"`
	Soil               string               `xml:"artikel_boden" json:"soil,omitempty"`
	StorableYears      string               `xml:"artikel_lagerfaehigkeit" json:"storable_years,omitempty"`
	VideoURL           string               `xml:"artikel_video" json:"video_url,omitempty"`
	Country            string               `xml:"artikel_land" json:"country,omitempty"`
	Region             string               `xml:"artikel_region" json:"region,omitempty"`
	Appellation        string               `xml:"artikel_anbaugebiet" json:"appellation,omitempty"`
	StockWarningLevel  int                  `xml:"artikel_bestand_warnung_ab" json:"stock_warning_level,omitempty"`
	ProducerID         int                  `xml:"artikel_erzeuger" json:"producer_id,omitempty"`
	ProducerName       string               `xml:"artikel_erzeuger_name" json:"producer_name,omitempty"`
	ProducerNumber     string               `xml:"artikel_erzeuger_nr" json:"producer_number,omitempty"`
	Category           string               `xml:"artikel_kategorie" json:"category,omitempty"`
	PackagingID        int                  `xml:"artikel_verpackung" json:"packaging_id,omitempty"`
	PackagingName      string               `xml:"artikel_verpackung_bezeichnung	" json:"packaging_name,omitempty"`
	PackagingQuantity  int                  `xml:"artikel_verpackung_inhalt" json:"packaging_quantity,omitempty"`
	EAN13              string               `xml:"artikel_ean13" json:"ean13,omitempty"`
	EAN13Packaging     string               `xml:"artikel_ean13_kiste" json:"ean13_packaging,omitempty"`
	Ingredients        string               `xml:"artikel_zutaten" json:"ingredients,omitempty"`
	BestByDate         string               `xml:"artikel_mhd" json:"best_by_date,omitempty"`
	Fat                string               `xml:"artikel_fett" json:"fat,omitempty"`
	UnsaturatedFat     string               `xml:"artikel_fetts" json:"unsaturated_fat,omitempty"`
	Carbohydrates      string               `xml:"artikel_kohlenhydrate" json:"carbohydrates,omitempty"`
	Salt               string               `xml:"artikel_salz" json:"salt,omitempty"`
	Fibre              string               `xml:"artikel_ballast" json:"fibre,omitempty"`
	Vitamins           string               `xml:"artikel_vitamine" json:"vitamins,omitempty"`
	SulfuricAcids      string               `xml:"artikel_gesamt_schwefelsaeure" json:"sulfuric_acids,omitempty"`
	FreeSulfuricAcids  string               `xml:"artikel_frei_schwefelsaeure" json:"free_sulfuric_acids,omitempty"`
	Histamine          string               `xml:"artikel_histamin" json:"histamine,omitempty"`
	Glycerin           string               `xml:"artikel_glycerin" json:"glycerin,omitempty"`
	ELabelText         string               `xml:"artikel_labeltext" json:"elabel_text,omitempty"`
	ELabelURL          string               `xml:"artikel_labellink" json:"elabel_url,omitempty"`
	ProductGroups      []string             `xml:"artikel_warengruppen>warengruppe" json:"product_groups,omitempty"`
	Stock              float32              `xml:"artikel_bestand" json:"stock,omitempty"`
	CompanyStock       float32              `xml:"artikel_bestand_firmenverbund" json:"company_stock,omitempty"`
	WebshopStock       float32              `xml:"artikel_bestand_webshop" json:"webshop_stock,omitempty"`
	LastModifiedDate   timestamp            `xml:"artikel_last_modified" json:"last_modified_date,omitempty"`
	NuanceItems        []ProductNuance      `xml:"artikel_nuancen_items>artikel_nuancen_item" json:"nuance_items,omitempty"`
	AwardItems         []ProductAward       `xml:"artikel_auszeichnungen>artikel_auszeichnungen_item" json:"award_items,omitempty"`
	FoodPairingItems   []ProductFoodPairing `xml:"artikel_speisen_items>artikel_speisen_item" json:"food_pairing_items,omitempty"`
}

func (p *Product) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	type productAlias Product
	var v productAlias
	err := d.DecodeElement(&v, &start)
	if err != nil {
		return err
	}

	v.SKU, err = url.QueryUnescape(v.SKU)
	if err != nil {
		return err
	}

	v.Name = html.UnescapeString(v.Name)
	v.Fat = html.UnescapeString(v.Fat)
	v.UnsaturatedFat = html.UnescapeString(v.UnsaturatedFat)
	v.Carbohydrates = html.UnescapeString(v.Carbohydrates)
	v.Salt = html.UnescapeString(v.Salt)
	v.Fibre = html.UnescapeString(v.Fibre)
	v.Vitamins = html.UnescapeString(v.Vitamins)
	v.SulfuricAcids = html.UnescapeString(v.SulfuricAcids)
	v.FreeSulfuricAcids = html.UnescapeString(v.FreeSulfuricAcids)
	v.Histamine = html.UnescapeString(v.Histamine)
	v.Glycerin = html.UnescapeString(v.Glycerin)
	v.Protein = html.UnescapeString(v.Protein)
	v.Calories = html.UnescapeString(v.Calories)

	*p = Product(v)
	return nil
}

type ProductResp struct {
	XMLName  xml.Name  `xml:"items"`
	Products []Product `xml:"item"`
}

type ProductOptions struct {
	GroupID             int    // Optional group ID to filter products by group
	Query               string // Optional search term for product name or description
	SKU                 string // Optional search term for product SKU
	IncludeCompanyStock bool
}

func (c *Client) FetchProducts(opt ProductOptions) ([]Product, error) {
	params := map[string]string{}
	if opt.GroupID > 0 {
		params["id_grp"] = fmt.Sprintf("%d", opt.GroupID)
	}
	if opt.Query != "" {
		params["suchstring"] = opt.Query
	}
	if opt.SKU != "" {
		params["artikelnr"] = opt.SKU
	}
	if opt.IncludeCompanyStock {
		params["firmenverbund"] = "true"
	}
	var products ProductResp
	err := c.do("getArtikel", params, &products)
	return products.Products, err
}
