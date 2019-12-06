package strand

func ToRNA(dna string) string {
	var rna []byte
	for _, c := range dna {
		switch c {
		case 'G':
			rna = append(rna, 'C')
		case 'C':
			rna = append(rna, 'G')
		case 'T':
			rna = append(rna, 'A')
		case 'A':
			rna = append(rna, 'U')
		}
	}
	return string(rna)
}
