package iwlist

import (
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

type AccessPoints []AccessPoint

type AccessPoint struct {
	Address       *string
	Channel       *int
	Frequency     *float64
	Quality       *int
	SignalLevel   *int
	EncryptionKey *bool
	ESSID         *string
	Mode          *string
}

func (this *AccessPoints) BestQuality() *AccessPoint {
	max := 0
	accessPoint := AccessPoint{}
	for _, ap := range *this {
		if *(ap.Quality) > max {
			max = *(ap.Quality)
			accessPoint = ap
		}
	}
	return &accessPoint
}

func (this *AccessPoints) BestSignalLevel() *AccessPoint {
	max := -500
	accessPoint := AccessPoint{}
	for _, ap := range *this {
		if *(ap.SignalLevel) > max {
			max = *(ap.SignalLevel)
			accessPoint = ap
		}
	}
	return &accessPoint
}

func (this *AccessPoints) Match(essids ...string) *AccessPoints {
	aps := AccessPoints{}
	for _, ap := range *this {
		for _, essid := range essids {
			if *(ap.ESSID) == essid {
				aps = append(aps, ap)
			}
		}
	}
	return &aps
}

func Scan(wInterface string) *AccessPoints {
	var accessPoints = AccessPoints{}

	result, _ := execute("sudo", "iwlist", wInterface, "scanning")
	elements := strings.Split(result, "Cell")

	for _, e := range elements {
		accessPoint := AccessPoint{
			Address:       getAddress(e),
			Channel:       getChannel(e),
			Frequency:     getFrequency(e),
			Quality:       getQuality(e),
			SignalLevel:   getSignalLevel(e),
			EncryptionKey: getEncryptionKey(e),
			ESSID:         getESSID(e),
			Mode:          getMode(e),
		}

		if accessPoint.Address != nil {
			accessPoints = append(accessPoints, accessPoint)
		}
	}

	return &accessPoints
}

func getMode(sample string) *string {
	modes := filter(sample, `Mode:.*`)
	if len(modes) == 1 {
		result := strings.TrimPrefix(modes[0], "Mode:")
		return &result
	}
	return nil
}

func getESSID(sample string) *string {
	ssids := filter(sample, `ESSID:.*`)
	if len(ssids) == 1 {
		result := strings.TrimPrefix(ssids[0], "ESSID:\"")
		result = strings.TrimSuffix(result, "\"")
		return &result
	}
	return nil
}

func getEncryptionKey(sample string) *bool {
	keys := filter(sample, `Encryption key:(on|off)`)
	if len(keys) == 1 {
		keyStr := strings.TrimPrefix(keys[0], `Encryption key:`)
		var result bool
		if keyStr == "on" {
			result = true
		}
		if keyStr == "off" {
			result = false
		}
		return &result
	}
	return nil
}

func getSignalLevel(sample string) *int {
	signalLevels := filter(sample, `Signal level=([-0-9])*`)
	if len(signalLevels) == 1 {
		signalLevelStr := strings.TrimPrefix(signalLevels[0], "Signal level=")
		num, err := strconv.Atoi(signalLevelStr)
		if err != nil {
			return nil
		}
		return &num
	}
	return nil
}

func getQuality(sample string) *int {
	qualities := filter(sample, `Quality=([0-9])*/`)
	if len(qualities) == 1 {
		result := strings.TrimPrefix(qualities[0], "Quality=")
		result = strings.TrimSuffix(result, "/")
		num, err := strconv.Atoi(result)
		if err != nil {
			return nil
		}
		return &num
	}
	return nil
}

func getFrequency(sample string) *float64 {
	frequencies := filter(sample, `Frequency:([0-9.])*`)
	if len(frequencies) == 1 {
		freqStr := strings.TrimPrefix(frequencies[0], "Frequency:")
		freq, err := strconv.ParseFloat(freqStr, 10)
		if err != nil {
			return nil
		}
		return &freq
	}
	return nil
}

func getChannel(sample string) *int {
	channels := filter(sample, `Channel:([0-9])*`)
	if len(channels) == 1 {
		channelNum := strings.TrimPrefix(channels[0], "Channel:")
		num, err := strconv.Atoi(channelNum)
		if err != nil {
			return nil
		}
		return &num
	}
	return nil
}

func getAddress(sample string) *string {
	addresses := filter(sample, `..:..:..:..:..:..`)
	if len(addresses) == 1 {
		return &addresses[0]
	}
	return nil
}

func filter(sample, pattern string) []string {
	rp := regexp.MustCompilePOSIX(pattern)
	return rp.FindAllString(sample, -1)
}

func execute(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	byteResult, err := cmd.CombinedOutput()
	return string(byteResult), err
}
