package rotationalcipher

func RotationalCipher(plain string, shiftKey int) string {
	positions := byte(shiftKey % 26)
	cipher := make([]byte, len(plain))

	for i := 0; i < len(plain); i++ {
		char := plain[i]

		switch {
		case char >= 'a' && char <= 'z':
			cipher[i] = 'a' + (char-'a'+positions)%26
		case char >= 'A' && char <= 'Z':
			cipher[i] = 'A' + (char-'A'+positions)%26
		default:
			cipher[i] = char
		}
	}

	return string(cipher)
}
