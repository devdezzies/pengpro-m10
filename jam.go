package main 

import "fmt"

func main() {
	var h1, m1, h2, m2 int 
	var jam, menit int
	var time int

	fmt.Scan(&h1, &m1, &h2, &m2) 
	if (h1 < h2) && (h2 - h1 == 1) {
		time = (60 - m1) + m2
        jam = time / 60
        menit = time % 60
	} else if (h1 < h2) {
		jam = h2 - h1
		menit = m1 - m2
		if (menit >= 0) {
			menit = menit
		} else {
			menit = -1 * menit
		}
	} else if (h1 > h2) {
		jam = (h2 + 12) - h1 
		menit = m1 - m2
		if (menit >= 0) {
			menit = menit
		} else {
			menit = -1 * menit
		}
	} else {
		jam = h1 - h2
		menit = m1 - m2
		if (menit >= 0) {
			menit = menit
		} else {
			menit = -1 * menit
		}
	}
	fmt.Println(jam, "jam", menit, "menit")
}