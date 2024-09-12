package main

import "time"

type ConfigEntry struct {
	AppToken    string
	PromoId     string
	EventType   string
	EventOrigin string
	Sleep       time.Duration
}

var Config = map[string]ConfigEntry{
	"BIKE":  {AppToken: "d28721be-fd2d-4b45-869e-9f253b554e50", PromoId: "43e35910-c168-4634-ad4f-52fd764a843f", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"CUBE":  {AppToken: "d1690a07-3780-4068-810f-9b5bbf2931b2", PromoId: "b4170868-cef0-424f-8eb9-be0622e8e8e3", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"CLONE": {AppToken: "74ee0b5b-775e-4bee-974f-63e7f4d5bacb", PromoId: "fe693b26-b342-4159-8808-15e3ff7f8767", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"TRAIN": {AppToken: "82647f43-3f87-402d-88dd-09a90025313f", PromoId: "c4480ac7-e178-4973-8061-9ed5b2e17954", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"MERGE": {AppToken: "8d1cc2ad-e097-4b86-90ef-7a27e19fb833", PromoId: "dc128d28-c45b-411c-98ff-ac7726fbaea4", EventType: "spend-energy", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"TWERK": {AppToken: "61308365-9d16-4040-8bb0-2f4a4c69074c", PromoId: "61308365-9d16-4040-8bb0-2f4a4c69074c", EventType: "StartLevel", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"POLY":  {AppToken: "2aaf5aee-2cbc-47ec-8a3f-0962cc14bc71", PromoId: "2aaf5aee-2cbc-47ec-8a3f-0962cc14bc71", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"TRIM":  {AppToken: "ef319a80-949a-492e-8ee0-424fb5fc20a6", PromoId: "ef319a80-949a-492e-8ee0-424fb5fc20a6", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"RACE":  {AppToken: "8814a785-97fb-4177-9193-ca4180ff9da8", PromoId: "8814a785-97fb-4177-9193-ca4180ff9da8", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"CAFE":  {AppToken: "bc0971b8-04df-4e72-8a3e-ec4dc663cd11", PromoId: "bc0971b8-04df-4e72-8a3e-ec4dc663cd11", EventType: "5visitorsChecks", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"GANGS": {AppToken: "b6de60a0-e030-48bb-a551-548372493523", PromoId: "c7821fa7-6632-482c-9635-2bd5798585f9", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"ZOO":   {AppToken: "b2436c89-e0aa-4aed-8046-9b0515e1c46b", PromoId: "b2436c89-e0aa-4aed-8046-9b0515e1c46b", EventType: "ZoopolisEvent", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"FLUF":  {AppToken: "112887b0-a8af-4eb2-ac63-d82df78283d9", PromoId: "112887b0-a8af-4eb2-ac63-d82df78283d9", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"TILE":  {AppToken: "e68b39d2-4880-4a31-b3aa-0393e7df10c7", PromoId: "e68b39d2-4880-4a31-b3aa-0393e7df10c7", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"STONE": {AppToken: "04ebd6de-69b7-43d1-9c4b-04a6ca3305af", PromoId: "04ebd6de-69b7-43d1-9c4b-04a6ca3305af", EventType: "undefined", EventOrigin: "undefined", Sleep: 20 * time.Second},
	"BOUNC": {AppToken: "bc72d3b9-8e91-4884-9c33-f72482f0db37", PromoId: "bc72d3b9-8e91-4884-9c33-f72482f0db37", EventType: "test", EventOrigin: "undefined", Sleep: 20 * time.Second},
}
