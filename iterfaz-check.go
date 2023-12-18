package structs

func CheckInterfaces(pkg_name string, struct_in any) (err string) {
	const e = "CheckInterfaces "
	nils := NilInterfaces(struct_in)
	// fmt.Println("NILS:", nils)
	names := GetFieldNames(struct_in)

	// fmt.Println("NAMES:", names)
	for _, name := range names {
		for _, n := range nils {
			if name == n {
				return e + "en " + pkg_name + "handler nil: " + name
			}
		}
	}

	return ""
}
