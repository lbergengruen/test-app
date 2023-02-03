package main

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type PageId struct {
	Ad struct {
		AdID                  string `json:"AdID"`
		LineID                string `json:"LineID"`
		CID                   string `json:"cID"`
		AdvertiserID          string `json:"AdvertiserID"`
		HDBannerURL           string `json:"HDBannerURL"`
		FHDBannerURL          string `json:"FHDBannerURL"`
		SDBannerURL           string `json:"SDBannerURL"`
		ClickAction           string `json:"clickAction"`
		ClickID               string `json:"clickID"`
		ClickParams           string `json:"clickParams"`
		Avsource              string `json:"avsource"`
		ClickURL              string `json:"clickURL"`
		ImpressionURL         string `json:"ImpressionURL"`
		InstallURL            string `json:"installURL"`
		PlaybackURL           string `json:"playbackURL"`
		ThirdPartyImpressions string `json:"ThirdPartyImpressions"`
		ThirdPartyClicks      string `json:"ThirdPartyClicks"`
		Refreshinterval       int    `json:"refreshinterval"`
		Refreshtime           int    `json:"refreshtime"`
		HDBackgroundImageURL  string `json:"HDBackgroundImageURL"`
		FHDBackgroundImageURL string `json:"FHDBackgroundImageURL"`
		SDBackgroundImageURL  string `json:"SDBackgroundImageURL"`
	} `json:"ad"`
}

type Response struct {
	Items []struct {
		Id                   string   `json:"id"`
		Name                 string   `json:"name"`
		CampaignClass        string   `json:"campaignClass"`
		Active               bool     `json:"active"`
		PageId               string   `json:"pageId"`
		IsDmpQualified       bool     `json:"isDmpQualified"`
		TargetedDevice       string   `json:"targetedDevice"`
		TargetedDeviceModel  string   `json:"targetedDeviceModel"`
		PromotedDeviceModel  string   `json:"promotedDeviceModel"`
		IgnoreDedupeByModel  bool     `json:"ignoreDedupeByModel"`
		Restriction          string   `json:"restriction"`
		StartDateUtc         string   `json:"startDateUtc"`
		EndDateUtc           string   `json:"endDateUtc"`
		ClientTypeArray      []string `json:"clientTypeArray"`
		Compatibility        []string `json:"compatibility"`
		OfferType            string   `json:"offerType"`
		StoreFront           string   `json:"storeFront"`
		TargetedDeviceModels string   `json:"targetedDeviceModels"`
		SegmentId            string   `json:"segmentId"`
		CampaignId           string   `json:"campaignId"`

		FreeCheckout         string    `json:"freeCheckout"`
		EndDateUtcReadable   string    `json:"endDateUtcReadable"`
		StartDateUtcReadable string    `json:"startDateUtcReadable"`
		PauseDateUtc         time.Time `json:"pauseDateUtc"`
		ResumeDateUtc        time.Time `json:"resumeDateUtc"`

		CampaignName  string `json:"campaignName"`
		Localizations struct {
			EnUs struct {
				Sku                 string   `json:"sku"`
				Discount            float32  `json:"discount"`
				Details             string   `json:"details"`
				Description         string   `json:"description"`
				FreeCheckout        bool     `json:"freeCheckout"`
				WwwOfferDescription string   `json:"wwwOfferDescription"`
				Disclaimer          string   `json:"disclaimer"`
				DisclaimerDeviceV2  []string `json:"disclaimerDeviceV2"`
				ExpirationDate      string   `json:"expirationDate"`
				CouponRule          string   `json:"couponRule"`
				Header              string   `json:"header"`
				VanityUrl           string   `json:"vanityUrl"`
				BulletFooter1       string   `json:"bullet_footer1"`
				BulletFooter2       string   `json:"bullet_footer2"`
				BulletFooter3       string   `json:"bullet_footer3"`
				MaxQty              int      `json:"maxQty"`
				Microsite           struct {
					EmailParams struct {
						EmailTriggerId string `json:"emailTriggerId"`
						EmailProduct   string `json:"emailProduct"`
						IsCheetah      bool   `json:"isCheetah"`
					} `json:"emailParams"`
					ButtonsToggle struct {
						BuyNow        bool `json:"buyNow"`
						SendMeDetails struct {
							SendMeEmail bool `json:"sendMeEmail"`
							QrCode      bool `json:"qrCode"`
							Sms         bool `json:"sms"`
						} `json:"sendMeDetails"`
						LearnMore bool `json:"learnMore"`
						PlayVideo bool `json:"playVideo"`
						NoThanks  bool `json:"noThanks"`
					} `json:"buttonsToggle"`
					Title    string `json:"title"`
					VideoUrl string `json:"videoUrl"`
					Ui       struct {
						Slideshow  []string `json:"slideshow"`
						Components []struct {
							Type   string `json:"type"`
							Name   string `json:"name"`
							Left   int    `json:"left,omitempty"`
							Top    int    `json:"top,omitempty"`
							Width  int    `json:"width,omitempty"`
							Height int    `json:"height,omitempty"`
							Text   string `json:"text,omitempty"`
						} `json:"components"`
					} `json:"ui"`
					MicrositeAssets          struct{} `json:"micrositeAssets"`
					IsCloudSDKEnabled        bool     `json:"isCloudSDKEnabled"`
					ShippingMsg              string   `json:"shippingMsg"`
					SortWeight               int      `json:"sortWeight"`
					LimitedOfferMsg          string   `json:"limitedOfferMsg"`
					DisplayProductCard       string   `json:"displayProductCard"`
					Notification             string   `json:"notification"`
					DeviceProductDescription string   `json:"device_product_description"`
					FirmwareAssets           struct {
						Text                  string `json:"text"`
						HdImage               string `json:"hdImage"`
						FhdImage              string `json:"fhdImage"`
						HdBackgroundImageURL  string `json:"hdBackgroundImageURL"`
						FhdBackgroundImageURL string `json:"fhdBackgroundImageURL"`
					} `json:"firmwareAssets"`
					Assets []struct {
						Type         string `json:"type"`
						HdImage      string `json:"hdImage"`
						SdImage      string `json:"sdImage"`
						DetailsImage string `json:"detailsImage"`
					} `json:"assets"`
					Upsell []string `json:"upsell"`
				} `json:"microsite"`
			} `json:"en-us"`
		} `json:"localizations"`
		MiniHuds struct {
			EnUs []struct {
				Context     string   `json:"context"`
				Conditions  []string `json:"conditions"`
				ContextData string   `json:"contextData"`
				Assets      struct {
					Duration int `json:"duration"`
					Frames   []struct {
						Delay     int    `json:"delay"`
						Type      string `json:"type,omitempty"`
						FhdUrl    string `json:"fhdUrl,omitempty"`
						HdUrl     string `json:"hdUrl,omitempty"`
						Title     string `json:"title,omitempty"`
						Subtitle  string `json:"subtitle,omitempty"`
						OkCta     string `json:"okCta,omitempty"`
						CancelCta string `json:"cancelCta,omitempty"`
						Footer    string `json:"footer,omitempty"`
					} `json:"frames"`
					Background struct {
						FhdUrl string `json:"fhdUrl"`
						HdUrl  string `json:"hdUrl"`
					} `json:"background,omitempty"`
				} `json:"assets"`
				Type    string `json:"type"`
				OfferId string `json:"offerId"`
				Id      string `json:"id"`
				Version int    `json:"version"`
			} `json:"en-us"`
		} `json:"miniHuds"`

		IsMLfilterable bool `json:"isMLfilterable"`
		MiniAds        struct {
		} `json:"miniAds"`
		IsCompatibilityFilterable bool   `json:"isCompatibilityFilterable"`
		ClientType                string `json:"clientType"`

		TargetedDeviceBlacklistHw  string `json:"targetedDeviceBlacklistHw"`
		TargetedPlatforms          string `json:"targetedPlatforms"`
		TargetedPlatformsBlacklist string `json:"targetedPlatformsBlacklist"`
		FilterBlackListByOfferIds  string `json:"filterBlackListByOfferIds"`
		BlacklistSegmentId         string `json:"blacklistSegmentId"`
	} `json:"items"`
}

type Assets struct {
	Background *Background `json:"background"`
	Duration   *int        `json:"duration"`
	Frames     []*Frame    `json:"frames"`
}

type Background struct {
	FhdURL *string `json:"fhdUrl"`
	HdURL  *string `json:"hdUrl"`
}

type ChannelActivities struct {
	AccountID       string             `json:"accountId"`
	ChannelActivity []*ChannelActivity `json:"channelActivity"`
}

type ChannelActivity struct {
	ChannelID           int                   `json:"channelId"`
	BillingProductNames []string              `json:"billingProductNames"`
	LastNDays           int                   `json:"lastNDays"`
	Action              ChannelActivityAction `json:"action"`
}

type ChannelFilter struct {
	ChannelID           int      `json:"channelId"`
	BillingProductNames []string `json:"billingProductNames"`
}

type ChannelStatus struct {
	ChannelSubscriptions *ChannelSubscriptions `json:"channelSubscriptions"`
	ChannelTrials        *ChannelTrials        `json:"channelTrials"`
}

type ChannelStatuses struct {
	AccountID     string           `json:"accountId"`
	ChannelStatus []*ChannelStatus `json:"channelStatus"`
}

type ChannelSubscriptions struct {
	ChannelID           int      `json:"channelId"`
	BillingProductNames []string `json:"billingProductNames"`
}

type ChannelTrials struct {
	ChannelID         int      `json:"channelId"`
	TrialProductNames []string `json:"trialProductNames"`
}

type CheckedItem struct {
	CaptionsAreOn              *ItemValue `json:"captionsAreOn"`
	HasPairedSpeakers          *ItemValue `json:"hasPairedSpeakers"`
	HasConnectedSoundbar       *ItemValue `json:"hasConnectedSoundbar"`
	HasPairedSubwoofer         *ItemValue `json:"hasPairedSubwoofer"`
	Is4kDevice                 *ItemValue `json:"is4kDevice"`
	IsConnectedToRokuTv        *ItemValue `json:"isConnectedToRokuTv"`
	IsRokuTv                   *ItemValue `json:"isRokuTv"`
	IsSoundbar                 *ItemValue `json:"isSoundbar"`
	MobileAppUsedRecently      *ItemValue `json:"mobileAppUsedRecently"`
	MadeFrequentTextSearches   *ItemValue `json:"madeFrequentTextSearches"`
	HasVoiceRemote             *ItemValue `json:"hasVoiceRemote"`
	HasVoiceRemotePro          *ItemValue `json:"hasVoiceRemotePro"`
	UsedPhysicalRemoteRecently *ItemValue `json:"usedPhysicalRemoteRecently"`
	MadeFrequentVolumeChanges  *ItemValue `json:"madeFrequentVolumeChanges"`
	HadConstantHighVolume      *ItemValue `json:"hadConstantHighVolume"`
	IsConnectedTo4kTv          *ItemValue `json:"isConnectedTo4kTv"`
	IsFrequentlyUsedApp        *ItemValue `json:"isFrequentlyUsedApp"`
	HasRemoteWithAppButton     *ItemValue `json:"hasRemoteWithAppButton"`
	HasPoorWifi                *ItemValue `json:"hasPoorWifi"`
	HadFrequentRebuffers       *ItemValue `json:"hadFrequentRebuffers"`
	RebufferHudShownRecently   *ItemValue `json:"rebufferHudShownRecently"`
	RebufferHudIgnoredRecently *ItemValue `json:"rebufferHudIgnoredRecently"`
	SuppressRebufferHud        *ItemValue `json:"suppressRebufferHud"`
	OnlyTargetedDevices        *ItemValue `json:"onlyTargetedDevices"`
}

type CreateAssets struct {
	Background *CreateBackground `json:"background"`
	Duration   *int              `json:"duration"`
	Frames     []*CreateFrame    `json:"frames"`
}

type CreateBackground struct {
	FhdURL *string `json:"fhdUrl"`
	HdURL  *string `json:"hdUrl"`
}

type CreateCheckedItem struct {
	CaptionsAreOn              *ItemValue `json:"captionsAreOn"`
	HasPairedSpeakers          *ItemValue `json:"hasPairedSpeakers"`
	HasConnectedSoundbar       *ItemValue `json:"hasConnectedSoundbar"`
	HasPairedSubwoofer         *ItemValue `json:"hasPairedSubwoofer"`
	Is4kDevice                 *ItemValue `json:"is4kDevice"`
	IsConnectedToRokuTv        *ItemValue `json:"isConnectedToRokuTv"`
	IsRokuTv                   *ItemValue `json:"isRokuTv"`
	IsSoundbar                 *ItemValue `json:"isSoundbar"`
	MobileAppUsedRecently      *ItemValue `json:"mobileAppUsedRecently"`
	MadeFrequentTextSearches   *ItemValue `json:"madeFrequentTextSearches"`
	HasVoiceRemote             *ItemValue `json:"hasVoiceRemote"`
	HasVoiceRemotePro          *ItemValue `json:"hasVoiceRemotePro"`
	UsedPhysicalRemoteRecently *ItemValue `json:"usedPhysicalRemoteRecently"`
	MadeFrequentVolumeChanges  *ItemValue `json:"madeFrequentVolumeChanges"`
	HadConstantHighVolume      *ItemValue `json:"hadConstantHighVolume"`
	IsConnectedTo4kTv          *ItemValue `json:"isConnectedTo4kTv"`
	IsFrequentlyUsedApp        *ItemValue `json:"isFrequentlyUsedApp"`
	HasRemoteWithAppButton     *ItemValue `json:"hasRemoteWithAppButton"`
	HasPoorWifi                *ItemValue `json:"hasPoorWifi"`
	HadFrequentRebuffers       *ItemValue `json:"hadFrequentRebuffers"`
	RebufferHudShownRecently   *ItemValue `json:"rebufferHudShownRecently"`
	RebufferHudIgnoredRecently *ItemValue `json:"rebufferHudIgnoredRecently"`
	SuppressRebufferHud        *ItemValue `json:"suppressRebufferHud"`
	OnlyTargetedDevices        *ItemValue `json:"onlyTargetedDevices"`
}

type CreateFirmwareOfferDetail struct {
	OfferID                 *string    `json:"offerId"`
	OnDeviceCheckoutImg     *string    `json:"onDeviceCheckoutImg"`
	FhdImg                  *string    `json:"fhdImg"`
	FwAssetHdImg            *string    `json:"fwAssetHdImg"`
	FhdBackgroundImg        *string    `json:"fhdBackgroundImg"`
	HdBackgroundImg         *string    `json:"hdBackgroundImg"`
	SlideShow               *string    `json:"slideShow"`
	MicrositeDescriptions   *string    `json:"micrositeDescriptions"`
	MicrositeConfigurations *string    `json:"micrositeConfigurations"`
	MicrositeDisclaimer     *string    `json:"micrositeDisclaimer"`
	IsCloudSdkMicrosite     *bool      `json:"isCloudSdkMicrosite"`
	CloudSdkSettings        *string    `json:"cloudSdkSettings"`
	HasBuyNow               *bool      `json:"hasBuyNow"`
	HasSendMeEmail          *bool      `json:"hasSendMeEmail"`
	HasQRCode               *bool      `json:"hasQrCode"`
	HasSms                  *bool      `json:"hasSms"`
	HasLearnMore            *bool      `json:"hasLearnMore"`
	HasPlayVideo            *bool      `json:"hasPlayVideo"`
	HasNoThanks             *bool      `json:"hasNoThanks"`
	EmailTriggerID          *string    `json:"emailTriggerId"`
	Title                   *string    `json:"title"`
	VideoID                 *string    `json:"videoId"`
	Disclaimer              []*string  `json:"disclaimer"`
	CreatedBy               *string    `json:"createdBy"`
	CreatedTimestamp        *time.Time `json:"createdTimestamp"`
	LastModifiedBy          *string    `json:"lastModifiedBy"`
	LastModifiedTimestamp   *time.Time `json:"lastModifiedTimestamp"`
}

type CreateFrame struct {
	FhdURL    *string `json:"fhdUrl"`
	HdURL     *string `json:"hdUrl"`
	Delay     *int    `json:"delay"`
	Type      *string `json:"type"`
	Title     *string `json:"title"`
	Subtitle  *string `json:"subtitle"`
	OkCta     *string `json:"okCta"`
	CancelCta *string `json:"cancelCta"`
	Footer    *string `json:"footer"`
}

type CreateHudOffer struct {
	OfferID               *string            `json:"offerId"`
	ContextType           *string            `json:"contextType"`
	HudType               *string            `json:"hudType"`
	ContextData           *string            `json:"contextData"`
	CheckedItem           *CreateCheckedItem `json:"checkedItem"`
	Assets                *CreateAssets      `json:"assets"`
	CreatedBy             *string            `json:"createdBy"`
	CreatedTimestamp      *time.Time         `json:"createdTimestamp"`
	LastModifiedBy        *string            `json:"lastModifiedBy"`
	LastModifiedTimestamp *time.Time         `json:"lastModifiedTimestamp"`
}

type CreateOfferDetail struct {
	OfferName                         *string    `json:"offerName"`
	OfferModel                        *string    `json:"offerModel"`
	OfferDeviceName                   *string    `json:"offerDeviceName"`
	AdBannerFlag                      *bool      `json:"adBannerFlag"`
	IsFreeCheckoutEnabled             *bool      `json:"isFreeCheckoutEnabled"`
	TargetedDevice                    *string    `json:"targetedDevice"`
	TargetedBlacklistModels           *string    `json:"targetedBlacklistModels"`
	IsDedupeEnabled                   *bool      `json:"isDedupeEnabled"`
	IsFilteringByAbTestEnabled        *bool      `json:"isFilteringByAbTestEnabled"`
	IsFilteringByCompatibilityEnabled *bool      `json:"isFilteringByCompatibilityEnabled"`
	IsFilteringByMlEnabled            *bool      `json:"isFilteringByMlEnabled"`
	Compatibility                     *string    `json:"compatibility"`
	Notes                             *string    `json:"notes"`
	Discount                          *float32   `json:"discount"`
	CouponRule                        *string    `json:"couponRule"`
	UpsellItems                       *string    `json:"upsellItems"`
	RedemptionRule                    *string    `json:"redemptionRule"`
	JiraTicket                        *string    `json:"jiraTicket"`
	CreatedBy                         *string    `json:"createdBy"`
	CreatedTimestamp                  *time.Time `json:"createdTimestamp"`
	LastModifiedBy                    *string    `json:"lastModifiedBy"`
	LastModifiedTimestamp             *time.Time `json:"lastModifiedTimestamp"`
}

type CreateWebOffer struct {
	OfferID                   *string    `json:"offerId"`
	LegalDisclaimer           *string    `json:"legalDisclaimer"`
	VanityURL                 *string    `json:"vanityUrl"`
	OfferDisplayOnTop         *bool      `json:"offerDisplayOnTop"`
	OfferDisplayOnUpgradePage *bool      `json:"offerDisplayOnUpgradePage"`
	Notification              *string    `json:"notification"`
	DetailsBulletFooter       *string    `json:"detailsBulletFooter"`
	DetailsBulletFooter1      *string    `json:"detailsBulletFooter1"`
	DetailsBulletFooter2      *string    `json:"detailsBulletFooter2"`
	HdImage                   *string    `json:"hdImage"`
	Descriptions              *string    `json:"descriptions"`
	ProdDescDetail            *string    `json:"prodDescDetail"`
	WwwOfferDesc              *string    `json:"wwwOfferDesc"`
	OfferCardNote             *string    `json:"offerCardNote"`
	MaxQuantity               *int       `json:"maxQuantity"`
	ExpirationDate            *time.Time `json:"expirationDate"`
	ShippingMessage           *string    `json:"shippingMessage"`
	DetailModalImage          *string    `json:"detailModalImage"`
	CreatedBy                 *string    `json:"createdBy"`
	CreatedTimestamp          *time.Time `json:"createdTimestamp"`
	LastModifiedBy            *string    `json:"lastModifiedBy"`
	LastModifiedTimestamp     *time.Time `json:"lastModifiedTimestamp"`
}

type FirmwareOfferDetail struct {
	FirmwareOfferID         *string    `json:"firmwareOfferId"`
	OnDeviceCheckoutImg     *string    `json:"onDeviceCheckoutImg"`
	FhdImg                  *string    `json:"fhdImg"`
	FwAssetHdImg            *string    `json:"fwAssetHdImg"`
	FhdBackgroundImg        *string    `json:"fhdBackgroundImg"`
	HdBackgroundImg         *string    `json:"hdBackgroundImg"`
	SlideShow               *string    `json:"slideShow"`
	MicrositeDescriptions   *string    `json:"micrositeDescriptions"`
	MicrositeConfigurations *string    `json:"micrositeConfigurations"`
	MicrositeDisclaimer     *string    `json:"micrositeDisclaimer"`
	IsCloudSdkMicrosite     *bool      `json:"isCloudSdkMicrosite"`
	CloudSdkSettings        *string    `json:"cloudSdkSettings"`
	HasBuyNow               *bool      `json:"hasBuyNow"`
	HasSendMeEmail          *bool      `json:"hasSendMeEmail"`
	HasQRCode               *bool      `json:"hasQrCode"`
	HasSms                  *bool      `json:"hasSms"`
	HasLearnMore            *bool      `json:"hasLearnMore"`
	HasPlayVideo            *bool      `json:"hasPlayVideo"`
	HasNoThanks             *bool      `json:"hasNoThanks"`
	EmailTriggerID          *string    `json:"emailTriggerId"`
	Title                   *string    `json:"title"`
	VideoID                 *string    `json:"videoId"`
	Disclaimer              []*string  `json:"disclaimer"`
	CreatedBy               *string    `json:"createdBy"`
	CreatedTimestamp        *time.Time `json:"createdTimestamp"`
	LastModifiedBy          *string    `json:"lastModifiedBy"`
	LastModifiedTimestamp   *time.Time `json:"lastModifiedTimestamp"`
}

type Frame struct {
	FhdURL    *string `json:"fhdUrl"`
	HdURL     *string `json:"hdUrl"`
	Delay     *int    `json:"delay"`
	Type      *string `json:"type"`
	Title     *string `json:"title"`
	Subtitle  *string `json:"subtitle"`
	OkCta     *string `json:"okCta"`
	CancelCta *string `json:"cancelCta"`
	Footer    *string `json:"footer"`
}

type HudOffer struct {
	HudOfferID            *string      `json:"hudOfferId"`
	OfferID               *string      `json:"offerId"`
	ContextType           *string      `json:"contextType"`
	HudType               *string      `json:"hudType"`
	ContextData           *string      `json:"contextData"`
	CheckedItem           *CheckedItem `json:"checkedItem"`
	Assets                *Assets      `json:"assets"`
	Deleted               *bool        `json:"deleted"`
	CreatedBy             *string      `json:"createdBy"`
	CreatedTimestamp      *time.Time   `json:"createdTimestamp"`
	LastModifiedBy        *string      `json:"lastModifiedBy"`
	LastModifiedTimestamp *time.Time   `json:"lastModifiedTimestamp"`
}

type OfferDetail struct {
	OfferID                           *string                `json:"offerId"`
	OfferName                         *string                `json:"offerName"`
	OfferModel                        *string                `json:"offerModel"`
	OfferDeviceName                   *string                `json:"offerDeviceName"`
	AdBannerFlag                      *bool                  `json:"adBannerFlag"`
	IsFreeCheckoutEnabled             *bool                  `json:"isFreeCheckoutEnabled"`
	TargetedDevice                    *string                `json:"targetedDevice"`
	TargetedBlacklistModels           *string                `json:"targetedBlacklistModels"`
	IsDedupeEnabled                   *bool                  `json:"isDedupeEnabled"`
	IsFilteringByAbTestEnabled        *bool                  `json:"isFilteringByAbTestEnabled"`
	IsFilteringByCompatibilityEnabled *bool                  `json:"isFilteringByCompatibilityEnabled"`
	IsFilteringByMlEnabled            *bool                  `json:"isFilteringByMlEnabled"`
	Compatibility                     *string                `json:"compatibility"`
	Notes                             *string                `json:"notes"`
	Discount                          *float32               `json:"discount"`
	CouponRule                        *string                `json:"couponRule"`
	UpsellItems                       *string                `json:"upsellItems"`
	RedemptionRule                    *string                `json:"redemptionRule"`
	JiraTicket                        *string                `json:"jiraTicket"`
	FirmwareOfferDetail               []*FirmwareOfferDetail `json:"firmwareOfferDetail"`
	WebOffer                          []*WebOffer            `json:"webOffer"`
	HudOfferDetail                    []*HudOffer            `json:"hudOfferDetail"`
	CreatedBy                         *string                `json:"createdBy"`
	CreatedTimestamp                  *time.Time             `json:"createdTimestamp"`
	LastModifiedBy                    *string                `json:"lastModifiedBy"`
	LastModifiedTimestamp             *time.Time             `json:"lastModifiedTimestamp"`
}

type UpsertFirmwareOfferOutput struct {
	OfferID         *string    `json:"offerId"`
	FirmwareOfferID *string    `json:"firmwareOfferId"`
	CreatedBy       *string    `json:"createdBy"`
	CreatedTime     *time.Time `json:"createdTime"`
	LastUpdatedBy   *string    `json:"lastUpdatedBy"`
	LastUpdatedTime *time.Time `json:"lastUpdatedTime"`
}

type UpsertHudOfferOutput struct {
	HudOfferID      *string      `json:"hudOfferId"`
	OfferID         *string      `json:"offerId"`
	ContextType     *string      `json:"contextType"`
	HudType         *string      `json:"hudType"`
	ContextData     *string      `json:"contextData"`
	CheckedItem     *CheckedItem `json:"checkedItem"`
	Assets          *Assets      `json:"assets"`
	CreatedBy       *string      `json:"createdBy"`
	CreatedTime     *time.Time   `json:"createdTime"`
	LastUpdatedBy   *string      `json:"lastUpdatedBy"`
	LastUpdatedTime *time.Time   `json:"lastUpdatedTime"`
}

type UpsertOfferOutput struct {
	OfferID         *string    `json:"offerId"`
	OfferName       *string    `json:"offerName"`
	CreatedBy       *string    `json:"createdBy"`
	CreatedTime     *time.Time `json:"createdTime"`
	LastUpdatedBy   *string    `json:"lastUpdatedBy"`
	LastUpdatedTime *time.Time `json:"lastUpdatedTime"`
}

type UpsertWebOfferOutput struct {
	OfferID         *string    `json:"offerId"`
	WebOfferID      *string    `json:"webOfferId"`
	CreatedBy       *string    `json:"createdBy"`
	CreatedTime     *time.Time `json:"createdTime"`
	LastUpdatedBy   *string    `json:"lastUpdatedBy"`
	LastUpdatedTime *time.Time `json:"lastUpdatedTime"`
}

type WebOffer struct {
	WebOfferID                *string    `json:"webOfferId"`
	OfferID                   *string    `json:"offerId"`
	VersionID                 *string    `json:"versionId"`
	LegalDisclaimer           *string    `json:"legalDisclaimer"`
	VanityURL                 *string    `json:"vanityUrl"`
	OfferDisplayOnTop         *bool      `json:"offerDisplayOnTop"`
	OfferDisplayOnUpgradePage *bool      `json:"offerDisplayOnUpgradePage"`
	Notification              *string    `json:"notification"`
	DetailsBulletFooter       *string    `json:"detailsBulletFooter"`
	DetailsBulletFooter1      *string    `json:"detailsBulletFooter1"`
	DetailsBulletFooter2      *string    `json:"detailsBulletFooter2"`
	HdImage                   *string    `json:"hdImage"`
	Descriptions              *string    `json:"descriptions"`
	ProdDescDetail            *string    `json:"prodDescDetail"`
	WwwOfferDesc              *string    `json:"wwwOfferDesc"`
	OfferCardNote             *string    `json:"offerCardNote"`
	MaxQuantity               *int       `json:"maxQuantity"`
	ExpirationDate            *time.Time `json:"expirationDate"`
	ShippingMessage           *string    `json:"shippingMessage"`
	DetailModalImage          *string    `json:"detailModalImage"`
	CreatedBy                 *string    `json:"createdBy"`
	CreatedTimestamp          *time.Time `json:"createdTimestamp"`
	LastModifiedBy            *string    `json:"lastModifiedBy"`
	LastModifiedTimestamp     *time.Time `json:"lastModifiedTimestamp"`
}

type ChannelActivityAction string

const (
	// Installed channel
	ChannelActivityActionInstall ChannelActivityAction = "INSTALL"
	// Uninstalled channel
	ChannelActivityActionUninstall ChannelActivityAction = "UNINSTALL"
)

var AllChannelActivityAction = []ChannelActivityAction{
	ChannelActivityActionInstall,
	ChannelActivityActionUninstall,
}

func (e ChannelActivityAction) IsValid() bool {
	switch e {
	case ChannelActivityActionInstall, ChannelActivityActionUninstall:
		return true
	}
	return false
}

func (e ChannelActivityAction) String() string {
	return string(e)
}

func (e *ChannelActivityAction) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ChannelActivityAction(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ChannelActivityAction", str)
	}
	return nil
}

func (e ChannelActivityAction) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ChannelStatusType string

const (
	// Not currently in subscription, but was in subscription in the past
	ChannelStatusTypePastSubscription7Day  ChannelStatusType = "PAST_SUBSCRIPTION_7_DAY"
	ChannelStatusTypePastSubscription30Day ChannelStatusType = "PAST_SUBSCRIPTION_30_DAY"
	ChannelStatusTypePastSubscription90Day ChannelStatusType = "PAST_SUBSCRIPTION_90_DAY"
	// Currently active in subsctiption of channel
	ChannelStatusTypeSubscription ChannelStatusType = "SUBSCRIPTION"
	// Currently active in trial of channel
	ChannelStatusTypeTrial ChannelStatusType = "TRIAL"
)

var AllChannelStatusType = []ChannelStatusType{
	ChannelStatusTypePastSubscription7Day,
	ChannelStatusTypePastSubscription30Day,
	ChannelStatusTypePastSubscription90Day,
	ChannelStatusTypeSubscription,
	ChannelStatusTypeTrial,
}

func (e ChannelStatusType) IsValid() bool {
	switch e {
	case ChannelStatusTypePastSubscription7Day, ChannelStatusTypePastSubscription30Day, ChannelStatusTypePastSubscription90Day, ChannelStatusTypeSubscription, ChannelStatusTypeTrial:
		return true
	}
	return false
}

func (e ChannelStatusType) String() string {
	return string(e)
}

func (e *ChannelStatusType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ChannelStatusType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ChannelStatusType", str)
	}
	return nil
}

func (e ChannelStatusType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type ItemValue string

const (
	ItemValueTrue  ItemValue = "True"
	ItemValueFalse ItemValue = "False"
	ItemValueNa    ItemValue = "NA"
)

var AllItemValue = []ItemValue{
	ItemValueTrue,
	ItemValueFalse,
	ItemValueNa,
}

func (e ItemValue) IsValid() bool {
	switch e {
	case ItemValueTrue, ItemValueFalse, ItemValueNa:
		return true
	}
	return false
}

func (e ItemValue) String() string {
	return string(e)
}

func (e *ItemValue) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = ItemValue(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid ItemValue", str)
	}
	return nil
}

func (e ItemValue) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OfferType string

const (
	OfferTypeNone OfferType = "NONE"
	OfferTypeDtc  OfferType = "DTC"
	OfferTypeTrc  OfferType = "TRC"
)

var AllOfferType = []OfferType{
	OfferTypeNone,
	OfferTypeDtc,
	OfferTypeTrc,
}

func (e OfferType) IsValid() bool {
	switch e {
	case OfferTypeNone, OfferTypeDtc, OfferTypeTrc:
		return true
	}
	return false
}

func (e OfferType) String() string {
	return string(e)
}

func (e *OfferType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OfferType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OfferType", str)
	}
	return nil
}

func (e OfferType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
