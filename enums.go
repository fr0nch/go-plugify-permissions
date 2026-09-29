package permissions

// Generated from permissions

type Status int32

const (
	Status_Success Status = 0
	Status_Allow Status = 1
	Status_Disallow Status = 2
	Status_PermNotFound Status = 3
	Status_TemporalGroup Status = 4
	Status_PermanentGroup Status = 5
	Status_GroupNotDefined Status = 6
	Status_ErrorsStart Status = 30
	Status_PermAlreadyGranted Status = 31
	Status_CookieNotFound Status = 32
	Status_OptionNotFound Status = 32
	Status_GroupNotFound Status = 33
	Status_ChildGroupNotFound Status = 34
	Status_ParentGroupNotFound Status = 35
	Status_GroupAlreadyExist Status = 36
	Status_GroupHierarchyCycle Status = 37
	Status_ActorUserNotFound Status = 38
	Status_TargetUserNotFound Status = 39
	Status_UserAlreadyExist Status = 40
	Status_CallbackInvalid Status = 41
	Status_CallbackAlreadyExist Status = 42
	Status_CallbackNotFound Status = 43
	Status_StorageError Status = 44
	Status_DBNotReady Status = 45
	Status_InvalidPermission Status = 46
)

type Action int32

const (
	Action_Add Action = 0
	Action_Remove Action = 1
	Action_Replace Action = 2
	Action_ReplaceToWC Action = 3
)

type PlayerState uint32

const (
	PlayerState_NotFound PlayerState = 0
	PlayerState_Online PlayerState = 1
	PlayerState_Offline PlayerState = 2
)

type PermSource uint32

const (
	PermSource_UserTemp PermSource = 0
	PermSource_User PermSource = 1
	PermSource_GroupTemp PermSource = 2
	PermSource_Group PermSource = 3
	PermSource_NotFound PermSource = 4
)


