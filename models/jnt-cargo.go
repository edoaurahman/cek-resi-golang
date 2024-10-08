package models

type JntCargoModel struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data []struct {
		Keyword string `json:"keyword"`
		Details []struct {
			Change                int      `json:"change"`
			Index                 int      `json:"index"`
			BillCode              string   `json:"billCode"`
			ScanTime              string   `json:"scanTime"`
			ScanTypeName          string   `json:"scanTypeName"`
			ScanNetworkName       string   `json:"scanNetworkName"`
			ScanNetworkID         int      `json:"scanNetworkId"`
			ScanByName            string   `json:"scanByName"`
			StaffName             string   `json:"staffName,omitempty"`
			StaffContact          string   `json:"staffContact,omitempty"`
			Remark1               string   `json:"remark1,omitempty"`
			Remark2               string   `json:"remark2,omitempty"`
			CustomerTracking      string   `json:"customerTracking"`
			Code                  int      `json:"code"`
			Status                string   `json:"status"`
			ScanNetworkCode       string   `json:"scanNetworkCode"`
			ScanNetworkCity       string   `json:"scanNetworkCity"`
			ScanNetworkProvince   string   `json:"scanNetworkProvince"`
			ScanNetworkTypeID     string   `json:"scanNetworkTypeId"`
			ScanNetworkTypeName   string   `json:"scanNetworkTypeName"`
			ScanNetworkContact    string   `json:"scanNetworkContact"`
			ScanNetworkArea       string   `json:"scanNetworkArea"`
			DeliveryTelephone     string   `json:"deliveryTelephone,omitempty"`
			UploadTime            string   `json:"uploadTime"`
			PicURL                []string `json:"picUrl,omitempty"`
			NextStopName          string   `json:"nextStopName,omitempty"`
			Length                string   `json:"length,omitempty"`
			Width                 string   `json:"width,omitempty"`
			High                  string   `json:"high,omitempty"`
			TaskCode              string   `json:"taskCode,omitempty"`
			NextStopTypeID        string   `json:"nextStopTypeId,omitempty"`
			NextStopTypeName      string   `json:"nextStopTypeName,omitempty"`
			NextNetworkID         string   `json:"nextNetworkId,omitempty"`
			NextNetworkCode       string   `json:"nextNetworkCode,omitempty"`
			NextNetworkProvinceID int      `json:"nextNetworkProvinceId,omitempty"`
			NextNetworkCityID     int      `json:"nextNetworkCityId,omitempty"`
			NextNetworkAreaID     int      `json:"nextNetworkAreaId,omitempty"`
			Remark4               string   `json:"remark4,omitempty"`
			Weight                string   `json:"weight,omitempty"`
		} `json:"details"`
		ExpressTypeName           string      `json:"expressTypeName"`
		PackageTotalWeight        float64     `json:"packageTotalWeight"`
		PackageTotalVolume        float64     `json:"packageTotalVolume"`
		CollectTime               string      `json:"collectTime"`
		SenderName                interface{} `json:"senderName"`
		SenderMobilePhone         interface{} `json:"senderMobilePhone"`
		SenderTelphone            interface{} `json:"senderTelphone"`
		SenderPostalCode          interface{} `json:"senderPostalCode"`
		SenderCountryID           interface{} `json:"senderCountryId"`
		SenderCountryName         interface{} `json:"senderCountryName"`
		SenderProvinceID          interface{} `json:"senderProvinceId"`
		SenderProvinceName        interface{} `json:"senderProvinceName"`
		SenderCityID              interface{} `json:"senderCityId"`
		SenderCityName            string      `json:"senderCityName"`
		SenderAreaID              interface{} `json:"senderAreaId"`
		SenderAreaName            interface{} `json:"senderAreaName"`
		SenderTownship            interface{} `json:"senderTownship"`
		SenderStreet              interface{} `json:"senderStreet"`
		SenderDetailedAddress     interface{} `json:"senderDetailedAddress"`
		ReceiverName              interface{} `json:"receiverName"`
		ReceiverMobilePhone       interface{} `json:"receiverMobilePhone"`
		ReceiverTelphone          interface{} `json:"receiverTelphone"`
		ReceiverPostalCode        interface{} `json:"receiverPostalCode"`
		ReceiverCountryID         interface{} `json:"receiverCountryId"`
		ReceiverCountryName       interface{} `json:"receiverCountryName"`
		ReceiverProvinceID        interface{} `json:"receiverProvinceId"`
		ReceiverProvinceName      interface{} `json:"receiverProvinceName"`
		ReceiverCityID            interface{} `json:"receiverCityId"`
		ReceiverCityName          string      `json:"receiverCityName"`
		ReceiverAreaID            interface{} `json:"receiverAreaId"`
		ReceiverAreaName          interface{} `json:"receiverAreaName"`
		ReceiverTownship          interface{} `json:"receiverTownship"`
		ReceiverStreet            interface{} `json:"receiverStreet"`
		ReceiverDetailedAddress   interface{} `json:"receiverDetailedAddress"`
		InsuredAmount             interface{} `json:"insuredAmount"`
		InsuredFee                interface{} `json:"insuredFee"`
		GoodsTypeCode             interface{} `json:"goodsTypeCode"`
		GoodsTypeName             interface{} `json:"goodsTypeName"`
		GoodsName                 interface{} `json:"goodsName"`
		PackageNumber             int         `json:"packageNumber"`
		ExpressTypeCode           interface{} `json:"expressTypeCode"`
		DispatchCode              interface{} `json:"dispatchCode"`
		DispatchName              interface{} `json:"dispatchName"`
		PaidModeCode              interface{} `json:"paidModeCode"`
		PaidModeName              interface{} `json:"paidModeName"`
		PackageChargeWeight       interface{} `json:"packageChargeWeight"`
		PackageType               interface{} `json:"packageType"`
		PackageTypeName           interface{} `json:"packageTypeName"`
		DispatchStaffCode         interface{} `json:"dispatchStaffCode"`
		DispatchStaffMobile       interface{} `json:"dispatchStaffMobile"`
		DispatchStaffName         interface{} `json:"dispatchStaffName"`
		CollectStaffCode          interface{} `json:"collectStaffCode"`
		CollectStaffCodeMobile    interface{} `json:"collectStaffCodeMobile"`
		CollectStaffName          interface{} `json:"collectStaffName"`
		DestinationID             interface{} `json:"destinationId"`
		DestinationCode           interface{} `json:"destinationCode"`
		DestinationName           interface{} `json:"destinationName"`
		OriginID                  interface{} `json:"originId"`
		OriginCode                interface{} `json:"originCode"`
		OriginName                interface{} `json:"originName"`
		DispatchNetworkID         interface{} `json:"dispatchNetworkId"`
		DispatchNetworkCode       interface{} `json:"dispatchNetworkCode"`
		DispatchNetworkCodeMobile interface{} `json:"dispatchNetworkCodeMobile"`
		DispatchNetworkName       interface{} `json:"dispatchNetworkName"`
		PickNetworkID             interface{} `json:"pickNetworkId"`
		PickNetworkCode           interface{} `json:"pickNetworkCode"`
		PickNetworkCodeMobile     interface{} `json:"pickNetworkCodeMobile"`
		PickNetworkName           interface{} `json:"pickNetworkName"`
		WaybillStatusCode         interface{} `json:"waybillStatusCode"`
		DispatchTime              string      `json:"dispatchTime"`
		DamageFlag                bool        `json:"damageFlag"`
		LostFlag                  bool        `json:"lostFlag"`
		ClaimFlag                 bool        `json:"claimFlag"`
		SendName                  string      `json:"sendName"`
		SendCode                  string      `json:"sendCode"`
	} `json:"data"`
	Fail bool `json:"fail"`
	Succ bool `json:"succ"`
}
