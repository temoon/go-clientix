package types

const AppointmentStatusCancelled = "cancelled"
const AppointmentStatusInProgress = "in_progress"
const AppointmentStatusScheduled = "scheduled"
const AppointmentStatusFinished = "finished"
const AppointmentStatusMissed = "missed"
const AppointmentStatusConfirmed = "confirmed"
const AppointmentStatusCancelledBySms = "cancelled_by_sms"
const AppointmentStatusSmsConfirmationSent = "sms_confirmation_sent"

type AppointmentStatus string

func (t AppointmentStatus) IsValid() bool {
	switch t {
	case AppointmentStatusCancelled, AppointmentStatusInProgress, AppointmentStatusScheduled, AppointmentStatusFinished,
		AppointmentStatusMissed, AppointmentStatusConfirmed, AppointmentStatusCancelledBySms, AppointmentStatusSmsConfirmationSent:
		return true
	default:
		return false
	}
}
