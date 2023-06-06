package main

import (
	"fmt"
)

const NMAX int = 2023

type dataBarang struct {
	namaBarang  string
	harga, stok int
}

type dataTransaksi struct {
	namaBarang, jenis string
	waktu             date
	nilai             int
	jumlahWaktu       int
}

type date struct {
	tanggal, bulan, tahun int
}

type arrBarang [NMAX]dataBarang
type arrTransaksi [NMAX]dataTransaksi

func inputDataBarang(T *arrBarang, N *int) {
	*T = arrBarang{
		{"beras", 10000, 18},
		{"sagu", 16000, 15},
		{"gulaPasir", 10500, 30},
		{"gulaMerah", 9500, 36},
		{"telur", 3000, 154},
		{"minyakGoreng", 30000, 34},
		{"margarin", 20000, 32},
		{"susu", 14500, 61},
		{"garam", 15000, 43},
		{"ayam", 40000, 23},
		{"ikan", 25000, 24},
		{"lada", 4000, 80},
		{"kecap", 5000, 58},
		{"saos", 10000, 74},
		{"terigu", 12000, 52},
	}
	*N = 15
}

func inputDataTransaksi(T *arrTransaksi, N *int) {
	*T = arrTransaksi{
		{"telur", "ubah", date{03, 01, 2023}, -14, 738428},
		{"susu", "ubah", date{04, 01, 2023}, 9, 738429},
		{"ikan", "ubah", date{04, 01, 2023}, 11, 738429},
		{"beras", "ubah", date{05, 01, 2023}, 22, 738430},
		{"beras", "ubah", date{05, 01, 2023}, -5, 738430},
		{"terigu", "tambah", date{07, 01, 2023}, 13, 738432},
		{"saos", "ubah", date{07, 01, 2023}, -24, 738432},
		{"kecap", "ubah", date{10, 01, 2023}, -8, 738435},
		{"lada", "ubah", date{10, 01, 2023}, 25, 738435},
		{"ayam", "ubah", date{11, 01, 2023}, -3, 738436},
		{"beras", "ubah", date{11, 01, 2023}, 15, 738436},
		{"minyakGoreng", "ubah", date{11, 01, 2023}, -4, 738436},
		{"ikan", "ubah", date{12, 01, 2023}, -4, 738437},
		{"telur", "ubah", date{12, 01, 2023}, -34, 738437},
		{"sagu", "ubah", date{12, 01, 2023}, -10, 738437},
	}
	*N = 15
}

func idxData_nama(T arrBarang, N int, x string) int {
	var found bool
	var i int

	for i < N && !found {
		found = T[i].namaBarang == x
		i++
	}
	return i - 1
}

func idxData_harga(T arrBarang, N int, x int) int {
	var found bool
	var i int

	for i < N && !found {
		found = T[i].harga == x
		i++
	}
	return i - 1
}

func idxData_stok(T arrBarang, N int, x int) int {
	var found bool
	var i int

	for i < N && !found {
		found = T[i].stok == x
		i++
	}
	return i - 1
}

func idxMaksimum(T arrBarang, X arrTransaksi, N, N1, x int) int {
	var idx int
	var j int = 1

	if x == 1 {
		for j < N {
			if T[idx].harga < T[j].harga {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 2 {
		for j < N {
			if T[idx].stok < T[j].stok {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 3 {
		for j < N {
			if X[idx].jumlahWaktu <= X[j].jumlahWaktu {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 4 {
		for j < N {
			if X[idx].nilai <= X[j].nilai {
				idx = j
			}
			j++
		}
		return idx
	}
	return 0
}

func idxMinimum(T arrBarang, X arrTransaksi, N, N1, x int) int {
	var idx int
	var j int = 1

	if x == 1 {
		for j < N {
			if T[idx].harga > T[j].harga {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 2 {
		for j < N {
			if T[idx].stok > T[j].stok {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 3 {
		for j < N {
			if X[idx].jumlahWaktu > X[j].jumlahWaktu {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 4 {
		for j < N {
			if X[idx].nilai > X[j].nilai {
				idx = j
			}
			j++
		}
		return idx
	}
	return 0
}

func ascending(T *arrBarang, X *arrTransaksi, N, N1, x int) {
	var i, j, temp, temp1, temp2, temp3, temp6 int
	var temp4, temp5 string
	i = 1

	if x == 1 {
		for i <= N-1 {
			j = i
			temp = T[j].harga
			temp2 = T[j].stok
			temp4 = T[j].namaBarang
			for j > 0 && temp < T[j-1].harga {
				T[j].harga = T[j-1].harga
				T[j].stok = T[j-1].stok
				T[j].namaBarang = T[j-1].namaBarang
				j--
			}
			T[j].harga = temp
			T[j].stok = temp2
			T[j].namaBarang = temp4
			i++
		}

	} else if x == 2 {
		for i <= N-1 {
			j = i
			temp = T[j].harga
			temp2 = T[j].stok
			temp4 = T[j].namaBarang
			for j > 0 && temp2 < T[j-1].stok {
				T[j].harga = T[j-1].harga
				T[j].stok = T[j-1].stok
				T[j].namaBarang = T[j-1].namaBarang
				j--
			}
			T[j].harga = temp
			T[j].stok = temp2
			T[j].namaBarang = temp4
			i++
		}

	} else if x == 3 {
		for i <= N1-1 {
			j = i
			temp = X[j].jumlahWaktu
			temp1 = X[j].waktu.bulan
			temp2 = X[j].waktu.tanggal
			temp3 = X[j].waktu.tahun
			temp4 = X[j].namaBarang
			temp5 = X[j].jenis
			temp6 = X[j].nilai
			for j > 0 && temp < X[j-1].jumlahWaktu {
				X[j].jumlahWaktu = X[j-1].jumlahWaktu
				X[j].waktu.bulan = X[j-1].waktu.bulan
				X[j].waktu.tanggal = X[j-1].waktu.tanggal
				X[j].waktu.tahun = X[j-1].waktu.tahun
				X[j].namaBarang = X[j-1].namaBarang
				X[j].jenis = X[j-1].jenis
				X[j].nilai = X[j-1].nilai
				j--
			}
			X[j].jumlahWaktu = temp
			X[j].waktu.bulan = temp1
			X[j].waktu.tanggal = temp2
			X[j].waktu.tahun = temp3
			X[j].namaBarang = temp4
			X[j].jenis = temp5
			X[j].nilai = temp6
			i++
		}

	} else if x == 4 {
		for i <= N1-1 {
			j = i
			temp = X[j].jumlahWaktu
			temp1 = X[j].waktu.bulan
			temp2 = X[j].waktu.tanggal
			temp3 = X[j].waktu.tahun
			temp4 = X[j].namaBarang
			temp5 = X[j].jenis
			temp6 = X[j].nilai
			for j > 0 && temp6 > X[j-1].nilai {
				X[j].jumlahWaktu = X[j-1].jumlahWaktu
				X[j].waktu.bulan = X[j-1].waktu.bulan
				X[j].waktu.tanggal = X[j-1].waktu.tanggal
				X[j].waktu.tahun = X[j-1].waktu.tahun
				X[j].namaBarang = X[j-1].namaBarang
				X[j].jenis = X[j-1].jenis
				X[j].nilai = X[j-1].nilai
				j--
			}
			X[j].jumlahWaktu = temp
			X[j].waktu.bulan = temp1
			X[j].waktu.tanggal = temp2
			X[j].waktu.tahun = temp3
			X[j].namaBarang = temp4
			X[j].jenis = temp5
			X[j].nilai = temp6
			i++
		}

	}
}

func descending(T *arrBarang, X *arrTransaksi, N, N1, x int) {
	var i, j, temp, temp1, temp2, temp3, temp6 int
	var temp4, temp5 string
	i = 1

	if x == 1 {
		for i <= N-1 {
			j = i
			temp = T[j].harga
			temp2 = T[j].stok
			temp4 = T[j].namaBarang
			for j > 0 && temp > T[j-1].harga {
				T[j].harga = T[j-1].harga
				T[j].stok = T[j-1].stok
				T[j].namaBarang = T[j-1].namaBarang
				j--
			}
			T[j].harga = temp
			T[j].stok = temp2
			T[j].namaBarang = temp4
			i++
		}

	} else if x == 2 {
		for i <= N-1 {
			j = i
			temp = T[j].harga
			temp2 = T[j].stok
			temp4 = T[j].namaBarang
			for j > 0 && temp2 > T[j-1].stok {
				T[j].harga = T[j-1].harga
				T[j].stok = T[j-1].stok
				T[j].namaBarang = T[j-1].namaBarang
				j--
			}
			T[j].harga = temp
			T[j].stok = temp2
			T[j].namaBarang = temp4
			i++
		}

	} else if x == 3 {
		for i <= N1-1 {
			j = i
			temp = X[j].jumlahWaktu
			temp1 = X[j].waktu.bulan
			temp2 = X[j].waktu.tanggal
			temp3 = X[j].waktu.tahun
			temp4 = X[j].namaBarang
			temp5 = X[j].jenis
			temp6 = X[j].nilai
			for j > 0 && temp > X[j-1].jumlahWaktu {
				X[j].jumlahWaktu = X[j-1].jumlahWaktu
				X[j].waktu.bulan = X[j-1].waktu.bulan
				X[j].waktu.tanggal = X[j-1].waktu.tanggal
				X[j].waktu.tahun = X[j-1].waktu.tahun
				X[j].namaBarang = X[j-1].namaBarang
				X[j].jenis = X[j-1].jenis
				X[j].nilai = X[j-1].nilai
				j--
			}
			X[j].jumlahWaktu = temp
			X[j].waktu.bulan = temp1
			X[j].waktu.tanggal = temp2
			X[j].waktu.tahun = temp3
			X[j].namaBarang = temp4
			X[j].jenis = temp5
			X[j].nilai = temp6

			i++
		}

	} else if x == 4 {
		for i <= N1-1 {
			j = i
			temp = X[j].jumlahWaktu
			temp1 = X[j].waktu.bulan
			temp2 = X[j].waktu.tanggal
			temp3 = X[j].waktu.tahun
			temp4 = X[j].namaBarang
			temp5 = X[j].jenis
			temp6 = X[j].nilai
			for j > 0 && temp6 > X[j-1].nilai {
				X[j].jumlahWaktu = X[j-1].jumlahWaktu
				X[j].waktu.bulan = X[j-1].waktu.bulan
				X[j].waktu.tanggal = X[j-1].waktu.tanggal
				X[j].waktu.tahun = X[j-1].waktu.tahun
				X[j].namaBarang = X[j-1].namaBarang
				X[j].jenis = X[j-1].jenis
				X[j].nilai = X[j-1].nilai
				j--
			}
			X[j].jumlahWaktu = temp
			X[j].waktu.bulan = temp1
			X[j].waktu.tanggal = temp2
			X[j].waktu.tahun = temp3
			X[j].namaBarang = temp4
			X[j].jenis = temp5
			X[j].nilai = temp6
			i++
		}

	}
}

func print(T arrBarang, X arrTransaksi, N, N1, x int) {
	if x == 1 {
		fmt.Println("nama", "\t\t", "harga", "\t", "stok")
		for i := 0; i < N; i++ {
			if len(T[i].namaBarang) <= 6 && len(T[i].namaBarang) >= 1 {
				fmt.Println(T[i].namaBarang, "\t\t", T[i].harga, "\t", T[i].stok)
			} else if len(T[i].namaBarang) >= 7 {
				fmt.Println(T[i].namaBarang, "\t", T[i].harga, "\t", T[i].stok)
			}
		}

	} else if x == 2 {
		fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
		for i := 0; i < N1; i++ {
			if len(X[i].namaBarang) <= 6 && len(X[i].namaBarang) >= 1 {
				fmt.Println(X[i].namaBarang, "\t\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
			} else if len(X[i].namaBarang) >= 7 {
				fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
			}
		}
	}
}

func menu(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("1. modifikasi data")
	fmt.Println("2. cari data")
	fmt.Println("3. tampilan data")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")
	if pilihan == 1 {
		modifikasi(&*T, &*X, &*N, &*N1)
	} else if pilihan == 2 {
		cari(&*T, &*X, &*N, &*N1)
	} else if pilihan == 3 {
		tampilan(&*T, &*X, &*N, &*N1)
	}

}

func modifikasi(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var idx, pilihan int
	var nama string

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("modifikasi:")
	fmt.Println("1. tambah data")
	fmt.Println("2. ubah data")
	fmt.Println("3. hapus data")
	fmt.Println("4. kembali")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	if pilihan == 1 {
		modifikasiTambah(&*T, &*X, &*N, &*N1)
	} else if pilihan == 2 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("nama barang: ")
		fmt.Scan(&nama)
		fmt.Print("\n")

		idx = idxData_nama(*T, *N, nama)
		modifikasiUbah(&*T, &*X, &*N, &*N1, idx)
	} else if pilihan == 3 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("nama barang: ")
		fmt.Scan(&nama)
		fmt.Print("\n")

		idx = idxData_nama(*T, *N, nama)
		modifikasiHapus(&*T, &*X, &*N, &*N1, idx)
	} else if pilihan == 4 {
		menu(&*T, &*X, &*N, &*N1)
	}
}

func modifikasiTambah(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var nama string
	var harga, stok, tanggal, bulan, tahun int

	fmt.Print("buat data")
	fmt.Print("\nnama: ")
	fmt.Scan(&nama)
	fmt.Print("harga: ")
	fmt.Scan(&harga)
	fmt.Print("stok: ")
	fmt.Scan(&stok)
	fmt.Print("waktu (DD MM YYYY): ")
	fmt.Scan(&tanggal, &bulan, &tahun)

	T[*N].namaBarang = nama
	T[*N].harga = harga
	T[*N].stok = stok
	fmt.Println("\nsukses ditambah")
	*N++

	X[*N1].namaBarang = nama
	X[*N1].waktu.tanggal = tanggal
	X[*N1].waktu.bulan = bulan
	X[*N1].waktu.tahun = tahun
	X[*N1].jenis = "tambah"
	X[*N1].nilai = stok
	X[*N1].jumlahWaktu = (365 * tahun) + (30 * bulan) + tanggal
	*N1++
}

func modifikasiHapus(T *arrBarang, X *arrTransaksi, N, N1 *int, idx int) {
	var tanggal, bulan, tahun int

	fmt.Print("waktu (DD MM YYYY): ")
	fmt.Scan(&tanggal, &bulan, &tahun)

	X[*N1].namaBarang = T[idx].namaBarang
	X[*N1].waktu.tanggal = tanggal
	X[*N1].waktu.bulan = bulan
	X[*N1].waktu.tahun = tahun
	X[*N1].jenis = "hapus"
	X[*N1].nilai = 0 - T[idx].stok
	X[*N1].jumlahWaktu = (365 * tahun) + (30 * bulan) + tanggal
	*N1++

	for i := idx; i < *N; i++ {
		T[i] = T[i+1]
	}
	fmt.Println("\nsukses dihapus")
	*N--
}

func modifikasiUbah(T *arrBarang, X *arrTransaksi, N, N1 *int, idx int) {
	var pilihan, tanggal, bulan, tahun, temp int
	var a string
	var b, c int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("modifikasi:")
	fmt.Println("1. nama barang")
	fmt.Println("2. harga barang")
	fmt.Println("3. stok")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	if pilihan == 1 {
		fmt.Print("ubah nama: ")
		fmt.Scan(&a)

		fmt.Println(idx)

		T[idx].namaBarang = a
		fmt.Println("\nsukses diubah")

	} else if pilihan == 2 {
		fmt.Print("ubah harga: ")
		fmt.Scan(&b)
		T[idx].harga = b
		fmt.Println("\nsukses diubah")

	} else if pilihan == 3 {
		fmt.Print("ubah stok: ")
		fmt.Scan(&c)
		fmt.Print("waktu (DD MM YYYY): ")
		fmt.Scan(&tanggal, &bulan, &tahun)
		temp = T[idx].stok
		T[idx].stok = c
		fmt.Println("\nsukses diubah")

		X[*N1].namaBarang = T[idx].namaBarang
		X[*N1].waktu.tanggal = tanggal
		X[*N1].waktu.bulan = bulan
		X[*N1].waktu.tahun = tahun
		X[*N1].jenis = "ubah"
		X[*N1].nilai = c - temp
		X[*N1].jumlahWaktu = (365 * tahun) + (30 * bulan) + tanggal
		*N1++
	}
}

func cari(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("cari:")
	fmt.Println("1. nama barang")
	fmt.Println("2. harga barang")
	fmt.Println("3. stok")
	fmt.Println("4. transaksi")
	fmt.Println("5. kembali")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	if pilihan == 1 {
		cariNama(&*T, &*X, &*N, &*N1)
	} else if pilihan == 2 {
		cariHarga(&*T, &*X, &*N, &*N1)
	} else if pilihan == 3 {
		cariStok(&*T, &*X, &*N, &*N1)
	} else if pilihan == 4 {
		cariTransaksi(&*T, &*X, &*N, &*N1)
	} else if pilihan == 5 {
		menu(&*T, &*X, &*N, &*N1)
	}
}

func cariNama(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var nama string
	var j int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("====================")
	fmt.Print("nama barang: ")
	fmt.Scan(&nama)
	fmt.Print("\n")

	for i := 0; i < *N; i++ {
		if T[i].namaBarang == nama {
			if len(T[i].namaBarang) <= 6 && len(T[i].namaBarang) >= 1 {
				if j == 0 {
					fmt.Println("nama", "\t\t", "harga", "\t", "stok")
				}
				fmt.Println(T[i].namaBarang, "\t\t", T[i].harga, "\t", T[i].stok)
			} else if len(T[i].namaBarang) >= 7 {
				if j == 0 {
					fmt.Println("nama", "\t\t", "harga", "\t", "stok")
				}
				fmt.Println(T[i].namaBarang, "\t", T[i].harga, "\t", T[i].stok)
			}
			j++
		}
	}

	if j == 0 {
		fmt.Print("\ntidak ada barang")
	}
}

func cariHarga(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan, j, harga, idx int
	var found bool

	pilihan = cariEkstrim()

	if pilihan == 1 {
		idx = idxMaksimum(*T, *X, *N, *N1, 1)
		if len(T[idx].namaBarang) <= 6 && len(T[idx].namaBarang) >= 1 {
			fmt.Println("nama", "\t", "harga", "\t", "stok")
			fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
			fmt.Print("\n")
		} else if len(T[idx].namaBarang) >= 7 {
			fmt.Println("nama", "\t\t", "harga", "\t", "stok")
			fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
			fmt.Print("\n")
		}

	} else if pilihan == 2 {
		idx = idxMinimum(*T, *X, *N, *N1, 1)
		if len(T[idx].namaBarang) <= 6 && len(T[idx].namaBarang) >= 1 {
			fmt.Println("nama", "\t", "harga", "\t", "stok")
			fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
			fmt.Print("\n")
		} else if len(T[idx].namaBarang) >= 7 {
			fmt.Println("nama", "\t\t", "harga", "\t", "stok")
			fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
			fmt.Print("\n")
		}

	} else if pilihan == 3 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("harga barang: ")
		fmt.Scan(&harga)
		fmt.Print("\n")

		for i := 0; i < *N; i++ {
			if T[i].harga == harga {
				found = true

			}
		}

		if found {
			fmt.Println("nama", "\t\t", "harga", "\t", "stok")
			for i := 0; i < *N; i++ {
				if T[i].harga == harga {
					if len(T[i].namaBarang) <= 6 && len(T[i].namaBarang) >= 1 {
						fmt.Println(T[i].namaBarang, "\t\t", T[i].harga, "\t", T[i].stok)
					} else if len(T[idx].namaBarang) >= 7 {
						fmt.Println(T[i].namaBarang, "\t", T[i].harga, "\t", T[i].stok)
					}
					j++
				}
			}
		}

		if j == 0 {
			fmt.Print("\ntidak ada barang")
		}
	} else if pilihan == 4 {
		cari(&*T, &*X, &*N, &*N1)
	}
}

func cariStok(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan, j, stok, idx int
	var found bool

	pilihan = cariEkstrim()

	if pilihan == 1 {
		idx = idxMaksimum(*T, *X, *N, *N1, 2)
		if len(T[idx].namaBarang) <= 6 && len(T[idx].namaBarang) >= 1 {
			fmt.Println("nama", "\t", "harga", "\t", "stok")
			fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
			fmt.Print("\n")
		} else if len(T[idx].namaBarang) >= 7 {
			fmt.Println("nama", "\t\t", "harga", "\t", "stok")
			fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
			fmt.Print("\n")
		}
	} else if pilihan == 2 {
		idx = idxMinimum(*T, *X, *N, *N1, 2)
		if len(T[idx].namaBarang) <= 6 && len(T[idx].namaBarang) >= 1 {
			fmt.Println("nama", "\t", "harga", "\t", "stok")
			fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
			fmt.Print("\n")
		} else if len(T[idx].namaBarang) >= 7 {
			fmt.Println("nama", "\t\t", "harga", "\t", "stok")
			fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
			fmt.Print("\n")
		}
	} else if pilihan == 3 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("stok: ")
		fmt.Scan(&stok)
		fmt.Print("\n")

		for i := 0; i < *N; i++ {
			if T[i].stok == stok {
				found = true
			}
		}

		if found {
			fmt.Println("nama", "\t\t", "harga", "\t", "stok")
			for i := 0; i < *N; i++ {
				if T[i].stok == stok {
					if len(T[i].namaBarang) <= 6 && len(T[i].namaBarang) >= 1 {
						fmt.Println(T[i].namaBarang, "\t\t", T[i].harga, "\t", T[i].stok)
					} else if len(T[i].namaBarang) >= 7 {
						fmt.Println(T[i].namaBarang, "\t", T[i].harga, "\t", T[i].stok)
					}
					j++
				}
			}
		}

		if j == 0 {
			fmt.Print("tidak ada barang\n")
		}
	} else if pilihan == 4 {
		cari(&*T, &*X, &*N, &*N1)
	}
}

func cariTransaksi(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("cari data transaksi:")
	fmt.Println("1. nama barang")
	fmt.Println("2. waktu transaksi")
	fmt.Println("3. jenis")
	fmt.Println("4. nilai")
	fmt.Println("5. kembali")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	if pilihan == 1 {
		cariTransaksiNama(&*T, &*X, &*N, &*N1)
	} else if pilihan == 2 {
		cariTransaksiWaktu(&*T, &*X, &*N, &*N1)
	} else if pilihan == 3 {
		cariTransaksiJenis(&*T, &*X, &*N, &*N1)
	} else if pilihan == 4 {
		cariTransaksiNilai(&*T, &*X, &*N, &*N1)
	} else if pilihan == 5 {
		cari(&*T, &*X, &*N, &*N1)
	}
}

func cariTransaksiNama(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var nama string
	var j int
	var found bool

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("====================")
	fmt.Print("nama barang: ")
	fmt.Scan(&nama)
	fmt.Print("\n")

	for i := 0; i < *N; i++ {
		if X[i].namaBarang == nama {
			found = true
		}
	}

	if found {
		fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
		for i := 0; i < *N; i++ {
			if X[i].namaBarang == nama {
				if len(X[i].namaBarang) <= 6 && len(X[i].namaBarang) >= 1 {
					fmt.Println(X[i].namaBarang, "\t\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
				} else if len(X[i].namaBarang) >= 7 {
					fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t1", X[i].waktu.tahun, "\t", X[i].nilai)
				}
				j++
			}
		}
	}

	if j == 0 {
		fmt.Print("\ntidak ada transaksi")
	}
}

func cariTransaksiWaktu(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan, idx, tanggal, bulan, tahun, waktu, j int
	var found bool

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("cari:")
	fmt.Println("1. nilai maksimum")
	fmt.Println("2. nilai minimum")
	fmt.Println("3. lainnya")
	fmt.Println("4.kembali")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	if pilihan == 1 {
		idx = idxMaksimum(*T, *X, *N, *N1, 3)
		if len(X[idx].namaBarang) <= 6 && len(X[idx].namaBarang) >= 1 {
			fmt.Println("\nnama", "\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].waktu.tanggal, "\t\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun, "\t", X[idx].nilai)
		} else if len(X[idx].namaBarang) >= 7 {
			fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].waktu.tanggal, "\t\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun, "\t", X[idx].nilai)
		}

	} else if pilihan == 2 {
		idx = idxMinimum(*T, *X, *N, *N1, 3)
		if len(X[idx].namaBarang) <= 6 && len(X[idx].namaBarang) >= 1 {
			fmt.Println("\nnama", "\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].waktu.tanggal, "\t\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun, "\t", X[idx].nilai)
		} else if len(X[idx].namaBarang) >= 7 {
			fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].waktu.tanggal, "\t\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun, "\t", X[idx].nilai)
		}

	} else if pilihan == 3 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("waktu (DD MM YYYY): ")
		fmt.Scan(&tanggal, &bulan, &tahun)
		fmt.Print("\n")

		waktu = (365 * tahun) + (30 * bulan) + tanggal

		for i := 0; i < *N; i++ {
			if X[i].jumlahWaktu == waktu {
				found = true
			}
		}

		if found {
			fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			for i := 0; i < *N; i++ {
				if X[i].jumlahWaktu == waktu {
					if len(X[i].namaBarang) <= 6 && len(X[i].namaBarang) >= 1 {
						fmt.Println(X[i].namaBarang, "\t\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
					} else if len(X[i].namaBarang) >= 7 {
						fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
					}
					j++
				}
			}
		}

		if j == 0 {
			fmt.Print("\ntidak ada barang")
		}
	} else if pilihan == 4 {
		cariTransaksi(&*T, &*X, &*N, &*N1)
	}
}

func cariTransaksiNilai(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan, idx, nilai, j int
	var found bool

	pilihan = cariEkstrim()

	if pilihan == 1 {
		idx = idxMaksimum(*T, *X, *N, *N1, 4)
		if len(X[idx].namaBarang) <= 6 && len(X[idx].namaBarang) >= 1 {
			fmt.Println("\nnama", "\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].waktu.tanggal, "\t\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun, "\t", X[idx].nilai)
		} else if len(X[idx].namaBarang) >= 7 {
			fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].waktu.tanggal, "\t\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun, "\t", X[idx].nilai)
		}

	} else if pilihan == 2 {
		idx = idxMinimum(*T, *X, *N, *N1, 4)
		if len(X[idx].namaBarang) <= 6 && len(X[idx].namaBarang) >= 1 {
			fmt.Println("\nnama", "\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].waktu.tanggal, "\t\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun, "\t", X[idx].nilai)
		} else if len(X[idx].namaBarang) >= 7 {
			fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].waktu.tanggal, "\t\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun, "\t", X[idx].nilai)
		}

	} else if pilihan == 3 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("nilai: ")
		fmt.Scan(&nilai)
		fmt.Print("\n")

		for i := 0; i < *N1; i++ {
			if X[i].nilai == nilai {
				found = true
			}
		}

		if found {
			fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			for i := 0; i < *N1; i++ {
				if X[i].nilai == nilai {
					if len(X[i].namaBarang) <= 6 && len(X[i].namaBarang) >= 1 {
						fmt.Println(X[i].namaBarang, "\t\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
					} else if len(X[i].namaBarang) >= 7 {
						fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
					}
					j++
				}
			}
		}

		if j == 0 {
			fmt.Print("\ntidak ada barang")
		}
	} else if pilihan == 4 {
		cariTransaksi(&*T, &*X, &*N, &*N1)
	}
}

func cariTransaksiJenis(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan int
	var found bool

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("cari jenis data transaksi:")
	fmt.Println("1. tambah data")
	fmt.Println("2. ubah data")
	fmt.Println("3. hapus data")
	fmt.Println("4. kembali")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	if pilihan == 1 {
		for i := 0; i < *N1; i++ {
			if X[i].jenis == "tambah" {
				found = true
			}
		}

		if found {
			fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			for i := 0; i < *N1; i++ {
				if X[i].jenis == "tambah" {
					if len(T[i].namaBarang) <= 6 && len(T[i].namaBarang) >= 1 {
						fmt.Println(X[i].namaBarang, "\t\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
					} else if len(T[i].namaBarang) >= 7 {
						fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
					}
				}
			}
		}
	} else if pilihan == 2 {
		for i := 0; i < *N1; i++ {
			if X[i].jenis == "ubah" {
				found = true
			}
		}

		if found {
			fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			for i := 0; i < *N1; i++ {
				if X[i].jenis == "ubah" {
					if len(X[i].namaBarang) <= 6 && len(X[i].namaBarang) >= 1 {
						fmt.Println(X[i].namaBarang, "\t\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
					} else if len(X[i].namaBarang) >= 7 {
						fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
					}
				}
			}
		}
	} else if pilihan == 3 {
		for i := 0; i < *N1; i++ {
			if X[i].jenis == "hapus" {
				found = true
			}
		}

		if found {
			fmt.Println("\nnama", "\t\t", "jenis", "\t", "tanggal", "\t", "bulan", "\t", "tahun", "\t", "nilai")
			for i := 0; i < *N1; i++ {
				if X[i].jenis == "hapus" {
					if len(X[i].namaBarang) <= 6 && len(X[i].namaBarang) >= 1 {
						fmt.Println(X[i].namaBarang, "\t\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
					} else if len(X[i].namaBarang) >= 7 {
						fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].waktu.tanggal, "\t\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun, "\t", X[i].nilai)
					}
				}
			}
		}
	} else if pilihan == 4 {
		cariTransaksi(&*T, &*X, &*N, &*N1)
	}
}

func cariEkstrim() int {
	var pilihan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("cari:")
	fmt.Println("1. nilai maksimum")
	fmt.Println("2. nilai minimum")
	fmt.Println("3. lainnya")
	fmt.Println("4. kembali")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	return pilihan
}

func tampilan(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan, pilihPengurutan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("tampilan:")
	fmt.Println("1. data barang")
	fmt.Println("2. transaksi")
	fmt.Println("3. kembali")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	if pilihan == 1 || pilihan == 2 {
		cariPengurutan(&*T, &*X, &*N, &*N1, &pilihPengurutan)
	}

	if pilihan == 1 {
		tampilanDataBarang(&*T, &*X, &*N, &*N1, pilihPengurutan)
	} else if pilihan == 2 {
		tampilanDataTransaksi(&*T, &*X, &*N, &*N1, pilihPengurutan)
	} else if pilihan == 3 {
		menu(&*T, &*X, &*N, &*N1)
	}
}

func cariPengurutan(T *arrBarang, X *arrTransaksi, N, N1 *int, pilihPengurutan *int) {

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("pengurutan:")
	fmt.Println("1. ascending")
	fmt.Println("2. descending")
	fmt.Println("3. kembali")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&*pilihPengurutan)
	fmt.Print("\n")
	if *pilihPengurutan == 3 {
		tampilan(&*T, &*X, &*N, &*N1)
	}
}

func tampilanDataBarang(T *arrBarang, X *arrTransaksi, N, N1 *int, x int) {
	var pilihan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("tampilan data barang:")
	fmt.Println("1. harga barang")
	fmt.Println("2. stok barang")
	fmt.Println("3. kembali")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	if pilihan == 1 {
		if x == 1 {
			ascending(&*T, &*X, *N, *N1, 1)
		} else if x == 2 {
			descending(&*T, &*X, *N, *N1, 1)
		}
		print(*T, *X, *N, *N1, 1)
	} else if pilihan == 2 {
		if x == 1 {
			ascending(&*T, &*X, *N, *N1, 2)
		} else if x == 2 {
			descending(&*T, &*X, *N, *N1, 2)
		}
		print(*T, *X, *N, *N1, 1)
	} else if pilihan == 3 {
		tampilan(&*T, &*X, &*N, &*N1)
	}
}

func tampilanDataTransaksi(T *arrBarang, X *arrTransaksi, N, N1 *int, x int) {
	var pilihan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("tampilan data transaksi:")
	fmt.Println("1. waktu")
	fmt.Println("2. nilai")
	fmt.Println("3. kembali")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	if pilihan == 1 {
		if x == 1 {
			ascending(&*T, &*X, *N, *N1, 3)
		} else if x == 2 {
			descending(&*T, &*X, *N, *N1, 3)
		}
		print(*T, *X, *N, *N1, 2)
	} else if pilihan == 2 {
		if x == 1 {
			ascending(&*T, &*X, *N, *N1, 4)
		} else if x == 2 {
			descending(&*T, &*X, *N, *N1, 4)
		}
		print(*T, *X, *N, *N1, 2)
	} else if pilihan == 3 {
		tampilan(&*T, &*X, &*N, &*N1)
	}
}

func exit() int {
	var pilihan int

	fmt.Print("\n")
	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("1. menu")
	fmt.Println("2. exit")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
	fmt.Print("\n")

	return pilihan
}

func main() {
	var dataBarang arrBarang
	var dataTransaksi arrTransaksi
	var NBarang, NTransaksi int
	inputDataBarang(&dataBarang, &NBarang)
	inputDataTransaksi(&dataTransaksi, &NTransaksi)
	menu(&dataBarang, &dataTransaksi, &NBarang, &NTransaksi)

	for exit() == 1 {
		menu(&dataBarang, &dataTransaksi, &NBarang, &NTransaksi)
	}
}
