package score

func CalculateScore(level int, trys int, time int, maxTime int) int64 {

	score := max(100, (level*200)+(max(100, (1000-maxTime)-(time*2)))-(trys*20))
	return int64(score)
}
