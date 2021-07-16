package common

import "math/rand"

const sevenDaysSeconds = 7*24*60*60
const oneDaySeconds = 24*60*60

func GetCommonRedisExpireSeconds()int{
	offsetSeconds := rand.Intn(5) * oneDaySeconds
	offsetSeconds = sevenDaysSeconds + offsetSeconds

	return offsetSeconds
}
