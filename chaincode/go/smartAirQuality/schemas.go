package main

var schemas = `
{
	"API": {
        "createAsset": {
               "description": "Create an asset. One argument, a JSON encoded event. AssetID is required with zero or more writable properties. Establishes an initial asset state.",
               "properties": {
                   "args": {
                       "description": "args are JSON encoded strings",
                       "items": {
                           "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                           "properties": {
                               "assetID": {
                               		"description": "The ID of a managed asset. The resource focal point for a smart contract.",
                               		"type": "string"
                               },
                               "roomtype": {

                               },
                               "locationname": {
                               		"type": "string"
                               },
                               "latitude": {
                               		"description": "A geographical coordinate",
                               		"type": "number"
                               },
                               "longitude": {
                               		"description": "A geographical coordinate",
                               		"type": "number"
                               },
                               "ownertype": {
                               		"type": "string"
                               }
				 								
                           },
                           "required": [
                   				"assetID"
               			   ],
               			   "type": "object"
                       }, //args
                       "maxItems": 1,
                       "minItems": 1,
                       "type": "array"
                   },
                   "function": {
                       "description": "createAsset function",
                       "enum": [
                           "createAsset"
                       ],
                       "type": "string"
                   },
                   "method": "invoke"
               },

        }, //createAsset
        "updateAsset": {
               "description": "Update the state of an asset. The one argument is a JSON encoded event. AssetID is required along with one or more writable properties. Establishes the next asset state. ",
               "properties": {
                   "args": {
                       "description": "args are JSON encoded strings",
                       "items": {
                           "description": "A set of fields that constitute the writable fields in an asset's state. AssetID is mandatory along with at least one writable field. In this contract pattern, a partial state is used as an event.",
                           "properties": {
                               "assetID": {
                               		"description": "The ID of a managed asset. The resource focal point for a smart contract.",
                               		"type": "string"
                               },
                               "roomtype": {

                               },
                               "locationname": {
                               		"type": "string"
                               },
                               "latitude": {
                               		"description": "A geographical coordinate",
                               		"type": "number"
                               },
                               "longitude": {
                               		"description": "A geographical coordinate",
                               		"type": "number"
                               },
                               "ownertype": {
                               		"type": "string"
                               }
				 								
                           },
                           "required": [
                   				"assetID"
               			   ],
               			   "type": "object"
                       }, //args
                       "maxItems": 1,
                       "minItems": 1,
                       "type": "array"
                   },
                   "function": {
                       "description": "updateAsset function",
                       "enum": [
                           "updateAsset"
                       ],
                       "type": "string"
                   },
                   "method": "invoke"
               },
               "type": "object"
        } //updateAsset

	} //API

}`





