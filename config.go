package main

type ConfigEntry struct {
	AppToken    string
	PromoId     string
	EventType   string
	EventOrigin string
}

var Config = map[string]ConfigEntry{
	"BIKE":  {AppToken: "d28721be-fd2d-4b45-869e-9f253b554e50", PromoId: "43e35910-c168-4634-ad4f-52fd764a843f", EventType: "", EventOrigin: ""},
	"CUBE":  {AppToken: "d1690a07-3780-4068-810f-9b5bbf2931b2", PromoId: "b4170868-cef0-424f-8eb9-be0622e8e8e3", EventType: "", EventOrigin: ""},
	"CLONE": {AppToken: "74ee0b5b-775e-4bee-974f-63e7f4d5bacb", PromoId: "fe693b26-b342-4159-8808-15e3ff7f8767", EventType: "", EventOrigin: ""},
	"TRAIN": {AppToken: "82647f43-3f87-402d-88dd-09a90025313f", PromoId: "c4480ac7-e178-4973-8061-9ed5b2e17954", EventType: "", EventOrigin: ""},
	"MERGE": {AppToken: "8d1cc2ad-e097-4b86-90ef-7a27e19fb833", PromoId: "dc128d28-c45b-411c-98ff-ac7726fbaea4", EventType: "spend-energy", EventOrigin: "undefined"},
	"TWERK": {AppToken: "61308365-9d16-4040-8bb0-2f4a4c69074c", PromoId: "61308365-9d16-4040-8bb0-2f4a4c69074c", EventType: "StartLevel", EventOrigin: "undefined"},
}
