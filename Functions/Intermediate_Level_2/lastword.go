package piscine


func LastWord(s string) string{
	arr:=Split1(s)
	if len(arr)==0{
		return ""+"\n"
	}
	return arr[len(arr)-1]+"\n"
}

func Split1(s string) []string{
	res:=""
	slice :=[]string{}
	for _,r:=range s {
		if r==' '{
			if res!=""{
				slice=append(slice, res)
				res=""
			}else{
				continue
			}
		}else{
			res+=string(r)
		}
	}
	if res!=""{
		slice=append(slice, res)
	}
	return slice
}