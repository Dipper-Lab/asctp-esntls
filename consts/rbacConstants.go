package consts

var roles = [4]string{"ADMIN", "FIELD-OFFICER", "ATTENDANT", ""}

//feature lists
/*
~ Purchasing ->  FO
~ Receive farmbags ->  FO
~ Receive zerofly bags -> Attendant
~ batch farmbags ->  FO
~ Attach recorded values ->  FO
~ attach batch to zerofly bags  ->  FO
~ Report sale and defect -> Attendant
~ Register NFC bag -> ADMIN
~ Initiate transport -> FO, Attendant
~ receive transport -> Attendant
*/

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

var RBACMap = map[string][]string{
	"ADMIN":         {PurchaseFeatures, ReceiveFarmBags, ReceiveZeroFlyBags, BatchFarmBags, BatchZeroFlyBags, AttachRecordedValues, Reports, RegisterNFC, TransportInitiate, TransportReceive},
	"FIELD-OFFICER": {PurchaseFeatures, ReceiveFarmBags, BatchFarmBags, BatchZeroFlyBags, AttachRecordedValues, TransportInitiate},
	"ATTENDANT":     {ReceiveZeroFlyBags, TransportInitiate, TransportReceive, Reports},
	"":              {PurchaseFeatures, ReceiveFarmBags, ReceiveZeroFlyBags, BatchFarmBags, BatchZeroFlyBags, AttachRecordedValues, Reports, RegisterNFC, TransportInitiate, TransportReceive},
}
