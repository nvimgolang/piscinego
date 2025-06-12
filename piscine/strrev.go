package piscine

func StrRev(s string) string {
	slovonaoborot := ""
	for _, pokajdoibukve := range s {
		slovonaoborot = string(pokajdoibukve) + slovonaoborot
	}
	return slovonaoborot
}
