package utils

//红包结果查询
type RedPackObj struct {
	/**
	 * 0: uint256: amount 50000000000000000000
	 * 1: uint256: balance 41248506822956664767
	 * 2: uint256: count 8
	 * 3: address[]: hunterInfos 0x3901952De2f16ad9B8646CF59C337d0b445A81Ca
	 * 4: uint256[]: pickAmounts 8751493177043335233
	 */
	Amount uint64  `json:"amount" gencodec:"required"`
	Balance uint64  `json:"balance" gencodec:"required"`
	Count uint64  `json:"count" gencodec:"required"`
	HunterInfos []string `json:"hunterInfos" gencodec:"required"`
	PickAmounts []uint64 `json:"pickAmounts" gencodec:"required"`
}


//查询授权可用额度
type ApproveRemain struct {
	Remaining uint64  `json:"remaining" gencodec:"required"`
}

// 商家信息
/**
  0: address: ow 0x3901952De2f16ad9B8646CF59C337d0b445A81Ca
  1: string: icon https://avatar.csdnimg.cn/6/F/5/0_qq_31708101.jpg
  2: string: name 商家1号
  3: uint256: status 0
  4: uint256: balance 80000000000000000000
  5: uint256: count 0
  6: uint256: fee 10
  7: uint256: totalFeeBalance 0
  8: uint256: keepBalance 500000000000000000000
  9: bool: used true
*/
type BusinessInfo struct {
	Address string  `json:"address" gencodec:"required"`
	Icon string  `json:"icon" gencodec:"required"`
	Name string  `json:"name" gencodec:"required"`
	Status uint64  `json:"status" gencodec:"required"`
	Balance uint64  `json:"balance" gencodec:"required"`
	Count uint64  `json:"count" gencodec:"required"`
	Fee uint64  `json:"fee" gencodec:"required"`
	TotalFeeBalance uint64  `json:"totalFeeBalance" gencodec:"required"`
	KeepBalance uint64  `json:"keepBalance" gencodec:"required"`
	Used bool  `json:"used" gencodec:"required"`
}

// 订单信息
/**
 * 0: string: oid
 * 1: uint256: amount 300000000000000000000
 * 2: uint256: status 1
 * 3: address: payUser 0x3901952De2f16ad9B8646CF59C337d0b445A81Ca
 * 4: uint256: block 128719
 * 5: bool: used true
 * 6: address: business 0x3901952De2f16ad9B8646CF59C337d0b445A81Ca
 */
type OrderInfo struct {
	Oid string  `json:"oid" gencodec:"required"`
	Amount uint64  `json:"amount" gencodec:"required"`
	Status uint64  `json:"status" gencodec:"required"`
	PayUser string  `json:"payUser" gencodec:"required"`
	Block uint64  `json:"block" gencodec:"required"`
	Used bool  `json:"used" gencodec:"required"`
	Address string  `json:"address" gencodec:"required"`
}