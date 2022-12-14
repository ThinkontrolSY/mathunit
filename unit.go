package mathunit

type Unit interface {
	// String returns the name of the unit
	UnitType() string
	String() string

	Coefficient() float64

	Values() []string

	Valid() bool

	Mul(Unit) (Unit, float64)
	Div(Unit) (Unit, float64)
}

func Group() map[string][]string {
	ug := make(map[string][]string)
	ug["angle"] = Angle("").Values()
	ug["area"] = Area("").Values()
	ug["energy"] = Energy("").Values()
	ug["digital_storage"] = DigitalStorage("").Values()
	ug["frequency"] = Frequency("").Values()
	ug["length"] = Length("").Values()
	ug["mass"] = Mass("").Values()
	ug["package_unit"] = PackageUnit("").Values()
	ug["volume"] = Volume("").Values()
	ug["duration"] = Duration("").Values()
	ug["speed"] = Speed("").Values()
	ug["pressure"] = Pressure("").Values()
	return ug
}

func Parse(src string) Unit {
	if u := Angle(src); u.Valid() {
		return u
	}
	if u := Area(src); u.Valid() {
		return u
	}
	if u := Energy(src); u.Valid() {
		return u
	}
	if u := DigitalStorage(src); u.Valid() {
		return u
	}
	if u := Frequency(src); u.Valid() {
		return u
	}
	if u := Length(src); u.Valid() {
		return u
	}
	if u := Mass(src); u.Valid() {
		return u
	}
	if u := PackageUnit(src); u.Valid() {
		return u
	}
	if u := Volume(src); u.Valid() {
		return u
	}
	if u := Duration(src); u.Valid() {
		return u
	}
	if u := Speed(src); u.Valid() {
		return u
	}
	return nil
}
