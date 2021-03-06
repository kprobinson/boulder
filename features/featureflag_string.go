// Code generated by "stringer -type=FeatureFlag"; DO NOT EDIT.

package features

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[unused-0]
	_ = x[PrecertificateRevocation-1]
	_ = x[CAAValidationMethods-2]
	_ = x[CAAAccountURI-3]
	_ = x[EnforceMultiVA-4]
	_ = x[MultiVAFullResults-5]
	_ = x[MandatoryPOSTAsGET-6]
	_ = x[AllowV1Registration-7]
	_ = x[V1DisableNewValidations-8]
	_ = x[StripDefaultSchemePort-9]
	_ = x[StoreIssuerInfo-10]
	_ = x[StoreRevokerInfo-11]
	_ = x[RestrictRSAKeySizes-12]
	_ = x[FasterNewOrdersRateLimit-13]
	_ = x[NonCFSSLSigner-14]
	_ = x[ECDSAForAll-15]
}

const _FeatureFlag_name = "unusedPrecertificateRevocationCAAValidationMethodsCAAAccountURIEnforceMultiVAMultiVAFullResultsMandatoryPOSTAsGETAllowV1RegistrationV1DisableNewValidationsStripDefaultSchemePortStoreIssuerInfoStoreRevokerInfoRestrictRSAKeySizesFasterNewOrdersRateLimitNonCFSSLSignerECDSAForAll"

var _FeatureFlag_index = [...]uint16{0, 6, 30, 50, 63, 77, 95, 113, 132, 155, 177, 192, 208, 227, 251, 265, 276}

func (i FeatureFlag) String() string {
	if i < 0 || i >= FeatureFlag(len(_FeatureFlag_index)-1) {
		return "FeatureFlag(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _FeatureFlag_name[_FeatureFlag_index[i]:_FeatureFlag_index[i+1]]
}
