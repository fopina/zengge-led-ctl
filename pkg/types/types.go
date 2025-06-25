package types

// ScanHandler handles Zengge advertisements.
type ScanHandler func(a ZenggeAdvertisement)

// NotificationHandler handles Zengge notifications.
type NotificationHandler func(a Notification)
