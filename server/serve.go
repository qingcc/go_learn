package main

func main() {
	//region Remark: 每日凌晨0点0分0秒更新模拟交易状态 Author; chijian
	//for {
	//	var flag = true
	//	if time.Now().Hour() == 0 && time.Now().Minute() == 0 && time.Now().Second() == 0 { //凌晨 00:00:00
	//		fmt.Println("")
	//		list := make([]*models.CostTradingContest, 0)
	//		err := databases.Orm.Where("status = 1").Find(&list)
	//		if err != nil {
	//			fmt.Println(err.Error())
	//		}
	//		fmt.Println("contest status = 1, num: ", len(list))
	//		var fail_num int64
	//		fail_num = 0
	//		for _, value := range list {
	//			if value.EndTime.Before(time.Now()) {
	//				value.Status = 2 //已结束
	//				has, err := databases.Orm.Id(value.Id).Cols("status").Update(value)
	//				if err != nil {
	//					fmt.Println(err.Error())
	//				}
	//				if has < 1 {
	//					fail_num++
	//					fmt.Println("update contest failed! num is: " + strconv.FormatInt(fail_num, 64))
	//				}
	//			}
	//			fmt.Println(value.Id, value.Title, "status:", value.Status)
	//		}
	//		if fail_num == 0 {
	//			fmt.Println("update contest success! time: " + time.Now().Format("2006-01-02 15:04:05"))
	//			fmt.Println("update contest success! updated num: " + strconv.Itoa(len(list)))
	//		}
	//		flag = false
	//	} else {
	//		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	//	}
	//	if flag {
	//		time.Sleep(time.Second)
	//	} else {
	//		fmt.Println("now sleep 1 day")
	//		time.Sleep(time.Hour * 24)
	//	}
	//}
	//endregion
}
