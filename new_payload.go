// generated by go:generate, DO NOT EDIT

package mdm

import "fmt"

func NewPayload(request *CommandRequest) (*Payload, error) {
	requestType := request.RequestType
	payload := newPayload(requestType)
	switch requestType {

	case "DeviceInformation":
		payload.Command.DeviceInformation = request.DeviceInformation

	case "InstallApplication":
		payload.Command.InstallApplication = request.InstallApplication

	case "AccountConfiguration":
		payload.Command.AccountConfiguration = request.AccountConfiguration

	case "ScheduleOSUpdateScan":
		payload.Command.ScheduleOSUpdateScan = request.ScheduleOSUpdateScan

	case "ScheduleOSUpdate":
		payload.Command.ScheduleOSUpdate = request.ScheduleOSUpdate

	case "InstallProfile":
		payload.Command.InstallProfile = request.InstallProfile

	case "RemoveProfile":
		payload.Command.RemoveProfile = request.RemoveProfile

	case "InstallProvisioningProfile":
		payload.Command.InstallProvisioningProfile = request.InstallProvisioningProfile

	case "RemoveProvisioningProfile":
		payload.Command.RemoveProvisioningProfile = request.RemoveProvisioningProfile

	case "InstalledApplicationList":
		payload.Command.InstalledApplicationList = request.InstalledApplicationList

	case "DeviceLock":
		payload.Command.DeviceLock = request.DeviceLock

	case "ClearPasscode":
		payload.Command.ClearPasscode = request.ClearPasscode

	case "EraseDevice":
		payload.Command.EraseDevice = request.EraseDevice

	case "RequestMirroring":
		payload.Command.RequestMirroring = request.RequestMirroring

	case "DeleteUser":
		payload.Command.DeleteUser = request.DeleteUser

	case "EnableLostMode":
		payload.Command.EnableLostMode = request.EnableLostMode

	case "ApplyRedemptionCode":
		payload.Command.ApplyRedemptionCode = request.ApplyRedemptionCode

	case "InstallMedia":
		payload.Command.InstallMedia = request.InstallMedia

	case "RemoveMedia":
		payload.Command.RemoveMedia = request.RemoveMedia

	case "Settings":
		payload.Command.Settings = request.Settings

	case "RefreshCellularPlans":
		payload.Command.RefreshCellularPlans = request.RefreshCellularPlans

	case "RequestUnlockToken":
		payload.Command.RequestUnlockToken = request.RequestUnlockToken

	case "ProfileList",
		"ProvisioningProfileList",
		"CertificateList",
		"SecurityInfo",
		"StopMirroring",
		"ClearRestrictionsPassword",
		"UserList",
		"LogOutUser",
		"DisableLostMode",
		"DeviceLocation",
		"ManagedMediaList",
		"OSUpdateStatus",
		"DeviceConfigured",
		"AvailableOSUpdates",
		"Restrictions",
		"ShutDownDevice",
		"RestartDevice":
		return payload, nil
	default:
		return nil, fmt.Errorf("Unsupported MDM RequestType %v", requestType)
	}
	return payload, nil
}
