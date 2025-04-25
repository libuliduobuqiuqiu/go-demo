package godemo

import "fmt"

type InfoInterface interface {
	GetName() string
}

type CompanyInfo struct{}

func (c CompanyInfo) GetName() string {
	return "company"
}

func GetOpts(opts ...InfoInterface) {
	for _, v := range opts {
		fmt.Println(v.GetName())
	}
}

func GetOptsByGenericasByGenericas[T InfoInterface](opts ...T) {
	for _, v := range opts {
		fmt.Println(v.GetName())
	}

}

func CompanyGetOpts() {
	c := []CompanyInfo{}

	// Not allow: GetOpts(c), because GetOpts allow []InfoInterface, not allow []CompanyInfo
	// Go not allow slice covariance

	// You can convert it to an interface type slice.
	b := []InfoInterface{}
	for _, v := range c {
		b = append(b, v)
	}

	// Other: you can directyl use generics.

}
