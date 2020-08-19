package model

const DefaultEsSize = 10000
const ComapnyEsIndex = "tyc_company"

type Point struct {
	Latitude  float64 `json:"lat,omitempty"`
	Longitude float64 `json:"lon,omitempty"`
}

type OfficeS struct {
	Area        string  `json:"area,omitempty"`
	Total       int32   `json:"total,omitempty"`
	CompanyName string  `json:"companyName,omitempty"`
	Cid         int64   `json:"cid,omitempty"`
	Score       float32 `json:"score,omitempty"`
	State       string  `json:"state,omitempty"`
}

type LegalInfo struct {
	Name            string    `json:"name,omitempty"`
	Hid             int64     `json:"hid,omitempty"`
	HeadUrl         string    `json:"headUrl,omitempty"`
	Introduction    string    `json:"introduction,omitempty"`
	Event           string    `json:"event,omitempty"`
	BossCertificate int32     `json:"bossCertificate,omitempty"`
	CompanyNum      int32     `json:"companyNum,omitempty"`
	Office          []OfficeS `json:"office,omitempty"`
	Companys        string    `json:"companys,omitempty"`
	PartnerNum      int32     `json:"partnerNum,omitempty"`
	Partners        string    `json:"partners,omitempty"`
	Cid             int64     `json:"cid,omitempty"`
	TypeJoin        string    `json:"typeJoin,omitempty"`
	Alias           string    `json:"alias,omitempty"`
	Pid             int64     `json:"pid,omitempty"`
	Role            string    `json:"role,omitempty"`
}

type TagS struct {
	Type int8   `json:"type"`
	Name string `json:"name"`
}

type CapitalS struct {
	Amomon  string `json:"amomon,omitempty"`
	Time    string `json:"time,omitempty"`
	Percent string `json:"percent,omitempty"`
	Paymet  string `json:"paymet,omitempty"`
}

type HolderS struct {
	Toco    int8       `json:"toco,omitempty"`
	Id      int64      `json:"id,omitempty"`
	Name    string     `json:"name,omitempty"`
	TagList []TagS     `json:"tagList,omitempty"`
	Type    int8       `json:"type,omitempty"`
	Capital []CapitalS `json:"capital,omitempty"`
}

type StuffS struct {
	Toco     int8     `json:"toco,omitempty"`
	Id       int64    `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	TypeSore string   `json:"typeSore,omitempty"`
	Logo     string   `json:"logo,omitempty"`
	Score    int32    `json:"score,omitempty"`
	TypeJoin []string `json:"typeJoin,omitempty"`
	Type     int8     `json:"type,omitempty"`
}

type CompanyInfo struct {
	CompanyId              int64     `json:"companyd,omitempty"`
	Name                   string    `json:"name,omitempty"`
	Alias                  string    `json:"alias,omitempty"`
	LegalPersonId          int64     `json:"legalPersonId,omitempty"`
	LegalPersonName        string    `json:"legalPersonName,omitempty"`
	LegalPersonShowStr     string    `json:"legalPersonShowStr,omitempty"`
	CompanyOrgType         string    `json:"companyOrgType,omitempty"`
	CompanyType            int8      `json:"companyType,omitempty"`
	Industry               string    `json:"industry,omitempty"`
	EstablishTime          string    `json:"establishTime,omitempty"`
	Province               string    `json:"province,omitempty"`
	City                   string    `json:"city,omitempty"`
	District               string    `json:"district,omitempty"`
	Address                string    `json:"address,omitempty"`
	Round                  string    `json:"round,omitempty"`
	BondName               string    `json:"bondName,omitempty"`
	BondNum                string    `json:"bondNum,omitempty"`
	BondType               string    `json:"bondType,omitempty"`
	Phone                  string    `json:"phone,omitempty"`
	Telephone              string    `json:"telephone,omitempty"`
	PhoneList              string    `json:"phoneList,omitempty"`
	Email                  string    `json:"email,omitempty"`
	Emails                 string    `json:"emails,omitempty"`
	Fax                    string    `json:"fax,omitempty"`
	Postcode               string    `json:"postcode,omitempty"`
	RegStatus              string    `json:"regStatus,omitempty"`
	RegTime                string    `json:"regTime,omitempty"`
	ApprovedTime           int64     `json:"approvedTime,omitempty"`
	RegNumber              string    `json:"regNumber,omitempty"`
	RegInstitute           string    `json:"regInstitute,omitempty"`
	RegLocation            string    `json:"regLocation,omitempty"`
	RegCapital             string    `json:"regCapital,omitempty"`
	RegCapitalCurrency     string    `json:"regCapitalCurrency,omitempty"`
	ActualCapital          string    `json:"actualCapital,omitempty"`
	ActualCapitalCurrency  string    `json:"actualCapitalCurrency,omitempty"`
	CreditCode             string    `json:"creditCode,omitempty"`
	TaxCode                string    `json:"taxCode,omitempty"`
	Logo                   string    `json:"logo,omitempty"`
	EquityUrl              string    `json:"equityUrl,omitempty"`
	StaffNumRange          string    `json:"staffNumRange,omitempty"`
	SocialSecurityStaffNum int32     `json:"socialSecurityStaff_num,omitempty"`
	SocialStaffNum         int32     `json:"socialStaffNum,omitempty"`
	IsBranch               int8      `json:"isBranch,omitempty"`
	IsClaimed              int8      `json:"isClaimed,omitempty"`
	IsHightTech            int8      `json:"isHightTech,omitempty"`
	IsMicroEnt             int8      `json:"isMicroEnt,omitempty"`
	InverstStatus          string    `json:"inverstStatus,omitempty"`
	P2p                    string    `json:"p2p,omitempty"`
	Score                  string    `json:"score,omitempty"`
	BaseInfo               string    `json:"baseInfo,omitempty"`
	BusinessScope          string    `json:"businessScope,omitempty"`
	Staff                  []StuffS  `json:"staff,omitempty"`
	Holder                 []HolderS `json:"holder,omitempty"`
	LegalInfo              LegalInfo `json:"legalInfo,omitempty"`
	Updatetime             int64     `json:"updatetime,omitempty"`
	GeoHash                string    `json:"geo_hash,omitempty"`
	GeoLatLon              Point     `json:"geo_lat_lon,omitempty"`
}
