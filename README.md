# jsonreader
json reader

    obj := ParseJsonObject([]byte("{\"contacts\":{ \"company\":{ \"address\":null, \"state_code\":null }, \"employees\":[ null] }}"))
	v := obj.V("contacts").ParseJsonObject().A("employees")
	if v.Size() == 0 {
		t.Error(" parse error")
	} else {
		t.Log(v.Get(0).String(), v.Get(0).IsNull())
	}
	
	
    arr := ParseJsonArray([]byte("[{\"data\" : {\"key\" : 123,\"abc\":-1021e5 } , \"value\":5 ,\"ars\":[1,2,3,4,{\"value\":5},\"6\",{\"value\":7}]},0,1,\"sfdada\"]"))
    arr.Each(func(i int, value JsonValue) {
        if i == 0 {
            key := value.ParseJsonObject().GetObject("data").GetValue("key")
            if key != "123" {
                t.Error("get key value!=123")
            }
        }
        t.Log(i, value)
    })