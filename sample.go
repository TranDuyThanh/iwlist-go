package iwlist

var (
	addr1 = "24:A2:E1:EB:76:C0"
	addr2 = "88:1F:A1:30:6A:2A"
	addr3 = "8A:DC:96:1E:C7:5C"
	addr4 = "4C:E6:76:E0:BD:0F"
	addr5 = "24:A4:3C:13:50:60"

	chan1 = 1
	chan3 = 11
	chan6 = 6

	freq1 = float64(2.412)
	freq3 = float64(2.422)
	freq6 = float64(2.437)

	qual1 = 10
	qual2 = 20
	qual3 = 30
	qual4 = 40
	qual5 = 50

	sign1 = -90
	sign2 = -70
	sign3 = -10
	sign4 = -85
	sign5 = -19

	encr1 = true
	cncr2 = false

	ssid1 = "hihi"
	ssid2 = "abc abc"
	ssid3 = "hoi lam gi"
	ssid4 = "qua Quan b3n canh"
	ssid5 = "khong co p@ss!"

	mode0 = "Master"
)

var (
	ap1 = AccessPoint{
		Address:       &addr1,
		Channel:       &chan1,
		Frequency:     &freq1,
		Quality:       &qual1,
		SignalLevel:   &sign1,
		EncryptionKey: &encr1,
		ESSID:         &ssid1,
		Mode:          &mode0,
	}
	ap2 = AccessPoint{
		Address:       &addr2,
		Channel:       &chan1,
		Frequency:     &freq1,
		Quality:       &qual2,
		SignalLevel:   &sign2,
		EncryptionKey: &encr1,
		ESSID:         &ssid2,
		Mode:          &mode0,
	}
	ap3 = AccessPoint{
		Address:       &addr3,
		Channel:       &chan3,
		Frequency:     &freq3,
		Quality:       &qual3,
		SignalLevel:   &sign3,
		EncryptionKey: &encr1,
		ESSID:         &ssid3,
		Mode:          &mode0,
	}
	ap4 = AccessPoint{
		Address:       &addr4,
		Channel:       &chan3,
		Frequency:     &freq3,
		Quality:       &qual4,
		SignalLevel:   &sign4,
		EncryptionKey: &encr1,
		ESSID:         &ssid4,
		Mode:          &mode0,
	}
	ap5 = AccessPoint{
		Address:       &addr5,
		Channel:       &chan6,
		Frequency:     &freq6,
		Quality:       &qual5,
		SignalLevel:   &sign5,
		EncryptionKey: &encr1,
		ESSID:         &ssid5,
		Mode:          &mode0,
	}

	aps = &AccessPoints{ap1, ap2, ap3, ap4, ap5}
)
