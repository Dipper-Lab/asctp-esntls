package consts

const PurchaseFeatures = "PURCHASE"
const ReceiveFarmBags = "RECEIVE-FB"
const ReceiveZeroFlyBags = "RECEIVE-ZFB"
const BatchFarmBags = "BATCH-FB"
const BatchZeroFlyBags = "BATCH-ZFB"
const AttachRecordedValues = "ATTACH-RV"
const Reports = "REPORTS"
const RegisterNFC = "REGISTER-NFC"
const TransportInitiate = "TRANSPORT-INIT"
const TransportReceive = "TRANSPORT-RECEIVE"

const PurchaseView = "PURCHASE-VIEW"
const BatchView = "BATCH-VIEW"
const FarmBagsView = "FB-VIEW"
const ZeroFlyBagView = "ZFB-VIEW"
const ReportsView = "REPORTS-VIEW"
const TransportView = "TRANSPORT-VIEW"

var RBACMap = map[string][]string{
	"ADMIN": {
		PurchaseFeatures,
		ReceiveFarmBags,
		ReceiveZeroFlyBags,
		BatchFarmBags,
		BatchZeroFlyBags,
		AttachRecordedValues,
		Reports,
		RegisterNFC,
		TransportInitiate,
		TransportReceive,
		BatchView,
		FarmBagsView,
		PurchaseView,
		ZeroFlyBagView,
		ReportsView,
		TransportView,
	},
	"FIELD-OFFICER": {
		PurchaseFeatures,
		ReceiveFarmBags,
		BatchFarmBags,
		BatchZeroFlyBags,
		AttachRecordedValues,
		TransportInitiate,
		BatchView,
		FarmBagsView,
		PurchaseView,
		ZeroFlyBagView,
		TransportView,
	},
	"ATTENDANT": {
		ReceiveZeroFlyBags,
		TransportInitiate,
		TransportReceive,
		Reports,
	},
	"": {
		PurchaseFeatures,
		ReceiveFarmBags,
		ReceiveZeroFlyBags,
		BatchFarmBags,
		BatchZeroFlyBags,
		AttachRecordedValues,
		Reports,
		RegisterNFC,
		TransportInitiate,
		TransportReceive,
		BatchView,
		FarmBagsView,
		PurchaseView,
		ZeroFlyBagView,
		ReportsView,
		TransportView,
	},
}
