package winestro

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

type Customer struct {
	ID               int32   `xml:"adr_id" json:"id"`
	Number           string  `xml:"adr_nr" json:"number"`
	FirstName        string  `xml:"adr_vorname" json:"first_name,omitempty"`
	LastName         string  `xml:"adr_nachname" json:"last_name,omitempty"`
	Company          string  `xml:"adr_firma" json:"company,omitempty"`
	ZIP              string  `xml:"adr_plz" json:"zip,omitempty"`
	City             string  `xml:"adr_ort" json:"city,omitempty"`
	WWW              string  `xml:"adr_www" json:"www,omitempty"`
	Email            string  `xml:"adr_email" json:"email,omitempty"`
	Street           string  `xml:"adr_str" json:"street,omitempty"`
	HouseNum         string  `xml:"adr_str_nr" json:"house_number,omitempty"`
	Country          string  `xml:"adr_land" json:"country,omitempty"`
	Landline         string  `xml:"adr_festnetz" json:"landline,omitempty"`
	Mobile           string  `xml:"adr_mobil" json:"mobile,omitempty"`
	Fax              string  `xml:"adr_fax" json:"fax,omitempty"`
	Note1            string  `xml:"adr_note1" json:"note1,omitempty"`
	Note2            string  `xml:"adr_note2" json:"note2,omitempty"`
	Note3            string  `xml:"adr_note3" json:"note3,omitempty"`
	Note4            string  `xml:"adr_note4" json:"note4,omitempty"`
	Discount         float64 `xml:"adr_rabatt" json:"discount,omitempty"`
	PriceCategory    int     `xml:"adr_id_preiskategorie" json:"price_category,omitempty"`
	NewsletterActive bool    `xml:"adr_newsletter_aktiv" json:"newsletter_active,omitempty"`
	TaxType          int     `xml:"adr_kunden_mwst" json:"tax_type,omitempty"`
	Salutation       string  `xml:"adr_anrede" json:"salutation,omitempty"`
	SalutationType   int     `xml:"adr_anredenart" json:"salutation_type,omitempty"`
	PaymentType      int     `xml:"adr_id_zahlungsart" json:"payment_type,omitempty"`
}

func (c Customer) FullName() string {
	if c.FirstName == "" && c.LastName == "" {
		return c.Company
	}
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

func (c Customer) FullSalutation() string {
	re := regexp.MustCompile(`,\s*$`)
	trimSaluation := strings.TrimSpace(re.ReplaceAllString(c.Salutation, ""))
	switch c.SalutationType {
	case 0:
		return fmt.Sprintf("%s %s", trimSaluation, strings.TrimSpace(c.FirstName))
	case 1:
		return fmt.Sprintf("%s %s", trimSaluation, strings.TrimSpace(c.LastName))
	default:
		return trimSaluation
	}
}

type CustomersResp struct {
	XMLName   xml.Name   `xml:"items"`
	Customers []Customer `xml:"item"`
}

func (c *Client) FetchCustomersForGroup(groupID int) ([]Customer, error) {
	params := map[string]string{
		"id_grp": fmt.Sprintf("%d", groupID),
	}
	var customers CustomersResp
	err := c.do("getKundenGruppe", params, &customers)
	return customers.Customers, err
}
