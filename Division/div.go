package div

func Division(a float32, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.new("No se puede dividir por 0")
	}
	return a / b, nil
	var res float32 = a / b
	return res, nil
}
