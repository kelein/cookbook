package wallet

import "time"

// Transaction Kind Enumrate
const (
	DEBIT    TransKind = "DEBIT"
	CREDIT   TransKind = "CREDIT"
	TRANSFER TransKind = "TRANSFER"
)

// TransKind Transaction Type
type TransKind string

// VirtualWallectController .
type VirtualWallectController struct{}

func (vc *VirtualWallectController) getBalance(walletID int64) int64 { return 0 }

func (vc *VirtualWallectController) debit() {}

func (vc *VirtualWallectController) credit() {}

func (vc *VirtualWallectController) transfer() {}

// VirtualWallet .
type VirtualWallet struct {
	ID         int64
	Balance    float64
	CreateTime time.Time
}

// VirtualWallet

// VirtualWallectService .
type VirtualWallectService struct{}
