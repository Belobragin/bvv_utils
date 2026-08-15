package mistake

import (
	"errors"
)

var (
	ErrStrangeOutput            = errors.New("this must not happen")
	ErrNullLogger               = errors.New("logger is null")
	ErrCommunicationChan        = errors.New("error channel is null")
	ErrInvalidUserUUID          = errors.New("user id invalid")
	ErrInvalidOtherTokenData    = errors.New("token data invalid")
	ErrAutorenewSessionIDNull   = errors.New("session is is null")
	ErrNullUserUUID             = errors.New("token user_uuid field is empty")
	ErrInvalidAud               = errors.New("Token's audience is out of service scope")
	ErrInvalidStatusCode        = errors.New("status code is not 200, 201")
	ErrDbConnect                = errors.New("failed to connect to database")
	ErrNullApiId                = errors.New("api id entity identifier is null")
	ErrNullApiGuid              = errors.New("api guid entity identifier is null")
	ErrMustBeAdmin              = errors.New("this api demands is_admin role")
	ErrUserAlreadyExist         = errors.New("user with the phone exists")
	ErrIncorrectNotFoundHandler = errors.New("custom not found handler fail")
	ErrNotUniqueSerial          = errors.New("not unique serial key in the psql table")
	ErrAllUpdateNull            = errors.New("one of update parameters must be not null")
	ErrDbTxConflict             = errors.New("either db or tx must be non-nil, the other one must be nil")
)

// control panel errors:
var (
	ErrNoBaseUrl                     = errors.New("no baseUrl data in service panel")
	ErrInvalidControlPanel           = errors.New("control panel does not contain must panel data")
	ErrInvalidMeltingForLogicClient  = errors.New("can not melt application client")
	ErrInvalidMeltingForMetricClient = errors.New("can not melt metric client")
	ErrInvalidMeltingForEventClient  = errors.New("can not melt event client")
)

// filter mistakes:
var (
	ErrFilterQuery            = errors.New("produce filter query err")
	ErrInvalidSortOrder       = errors.New("sort order invalid: must be asc, desc")
	ErrNillFilterData         = errors.New("standard filter data input must not be nil")
	ErrInvalidTableFilter     = errors.New("table name input is invalid")
	ErrInvalidParameterFilter = errors.New("filter parameters invalid")
	ErrorSortOnGetItems       = errors.New("sorting parameters are invalid")
	ErrUrlPhoneParameterNull  = errors.New("url phone parameter is a must, but not set")
	ErrUrlEmailParameterNull  = errors.New("url email parameter is a must, but not set")
)

// validation mistake:
var (
	ErrEventBodyNull       = errors.New("event body null")
	ErrStrangeConversion   = errors.New("unexpected error on type conversion")
	ErrIDValidation        = errors.New("id validation failed")
	ErrUserValidation      = errors.New("user validation failed")
	ErrFUserValidation     = errors.New("fuser (temporary user) validation failed")
	ErrPartnerValidation   = errors.New("partner validation failed")
	ErrMistakeArgNum       = errors.New("incorrect arguments length for get query")
	ErrSitePrefixInvalid   = errors.New("invalid site prefix")
	ErrPrefixFormatInvalid = errors.New("invalid country prefix: illegal formst")
)

// api mistakes:
var (
	ErrInvalidApiPathFormat = errors.New("invalid api path formst")
	ErrIdInPath             = errors.New("path object id error")
	ErrPointIDNull          = errors.New("point ID must not be 0")
	ErrInvalidHeader        = errors.New("invalid Authorization header")
	ErrNoAuthHeader         = errors.New("no Authorization header")
	ErrNulUser              = errors.New("user uuid must not be nil")
)

// input-output mistakes:
var (
	ErrInvalidGRPCConn            = errors.New("grpc connection failed")
	ErrApiUnknown                 = errors.New("unexpected unknown error")
	ErrTimeDbOutputTable          = errors.New("database timestamp invalid")
	ErrFailUseCase                = errors.New("usecase init failed")
	ErrConvertFeatureToMap        = errors.New("mistake when convert output structure to map")
	ErrPostFailed                 = errors.New("operation POST failed")
	ErrGetUserData                = errors.New("can not pull data about user")
	ErrGetFUserData               = errors.New("can not pull data about fuser(temporary user)")
	ErrGetPartnerData             = errors.New("can not pull data about partner")
	ErrNoSuchPartner              = errors.New("partner record with a given id does not exist")
	ErrNoSuchUser                 = errors.New("user record with a given id does not exist")
	ErrNoSuchFUser                = errors.New("temporary f_user record with a given id does not exist")
	ErrLoginIsNotAPhone           = errors.New("the user login is not a phone")
	ErrPasswordLen                = errors.New("password length must be 72 bytes or less")
	ErrSmsCodeLen                 = errors.New("sms code must be 4 digits")
	ErrEmailCodeLen               = errors.New("sms code must be 6 digits/letters")
	ErrNoAudience                 = errors.New("target audience undefined")
	ErrInvalidBothUserCredentials = errors.New("must submit either password or code, not both")
	ErrNoValidLogin               = errors.New("No valid login or phone value")
	ErrNoValidUserPsw             = errors.New("invalid credentials: niether password nor code")
	ErrNoValidPartnerPsw          = errors.New("invalid credentials: no password")
	ErrNeverCodeAndPsw            = errors.New("invalid credentials: both code and password")
	ErrPaswordIsNotALegalCode     = errors.New("this password is not a code")
	ErrPswCode                    = errors.New("control code input invalid")
	ErrCodeIsExpired              = errors.New("verification code expired")
	ErrNoValidRambler             = errors.New("invalid subject type (nomad) ")
	ErrServiceUnIntended          = errors.New("token was not intended for the service")
	ErrNoPermissionToDelete       = errors.New("no permission for delete operation")
	ErrFUserNotAllowedApi         = errors.New("this api is not allowed for FUser")
	ErrNoAlienSession             = errors.New("session_id for input token user uuid does nor exist")
	ErrInvalidNameLength          = errors.New("name length incorrect")
	ErrInvalidAliasLength         = errors.New("alias length incorrect")
	ErrInvalidGuid                = errors.New("field must be uuid")
	ErrInvalidId                  = errors.New("invalid id")
	ErrNilName                    = errors.New("name is nil invalid")
	ErrInvalidActivity            = errors.New("activity invalid id")
)

// handler mistakes
var (
	ErrInvalidUserLogin    = errors.New("invalid user login")
	ErrInvalidUserPassword = errors.New("invalid user password")
	ErrCacheRequest        = errors.New("cache request fail")
	ErrNoCacheValue        = errors.New("cache respond no value")
)

// database process mistakes:
var (
	ErrTxFailed         = errors.New("can not start transaction")
	ErrCloseTransaction = errors.New("can not close transaction")
	ErrDbPing           = errors.New("failed to ping database")
	ErrDbIsNotPsql      = errors.New("database object is not postgres")
)

// background process errors:
var (
	ErrMistakeCacheRenewCtx   = errors.New("context cancel on error")
	ErrMistakeCacheRenewOther = errors.New("background process other error")
)

// mocking errors
var (
	ErrMock = errors.New("mocking err")
)

// nats mistakes
var (
	ErrNatsClientNameAbsent = errors.New("nats client name not presented")
	ErrNatsUriAbsent        = errors.New("nats connection uri absent")
	ErrInvalidNatsValue     = errors.New("nats message is nil")
)
