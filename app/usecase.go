package app

import "log"

type Usecase struct {
	repo IRepository
}

func NewUsecase(repo IRepository) *Usecase {
	return &Usecase{repo: repo}
}

func (u *Usecase) GetProvince() (map[string]interface{}, error) {
	infos, _ := u.repo.DecodeJSONUrl()
	//count infected ppl in each province
	pvncs := make(map[string]int)
	for _, data := range infos.Data {
		_, exist := pvncs[data.ProvinceEn]
		if exist {
			pvncs[data.ProvinceEn]++
		} else {
			pvncs[data.ProvinceEn] = 1
		}

	}

	//reformat data
	var provinces []map[string]int
	for k, j := range pvncs {
		if k == "" {
			k = "N/A"
		}
		province := map[string]int{
			k: j,
		}
		provinces = append(provinces, province)
	}

	info := map[string]interface{}{
		"Province": provinces,
	}

	log.Printf("info: %+v", info)

	return info, nil
}

//count infected ages
func (u *Usecase) GetAges() (*Ages, error) {
	infos, _ := u.repo.DecodeJSONUrl()
	//group the infected ages
	var a60, a3060, a30, na int
	for _, info := range infos.Data {
		switch {
		case info.Age >= 61:
			a60++
		case info.Age >= 31 && info.Age <= 60:
			a3060++
		case info.Age <= 30:
			a30++
		default:
			na++
		}
	}

	//reformat data
	ages := Ages{
		Below30:      a30,
		Above30lt60:  a3060,
		Above60:      a60,
		NotAvailable: na,
	}

	return &ages, nil
}
