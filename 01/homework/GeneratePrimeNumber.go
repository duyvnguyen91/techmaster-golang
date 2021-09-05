package homework

// GeneratePrimeNumber 4.4 Lập dãy số nguyên tố
func GeneratePrimeNumber() (primes []int) {
	const N = 100000
	b := make([]bool, N)
	for i := 2; i < N; i++ {
		if b[i] == true { continue }
		primes = append(primes, i)
		for k := i * i; k < N; k += i {
			b[k] = true
		}
	}
	return primes
}