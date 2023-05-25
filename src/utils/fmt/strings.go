package fmt

func (s *Str) Listify(str string) []string {

	NewStr := make([]string, 0)
	
	for i := 0; i < len(str); i++ {
		
		NewStr = append(NewStr, string(str[i]))

	}

	return NewStr

}

func (s *Str) SplitAtChar(str string, seperator string) []string {

	list := s.Listify(str)
	NewList := make([]string, 0)
	NewStr := ""

	for _, v := range list {
		
		if (v != seperator){
			NewStr += v
		} else {
			NewList = append(NewList, NewStr)
			NewStr = ""
		}

	}

	return NewList

}