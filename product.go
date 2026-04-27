package winestro

import (
	"encoding/xml"
	"fmt"
	"html"
	"net/url"
	"time"
)

type timestamp time.Time

func (t *timestamp) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s string
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}
	parsedTime, err := time.Parse("2006-01-02 15:04:05", s)
	if err != nil {
		return err
	}
	*t = timestamp(parsedTime)
	return nil
}

type BundleItem struct {
	SKU      string  `xml:"artikel_sort_item_weinnr" json:"sku"`
	Quantity float32 `xml:"artikel_sort_item_anzahl" json:"quantity"`
}

type ProductNuance struct {
	Name  string `xml:"artikel_nuancen_item_name" json:"name"`
	Image string `xml:"artikel_nuancen_item_image" json:"image"`
}

type ProductAward struct {
	Name  string `xml:"artikel_auszeichnungen_name" json:"name"`
	Image string `xml:"artikel_auszeichnungen_image" json:"image"`
}

type ProductFoodPairing struct {
	Food string `xml:"artikel_essen" json:"food"`
}

type Product struct {
	SKU                string               `xml:"artikel_nr" json:"sku"`
	Name               string               `xml:"artikel_name" json:"name"`
	Description        string               `xml:"artikel_beschreibung" json:"description"`
	Vintage            string               `xml:"artikel_jahrgang" json:"vintage"`
	Variety            string               `xml:"artikel_sorte" json:"variety"`
	Quality            string               `xml:"artikel_qualitaet" json:"quality"`
	Taste              string               `xml:"artikel_geschmack" json:"taste"`
	Sugar              string               `xml:"artikel_zucker" json:"sugar"`
	Alcohol            string               `xml:"artikel_alkohol" json:"alcohol"`
	Acid               string               `xml:"artikel_saeure" json:"acid"`
	VolumeLiter        float32              `xml:"artikel_liter" json:"volume_liter"`
	WeightKg           float32              `xml:"artikel_gewicht" json:"weight"`
	Sulfites           bool                 `xml:"artikel_sulfite" json:"sulfites"`
	Image              string               `xml:"artikel_bild" json:"image"`
	ImageLarge         string               `xml:"artikel_bild_big" json:"image_large"`
	Image2             string               `xml:"artikel_bild_2" json:"image_2"`
	Image2Large        string               `xml:"artikel_bild_big_2" json:"image_2_large"`
	Image3             string               `xml:"artikel_bild_3" json:"image_3"`
	Image3Large        string               `xml:"artikel_bild_big_3" json:"image_3_large"`
	Image4             string               `xml:"artikel_bild_4" json:"image_4"`
	Image4Large        string               `xml:"artikel_bild_big_4" json:"image_4_large"`
	ShippingQuantity   int                  `xml:"artikel_versandmenge" json:"shipping_quantity"`
	BundleItems        []BundleItem         `xml:"artikel_sort_items>artikel_sort_item" json:"bundle_items"`
	Price              float32              `xml:"artikel_preis" json:"price"`
	PriceLiter         float32              `xml:"artikel_literpreis" json:"price_liter"`
	VAT                float32              `xml:"artikel_mwst" json:"vat"`
	Calories           string               `xml:"artikel_brennwert" json:"calories"`
	Protein            string               `xml:"artikel_eiweiss" json:"protein"`
	FreeShipping       bool                 `xml:"artikel_versandfrei" json:"free_shipping"`
	HidePriceLiter     bool                 `xml:"artikel_keinliterpreis" json:"hide_price_liter"`
	FillWeightGram     int                  `xml:"artikel_fuellgewicht" json:"fill_weight_gram"`
	PriceKg            float32              `xml:"artikel_kilopreis" json:"price_kg"`
	OutOfStock         bool                 `xml:"artikel_ausgetrunken" json:"out_of_stock"`
	APNR               string               `xml:"artikel_apnr" json:"apnr"`
	Vineyard           string               `xml:"artikel_lage" json:"vineyardn"`
	Expertise          string               `xml:"artikel_expertise" json:"expertise"`
	TypeName           string               `xml:"artikel_typ" json:"type_name"`
	TypeID             int                  `xml:"artikel_typ_id" json:"type_id"`
	Color              string               `xml:"artikel_farbe" json:"color"`
	DrinkTemperature   string               `xml:"artikel_trinktemperatur" json:"drink_temperature"`
	StorageTemperature string               `xml:"artikel_lagertemperatur" json:"storage_temperature"`
	Elaboration        string               `xml:"artikel_ausbau" json:"elaboration"`
	Soil               string               `xml:"artikel_boden" json:"soil"`
	StorableYears      string               `xml:"artikel_lagerfaehigkeit" json:"storable_years"`
	VideoURL           string               `xml:"artikel_video" json:"video_url"`
	Country            string               `xml:"artikel_land" json:"country"`
	Region             string               `xml:"artikel_region" json:"region"`
	Appellation        string               `xml:"artikel_anbaugebiet" json:"appellation"`
	StockWarningLevel  int                  `xml:"artikel_bestand_warnung_ab" json:"stock_warning_level"`
	ProducerID         int                  `xml:"artikel_erzeuger" json:"producer_id"`
	ProducerName       string               `xml:"artikel_erzeuger_name" json:"producer_name"`
	ProducerNumber     string               `xml:"artikel_erzeuger_nr" json:"producer_number"`
	Category           string               `xml:"artikel_kategorie" json:"category"`
	PackagingID        int                  `xml:"artikel_verpackung" json:"packaging_id"`
	PackagingName      string               `xml:"artikel_verpackung_bezeichnung	" json:"packaging_name"`
	PackagingQuantity  int                  `xml:"artikel_verpackung_inhalt" json:"packaging_quantity"`
	EAN13              string               `xml:"artikel_ean13" json:"ean13"`
	EAN13Packaging     string               `xml:"artikel_ean13_kiste" json:"ean13_packaging"`
	Ingredients        string               `xml:"artikel_zutaten" json:"ingredients"`
	BestByDate         string               `xml:"artikel_mhd" json:"best_by_date"`
	Fat                string               `xml:"artikel_fett" json:"fat"`
	UnsaturatedFat     string               `xml:"artikel_fetts" json:"unsaturated_fat"`
	Carbohydrates      string               `xml:"artikel_kohlenhydrate" json:"carbohydrates"`
	Salt               string               `xml:"artikel_salz" json:"salt"`
	Fibre              string               `xml:"artikel_ballast" json:"fibre"`
	Vitamins           string               `xml:"artikel_vitamine" json:"vitamins"`
	SulfuricAcids      string               `xml:"artikel_gesamt_schwefelsaeure" json:"sulfuric_acids"`
	FreeSulfuricAcids  string               `xml:"artikel_frei_schwefelsaeure" json:"free_sulfuric_acids"`
	Histamine          string               `xml:"artikel_histamin" json:"histamine"`
	Glycerin           string               `xml:"artikel_glycerin" json:"glycerin"`
	ELabelText         string               `xml:"artikel_labeltext" json:"elabel_text"`
	ELabelURL          string               `xml:"artikel_labellink" json:"elabel_url"`
	ProductGroups      []string             `xml:"artikel_warengruppen>warengruppe" json:"product_groups"`
	Stock              float32              `xml:"artikel_bestand" json:"stock"`
	CompanyStock       float32              `xml:"artikel_bestand_firmenverbund" json:"company_stock"`
	WebshopStock       float32              `xml:"artikel_bestand_webshop" json:"webshop_stock"`
	LastModifiedDate   timestamp            `xml:"artikel_last_modified" json:"last_modified_date"`
	NuanceItems        []ProductNuance      `xml:"artikel_nuancen_items>artikel_nuancen_item" json:"nuance_items"`
	AwardItems         []ProductAward       `xml:"artikel_auszeichnungen>artikel_auszeichnungen_item" json:"award_items"`
	FoodPairingItems   []ProductFoodPairing `xml:"artikel_speisen_items>artikel_speisen_item" json:"food_pairing_items"`
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
