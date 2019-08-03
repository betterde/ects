package utils

import "net"

// Get all network card IP
func GetIPs() (ips []string) {
	cards, err := net.Interfaces()
	if err != nil {

	}

	for _, card := range cards {
		if (card.Flags & net.FlagUp) != 0 {
			addrs, err := card.Addrs()
			if err != nil {

			}
			for _, addr := range addrs {
				if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() {
					if ipnet.IP.To4() != nil {
						ips = append(ips, ipnet.IP.String())
					}
				}
			}
		}
	}

	return ips
}
