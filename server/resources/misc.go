package resources

import "mytholojam/server/types"

var moveTypes = [4]string{"STRN", "FLAM", "WEAR", "NAME"}

const selfTarget = "self"
const opTarget = "other"

const lowHP = 2
const avgHP = 3
const highHP = 4

const lowATK = 1
const avgATK = 2
const highATK = 3

const lowDEF = 1
const avgDEF = 2
const highDEF = 3 // 1080p lmao

const lowWeight = 1
const avgWeight = 2
const highWeight = 3

const lowSPD = 1
const avgSPD = 2
const highSPD = 3

const lowPWR = 1
const avgPWR = 2
const highPWR = 3

const lowPri = -1
const avgPri = 0
const highPri = 1

// special constants
const switchType = "switch"
const switchPower = -1
const switchPri = 100

func noop(_ *types.Spirit, _ *types.Equipment, _ *types.Move) {

}
