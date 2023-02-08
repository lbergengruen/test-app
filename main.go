package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func GetWonkaOffers(url string) Response {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	// Parse Jsons to Structs
	var r Response
	err = json.Unmarshal(cleanJson(body), &r)
	if err != nil {
		log.Fatalln(err)
	}
	return r
}

func cleanJson(data []byte) []byte {
	asString := string(data)
	asString = strings.Replace(asString, "\"expirationDate\": \"none\"", "\"expirationDate\": \"\"", -1)
	asString = strings.Replace(asString, "\"expirationDate\": \"\"", "\"expirationDate\": null", -1)
	asString = strings.Replace(asString, "\"disclaimerDeviceV2\": \"\"", "\"disclaimerDeviceV2\": []", -1)
	asString = strings.Replace(asString, "\"left\": \"189\"", "\"left\": 189", -1)
	asString = strings.Replace(asString, "\"width\": \"1166\"", "\"width\": 1166", -1)
	asString = strings.Replace(asString, "\"height\": \"275\"", "\"height\": 275", -1)
	asString = strings.Replace(asString, "\"false\"", "false", -1)
	asString = strings.Replace(asString, "\"true\"", "true", -1)
	asString = strings.Replace(asString, "\"background\": \"none\"", "\"background\": {}", -1)
	asString = strings.Replace(asString, "\\\"background\\\":\\\"none\\\"", "\\\"background\\\":{}", -1)
	//asString = strings.Replace(asString, "\"pageId\": \"{\\n\\\"\\\"ad\\\"\\\": {\\n\\\"\\\"AdID\\\"\\\": \\\"\\\"138305732180\\\"\\\",\\n\\\"\\\"LineID\\\"\\\": \\\"\\\"0\\\"\\\",\\n\\\"\\\"AdvertiserID\\\"\\\": \\\"\\\"0\\\"\\\",\\n\\\"\\\"HDBannerURL\\\"\\\": \\\"\\\"\\\"\\\",\\n\\\"\\\"FHDBannerURL\\\"\\\": \\\"\\\"https://tpc.googlesyndication.com/pagead/imgad?id=CICAgKCXzPe7KhABGAEyCIfKdyRcV4rX\\\"\\\",\\n\\\"\\\"SDBannerURL\\\"\\\": \\\"\\\"\\\"\\\",\\n\\\"\\\"clickAction\\\"\\\": \\\"\\\"DaVinci\\\"\\\",\\n\\\"\\\"clickID\\\"\\\": \\\"\\\"0\\\"\\\",\\n\\\"\\\"clickParams\\\"\\\": \\\"\\\"contentid=rpes-soundbar-marpromo2ndtarget-en-current;mediatype=page;smtptriggerid=31719;offerid=c1f8d563bfd301b5d4373137aba7bc1b\\\"\\\",\\n\\\"\\\"avsource\\\"\\\": \\\"\\\"\\\"\\\",\\n\\\"\\\"clickURL\\\"\\\": \\\"\\\"https://adclick.g.doubleclick.net/pcs/click%253Fxai%253DAKAOjstIStjsdBOIX8zdvqZmO6QUT5Xl9MASHu_x_c9WdOJTxAsLw1rkSHIU5GeJwkrNQU-QqrLCOypuRK0a6vvxAfgu7tWQERqdPaRDwjTXwxr_j7kEm1cAJ-2FAVrRw15po-j2Htfwz3a93AwSCGl4xMCeg2CwH9mU5-ngx2judwGeR9ccbdK258qd6Sef%2526sig%253DCg0ArKJSzCluG4GPzhbPEAE%2526urlfix%253D1%2526adurl%253D\\\"\\\",\\n\\\"\\\"ImpressionURL\\\"\\\": \\\"\\\"\\\"\\\",\\n\\\"\\\"installURL\\\"\\\": \\\"\\\"https://pubads.g.doubleclick.net/gampad/adx?sz=1x1&iu=/82114269/install_tracking&c=1557701928&dpt=1&d_imp=1&d_imp_hdr=1&t=lineitemid%253D0%2526creativeid%253D138305732180%2526adguid%253D\\\"\\\",\\n\\\"\\\"playbackURL\\\"\\\": \\\"\\\"\\\"\\\",\\n\\\"\\\"ThirdPartyImpressions\\\"\\\": \\\"\\\"https://ravm.tv/pixel/display/v1?meType=0&evType=0&adv_id=0&cID=0&plID=0&crID=138305732180&sw_version=&dev_model=&ott_id=&ip=&adguid=&last_channel=&bnr_loc=&locale=&is_lat=&country_code=&trc_version=&trc_channel_version=&grandcentral_version=&davinci_version=&ad_srv=&libversion=&platform=&app=&clientversion=&trc_genre_row=&trc_category_id=&abrejectcount=&guest_mode=&theme=&providerproductids=&adunitpath=/82114269&adunitID=81114389&width=350&height=490&screensaver=&ria_adid=&ria_altid=\\\"\\\",\\n\\\"\\\"ThirdPartyClicks\\\"\\\":\\\"\\\"\\\"\\\",\\n\\\"\\\"refreshinterval\\\"\\\":0,\\n\\\"\\\"refreshtime\\\"\\\":0,\\n\\\"\\\"HDBackgroundImageURL\\\"\\\": \\\"\\\"\\\"\\\",\\n\\\"\\\"FHDBackgroundImageURL\\\"\\\": \\\"\\\"\\\"\\\",\\n\\\"\\\"SDBackgroundImageURL\\\"\\\": \\\"\\\"\\\"\\\"\\n}\\n}\",",
	//	"\"pageId\": \"{\\\"ad\\\": {\\\"AdID\\\": \\\"138305732180\\\",\\\"LineID\\\": \\\"0\\\",\\\"AdvertiserID\\\": \\\"0\\\",\\\"HDBannerURL\\\": \\\"\\\",\\\"FHDBannerURL\\\": \\\"https://tpc.googlesyndication.com/pagead/imgad?id=CICAgKCXzPe7KhABGAEyCIfKdyRcV4rX\\\",\\\"SDBannerURL\\\": \\\"\\\",\\\"clickAction\\\": \\\"DaVinci\\\",\\\"clickID\\\": \\\"0\\\",\\\"clickParams\\\": \\\"contentid=rpes-soundbar-marpromo2ndtarget-en-current;mediatype=page;smtptriggerid=31719;offerid=c1f8d563bfd301b5d4373137aba7bc1b\\\",\\\"avsource\\\": \\\"\\\",\\\"clickURL\\\": \\\"https://adclick.g.doubleclick.net/pcs/click%253Fxai%253DAKAOjstIStjsdBOIX8zdvqZmO6QUT5Xl9MASHu_x_c9WdOJTxAsLw1rkSHIU5GeJwkrNQU-QqrLCOypuRK0a6vvxAfgu7tWQERqdPaRDwjTXwxr_j7kEm1cAJ-2FAVrRw15po-j2Htfwz3a93AwSCGl4xMCeg2CwH9mU5-ngx2judwGeR9ccbdK258qd6Sef%2526sig%253DCg0ArKJSzCluG4GPzhbPEAE%2526urlfix%253D1%2526adurl%253D\\\",\\\"ImpressionURL\\\": \\\"\\\",\\\"installURL\\\": \\\"https://pubads.g.doubleclick.net/gampad/adx?sz=1x1&iu=/82114269/install_tracking&c=1557701928&dpt=1&d_imp=1&d_imp_hdr=1&t=lineitemid%253D0%2526creativeid%253D138305732180%2526adguid%253D\\\",\\\"playbackURL\\\": \\\"\\\",\\\"ThirdPartyImpressions\\\": \\\"https://ravm.tv/pixel/display/v1?meType=0&evType=0&adv_id=0&cID=0&plID=0&crID=138305732180&sw_version=&dev_model=&ott_id=&ip=&adguid=&last_channel=&bnr_loc=&locale=&is_lat=&country_code=&trc_version=&trc_channel_version=&grandcentral_version=&davinci_version=&ad_srv=&libversion=&platform=&app=&clientversion=&trc_genre_row=&trc_category_id=&abrejectcount=&guest_mode=&theme=&providerproductids=&adunitpath=/82114269&adunitID=81114389&width=350&height=490&screensaver=&ria_adid=&ria_altid=\\\",\\\"ThirdPartyClicks\\\":\\\"\\\",\\\"refreshinterval\\\":0,\\\"refreshtime\\\":0,\\\"HDBackgroundImageURL\\\": \\\"\\\",\\\"FHDBackgroundImageURL\\\": \\\"\\\",\\\"SDBackgroundImageURL\\\": \\\"\\\"}}\",",
	//	-1)
	//asString = strings.Replace(asString, "\\\"smtptriggerid=3940", "smtptriggerid=3940", -1)

	return []byte(asString)
}

func main() {
	// Request Get Offers from Wonka
	r := GetWonkaOffers("https://wonka.web.roku.com/api/v1/offers?fetchAllOffers=true")

	// Map into Wonka's offers into Offersvs' offers
	var offers []OfferDetail
	for _, item := range r.Items {
		fmt.Println(item.Id)

		var detailsImage, slideshow, hdImage, microsite string
		var upsell, compatibility, disclaimerDeviceV2 []*string
		var sortWeight bool

		if len(item.Localizations.EnUs.Assets) == 1 {
			detailsImage = item.Localizations.EnUs.Assets[0].DetailsImage
			hdImage = item.Localizations.EnUs.Assets[0].HdImage
		} else if len(item.Localizations.EnUs.Assets) > 1 {
			log.Fatalln("More than one! Assets")
		}

		byteSlideshow, err := json.Marshal(item.Localizations.EnUs.Microsite.Ui.Slideshow)
		if err != nil {
			log.Println(err)
		}
		slideshow = string(byteSlideshow)

		for _, object := range item.Localizations.EnUs.DisclaimerDeviceV2 {
			disclaimerDeviceV2 = append(disclaimerDeviceV2, &object)
		}

		for _, object := range item.Compatibility {
			compatibility = append(compatibility, &object)
		}

		for _, object := range item.Localizations.EnUs.Upsell {
			upsell = append(upsell, &object)
		}

		byteMicrosite, err := json.Marshal(item.Localizations.EnUs.Microsite)
		if err != nil {
			log.Println(err)
		}
		microsite = string(byteMicrosite)

		if item.Localizations.EnUs.SortWeight == 1 {
			sortWeight = true
		}

		var expirationDate time.Time
		if item.Localizations.EnUs.ExpirationDate != "" {
			expirationDate, _ = time.Parse("DD/MM/YYYY", item.Localizations.EnUs.ExpirationDate)
		}

		firmwareOffersDetails := []*FirmwareOfferDetail{{
			//FirmwareOfferID:         _PLACEHOLDER_STRING_, // Created by Offersvs API
			OnDeviceCheckoutImg:     &detailsImage,
			FhdImg:                  &item.Localizations.EnUs.FirmwareAssets.FhdImage,
			FwAssetHdImg:            &item.Localizations.EnUs.FirmwareAssets.HdImage,
			FhdBackgroundImg:        &item.Localizations.EnUs.FirmwareAssets.FhdBackgroundImageURL,
			HdBackgroundImg:         &item.Localizations.EnUs.FirmwareAssets.HdBackgroundImageURL,
			SlideShow:               &slideshow,
			MicrositeDescriptions:   &item.Localizations.EnUs.DeviceProductDescription,
			MicrositeConfigurations: &item.PageId,
			MicrositeDisclaimer:     disclaimerDeviceV2,
			IsCloudSdkMicrosite:     &item.Localizations.EnUs.IsCloudSDKEnabled,
			CloudSdkSettings:        &microsite,
			HasBuyNow:               &item.Localizations.EnUs.Microsite.ButtonsToggle.BuyNow,
			HasSendMeEmail:          &item.Localizations.EnUs.Microsite.ButtonsToggle.SendMeDetails.SendMeEmail,
			HasQRCode:               &item.Localizations.EnUs.Microsite.ButtonsToggle.SendMeDetails.QrCode,
			HasSms:                  &item.Localizations.EnUs.Microsite.ButtonsToggle.SendMeDetails.Sms,
			HasLearnMore:            &item.Localizations.EnUs.Microsite.ButtonsToggle.LearnMore,
			HasPlayVideo:            &item.Localizations.EnUs.Microsite.ButtonsToggle.PlayVideo,
			HasNoThanks:             &item.Localizations.EnUs.Microsite.ButtonsToggle.NoThanks,
			EmailTriggerID:          &item.Localizations.EnUs.Microsite.EmailParams.EmailTriggerId,
			Title:                   &item.Localizations.EnUs.Microsite.Title,
			VideoID:                 &item.Localizations.EnUs.Microsite.VideoUrl,
			Disclaimer:              []*string{&item.Localizations.EnUs.Disclaimer},
		}}

		webOffers := []*WebOffer{{
			//WebOfferID:                _PLACEHOLDER_STRING_, // Created by OfferSVS API
			OfferID: &item.Id,
			// VersionID:                 _PLACEHOLDER_STRING_, // Not used in Wonka
			LegalDisclaimer:           &item.Localizations.EnUs.Disclaimer,
			VanityURL:                 &item.Localizations.EnUs.VanityUrl,
			OfferDisplayOnTop:         &sortWeight,
			OfferDisplayOnUpgradePage: &sortWeight,
			Notification:              &item.Localizations.EnUs.Notification,
			DetailsBulletFooter:       &item.Localizations.EnUs.BulletFooter1,
			DetailsBulletFooter1:      &item.Localizations.EnUs.BulletFooter2,
			DetailsBulletFooter2:      &item.Localizations.EnUs.BulletFooter3,
			HdImage:                   &hdImage,
			Descriptions:              &item.Localizations.EnUs.Description,
			ProdDescDetail:            &item.Localizations.EnUs.DeviceProductDescription,
			WwwOfferDesc:              &item.Localizations.EnUs.WwwOfferDescription,
			OfferCardNote:             &item.Localizations.EnUs.LimitedOfferMsg,
			MaxQuantity:               &item.Localizations.EnUs.MaxQty,
			ExpirationDate:            &expirationDate,
			ShippingMessage:           &item.Localizations.EnUs.ShippingMsg,
			DetailModalImage:          &detailsImage,
		}}

		var miniHuds MiniHuds
		if item.MiniHuds != "none" {
			err := json.Unmarshal([]byte(item.MiniHuds), &miniHuds)
			if err != nil {
				log.Fatalln(err)
			}
		}

		var hudOfferDetails []*HudOffer
		for _, hudOffer := range miniHuds.EnUs {

			var frames []*Frame
			for _, frame := range hudOffer.Assets.Frames {
				frames = append(frames, &Frame{
					FhdURL:    &frame.FhdUrl,
					HdURL:     &frame.HdUrl,
					Delay:     &frame.Delay,
					Type:      &frame.Type,
					Title:     &frame.Title,
					Subtitle:  &frame.Subtitle,
					OkCta:     &frame.OkCta,
					CancelCta: &frame.CancelCta,
					Footer:    &frame.Footer,
				})
			}

			NaItemValue := ItemValue("NA")
			conditions := map[string]*ItemValue{
				"captions-are-on":               &NaItemValue,
				"has-paired-speakers":           &NaItemValue,
				"has-connected-soundbar":        &NaItemValue,
				"has-paired-subwoofer":          &NaItemValue,
				"is-4k-device":                  &NaItemValue,
				"is-connected-to-roku-tv":       &NaItemValue,
				"is-roku-tv":                    &NaItemValue,
				"is-soundbar":                   &NaItemValue,
				"mobile-app-used-recently":      &NaItemValue,
				"made-frequent-text-searches":   &NaItemValue,
				"has-voice-remote":              &NaItemValue,
				"has-voice-remote-pro":          &NaItemValue,
				"used-physical-remote-recently": &NaItemValue,
				"made-frequent-volume-changes":  &NaItemValue,
				"had-constant-high-volume":      &NaItemValue,
				"is-connected-to-4k-tv":         &NaItemValue,
				"is-frequently-used-app":        &NaItemValue,
				"has-remote-with-app-button":    &NaItemValue,
				"has-poor-wifi":                 &NaItemValue,
				"had-frequent-rebuffers":        &NaItemValue,
				"rebuffer-hud-shown-recently":   &NaItemValue,
				"rebuffer-hud-ignore-recently":  &NaItemValue,
				"suppress-rebuffer-hud":         &NaItemValue,
				"only-targeted-devices":         &NaItemValue,
			}

			for _, condition := range hudOffer.Conditions {
				keyValue := strings.Split(condition, "=")

				value := ItemValue(strings.ToUpper(keyValue[1]))
				conditions[keyValue[0]] = &value
			}

			hudOfferDetails = append(hudOfferDetails, &HudOffer{
				//HudOfferID:  _PLACEHOLDER_STRING_, // Created by Offersvs API
				OfferID:     &item.Id,
				ContextType: &hudOffer.Context,
				HudType:     &hudOffer.Type,
				ContextData: &hudOffer.ContextData,
				CheckedItem: &CheckedItem{
					CaptionsAreOn:              conditions["captions-are-on"],
					HasPairedSpeakers:          conditions["has-paired-speakers"],
					HasConnectedSoundbar:       conditions["has-connected-soundbar"],
					HasPairedSubwoofer:         conditions["has-paired-subwoofer"],
					Is4kDevice:                 conditions["is-4k-device"],
					IsConnectedToRokuTv:        conditions["is-connected-to-roku-tv"],
					IsRokuTv:                   conditions["is-roku-tv"],
					IsSoundbar:                 conditions["is-soundbar"],
					MobileAppUsedRecently:      conditions["mobile-app-used-recently"],
					MadeFrequentTextSearches:   conditions["made-frequent-text-searches"],
					HasVoiceRemote:             conditions["has-voice-remote"],
					HasVoiceRemotePro:          conditions["has-voice-remote-pro"],
					UsedPhysicalRemoteRecently: conditions["used-physical-remote-recently"],
					MadeFrequentVolumeChanges:  conditions["made-frequent-volume-changes"],
					HadConstantHighVolume:      conditions["had-constant-high-volume"],
					IsConnectedTo4kTv:          conditions["is-connected-to-4k-tv"],
					IsFrequentlyUsedApp:        conditions["is-frequently-used-app"],
					HasRemoteWithAppButton:     conditions["has-remote-with-app-button"],
					HasPoorWifi:                conditions["has-poor-wifi"],
					HadFrequentRebuffers:       conditions["had-frequent-rebuffers"],
					RebufferHudShownRecently:   conditions["rebuffer-hud-shown-recently"],
					RebufferHudIgnoredRecently: conditions["rebuffer-hud-ignore-recently"],
					SuppressRebufferHud:        conditions["suppress-rebuffer-hud"],
					OnlyTargetedDevices:        conditions["only-targeted-devices"],
				},
				Assets: &Assets{
					Background: &Background{
						FhdURL: &hudOffer.Assets.Background.FhdUrl,
						HdURL:  &hudOffer.Assets.Background.HdUrl,
					},
					Duration: &hudOffer.Assets.Duration,
					Frames:   frames,
				},
			})
		}

		offerDetail := OfferDetail{
			OfferID:         &item.Id,
			OfferName:       &item.Name,
			OfferModel:      &item.PromotedDeviceModel,
			OfferDeviceName: &item.Localizations.EnUs.Header,
			//AdBannerFlag:                      _PLACEHOLDER_BOOL_, // No prop in Wonka for indicating it's an AdBanner
			IsFreeCheckoutEnabled:             &item.Localizations.EnUs.FreeCheckout,
			TargetedDevice:                    &item.TargetedDevice,
			TargetedBlacklistModels:           &item.TargetedDeviceBlacklistHw,
			IsDedupeEnabled:                   &item.IgnoreDedupeByModel,
			IsFilteringByAbTestEnabled:        &item.IsDmpQualified,
			IsFilteringByCompatibilityEnabled: &item.IsCompatibilityFilterable,
			IsFilteringByMlEnabled:            &item.IsMLfilterable,
			Compatibility:                     compatibility,
			//Notes:                             _PLACEHOLDER_STRING_, // Stored on Campaign Object, should be stored on CMS
			Discount:       &item.Localizations.EnUs.Discount,
			CouponRule:     &item.Localizations.EnUs.CouponRule,
			UpsellItems:    upsell,
			RedemptionRule: &item.Restriction,
			//JiraTicket:                        _PLACEHOLDER_STRING_, // Not stored on Wonka Offer. it's a campaign-level setting on Comet
			FirmwareOfferDetail: firmwareOffersDetails,
			WebOffer:            webOffers,
			HudOfferDetail:      hudOfferDetails,
		}

		offers = append(offers, offerDetail)
	}
	fmt.Println(len(offers))

	// Write on Offersvs Database
	//https://offersvs-qa.us-east-1.dev.msc.rokulabs.net/
}
