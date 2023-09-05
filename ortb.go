package main

type RequestInfo struct {
	Width         int
	Height        int
	TMax          int
	GoogleQueryId string

	BidFloor          float64 // global bidfloor for post call
	SupplierAuctionID string
	SupplierImpID     string
	Secure            int // whether add call is secure or non secure
	ViewabilityExt    string
	Test              string
	// PmpAuction        *PmpAuction // Added for PMP
	ChocoCookie string // Chocolate cookie id

	FinalSaleDecision int
	TransactionID     string
	PaymentChain      string

	// Banner Parameters

	BType               []int
	BAttr               []int
	TopFrame            int
	ExpandableDirection []int

	// Native Parameters
	Ver            string
	Context        int
	ContextSubType int
	PlcmtType      int
	PlcmtCnt       int
	Seq            int

	AUrlSupport int
	DUrlSupport int

	Privacy int
	API     []int

	// Source Extension Parameters
	OmIDpn string
	OmIDpv string

	// Is Publisher sending truncated IP
	IsTruncatedIPEnabled bool

	// Is Publisher Seller Enabled
	IsPublisherSellerEnabled bool
	// Is Publisher and Site/App Ads.txt enabled for Chocolateplatform.com
	IsAdstxtEnabled bool

	// Performance Metric
	PerformanceMetric map[string]float64
	NativeRequest     string

	// Cookie drop check for GDPR
	GDPRCookieFlag int

	// Supply and Advertiser Deals Mapping
	SupplyAdvDealsMap map[string]string

	Eids string //optional
	// Skadn string //optional

	// CustomEvent []*CustomEvent
	// // ExchangeInternal       *ExchangeInternal        `json:"exchangeInternal,omitempty"`       //Custom Vdopia Field
	// DFilterRejectedBidders []DFilterRejectedBidders `json:"dFilterRejectedBidders,omitempty"` //Custom Vdopia Field
	// Status                 string                   `json:"status,omitempty"`                 //Status of auction, exchange related
	// SupplyType             string                   `json:"supply_type,omitempty"`            //SupplyType of auction, exchange related
	// Debug                  string                   `json:"debug,omitempty"`                  //Debug of auction, exchange related
	// FloorPriceBucket       *FloorPriceBucket        `json:"floorpricebucket,omitempty"`       // Floor Price Bucket which applied on incoming request
	// AdvertiserIdBuyerIdMap []AdvertiserIdBuyerIdMap `json:"advertiserIdBuyerIdMap,omitempty"` // Cookie Sync Advertiser BQ Logging which applied on incoming request
	// PublisherIdMap         []PublisherIdMap         `json:"publisherIdMap,omitempty"`         // Cookie Sync Publisher BQ Logging which applied on incoming request
	// DynamicBidFloorFactor  float64                  `json:"bidfloor_factor,omitempty"`        // Dynamic bid floor factor
	// AdType                 string                   `json:"adtype,omitempty"`                 // Ad type value:- video or banner
	// IsSellerEnabled        bool                     `json:"sellerenabled,omitempty"`          // Inventory is Seller Enabled - true/false
	// SupplyCookieMatch      string                   `json:"supplycookiematch,omitempty"`
	// SupplyCookieLUT        int64                    `json:"supplycookielut,omitempty"`    // LUT = Last Update Time of Supply Cookie
	// ContextualPackages     []string                 `json:"contextualPackages,omitempty"` // array of contextual package id
}
