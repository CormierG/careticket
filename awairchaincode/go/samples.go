package main

var samples = `
{
	"event": {
		"assetID": "The ID of an AWAIR device. The resource focal point for a smart contract.",
		"co2": 504.6,
		"dust": 13.958333333333334,
		"temp": 14.628,
		"humid": 34.514,
		"voc": 489.8,
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "alert": "Sent",
        "alerttime": '"2006-01-02T15:04:05.000Z"'	
	},
	"initEvent": {
        "nickname": "AWAIR",
        "version": "The ID of an AWAIR device. The resource focal point for a smart contract."
    },
    "state": {
		"assetID": "The ID of an AWAIR device. The resource focal point for a smart contract.",
		"co2": 504.6,
		"dust": 13.958333333333334,
		"temp": 14.628,
		"humid": 34.514,
		"voc": 489.8
        "location": {
            "latitude": 123.456,
            "longitude": 123.456
        },
        "alert": "Sent",
        "alerttime": '"2006-01-02T15:04:05.000Z"'
    }
}`



/*
// AssetState stores current state for any asset
type AssetState struct {
	AssetID		*string 	    `json:assetID,omitempty"`
	Co2        	*float64	    `json:co2,omitempty"`
	Dust    	*float64	    `json:dust,omitempty"`
	Temp 		*float64        `json:temp,omitempty"`
	Humid  		*float64	    `json:humid,omitempty"`
	Voc 		*float64	    `json:voc,omitempty"`
    Location    *Geolocation    `json:"location,omitempty"`    // current asset location
    Alert       *string         `json:alert,omitempty"`
    AlertTime   *string         `json:alerttime,omitempty"`      
}
*/