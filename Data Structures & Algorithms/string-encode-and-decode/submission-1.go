

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
	packet := []byte{}

    for _, s := range strs {
        packet = append(packet, byte(len(s)))
        packet = append(packet, []byte(s)...)
    }

    return string(packet)
}

func (s *Solution) Decode(encoded string) []string {
    packet := []byte(encoded)

    i := 0
    strings := []string{}

    if len(packet) == 0 {
        return []string{}
    }
    for {
        l := int(packet[i])

        if l == 0 {
            strings = append(strings, "")
        } else {
            strings = append(strings, string(packet[i+1:i+1+l]))
        }

        i += 1+l

        if i >= len(packet) {
            break
        }
    }

    return strings

}
