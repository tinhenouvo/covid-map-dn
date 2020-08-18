package config

// Constant vars
const (
	// Context key
	ContextKeyCurrentUserID   = "currentUserID"
	ContextKeyCurrentUserRole = "currentUserRole"
	ContextKeyLanguage        = "lang"

	// Language
	LangVi = "vi"
	LangEn = "en"

	Limit20 = 20

	// Queue name
	QueueNameNewNotification = "new_notification"
)

// Transaction
const (
	TransactionHistoryBuy      = "buy"
	TransactionHistoryPending  = "pending"
	TransactionHistoryApproved = "approved"
	TransactionHistoryRejected = "rejected"
	TransactionHistoryCashback = "cashback"
)

//Transcation Offline
const (
	TransactionOfflineSuccess  = "cashback"
	TransactionOfflineRejected = "rejected"
	TransactionOfflinePending  = "pending"
)

// Time
const (
	TimezoneHCM = "Asia/Ho_Chi_Minh"

	DateFormat = "02/01/2006"
)

// Withdraw cash
const (
	WithdrawCashStatusPending  = "pending"
	WithdrawCashStatusApproved = "approved"
	WithdrawCashStatusRejected = "rejected"
)

// Notification
const (
	NotificationTypeNewTransactionHistory    = "new_transaction_history"
	NotificationTypeWithdrawSuccess          = "withdraw_success"
	NotificationTypeWithdrawRejected         = "withdraw_rejected"
	NotificationTypeAdmin                    = "admin_notification"
	NotificationTypeBonusTransactionCashback = "bonus_transaction_cashback"
)

// Referral
const (
	ReferralActivityInputSuccess                   = "input_success"
	ReferralActivityTransactionApproved            = "transaction_approved"
	ReferralActivityTransactionCashback            = "transaction_cashback"
	ReferralActivityTransactionPending             = "transaction_pending"
	ReferralActivityTransactionRejected            = "transaction_rejected"
	ReferralActivityEventRewardPending             = "event_reward_pending"
	ReferralActivityEventRewardCompleted           = "event_reward_completed"
	ReferralActivityEventBrandTransactionPending   = "event_brand_transaction_pending"
	ReferralActivityEventBrandTransactionCompleted = "event_brand_transaction_completed"
)
