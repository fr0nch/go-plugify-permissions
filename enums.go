package permissions

// Generated from permissions

// Status - Result codes returned by the API. Values after ErrorsStart are errors.
type Status int32

const (
	// Success - Operation completed successfully.
	Status_Success Status = 0
	// Allow - Permission is allowed.
	Status_Allow Status = 1
	// Disallow - Permission is explicitly denied.
	Status_Disallow Status = 2
	// PermNotFound - Permission is not defined.
	Status_PermNotFound Status = 3
	// TemporalGroup - User has the group temporarily.
	Status_TemporalGroup Status = 4
	// PermanentGroup - User has the group permanently.
	Status_PermanentGroup Status = 5
	// GroupNotDefined - User doesn't have the group.
	Status_GroupNotDefined Status = 6
	// ErrorsStart - Marker: every value greater than this is an error.
	Status_ErrorsStart Status = 30
	// PermAlreadyGranted - Permission already exists with the same state, or conflicts with an existing wildcard.
	Status_PermAlreadyGranted Status = 31
	// CookieNotFound - Cookie (user) or option (group) is not set.
	Status_CookieNotFound Status = 32
	// OptionNotFound - Cookie (user) or option (group) is not set.
	Status_OptionNotFound Status = 32
	// GroupNotFound - Group doesn't exist.
	Status_GroupNotFound Status = 33
	// ChildGroupNotFound - Child group doesn't exist.
	Status_ChildGroupNotFound Status = 34
	// ParentGroupNotFound - Parent group doesn't exist, or the group has no parent.
	Status_ParentGroupNotFound Status = 35
	// GroupAlreadyExist - Group already exists, or the user already has it.
	Status_GroupAlreadyExist Status = 36
	// GroupHierarchyCycle - Setting the parent would create a cycle in the group hierarchy.
	Status_GroupHierarchyCycle Status = 37
	// ActorUserNotFound - Actor user is not loaded.
	Status_ActorUserNotFound Status = 38
	// TargetUserNotFound - Target user is not loaded.
	Status_TargetUserNotFound Status = 39
	// UserAlreadyExist - User is already loaded.
	Status_UserAlreadyExist Status = 40
	// CallbackInvalid - Callback is null.
	Status_CallbackInvalid Status = 41
	// CallbackAlreadyExist - Callback is already registered.
	Status_CallbackAlreadyExist Status = 42
	// CallbackNotFound - Callback is not registered.
	Status_CallbackNotFound Status = 43
	// StorageError - Storage error.
	Status_StorageError Status = 44
	// DBNotReady - A storage callback rejected the change (storage not ready or failed). The change wasn't applied.
	Status_DBNotReady Status = 45
	// InvalidPermission - Permission line is malformed (empty, empty segment, or '*' not as the last segment).
	Status_InvalidPermission Status = 46
	// InvalidGroupName - Group name is empty.
	Status_InvalidGroupName Status = 47
	// InvalidCookieName - Cookie or option name is empty.
	Status_InvalidCookieName Status = 48
	// InvalidOptionName - Cookie or option name is empty.
	Status_InvalidOptionName Status = 48
)

// Action - Kind of change reported to permission and group callbacks.
type Action int32

const (
	// Add - Added.
	Action_Add Action = 0
	// Remove - Removed.
	Action_Remove Action = 1
	// Replace - Replaced (state or duration changed).
	Action_Replace Action = 2
	// ReplaceToWC - A plain permission was replaced with its wildcard ("a" -> "a.*").
	Action_ReplaceToWC Action = 3
)

// PlayerState - User presence state.
type PlayerState uint32

const (
	// NotFound - User is not loaded.
	PlayerState_NotFound PlayerState = 0
	// Online - User is loaded and on the server.
	PlayerState_Online PlayerState = 1
	// Offline - User is loaded without presence on the server.
	PlayerState_Offline PlayerState = 2
)

// PermSource - Where a permission was found.
type PermSource uint32

const (
	// UserTemp - Temporary user permission.
	PermSource_UserTemp PermSource = 0
	// User - Permanent user permission.
	PermSource_User PermSource = 1
	// GroupTemp - From a temporary group.
	PermSource_GroupTemp PermSource = 2
	// Group - From a permanent group.
	PermSource_Group PermSource = 3
	// NotFound - Permission not found.
	PermSource_NotFound PermSource = 4
)


