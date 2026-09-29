package permissions

import "github.com/untrustedmodders/go-plugify"

var _ = plugify.ApiVersion

// Generated from permissions

// GroupCreateStorageCallback - Storage callback invoked before a group is created. Return `false` to cancel the change, the method then returns DBNotReady.
type GroupCreateStorageCallback func(pluginID int64, name string, perms []string, priority int32, parent string) bool


// GroupCreateCallback - Callback invoked after a group is successfully created.
type GroupCreateCallback func(pluginID int64, name string, perms []string, priority int32, parent string)


// GroupDeleteStorageCallback - Storage callback invoked before a group is deleted. Return `false` to cancel the change, the method then returns DBNotReady.
type GroupDeleteStorageCallback func(pluginID int64, name string) bool


// GroupDeleteCallback - Callback invoked after a group is deleted.
type GroupDeleteCallback func(pluginID int64, name string)


// GroupExpirationCallback - Callback invoked when a group in user has been expired.
type GroupExpirationCallback func(targetID uint64, group string)


// GroupOptionStorageCallback - Storage callback invoked before an option value is set for a group. Return `false` to cancel the change, the method then returns DBNotReady.
type GroupOptionStorageCallback func(pluginID int64, groupName string, optionName string, value any) bool


// GroupOptionCallback - Callback invoked when an option value is set for a group.
type GroupOptionCallback func(pluginID int64, groupName string, optionName string, value any)


// GroupPermissionStorageCallback - Storage callback invoked before a permission is added, replaced, or removed from a group. Return `false` to cancel the change, the method then returns DBNotReady.
type GroupPermissionStorageCallback func(pluginID int64, action Action, groupName string, perm string, oldState Status, newState Status) bool


// GroupPermissionCallback - Callback invoked when a permission is added, replaced, or removed from a group.
type GroupPermissionCallback func(pluginID int64, action Action, groupName string, perm string, oldState Status, newState Status)


// LoadGroupsCallback - Called when the core requests loading of server groups.
type LoadGroupsCallback func(pluginID int64) bool


// GroupsLoadedCallback - Called when server groups have been loaded.
type GroupsLoadedCallback func(pluginID int64)


// PermExpirationCallback - Callback invoked when a permission in user has been expired.
type PermExpirationCallback func(targetID uint64, perm string, state Status)


// SetParentStorageCallback - Storage callback invoked before a parent group is set for a child group. Return `false` to cancel the change, the method then returns DBNotReady.
type SetParentStorageCallback func(pluginID int64, childName string, parentName string) bool


// SetParentCallback - Callback invoked when a parent group is set for a child group.
type SetParentCallback func(pluginID int64, childName string, parentName string)


// UserCookieStorageCallback - Storage callback invoked before a cookie is set for a user. Return `false` to cancel the change, the method then returns DBNotReady.
type UserCookieStorageCallback func(pluginID int64, targetID uint64, name string, cookie any) bool


// UserCookieCallback - Callback invoked when a cookie is set for a user.
type UserCookieCallback func(pluginID int64, targetID uint64, name string, cookie any)


// UserCreateCallback - Callback invoked after a user is successfully created.
type UserCreateCallback func(pluginID int64, targetID uint64, immunity int32, offline bool, groupNames []string)


// UserDeleteCallback - Callback invoked after a user is deleted.
type UserDeleteCallback func(pluginID int64, targetID uint64)


// UserGroupStorageCallback - Storage callback invoked before a group is added, replaced, or removed for a user. Return `false` to cancel the change, the method then returns DBNotReady.
type UserGroupStorageCallback func(pluginID int64, action Action, targetID uint64, group string, oldTimestamp int64, newTimestamp int64) bool


// UserGroupCallback - Callback invoked when a group is added, replaced, or removed for a user.
type UserGroupCallback func(pluginID int64, action Action, targetID uint64, group string, oldTimestamp int64, newTimestamp int64)


// UserImmunityStorageCallback - Storage callback invoked before immunity is set for a user. Return `false` to cancel the change, the method then returns DBNotReady.
type UserImmunityStorageCallback func(pluginID int64, targetID uint64, immunity int32) bool


// UserImmunityCallback - Callback invoked when immunity is set for a user.
type UserImmunityCallback func(pluginID int64, targetID uint64, immunity int32)


// UserLoadedCallback - Called when a user's data has been fully loaded.
type UserLoadedCallback func(pluginID int64, targetID uint64, playerState PlayerState)


// UserPermissionStorageCallback - Storage callback invoked before a permission is added, replaced, or removed for a user. Return `false` to cancel the change, the method then returns DBNotReady.
type UserPermissionStorageCallback func(pluginID int64, action Action, targetID uint64, perm string, oldState Status, newState Status, oldTimestamp int64, newTimestamp int64) bool


// UserPermissionCallback - Callback invoked when a permission is added, removed, or replaced for a user.
type UserPermissionCallback func(pluginID int64, action Action, targetID uint64, perm string, oldState Status, newState Status, oldTimestamp int64, newTimestamp int64)


// UserRequestCallback - Called when a user data load is requested.
type UserRequestCallback func(pluginID int64, targetID uint64, username string, offline bool) bool


