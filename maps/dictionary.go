package maps



type Dictionary map[string]string
type DicErrString string

func (e DicErrString) Error() string {
    return string(e)
}
const (
	ErrorNotFound = DicErrString("couldn't konw what you want search")
	ErrorWordExist = DicErrString("can not add new word because words is exist in dic")
	ErrNonExistWord = DicErrString("can not update because words is not exist in dic")
)
// var (
// 	ErrorNotFound = errors.New(DicErrString)
// 	ErrorWordExist = errors.New()
// )

func (d Dictionary) Search(seachStr string) (res string, err error) {
	res, ok := d[seachStr]
	if !ok {
		return res, ErrorNotFound
	}
	return res, nil
}

func (d Dictionary) Add(key, val string) (err error) {
	_, err= d.Search(key)
	switch err {
	case ErrorNotFound:
		d[key] = val
		err = nil
	case nil:
		err = ErrorWordExist
	default:
		err = nil
	}
	return err
}

func (d Dictionary) Update(key, val string) (err error) {
	_, err= d.Search(key)
	switch err {
	case ErrorNotFound:
		err = ErrNonExistWord
	case nil:
		d[key] = val
	default:
		err = ErrNonExistWord
	}
	return err
}


func (d Dictionary) Delete(key string) {
	delete(d, key)
}

